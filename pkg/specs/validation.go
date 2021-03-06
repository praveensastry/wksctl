package specs

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/launcher/pkg/kubectl"
	"github.com/weaveworks/wksctl/pkg/addons"
	baremetalspecv1 "github.com/weaveworks/wksctl/pkg/baremetalproviderspec/v1alpha1"
	"github.com/weaveworks/wksctl/pkg/utilities/path"
	"k8s.io/apimachinery/pkg/util/validation/field"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

func clusterPath(args ...string) *field.Path {
	return field.NewPath("cluster", args...)
}

func clusterProviderPath(args ...string) *field.Path {
	allArgs := []string{"spec", "providerSpec", "value"}
	allArgs = append(allArgs, args...)
	return clusterPath(allArgs...)
}

func clusterSpec(cluster *clusterv1.Cluster) (*baremetalspecv1.BareMetalClusterProviderSpec, error) {
	codec, err := baremetalspecv1.NewCodec()
	if err != nil {
		return nil, err
	}
	clusterSpec, err := codec.ClusterProviderFromProviderSpec(cluster.Spec.ProviderSpec)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal cluster's providerSpec field: %v", err)
	}

	return clusterSpec, err
}

func populateNetwork(cluster *clusterv1.Cluster) {
	if cluster.Spec.ClusterNetwork.ServiceDomain == "" {
		cluster.Spec.ClusterNetwork.ServiceDomain = "cluster.local"
	}
}

func updateClusterProviderSpec(cluster *clusterv1.Cluster, updateFunc func(spec *baremetalspecv1.BareMetalClusterProviderSpec)) {
	spec, err := clusterSpec(cluster)
	if err != nil {
		return
	}

	updateFunc(spec)

	// We are changing ClusterProviderSpecs, the per-provider part of the
	// manifest. `spec` is a decoded copy of that portion of the manifest and
	// needs to be re-encoded and replace the original version.
	codec, err := baremetalspecv1.NewCodec()
	if err != nil {
		log.Fatal("couldn't create codec: ", err)
	}
	encodedSpec, err := codec.EncodeToProviderSpec(spec)
	if err != nil {
		log.Fatal("error encoding spec: ", err)
	}
	cluster.Spec.ProviderSpec = *encodedSpec
}

func resolveFilePath(filePath string) (string, error) {
	if strings.HasPrefix(filePath, "~/") {
		home, err := path.UserHomeDirectory()
		if err != nil {
			return "", errors.Wrap(err, "error getting home directory")
		}
		return filepath.Join(home, filePath[2:]), nil
	}

	return filepath.Abs(filePath)
}

func expandSSHKeyPath(cluster *clusterv1.Cluster) {
	updateClusterProviderSpec(cluster, func(spec *baremetalspecv1.BareMetalClusterProviderSpec) {
		sshKeyPath, err := resolveFilePath(spec.SSHKeyPath)
		if err != nil {
			log.WithError(err).Fatalf("Failed to resolve sshKeyPath")
		}
		spec.SSHKeyPath = sshKeyPath
	})
}

type clusterValidationFunc func(*clusterv1.Cluster, string) field.ErrorList

func isValidCIDR(s string) (*net.IPNet, error) {
	ip, cidr, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	// Check we don't have a host address.
	if !ip.Equal(cidr.IP) {
		return nil, fmt.Errorf("network CIDR IP required but host IP specified")
	}
	return cidr, nil
}

func networksIntersect(n1, n2 *net.IPNet) bool {
	return n2.Contains(n1.IP) || n1.Contains(n2.IP)
}

func validateCIDRBlocks(cluster *clusterv1.Cluster, manifestPath string) field.ErrorList {
	var errors field.ErrorList
	const (
		services = 0
		pods     = 1
	)

	blocks := []struct {
		path    []string
		field   []string
		network *net.IPNet
	}{{
		path:  []string{"spec", "clusterNetwork", "services", "cidrBlocks"},
		field: cluster.Spec.ClusterNetwork.Services.CIDRBlocks,
	}, {
		path:  []string{"spec", "clusterNetwork", "pods", "cidrBlocks"},
		field: cluster.Spec.ClusterNetwork.Pods.CIDRBlocks,
	}}

	for i := range blocks {
		block := &blocks[i]
		if len(block.field) != 1 {
			errors = append(errors, field.Invalid(
				clusterPath(block.path...),
				block.field,
				"CIDR blocks must contain exactly one IP range"),
			)
			continue
		}

		network, err := isValidCIDR(block.field[0])
		if err != nil {
			errors = append(errors, field.Invalid(
				clusterPath(block.path...),
				block.field,
				fmt.Sprintf("invalid CIDR: \"%s\": %v", block.field[0], err)),
			)
			continue
		}
		block.network = network
	}

	if len(errors) > 0 {
		return errors
	}

	if networksIntersect(blocks[services].network, blocks[pods].network) {
		return field.ErrorList{
			field.Invalid(
				clusterPath("spec", "clusterNetwork", "services", "cidrBlocks"),
				blocks[services].field[0],
				fmt.Sprintf("services network overlaps with pod network (\"%s\")", blocks[pods].field[0])),
		}
	}

	return field.ErrorList{}
}

func validateServiceDomain(cluster *clusterv1.Cluster, manifestPath string) field.ErrorList {
	f := cluster.Spec.ClusterNetwork.ServiceDomain
	if f != "cluster.local" {
		return field.ErrorList{
			field.Invalid(
				clusterPath("spec", "clusterNetwork", "serviceDomain"), f,
				fmt.Sprintf("unsupported non-default service domain \"%s\", use \"cluster.local\"", f)),
		}
	}

	return field.ErrorList{}
}

func fileExists(s string) bool {
	_, err := os.Stat(s)
	return err == nil
}

func validateSSHKey(cluster *clusterv1.Cluster, manifestPath string) field.ErrorList {
	spec, err := clusterSpec(cluster)
	if err != nil {
		log.Fatalf("Failed to parse the ClusterSpec -  %v", err)
	}
	f := spec.SSHKeyPath
	if !fileExists(f) {
		return field.ErrorList{
			field.Invalid(
				clusterProviderPath("sshKeyPath"), f,
				fmt.Sprintf("could not stat file: \"%s\"", f),
			),
		}
	}

	return field.ErrorList{}
}

func isDuration(s string) bool {
	_, err := time.ParseDuration(s)
	return err == nil
}

func addonPath(i int, args ...string) *field.Path {
	allArgs := []string{fmt.Sprintf("addons[%d]", i)}
	allArgs = append(allArgs, args...)
	return clusterProviderPath(allArgs...)
}

func validateAddons(cluster *clusterv1.Cluster, manifestPath string) field.ErrorList {
	spec, _ := clusterSpec(cluster)

	// Addons require kubectl for their manifests to be applied.
	kubectl := kubectl.LocalClient{}
	if len(spec.Addons) > 0 && !kubectl.IsPresent() {
		return field.ErrorList{
			field.Invalid(clusterProviderPath("addons"), "", "addons require kubectl to be installed"),
		}
	}

	// Validate addons and their parameters.
	for i, addonDesc := range spec.Addons {
		addon, err := addons.Get(addonDesc.Name)
		if err != nil {
			return field.ErrorList{
				field.Invalid(addonPath(i, addonDesc.Name), addonDesc.Name, err.Error()),
			}
		}
		if err := addon.ValidateOptions(&addons.BuildOptions{
			BasePath: filepath.Dir(manifestPath),
			Params:   addonDesc.Params,
		}); err != nil {
			if e, ok := err.(*addons.ParamError); ok {
				return field.ErrorList{
					field.Invalid(addonPath(i, addonDesc.Name, e.Param), addonDesc.Params[e.Param], err.Error()),
				}
			}
			return field.ErrorList{
				field.Invalid(addonPath(i, addonDesc.Name), addonDesc.Name, err.Error()),
			}
		}
	}

	return field.ErrorList{}
}

// populateCluster mutates the cluster manifest:
//   - fill in default values
//   - expand ~ and resolve relative path in SSH key path
func populateCluster(cluster *clusterv1.Cluster) {
	populateNetwork(cluster)
	expandSSHKeyPath(cluster)
}

func validateCluster(cluster *clusterv1.Cluster, manifestPath string) field.ErrorList {
	var errors field.ErrorList

	for _, f := range []clusterValidationFunc{
		validateCIDRBlocks,
		validateServiceDomain,
		validateSSHKey,
		validateAddons,
	} {
		errors = append(errors, f(cluster, manifestPath)...)
	}

	return errors
}

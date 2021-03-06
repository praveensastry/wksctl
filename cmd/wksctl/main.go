package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/weaveworks/go-checkpoint"
	"github.com/weaveworks/wksctl/cmd/wksctl/addon"
	"github.com/weaveworks/wksctl/cmd/wksctl/apply"
	"github.com/weaveworks/wksctl/cmd/wksctl/applyaddons"
	"github.com/weaveworks/wksctl/cmd/wksctl/bashcompletions"
	initpkg "github.com/weaveworks/wksctl/cmd/wksctl/init"
	"github.com/weaveworks/wksctl/cmd/wksctl/kubeconfig"
	"github.com/weaveworks/wksctl/cmd/wksctl/plan"
	"github.com/weaveworks/wksctl/cmd/wksctl/registrysynccommands"
	"github.com/weaveworks/wksctl/cmd/wksctl/version"
	"github.com/weaveworks/wksctl/cmd/wksctl/zshcompletions"
	v "github.com/weaveworks/wksctl/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:   "wksctl",
	Short: "Weave Kubernetes System CLI",

	PersistentPreRun: configureLogger,
}

var options struct {
	verbose bool
}

func configureLogger(cmd *cobra.Command, args []string) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if options.verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	rootCmd.PersistentFlags().BoolVarP(&options.verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.AddCommand(addon.Cmd)
	rootCmd.AddCommand(apply.Cmd)
	rootCmd.AddCommand(applyaddons.Cmd)
	rootCmd.AddCommand(initpkg.Cmd)
	rootCmd.AddCommand(kubeconfig.Cmd)
	rootCmd.AddCommand(plan.Cmd)
	rootCmd.AddCommand(registrysynccommands.Cmd)
	rootCmd.AddCommand(version.Cmd)

	rootCmd.AddCommand(bashcompletions.Cmd)
	rootCmd.AddCommand(zshcompletions.Cmd)

	if checkResponse, err := checkpoint.Check(&checkpoint.CheckParams{
		Product: "wksctl",
		Version: v.Version,
	}); err == nil && checkResponse.Outdated {
		log.Infof("wksctl version %s is available; please update at %s",
			checkResponse.CurrentVersion, checkResponse.CurrentDownloadURL)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

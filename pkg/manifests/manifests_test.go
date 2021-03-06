package manifests

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func TestNoDeployKey(t *testing.T) {
	co, err := cloneOptions("foo", "", "")
	assert.NoError(t, err)
	assert.Equal(t, gogit.CloneOptions{URL: "foo"}, co)

}

func TestSSHDeployKey(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.NoError(t, err)
	defer os.Remove(f.Name())
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	})
	f.Write(keyPem)
	f.Close()
	co, err := cloneOptions("url", f.Name(), "")
	assert.NoError(t, err)
	assert.NotNil(t, co.Auth)
}

func TestBranchClone(t *testing.T) {
	co, err := cloneOptions("foo", "", "develop")
	assert.NoError(t, err)
	assert.Equal(t, gogit.CloneOptions{URL: "foo", SingleBranch: true, ReferenceName: plumbing.NewBranchReferenceName("develop")}, co)
}

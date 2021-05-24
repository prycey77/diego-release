package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var builderPath string

func TestDockerLifecycleBuilder(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		var err error

		builderPath, err = gexec.Build("code.cloudfoundry.org/diego-release/dockerapplifecycle/builder")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Docker-App-Lifecycle-Builder Suite")
}

package main_test

import (
	"os/exec"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration Tests", func() {
	defer GinkgoRecover()

	if env := os.Getenv("INTEGRATION"); env != "1" && env != "true" {
		Skip("Skipping Integration Tests. To Enable, set the INTEGRATION env variable to 'true'")
	}

	var patchToRatchetCLI string

	BeforeSuite(func() {
		var err error
		patchToRatchetCLI, err = gexec.Build("github.com/chendrix/ratchet")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("golint", func() {
		fixturePath := "integrationtest/fixtures/golint"

		XIt("exits unsucessfully when the command ", func() {
			command := exec.Command(patchToRatchetCLI, fixturePath)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())

			Eventually(session.Out).Should(gbytes.Say("Ratchet `golint --set_exit_status` exited unsuccessfully"))
			Eventually(session).Should(gexec.Exit(1))
		})
	})

})

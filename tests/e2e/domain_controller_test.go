package e2e_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	digitaloceandomain "kubeform.dev/provider-digitalocean-api/apis/domain/v1alpha1"
	"kubeform.dev/provider-digitalocean-controller/tests/e2e/framework"
)

var _ = Describe("Test", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.DOMAIN, whichController) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichController flag is either `all` or `%s` but got `%s`", framework.DOMAIN, framework.DOMAIN, whichController))
		}
	})
	Describe("DigitalOcean", func() {
		Context("DomainController", func() {
			var (
				providerRef *core.Secret
				secretName  string
				domainName  string
				domain      *digitaloceandomain.Domain
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				domainName = f.GetRandomName("")
				providerRef = f.DigitalOceanProviderRef(secretName)
				domain = f.Domain(domainName, secretName)
			})

			AfterEach(func() {
				By("Deleting Domain")
				err = f.DeleteDomain(domain.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting Domain")
				f.EventuallyDomainDeleted(domain.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})

			It("should create and delete Domain successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Domain")
				err = f.CreateDomain(domain)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Domain")
				f.EventuallyDomainRunning(domain.ObjectMeta).Should(BeTrue())
			})

			It("should create, update and delete Domain successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Domain")
				err = f.CreateDomain(domain)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Domain")
				f.EventuallyDomainRunning(domain.ObjectMeta).Should(BeTrue())

				By("Updating Domain")
				err = f.UpdateDomain(domain)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Domain")
				f.EventuallyDomainRunning(domain.ObjectMeta).Should(BeTrue())
			})
		})
	})
})

/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

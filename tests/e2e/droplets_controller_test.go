/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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
	digitaloceandroplet "kubeform.dev/provider-digitalocean-api/apis/droplet/v1alpha1"
	"kubeform.dev/provider-digitalocean-controller/tests/e2e/framework"
)

var _ = Describe("Test", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.DROPLET, whichController) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichController flag is either `all` or `%s` but got `%s`", framework.DROPLET, framework.DROPLET, whichController))
		}
	})
	Describe("DigitalOcean", func() {
		Context("DropletController", func() {
			var (
				providerRef *core.Secret
				secretName  string
				dropletName string
				droplet     *digitaloceandroplet.Droplet
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				dropletName = f.GetRandomName("")
				providerRef = f.DigitalOceanProviderRef(secretName)
				droplet = f.Droplets(dropletName, secretName)
			})

			AfterEach(func() {
				By("Deleting droplet")
				err = f.DeleteDroplet(droplet.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting droplet")
				f.EventuallyDropletDeleted(droplet.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})

			It("should create and delete Droplet successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating droplet")
				err = f.CreateDroplet(droplet)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running droplet")
				f.EventuallyDropletRunning(droplet.ObjectMeta).Should(BeTrue())
			})

			It("should create, update and delete Droplet successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating droplet")
				err = f.CreateDroplet(droplet)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running droplet")
				f.EventuallyDropletRunning(droplet.ObjectMeta).Should(BeTrue())

				By("Updating Droplet")
				err = f.UpdateDroplet(droplet)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Droplet")
				f.EventuallyDropletRunning(droplet.ObjectMeta).Should(BeTrue())
			})
		})
	})
})

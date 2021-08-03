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

package framework

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCRD(whichController string) GomegaAsyncAssertion {
	return Eventually(
		func() error {
			if whichController == "all" || whichController == "droplet" {
				// Check Droplets CRD
				if _, err := f.kfClient.DropletV1alpha1().Droplets(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD Droplets is not ready")
				}
			}

			if whichController == "all" || whichController == "domain" {
				// Check Domain CRD
				if _, err := f.kfClient.DomainV1alpha1().Domains(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD Domain is not ready")
				}
			}

			return nil
		},
		time.Minute*4,
		time.Second*20,
	)
}

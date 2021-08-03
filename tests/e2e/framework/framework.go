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
	"gomodules.xyz/x/crypto/rand"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	kfclient "kubeform.dev/provider-digitalocean-api/client/clientset/versioned"
)

const (
	ALL     = "all"
	DROPLET = "droplet"
	DOMAIN  = "domain"
)

type Framework struct {
	restConfig *rest.Config
	kubeClient kubernetes.Interface
	kfClient   kfclient.Interface
	namespace  string
	name       string
}

func New(
	restConfig *rest.Config,
	kubeClient kubernetes.Interface,
	kfClient kfclient.Interface,
) *Framework {
	return &Framework{
		restConfig: restConfig,
		kubeClient: kubeClient,
		kfClient:   kfClient,
		name:       "kfc",
		namespace:  rand.WithUniqSuffix("kubeform"),
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("kfc-e2e"),
	}
}

func (fi *Invocation) GetRandomName(extraSuffix string) string {
	return rand.WithUniqSuffix(fi.name + extraSuffix)
}

func RunTest(controller, whichController string) bool {
	if whichController == ALL || controller == whichController {
		return true
	} else {
		return false
	}
}

type Invocation struct {
	*Framework
	app string
}

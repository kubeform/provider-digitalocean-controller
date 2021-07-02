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

package framework

import (
	"context"
	"time"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kubeform.dev/provider-digitalocean-api/apis/droplet/v1alpha1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

func (i *Invocation) Droplets(name string, secretName string) *v1alpha1.Droplet {
	region := "nyc1"
	size := "s-1vcpu-1gb"
	image := "ubuntu-18-04-x64"
	return &v1alpha1.Droplet{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.DropletSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: secretName,
			},
			Resource: v1alpha1.DropletSpecResource{
				Name:   &name,
				Region: &region,
				Size:   &size,
				Image:  &image,
			},
		},
	}
}

func (f *Framework) CreateDroplet(obj *v1alpha1.Droplet) error {
	_, err := f.kfClient.DropletV1alpha1().Droplets(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) UpdateDroplet(obj *v1alpha1.Droplet) error {
	obj, err := f.kfClient.DropletV1alpha1().Droplets(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	updateName := *obj.Spec.Resource.Name + "-update"
	obj.Spec.Resource.Name = &updateName

	_, err = f.kfClient.DropletV1alpha1().Droplets(obj.Namespace).Update(context.TODO(), obj, metav1.UpdateOptions{})
	return err
}

func (f *Framework) DeleteDroplet(meta metav1.ObjectMeta) error {
	return f.kfClient.DropletV1alpha1().Droplets(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInBackground())
}

func (f *Framework) EventuallyDropletRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			droplet, err := f.kfClient.DropletV1alpha1().Droplets(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return droplet.Status.Phase == status.CurrentStatus
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyDropletDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kfClient.DropletV1alpha1().Droplets(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}

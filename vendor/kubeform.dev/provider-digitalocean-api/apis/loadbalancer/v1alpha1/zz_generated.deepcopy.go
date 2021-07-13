// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loadbalancer) DeepCopyInto(out *Loadbalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loadbalancer.
func (in *Loadbalancer) DeepCopy() *Loadbalancer {
	if in == nil {
		return nil
	}
	out := new(Loadbalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Loadbalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerList) DeepCopyInto(out *LoadbalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Loadbalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerList.
func (in *LoadbalancerList) DeepCopy() *LoadbalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadbalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpec) DeepCopyInto(out *LoadbalancerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadbalancerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpec.
func (in *LoadbalancerSpec) DeepCopy() *LoadbalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpecForwardingRule) DeepCopyInto(out *LoadbalancerSpecForwardingRule) {
	*out = *in
	if in.CertificateID != nil {
		in, out := &in.CertificateID, &out.CertificateID
		*out = new(string)
		**out = **in
	}
	if in.CertificateName != nil {
		in, out := &in.CertificateName, &out.CertificateName
		*out = new(string)
		**out = **in
	}
	if in.EntryPort != nil {
		in, out := &in.EntryPort, &out.EntryPort
		*out = new(int64)
		**out = **in
	}
	if in.EntryProtocol != nil {
		in, out := &in.EntryProtocol, &out.EntryProtocol
		*out = new(string)
		**out = **in
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int64)
		**out = **in
	}
	if in.TargetProtocol != nil {
		in, out := &in.TargetProtocol, &out.TargetProtocol
		*out = new(string)
		**out = **in
	}
	if in.TlsPassthrough != nil {
		in, out := &in.TlsPassthrough, &out.TlsPassthrough
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpecForwardingRule.
func (in *LoadbalancerSpecForwardingRule) DeepCopy() *LoadbalancerSpecForwardingRule {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpecForwardingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpecHealthcheck) DeepCopyInto(out *LoadbalancerSpecHealthcheck) {
	*out = *in
	if in.CheckIntervalSeconds != nil {
		in, out := &in.CheckIntervalSeconds, &out.CheckIntervalSeconds
		*out = new(int64)
		**out = **in
	}
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ResponseTimeoutSeconds != nil {
		in, out := &in.ResponseTimeoutSeconds, &out.ResponseTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpecHealthcheck.
func (in *LoadbalancerSpecHealthcheck) DeepCopy() *LoadbalancerSpecHealthcheck {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpecHealthcheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpecResource) DeepCopyInto(out *LoadbalancerSpecResource) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.DropletIDS != nil {
		in, out := &in.DropletIDS, &out.DropletIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.DropletTag != nil {
		in, out := &in.DropletTag, &out.DropletTag
		*out = new(string)
		**out = **in
	}
	if in.EnableBackendKeepalive != nil {
		in, out := &in.EnableBackendKeepalive, &out.EnableBackendKeepalive
		*out = new(bool)
		**out = **in
	}
	if in.EnableProxyProtocol != nil {
		in, out := &in.EnableProxyProtocol, &out.EnableProxyProtocol
		*out = new(bool)
		**out = **in
	}
	if in.ForwardingRule != nil {
		in, out := &in.ForwardingRule, &out.ForwardingRule
		*out = make([]LoadbalancerSpecForwardingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Healthcheck != nil {
		in, out := &in.Healthcheck, &out.Healthcheck
		*out = new(LoadbalancerSpecHealthcheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Ip != nil {
		in, out := &in.Ip, &out.Ip
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RedirectHTTPToHTTPS != nil {
		in, out := &in.RedirectHTTPToHTTPS, &out.RedirectHTTPToHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StickySessions != nil {
		in, out := &in.StickySessions, &out.StickySessions
		*out = new(LoadbalancerSpecStickySessions)
		(*in).DeepCopyInto(*out)
	}
	if in.Urn != nil {
		in, out := &in.Urn, &out.Urn
		*out = new(string)
		**out = **in
	}
	if in.VpcUUID != nil {
		in, out := &in.VpcUUID, &out.VpcUUID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpecResource.
func (in *LoadbalancerSpecResource) DeepCopy() *LoadbalancerSpecResource {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpecStickySessions) DeepCopyInto(out *LoadbalancerSpecStickySessions) {
	*out = *in
	if in.CookieName != nil {
		in, out := &in.CookieName, &out.CookieName
		*out = new(string)
		**out = **in
	}
	if in.CookieTtlSeconds != nil {
		in, out := &in.CookieTtlSeconds, &out.CookieTtlSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpecStickySessions.
func (in *LoadbalancerSpecStickySessions) DeepCopy() *LoadbalancerSpecStickySessions {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpecStickySessions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerStatus) DeepCopyInto(out *LoadbalancerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerStatus.
func (in *LoadbalancerStatus) DeepCopy() *LoadbalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerStatus)
	in.DeepCopyInto(out)
	return out
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (c) 2022 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
	ipamv1alpha1 "github.com/onmetal/onmetal-api/api/ipam/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefix) DeepCopyInto(out *AliasPrefix) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefix.
func (in *AliasPrefix) DeepCopy() *AliasPrefix {
	if in == nil {
		return nil
	}
	out := new(AliasPrefix)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasPrefix) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefixList) DeepCopyInto(out *AliasPrefixList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AliasPrefix, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefixList.
func (in *AliasPrefixList) DeepCopy() *AliasPrefixList {
	if in == nil {
		return nil
	}
	out := new(AliasPrefixList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasPrefixList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefixRouting) DeepCopyInto(out *AliasPrefixRouting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]commonv1alpha1.LocalUIDReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefixRouting.
func (in *AliasPrefixRouting) DeepCopy() *AliasPrefixRouting {
	if in == nil {
		return nil
	}
	out := new(AliasPrefixRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasPrefixRouting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefixRoutingList) DeepCopyInto(out *AliasPrefixRoutingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AliasPrefixRouting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefixRoutingList.
func (in *AliasPrefixRoutingList) DeepCopy() *AliasPrefixRoutingList {
	if in == nil {
		return nil
	}
	out := new(AliasPrefixRoutingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasPrefixRoutingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefixSpec) DeepCopyInto(out *AliasPrefixSpec) {
	*out = *in
	out.NetworkRef = in.NetworkRef
	if in.NetworkInterfaceSelector != nil {
		in, out := &in.NetworkInterfaceSelector, &out.NetworkInterfaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Prefix.DeepCopyInto(&out.Prefix)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefixSpec.
func (in *AliasPrefixSpec) DeepCopy() *AliasPrefixSpec {
	if in == nil {
		return nil
	}
	out := new(AliasPrefixSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasPrefixStatus) DeepCopyInto(out *AliasPrefixStatus) {
	*out = *in
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasPrefixStatus.
func (in *AliasPrefixStatus) DeepCopy() *AliasPrefixStatus {
	if in == nil {
		return nil
	}
	out := new(AliasPrefixStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralPrefixSource) DeepCopyInto(out *EphemeralPrefixSource) {
	*out = *in
	if in.PrefixTemplate != nil {
		in, out := &in.PrefixTemplate, &out.PrefixTemplate
		*out = new(ipamv1alpha1.PrefixTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralPrefixSource.
func (in *EphemeralPrefixSource) DeepCopy() *EphemeralPrefixSource {
	if in == nil {
		return nil
	}
	out := new(EphemeralPrefixSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EphemeralVirtualIPSource) DeepCopyInto(out *EphemeralVirtualIPSource) {
	*out = *in
	if in.VirtualIPTemplate != nil {
		in, out := &in.VirtualIPTemplate, &out.VirtualIPTemplate
		*out = new(VirtualIPTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EphemeralVirtualIPSource.
func (in *EphemeralVirtualIPSource) DeepCopy() *EphemeralVirtualIPSource {
	if in == nil {
		return nil
	}
	out := new(EphemeralVirtualIPSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPSource) DeepCopyInto(out *IPSource) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = (*in).DeepCopy()
	}
	if in.Ephemeral != nil {
		in, out := &in.Ephemeral, &out.Ephemeral
		*out = new(EphemeralPrefixSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPSource.
func (in *IPSource) DeepCopy() *IPSource {
	if in == nil {
		return nil
	}
	out := new(IPSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerList.
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerPort) DeepCopyInto(out *LoadBalancerPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.EndPort != nil {
		in, out := &in.EndPort, &out.EndPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerPort.
func (in *LoadBalancerPort) DeepCopy() *LoadBalancerPort {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerRouting) DeepCopyInto(out *LoadBalancerRouting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.NetworkRef = in.NetworkRef
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]commonv1alpha1.LocalUIDReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerRouting.
func (in *LoadBalancerRouting) DeepCopy() *LoadBalancerRouting {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerRouting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerRoutingList) DeepCopyInto(out *LoadBalancerRoutingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerRouting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerRoutingList.
func (in *LoadBalancerRoutingList) DeepCopy() *LoadBalancerRoutingList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerRoutingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerRoutingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.IPFamilies != nil {
		in, out := &in.IPFamilies, &out.IPFamilies
		*out = make([]corev1.IPFamily, len(*in))
		copy(*out, *in)
	}
	out.NetworkRef = in.NetworkRef
	if in.NetworkInterfaceSelector != nil {
		in, out := &in.NetworkInterfaceSelector, &out.NetworkInterfaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]LoadBalancerPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]commonv1alpha1.IP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGateway) DeepCopyInto(out *NATGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGateway.
func (in *NATGateway) DeepCopy() *NATGateway {
	if in == nil {
		return nil
	}
	out := new(NATGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NATGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayDestination) DeepCopyInto(out *NATGatewayDestination) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]NATGatewayDestinationIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayDestination.
func (in *NATGatewayDestination) DeepCopy() *NATGatewayDestination {
	if in == nil {
		return nil
	}
	out := new(NATGatewayDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayDestinationIP) DeepCopyInto(out *NATGatewayDestinationIP) {
	*out = *in
	in.IP.DeepCopyInto(&out.IP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayDestinationIP.
func (in *NATGatewayDestinationIP) DeepCopy() *NATGatewayDestinationIP {
	if in == nil {
		return nil
	}
	out := new(NATGatewayDestinationIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayIP) DeepCopyInto(out *NATGatewayIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayIP.
func (in *NATGatewayIP) DeepCopy() *NATGatewayIP {
	if in == nil {
		return nil
	}
	out := new(NATGatewayIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayIPStatus) DeepCopyInto(out *NATGatewayIPStatus) {
	*out = *in
	in.IP.DeepCopyInto(&out.IP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayIPStatus.
func (in *NATGatewayIPStatus) DeepCopy() *NATGatewayIPStatus {
	if in == nil {
		return nil
	}
	out := new(NATGatewayIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayList) DeepCopyInto(out *NATGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NATGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayList.
func (in *NATGatewayList) DeepCopy() *NATGatewayList {
	if in == nil {
		return nil
	}
	out := new(NATGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NATGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayRouting) DeepCopyInto(out *NATGatewayRouting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.NetworkRef = in.NetworkRef
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]NATGatewayDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayRouting.
func (in *NATGatewayRouting) DeepCopy() *NATGatewayRouting {
	if in == nil {
		return nil
	}
	out := new(NATGatewayRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NATGatewayRouting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayRoutingList) DeepCopyInto(out *NATGatewayRoutingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NATGatewayRouting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayRoutingList.
func (in *NATGatewayRoutingList) DeepCopy() *NATGatewayRoutingList {
	if in == nil {
		return nil
	}
	out := new(NATGatewayRoutingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NATGatewayRoutingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewaySpec) DeepCopyInto(out *NATGatewaySpec) {
	*out = *in
	if in.IPFamilies != nil {
		in, out := &in.IPFamilies, &out.IPFamilies
		*out = make([]corev1.IPFamily, len(*in))
		copy(*out, *in)
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]NATGatewayIP, len(*in))
		copy(*out, *in)
	}
	out.NetworkRef = in.NetworkRef
	if in.NetworkInterfaceSelector != nil {
		in, out := &in.NetworkInterfaceSelector, &out.NetworkInterfaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PortsPerNetworkInterface != nil {
		in, out := &in.PortsPerNetworkInterface, &out.PortsPerNetworkInterface
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewaySpec.
func (in *NATGatewaySpec) DeepCopy() *NATGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(NATGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATGatewayStatus) DeepCopyInto(out *NATGatewayStatus) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]NATGatewayIPStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PortsUsed != nil {
		in, out := &in.PortsUsed, &out.PortsUsed
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATGatewayStatus.
func (in *NATGatewayStatus) DeepCopy() *NATGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(NATGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterface) DeepCopyInto(out *NetworkInterface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterface.
func (in *NetworkInterface) DeepCopy() *NetworkInterface {
	if in == nil {
		return nil
	}
	out := new(NetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceList) DeepCopyInto(out *NetworkInterfaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceList.
func (in *NetworkInterfaceList) DeepCopy() *NetworkInterfaceList {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInterfaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceSpec) DeepCopyInto(out *NetworkInterfaceSpec) {
	*out = *in
	out.NetworkRef = in.NetworkRef
	if in.MachineRef != nil {
		in, out := &in.MachineRef, &out.MachineRef
		*out = new(commonv1alpha1.LocalUIDReference)
		**out = **in
	}
	if in.IPFamilies != nil {
		in, out := &in.IPFamilies, &out.IPFamilies
		*out = make([]corev1.IPFamily, len(*in))
		copy(*out, *in)
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]IPSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualIP != nil {
		in, out := &in.VirtualIP, &out.VirtualIP
		*out = new(VirtualIPSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceSpec.
func (in *NetworkInterfaceSpec) DeepCopy() *NetworkInterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceStatus) DeepCopyInto(out *NetworkInterfaceStatus) {
	*out = *in
	if in.LastStateTransitionTime != nil {
		in, out := &in.LastStateTransitionTime, &out.LastStateTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]commonv1alpha1.IP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualIP != nil {
		in, out := &in.VirtualIP, &out.VirtualIP
		*out = (*in).DeepCopy()
	}
	if in.LastPhaseTransitionTime != nil {
		in, out := &in.LastPhaseTransitionTime, &out.LastPhaseTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceStatus.
func (in *NetworkInterfaceStatus) DeepCopy() *NetworkInterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceTemplateSpec) DeepCopyInto(out *NetworkInterfaceTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceTemplateSpec.
func (in *NetworkInterfaceTemplateSpec) DeepCopy() *NetworkInterfaceTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrefixSource) DeepCopyInto(out *PrefixSource) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = (*in).DeepCopy()
	}
	if in.Ephemeral != nil {
		in, out := &in.Ephemeral, &out.Ephemeral
		*out = new(EphemeralPrefixSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrefixSource.
func (in *PrefixSource) DeepCopy() *PrefixSource {
	if in == nil {
		return nil
	}
	out := new(PrefixSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIP) DeepCopyInto(out *VirtualIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIP.
func (in *VirtualIP) DeepCopy() *VirtualIP {
	if in == nil {
		return nil
	}
	out := new(VirtualIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIPList) DeepCopyInto(out *VirtualIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIPList.
func (in *VirtualIPList) DeepCopy() *VirtualIPList {
	if in == nil {
		return nil
	}
	out := new(VirtualIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIPSource) DeepCopyInto(out *VirtualIPSource) {
	*out = *in
	if in.VirtualIPRef != nil {
		in, out := &in.VirtualIPRef, &out.VirtualIPRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Ephemeral != nil {
		in, out := &in.Ephemeral, &out.Ephemeral
		*out = new(EphemeralVirtualIPSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIPSource.
func (in *VirtualIPSource) DeepCopy() *VirtualIPSource {
	if in == nil {
		return nil
	}
	out := new(VirtualIPSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIPSpec) DeepCopyInto(out *VirtualIPSpec) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(commonv1alpha1.LocalUIDReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIPSpec.
func (in *VirtualIPSpec) DeepCopy() *VirtualIPSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIPStatus) DeepCopyInto(out *VirtualIPStatus) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = (*in).DeepCopy()
	}
	if in.LastPhaseTransitionTime != nil {
		in, out := &in.LastPhaseTransitionTime, &out.LastPhaseTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIPStatus.
func (in *VirtualIPStatus) DeepCopy() *VirtualIPStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualIPTemplateSpec) DeepCopyInto(out *VirtualIPTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualIPTemplateSpec.
func (in *VirtualIPTemplateSpec) DeepCopy() *VirtualIPTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualIPTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

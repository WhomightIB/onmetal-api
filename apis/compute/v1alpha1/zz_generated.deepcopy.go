// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonv1alpha1 "github.com/onmetal/onmetal-api/apis/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZoneQuantity) DeepCopyInto(out *AvailabilityZoneQuantity) {
	*out = *in
	if in.Classes != nil {
		in, out := &in.Classes, &out.Classes
		*out = make([]MachineClassQuantity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZoneQuantity.
func (in *AvailabilityZoneQuantity) DeepCopy() *AvailabilityZoneQuantity {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZoneQuantity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capability) DeepCopyInto(out *Capability) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capability.
func (in *Capability) DeepCopy() *Capability {
	if in == nil {
		return nil
	}
	out := new(Capability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hash) DeepCopyInto(out *Hash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hash.
func (in *Hash) DeepCopy() *Hash {
	if in == nil {
		return nil
	}
	out := new(Hash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HashStatus) DeepCopyInto(out *HashStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HashStatus.
func (in *HashStatus) DeepCopy() *HashStatus {
	if in == nil {
		return nil
	}
	out := new(HashStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	in.ExpirationTime.DeepCopyInto(&out.ExpirationTime)
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]SourceAttribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	out.StateFields = in.StateFields
	if in.Hashes != nil {
		in, out := &in.Hashes, &out.Hashes
		*out = make([]HashStatus, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]RegionState, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClass) DeepCopyInto(out *MachineClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClass.
func (in *MachineClass) DeepCopy() *MachineClass {
	if in == nil {
		return nil
	}
	out := new(MachineClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClassList) DeepCopyInto(out *MachineClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClassList.
func (in *MachineClassList) DeepCopy() *MachineClassList {
	if in == nil {
		return nil
	}
	out := new(MachineClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClassQuantity) DeepCopyInto(out *MachineClassQuantity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClassQuantity.
func (in *MachineClassQuantity) DeepCopy() *MachineClassQuantity {
	if in == nil {
		return nil
	}
	out := new(MachineClassQuantity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClassSpec) DeepCopyInto(out *MachineClassSpec) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]Capability, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClassSpec.
func (in *MachineClassSpec) DeepCopy() *MachineClassSpec {
	if in == nil {
		return nil
	}
	out := new(MachineClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClassStatus) DeepCopyInto(out *MachineClassStatus) {
	*out = *in
	if in.Availability != nil {
		in, out := &in.Availability, &out.Availability
		*out = make(commonv1alpha1.Availability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClassStatus.
func (in *MachineClassStatus) DeepCopy() *MachineClassStatus {
	if in == nil {
		return nil
	}
	out := new(MachineClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePool) DeepCopyInto(out *MachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePool.
func (in *MachinePool) DeepCopy() *MachinePool {
	if in == nil {
		return nil
	}
	out := new(MachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePoolList) DeepCopyInto(out *MachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePoolList.
func (in *MachinePoolList) DeepCopy() *MachinePoolList {
	if in == nil {
		return nil
	}
	out := new(MachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePoolSpec) DeepCopyInto(out *MachinePoolSpec) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make([]AvailabilityZoneQuantity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePoolSpec.
func (in *MachinePoolSpec) DeepCopy() *MachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(MachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePoolStatus) DeepCopyInto(out *MachinePoolStatus) {
	*out = *in
	out.StateFields = in.StateFields
	in.Used.DeepCopyInto(&out.Used)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePoolStatus.
func (in *MachinePoolStatus) DeepCopy() *MachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(MachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionState) DeepCopyInto(out *RegionState) {
	*out = *in
	out.StateFields = in.StateFields
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionState.
func (in *RegionState) DeepCopy() *RegionState {
	if in == nil {
		return nil
	}
	out := new(RegionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKey) DeepCopyInto(out *SSHPublicKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKey.
func (in *SSHPublicKey) DeepCopy() *SSHPublicKey {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHPublicKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyList) DeepCopyInto(out *SSHPublicKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHPublicKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyList.
func (in *SSHPublicKeyList) DeepCopy() *SSHPublicKeyList {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHPublicKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeySpec) DeepCopyInto(out *SSHPublicKeySpec) {
	*out = *in
	in.ExpirationDate.DeepCopyInto(&out.ExpirationDate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeySpec.
func (in *SSHPublicKeySpec) DeepCopy() *SSHPublicKeySpec {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyStatus) DeepCopyInto(out *SSHPublicKeyStatus) {
	*out = *in
	out.StateFields = in.StateFields
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyStatus.
func (in *SSHPublicKeyStatus) DeepCopy() *SSHPublicKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceAttribute) DeepCopyInto(out *SourceAttribute) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(Hash)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceAttribute.
func (in *SourceAttribute) DeepCopy() *SourceAttribute {
	if in == nil {
		return nil
	}
	out := new(SourceAttribute)
	in.DeepCopyInto(out)
	return out
}

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	commonv1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
	v1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	storage "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis/storage"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Volume)(nil), (*storage.Volume)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Volume_To_storage_Volume(a.(*v1alpha1.Volume), b.(*storage.Volume), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.Volume)(nil), (*v1alpha1.Volume)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_Volume_To_v1alpha1_Volume(a.(*storage.Volume), b.(*v1alpha1.Volume), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeAccess)(nil), (*storage.VolumeAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeAccess_To_storage_VolumeAccess(a.(*v1alpha1.VolumeAccess), b.(*storage.VolumeAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeAccess)(nil), (*v1alpha1.VolumeAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeAccess_To_v1alpha1_VolumeAccess(a.(*storage.VolumeAccess), b.(*v1alpha1.VolumeAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeClass)(nil), (*storage.VolumeClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeClass_To_storage_VolumeClass(a.(*v1alpha1.VolumeClass), b.(*storage.VolumeClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeClass)(nil), (*v1alpha1.VolumeClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeClass_To_v1alpha1_VolumeClass(a.(*storage.VolumeClass), b.(*v1alpha1.VolumeClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeClassList)(nil), (*storage.VolumeClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeClassList_To_storage_VolumeClassList(a.(*v1alpha1.VolumeClassList), b.(*storage.VolumeClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeClassList)(nil), (*v1alpha1.VolumeClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeClassList_To_v1alpha1_VolumeClassList(a.(*storage.VolumeClassList), b.(*v1alpha1.VolumeClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeCondition)(nil), (*storage.VolumeCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeCondition_To_storage_VolumeCondition(a.(*v1alpha1.VolumeCondition), b.(*storage.VolumeCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeCondition)(nil), (*v1alpha1.VolumeCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeCondition_To_v1alpha1_VolumeCondition(a.(*storage.VolumeCondition), b.(*v1alpha1.VolumeCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeList)(nil), (*storage.VolumeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeList_To_storage_VolumeList(a.(*v1alpha1.VolumeList), b.(*storage.VolumeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeList)(nil), (*v1alpha1.VolumeList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeList_To_v1alpha1_VolumeList(a.(*storage.VolumeList), b.(*v1alpha1.VolumeList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumePool)(nil), (*storage.VolumePool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumePool_To_storage_VolumePool(a.(*v1alpha1.VolumePool), b.(*storage.VolumePool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumePool)(nil), (*v1alpha1.VolumePool)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumePool_To_v1alpha1_VolumePool(a.(*storage.VolumePool), b.(*v1alpha1.VolumePool), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumePoolCondition)(nil), (*storage.VolumePoolCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumePoolCondition_To_storage_VolumePoolCondition(a.(*v1alpha1.VolumePoolCondition), b.(*storage.VolumePoolCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumePoolCondition)(nil), (*v1alpha1.VolumePoolCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumePoolCondition_To_v1alpha1_VolumePoolCondition(a.(*storage.VolumePoolCondition), b.(*v1alpha1.VolumePoolCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumePoolList)(nil), (*storage.VolumePoolList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumePoolList_To_storage_VolumePoolList(a.(*v1alpha1.VolumePoolList), b.(*storage.VolumePoolList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumePoolList)(nil), (*v1alpha1.VolumePoolList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumePoolList_To_v1alpha1_VolumePoolList(a.(*storage.VolumePoolList), b.(*v1alpha1.VolumePoolList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumePoolSpec)(nil), (*storage.VolumePoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec(a.(*v1alpha1.VolumePoolSpec), b.(*storage.VolumePoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumePoolSpec)(nil), (*v1alpha1.VolumePoolSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec(a.(*storage.VolumePoolSpec), b.(*v1alpha1.VolumePoolSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumePoolStatus)(nil), (*storage.VolumePoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus(a.(*v1alpha1.VolumePoolStatus), b.(*storage.VolumePoolStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumePoolStatus)(nil), (*v1alpha1.VolumePoolStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus(a.(*storage.VolumePoolStatus), b.(*v1alpha1.VolumePoolStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeSpec)(nil), (*storage.VolumeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(a.(*v1alpha1.VolumeSpec), b.(*storage.VolumeSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeSpec)(nil), (*v1alpha1.VolumeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(a.(*storage.VolumeSpec), b.(*v1alpha1.VolumeSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeStatus)(nil), (*storage.VolumeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeStatus_To_storage_VolumeStatus(a.(*v1alpha1.VolumeStatus), b.(*storage.VolumeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeStatus)(nil), (*v1alpha1.VolumeStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeStatus_To_v1alpha1_VolumeStatus(a.(*storage.VolumeStatus), b.(*v1alpha1.VolumeStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.VolumeTemplateSpec)(nil), (*storage.VolumeTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VolumeTemplateSpec_To_storage_VolumeTemplateSpec(a.(*v1alpha1.VolumeTemplateSpec), b.(*storage.VolumeTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*storage.VolumeTemplateSpec)(nil), (*v1alpha1.VolumeTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_storage_VolumeTemplateSpec_To_v1alpha1_VolumeTemplateSpec(a.(*storage.VolumeTemplateSpec), b.(*v1alpha1.VolumeTemplateSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Volume_To_storage_Volume(in *v1alpha1.Volume, out *storage.Volume, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VolumeStatus_To_storage_VolumeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Volume_To_storage_Volume is an autogenerated conversion function.
func Convert_v1alpha1_Volume_To_storage_Volume(in *v1alpha1.Volume, out *storage.Volume, s conversion.Scope) error {
	return autoConvert_v1alpha1_Volume_To_storage_Volume(in, out, s)
}

func autoConvert_storage_Volume_To_v1alpha1_Volume(in *storage.Volume, out *v1alpha1.Volume, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_storage_VolumeStatus_To_v1alpha1_VolumeStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_Volume_To_v1alpha1_Volume is an autogenerated conversion function.
func Convert_storage_Volume_To_v1alpha1_Volume(in *storage.Volume, out *v1alpha1.Volume, s conversion.Scope) error {
	return autoConvert_storage_Volume_To_v1alpha1_Volume(in, out, s)
}

func autoConvert_v1alpha1_VolumeAccess_To_storage_VolumeAccess(in *v1alpha1.VolumeAccess, out *storage.VolumeAccess, s conversion.Scope) error {
	out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	out.Driver = in.Driver
	out.Handle = in.Handle
	out.VolumeAttributes = *(*map[string]string)(unsafe.Pointer(&in.VolumeAttributes))
	return nil
}

// Convert_v1alpha1_VolumeAccess_To_storage_VolumeAccess is an autogenerated conversion function.
func Convert_v1alpha1_VolumeAccess_To_storage_VolumeAccess(in *v1alpha1.VolumeAccess, out *storage.VolumeAccess, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeAccess_To_storage_VolumeAccess(in, out, s)
}

func autoConvert_storage_VolumeAccess_To_v1alpha1_VolumeAccess(in *storage.VolumeAccess, out *v1alpha1.VolumeAccess, s conversion.Scope) error {
	out.SecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	out.Driver = in.Driver
	out.Handle = in.Handle
	out.VolumeAttributes = *(*map[string]string)(unsafe.Pointer(&in.VolumeAttributes))
	return nil
}

// Convert_storage_VolumeAccess_To_v1alpha1_VolumeAccess is an autogenerated conversion function.
func Convert_storage_VolumeAccess_To_v1alpha1_VolumeAccess(in *storage.VolumeAccess, out *v1alpha1.VolumeAccess, s conversion.Scope) error {
	return autoConvert_storage_VolumeAccess_To_v1alpha1_VolumeAccess(in, out, s)
}

func autoConvert_v1alpha1_VolumeClass_To_storage_VolumeClass(in *v1alpha1.VolumeClass, out *storage.VolumeClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Capabilities = *(*v1.ResourceList)(unsafe.Pointer(&in.Capabilities))
	return nil
}

// Convert_v1alpha1_VolumeClass_To_storage_VolumeClass is an autogenerated conversion function.
func Convert_v1alpha1_VolumeClass_To_storage_VolumeClass(in *v1alpha1.VolumeClass, out *storage.VolumeClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeClass_To_storage_VolumeClass(in, out, s)
}

func autoConvert_storage_VolumeClass_To_v1alpha1_VolumeClass(in *storage.VolumeClass, out *v1alpha1.VolumeClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Capabilities = *(*v1.ResourceList)(unsafe.Pointer(&in.Capabilities))
	return nil
}

// Convert_storage_VolumeClass_To_v1alpha1_VolumeClass is an autogenerated conversion function.
func Convert_storage_VolumeClass_To_v1alpha1_VolumeClass(in *storage.VolumeClass, out *v1alpha1.VolumeClass, s conversion.Scope) error {
	return autoConvert_storage_VolumeClass_To_v1alpha1_VolumeClass(in, out, s)
}

func autoConvert_v1alpha1_VolumeClassList_To_storage_VolumeClassList(in *v1alpha1.VolumeClassList, out *storage.VolumeClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.VolumeClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VolumeClassList_To_storage_VolumeClassList is an autogenerated conversion function.
func Convert_v1alpha1_VolumeClassList_To_storage_VolumeClassList(in *v1alpha1.VolumeClassList, out *storage.VolumeClassList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeClassList_To_storage_VolumeClassList(in, out, s)
}

func autoConvert_storage_VolumeClassList_To_v1alpha1_VolumeClassList(in *storage.VolumeClassList, out *v1alpha1.VolumeClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.VolumeClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_VolumeClassList_To_v1alpha1_VolumeClassList is an autogenerated conversion function.
func Convert_storage_VolumeClassList_To_v1alpha1_VolumeClassList(in *storage.VolumeClassList, out *v1alpha1.VolumeClassList, s conversion.Scope) error {
	return autoConvert_storage_VolumeClassList_To_v1alpha1_VolumeClassList(in, out, s)
}

func autoConvert_v1alpha1_VolumeCondition_To_storage_VolumeCondition(in *v1alpha1.VolumeCondition, out *storage.VolumeCondition, s conversion.Scope) error {
	out.Type = storage.VolumeConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.ObservedGeneration = in.ObservedGeneration
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert_v1alpha1_VolumeCondition_To_storage_VolumeCondition is an autogenerated conversion function.
func Convert_v1alpha1_VolumeCondition_To_storage_VolumeCondition(in *v1alpha1.VolumeCondition, out *storage.VolumeCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeCondition_To_storage_VolumeCondition(in, out, s)
}

func autoConvert_storage_VolumeCondition_To_v1alpha1_VolumeCondition(in *storage.VolumeCondition, out *v1alpha1.VolumeCondition, s conversion.Scope) error {
	out.Type = v1alpha1.VolumeConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.ObservedGeneration = in.ObservedGeneration
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert_storage_VolumeCondition_To_v1alpha1_VolumeCondition is an autogenerated conversion function.
func Convert_storage_VolumeCondition_To_v1alpha1_VolumeCondition(in *storage.VolumeCondition, out *v1alpha1.VolumeCondition, s conversion.Scope) error {
	return autoConvert_storage_VolumeCondition_To_v1alpha1_VolumeCondition(in, out, s)
}

func autoConvert_v1alpha1_VolumeList_To_storage_VolumeList(in *v1alpha1.VolumeList, out *storage.VolumeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.Volume)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VolumeList_To_storage_VolumeList is an autogenerated conversion function.
func Convert_v1alpha1_VolumeList_To_storage_VolumeList(in *v1alpha1.VolumeList, out *storage.VolumeList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeList_To_storage_VolumeList(in, out, s)
}

func autoConvert_storage_VolumeList_To_v1alpha1_VolumeList(in *storage.VolumeList, out *v1alpha1.VolumeList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.Volume)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_VolumeList_To_v1alpha1_VolumeList is an autogenerated conversion function.
func Convert_storage_VolumeList_To_v1alpha1_VolumeList(in *storage.VolumeList, out *v1alpha1.VolumeList, s conversion.Scope) error {
	return autoConvert_storage_VolumeList_To_v1alpha1_VolumeList(in, out, s)
}

func autoConvert_v1alpha1_VolumePool_To_storage_VolumePool(in *v1alpha1.VolumePool, out *storage.VolumePool, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VolumePool_To_storage_VolumePool is an autogenerated conversion function.
func Convert_v1alpha1_VolumePool_To_storage_VolumePool(in *v1alpha1.VolumePool, out *storage.VolumePool, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumePool_To_storage_VolumePool(in, out, s)
}

func autoConvert_storage_VolumePool_To_v1alpha1_VolumePool(in *storage.VolumePool, out *v1alpha1.VolumePool, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_VolumePool_To_v1alpha1_VolumePool is an autogenerated conversion function.
func Convert_storage_VolumePool_To_v1alpha1_VolumePool(in *storage.VolumePool, out *v1alpha1.VolumePool, s conversion.Scope) error {
	return autoConvert_storage_VolumePool_To_v1alpha1_VolumePool(in, out, s)
}

func autoConvert_v1alpha1_VolumePoolCondition_To_storage_VolumePoolCondition(in *v1alpha1.VolumePoolCondition, out *storage.VolumePoolCondition, s conversion.Scope) error {
	out.Type = storage.VolumePoolConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.ObservedGeneration = in.ObservedGeneration
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert_v1alpha1_VolumePoolCondition_To_storage_VolumePoolCondition is an autogenerated conversion function.
func Convert_v1alpha1_VolumePoolCondition_To_storage_VolumePoolCondition(in *v1alpha1.VolumePoolCondition, out *storage.VolumePoolCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumePoolCondition_To_storage_VolumePoolCondition(in, out, s)
}

func autoConvert_storage_VolumePoolCondition_To_v1alpha1_VolumePoolCondition(in *storage.VolumePoolCondition, out *v1alpha1.VolumePoolCondition, s conversion.Scope) error {
	out.Type = v1alpha1.VolumePoolConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.ObservedGeneration = in.ObservedGeneration
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert_storage_VolumePoolCondition_To_v1alpha1_VolumePoolCondition is an autogenerated conversion function.
func Convert_storage_VolumePoolCondition_To_v1alpha1_VolumePoolCondition(in *storage.VolumePoolCondition, out *v1alpha1.VolumePoolCondition, s conversion.Scope) error {
	return autoConvert_storage_VolumePoolCondition_To_v1alpha1_VolumePoolCondition(in, out, s)
}

func autoConvert_v1alpha1_VolumePoolList_To_storage_VolumePoolList(in *v1alpha1.VolumePoolList, out *storage.VolumePoolList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]storage.VolumePool)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VolumePoolList_To_storage_VolumePoolList is an autogenerated conversion function.
func Convert_v1alpha1_VolumePoolList_To_storage_VolumePoolList(in *v1alpha1.VolumePoolList, out *storage.VolumePoolList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumePoolList_To_storage_VolumePoolList(in, out, s)
}

func autoConvert_storage_VolumePoolList_To_v1alpha1_VolumePoolList(in *storage.VolumePoolList, out *v1alpha1.VolumePoolList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.VolumePool)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_storage_VolumePoolList_To_v1alpha1_VolumePoolList is an autogenerated conversion function.
func Convert_storage_VolumePoolList_To_v1alpha1_VolumePoolList(in *storage.VolumePoolList, out *v1alpha1.VolumePoolList, s conversion.Scope) error {
	return autoConvert_storage_VolumePoolList_To_v1alpha1_VolumePoolList(in, out, s)
}

func autoConvert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec(in *v1alpha1.VolumePoolSpec, out *storage.VolumePoolSpec, s conversion.Scope) error {
	out.ProviderID = in.ProviderID
	out.Taints = *(*[]commonv1alpha1.Taint)(unsafe.Pointer(&in.Taints))
	return nil
}

// Convert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec is an autogenerated conversion function.
func Convert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec(in *v1alpha1.VolumePoolSpec, out *storage.VolumePoolSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumePoolSpec_To_storage_VolumePoolSpec(in, out, s)
}

func autoConvert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec(in *storage.VolumePoolSpec, out *v1alpha1.VolumePoolSpec, s conversion.Scope) error {
	out.ProviderID = in.ProviderID
	out.Taints = *(*[]commonv1alpha1.Taint)(unsafe.Pointer(&in.Taints))
	return nil
}

// Convert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec is an autogenerated conversion function.
func Convert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec(in *storage.VolumePoolSpec, out *v1alpha1.VolumePoolSpec, s conversion.Scope) error {
	return autoConvert_storage_VolumePoolSpec_To_v1alpha1_VolumePoolSpec(in, out, s)
}

func autoConvert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus(in *v1alpha1.VolumePoolStatus, out *storage.VolumePoolStatus, s conversion.Scope) error {
	out.State = storage.VolumePoolState(in.State)
	out.Conditions = *(*[]storage.VolumePoolCondition)(unsafe.Pointer(&in.Conditions))
	out.AvailableVolumeClasses = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.AvailableVolumeClasses))
	out.Available = *(*v1.ResourceList)(unsafe.Pointer(&in.Available))
	out.Used = *(*v1.ResourceList)(unsafe.Pointer(&in.Used))
	return nil
}

// Convert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus is an autogenerated conversion function.
func Convert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus(in *v1alpha1.VolumePoolStatus, out *storage.VolumePoolStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumePoolStatus_To_storage_VolumePoolStatus(in, out, s)
}

func autoConvert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus(in *storage.VolumePoolStatus, out *v1alpha1.VolumePoolStatus, s conversion.Scope) error {
	out.State = v1alpha1.VolumePoolState(in.State)
	out.Conditions = *(*[]v1alpha1.VolumePoolCondition)(unsafe.Pointer(&in.Conditions))
	out.AvailableVolumeClasses = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.AvailableVolumeClasses))
	out.Available = *(*v1.ResourceList)(unsafe.Pointer(&in.Available))
	out.Used = *(*v1.ResourceList)(unsafe.Pointer(&in.Used))
	return nil
}

// Convert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus is an autogenerated conversion function.
func Convert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus(in *storage.VolumePoolStatus, out *v1alpha1.VolumePoolStatus, s conversion.Scope) error {
	return autoConvert_storage_VolumePoolStatus_To_v1alpha1_VolumePoolStatus(in, out, s)
}

func autoConvert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(in *v1alpha1.VolumeSpec, out *storage.VolumeSpec, s conversion.Scope) error {
	out.VolumeClassRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.VolumeClassRef))
	out.VolumePoolSelector = *(*map[string]string)(unsafe.Pointer(&in.VolumePoolSelector))
	out.VolumePoolRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.VolumePoolRef))
	out.ClaimRef = (*commonv1alpha1.LocalUIDReference)(unsafe.Pointer(in.ClaimRef))
	out.Resources = *(*v1.ResourceList)(unsafe.Pointer(&in.Resources))
	out.Image = in.Image
	out.ImagePullSecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.ImagePullSecretRef))
	out.Unclaimable = in.Unclaimable
	out.Tolerations = *(*[]commonv1alpha1.Toleration)(unsafe.Pointer(&in.Tolerations))
	return nil
}

// Convert_v1alpha1_VolumeSpec_To_storage_VolumeSpec is an autogenerated conversion function.
func Convert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(in *v1alpha1.VolumeSpec, out *storage.VolumeSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(in, out, s)
}

func autoConvert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(in *storage.VolumeSpec, out *v1alpha1.VolumeSpec, s conversion.Scope) error {
	out.VolumeClassRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.VolumeClassRef))
	out.VolumePoolSelector = *(*map[string]string)(unsafe.Pointer(&in.VolumePoolSelector))
	out.VolumePoolRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.VolumePoolRef))
	out.ClaimRef = (*commonv1alpha1.LocalUIDReference)(unsafe.Pointer(in.ClaimRef))
	out.Resources = *(*v1.ResourceList)(unsafe.Pointer(&in.Resources))
	out.Image = in.Image
	out.ImagePullSecretRef = (*v1.LocalObjectReference)(unsafe.Pointer(in.ImagePullSecretRef))
	out.Unclaimable = in.Unclaimable
	out.Tolerations = *(*[]commonv1alpha1.Toleration)(unsafe.Pointer(&in.Tolerations))
	return nil
}

// Convert_storage_VolumeSpec_To_v1alpha1_VolumeSpec is an autogenerated conversion function.
func Convert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(in *storage.VolumeSpec, out *v1alpha1.VolumeSpec, s conversion.Scope) error {
	return autoConvert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(in, out, s)
}

func autoConvert_v1alpha1_VolumeStatus_To_storage_VolumeStatus(in *v1alpha1.VolumeStatus, out *storage.VolumeStatus, s conversion.Scope) error {
	out.State = storage.VolumeState(in.State)
	out.LastStateTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastStateTransitionTime))
	out.Phase = storage.VolumePhase(in.Phase)
	out.LastPhaseTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastPhaseTransitionTime))
	out.Access = (*storage.VolumeAccess)(unsafe.Pointer(in.Access))
	out.Conditions = *(*[]storage.VolumeCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_VolumeStatus_To_storage_VolumeStatus is an autogenerated conversion function.
func Convert_v1alpha1_VolumeStatus_To_storage_VolumeStatus(in *v1alpha1.VolumeStatus, out *storage.VolumeStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeStatus_To_storage_VolumeStatus(in, out, s)
}

func autoConvert_storage_VolumeStatus_To_v1alpha1_VolumeStatus(in *storage.VolumeStatus, out *v1alpha1.VolumeStatus, s conversion.Scope) error {
	out.State = v1alpha1.VolumeState(in.State)
	out.LastStateTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastStateTransitionTime))
	out.Phase = v1alpha1.VolumePhase(in.Phase)
	out.LastPhaseTransitionTime = (*metav1.Time)(unsafe.Pointer(in.LastPhaseTransitionTime))
	out.Access = (*v1alpha1.VolumeAccess)(unsafe.Pointer(in.Access))
	out.Conditions = *(*[]v1alpha1.VolumeCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_storage_VolumeStatus_To_v1alpha1_VolumeStatus is an autogenerated conversion function.
func Convert_storage_VolumeStatus_To_v1alpha1_VolumeStatus(in *storage.VolumeStatus, out *v1alpha1.VolumeStatus, s conversion.Scope) error {
	return autoConvert_storage_VolumeStatus_To_v1alpha1_VolumeStatus(in, out, s)
}

func autoConvert_v1alpha1_VolumeTemplateSpec_To_storage_VolumeTemplateSpec(in *v1alpha1.VolumeTemplateSpec, out *storage.VolumeTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VolumeSpec_To_storage_VolumeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VolumeTemplateSpec_To_storage_VolumeTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha1_VolumeTemplateSpec_To_storage_VolumeTemplateSpec(in *v1alpha1.VolumeTemplateSpec, out *storage.VolumeTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeTemplateSpec_To_storage_VolumeTemplateSpec(in, out, s)
}

func autoConvert_storage_VolumeTemplateSpec_To_v1alpha1_VolumeTemplateSpec(in *storage.VolumeTemplateSpec, out *v1alpha1.VolumeTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_storage_VolumeSpec_To_v1alpha1_VolumeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_storage_VolumeTemplateSpec_To_v1alpha1_VolumeTemplateSpec is an autogenerated conversion function.
func Convert_storage_VolumeTemplateSpec_To_v1alpha1_VolumeTemplateSpec(in *storage.VolumeTemplateSpec, out *v1alpha1.VolumeTemplateSpec, s conversion.Scope) error {
	return autoConvert_storage_VolumeTemplateSpec_To_v1alpha1_VolumeTemplateSpec(in, out, s)
}

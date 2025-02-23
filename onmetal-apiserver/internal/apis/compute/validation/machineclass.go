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

package validation

import (
	onmetalapivalidation "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/api/validation"
	"github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis/compute"
	corev1 "k8s.io/api/core/v1"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateMachineClass validates a MachineClass object.
func ValidateMachineClass(machineClass *compute.MachineClass) field.ErrorList {
	var allErrs field.ErrorList

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaAccessor(machineClass, false, apivalidation.NameIsDNSLabel, field.NewPath("metadata"))...)

	allErrs = append(allErrs, validateMachineClassCapabilities(machineClass.Capabilities, field.NewPath("capabilities"))...)

	return allErrs
}

func validateMachineClassCapabilities(capabilities corev1.ResourceList, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	cpu := capabilities.Cpu()
	allErrs = append(allErrs, onmetalapivalidation.ValidatePositiveQuantity(*cpu, fldPath.Key(string(corev1.ResourceCPU)))...)

	memory := capabilities.Memory()
	allErrs = append(allErrs, onmetalapivalidation.ValidatePositiveQuantity(*memory, fldPath.Key(string(corev1.ResourceMemory)))...)

	return allErrs
}

// ValidateMachineClassUpdate validates a MachineClass object before an update.
func ValidateMachineClassUpdate(newMachineClass, oldMachineClass *compute.MachineClass) field.ErrorList {
	var allErrs field.ErrorList

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaAccessorUpdate(newMachineClass, oldMachineClass, field.NewPath("metadata"))...)
	allErrs = append(allErrs, onmetalapivalidation.ValidateImmutableField(newMachineClass.Capabilities, oldMachineClass.Capabilities, field.NewPath("capabilities"))...)
	allErrs = append(allErrs, ValidateMachineClass(newMachineClass)...)

	return allErrs
}

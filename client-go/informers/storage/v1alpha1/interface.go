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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/onmetal/onmetal-api/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Volumes returns a VolumeInformer.
	Volumes() VolumeInformer
	// VolumeClasses returns a VolumeClassInformer.
	VolumeClasses() VolumeClassInformer
	// VolumePools returns a VolumePoolInformer.
	VolumePools() VolumePoolInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Volumes returns a VolumeInformer.
func (v *version) Volumes() VolumeInformer {
	return &volumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeClasses returns a VolumeClassInformer.
func (v *version) VolumeClasses() VolumeClassInformer {
	return &volumeClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VolumePools returns a VolumePoolInformer.
func (v *version) VolumePools() VolumePoolInformer {
	return &volumePoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

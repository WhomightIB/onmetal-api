// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis/networking"
	"github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis/storage"
)

// EphemeralNetworkInterfaceSource is a definition for an ephemeral (i.e. coupled to the lifetime of the surrounding
// object) networking.NetworkInterface.
type EphemeralNetworkInterfaceSource struct {
	// NetworkInterfaceTemplate is the template definition of the networking.NetworkInterface.
	NetworkInterfaceTemplate *networking.NetworkInterfaceTemplateSpec
}

// EphemeralVolumeSource is a definition for an ephemeral (i.e. coupled to the lifetime of the surrounding object)
// storage.Volume.
type EphemeralVolumeSource struct {
	// VolumeTemplate is the template definition of the storage.Volume.
	VolumeTemplate *storage.VolumeTemplateSpec
}

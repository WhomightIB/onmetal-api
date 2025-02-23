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

package mcm

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-logr/logr"
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1alpha1"
	"golang.org/x/exp/maps"
	"k8s.io/apimachinery/pkg/util/wait"
	ctrl "sigs.k8s.io/controller-runtime"
)

type capabilities struct {
	cpuMillis   int64
	memoryBytes uint64
}

func getCapabilities(oriCaps *ori.MachineClassCapabilities) capabilities {
	return capabilities{
		cpuMillis:   oriCaps.CpuMillis,
		memoryBytes: oriCaps.MemoryBytes,
	}
}

type Generic struct {
	mu sync.RWMutex

	sync   bool
	synced chan struct{}

	machineClassByName         map[string]*ori.MachineClass
	machineClassByCapabilities map[capabilities][]*ori.MachineClass

	machineRuntime ori.MachineRuntimeClient

	relistPeriod time.Duration
}

func (g *Generic) relist(ctx context.Context, log logr.Logger) error {
	log.V(1).Info("Relisting machine classes")
	res, err := g.machineRuntime.ListMachineClasses(ctx, &ori.ListMachineClassesRequest{})
	if err != nil {
		return fmt.Errorf("error listing machine classes: %w", err)
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	maps.Clear(g.machineClassByName)
	maps.Clear(g.machineClassByCapabilities)

	for _, machineClass := range res.MachineClasses {
		caps := capabilities{
			cpuMillis:   machineClass.Capabilities.CpuMillis,
			memoryBytes: machineClass.Capabilities.MemoryBytes,
		}
		g.machineClassByName[machineClass.Name] = machineClass
		g.machineClassByCapabilities[caps] = append(g.machineClassByCapabilities[caps], machineClass)
	}

	if !g.sync {
		g.sync = true
		close(g.synced)
	}

	return nil
}

func (g *Generic) Start(ctx context.Context) error {
	log := ctrl.LoggerFrom(ctx).WithName("mcm")
	wait.UntilWithContext(ctx, func(ctx context.Context) {
		if err := g.relist(ctx, log); err != nil {
			log.Error(err, "Error relisting")
		}
	}, g.relistPeriod)
	return nil
}

func (g *Generic) GetMachineClassFor(ctx context.Context, name string, caps *ori.MachineClassCapabilities) (*ori.MachineClass, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	expected := getCapabilities(caps)
	if byName, ok := g.machineClassByName[name]; ok && getCapabilities(byName.Capabilities) == expected {
		return byName, nil
	}

	if byCaps, ok := g.machineClassByCapabilities[expected]; ok {
		switch len(byCaps) {
		case 0:
			return nil, ErrNoMatchingMachineClass
		case 1:
			class := *byCaps[0]
			return &class, nil
		default:
			return nil, ErrAmbiguousMatchingMachineClass
		}
	}

	return nil, ErrNoMatchingMachineClass
}

func (g *Generic) WaitForSync(ctx context.Context) error {
	select {
	case <-g.synced:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

type GenericOptions struct {
	RelistPeriod time.Duration
}

func setGenericOptionsDefaults(o *GenericOptions) {
	if o.RelistPeriod == 0 {
		o.RelistPeriod = 1 * time.Hour
	}
}

func NewGeneric(runtime ori.MachineRuntimeClient, opts GenericOptions) MachineClassMapper {
	setGenericOptionsDefaults(&opts)
	return &Generic{
		synced:                     make(chan struct{}),
		machineClassByName:         map[string]*ori.MachineClass{},
		machineClassByCapabilities: map[capabilities][]*ori.MachineClass{},
		machineRuntime:             runtime,
		relistPeriod:               opts.RelistPeriod,
	}
}

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

package storage

import (
	"context"
	"sort"
	"strings"

	"github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis/compute"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/meta/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type convertor struct{}

var (
	objectMetaSwaggerDoc = metav1.ObjectMeta{}.SwaggerDoc()

	headers = []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: objectMetaSwaggerDoc["name"]},
		{Name: "MachineClasses", Type: "string", Description: "Machine classes available on this machine pool."},
		{Name: "Age", Type: "string", Format: "date", Description: objectMetaSwaggerDoc["creationTimestamp"]},
	}
)

func newTableConvertor() *convertor {
	return &convertor{}
}

func machinePoolMachineClassNames(machinePool *compute.MachinePool) []string {
	machineClassNames := make([]string, 0, len(machinePool.Status.AvailableMachineClasses))
	for _, machineClassRef := range machinePool.Status.AvailableMachineClasses {
		machineClassNames = append(machineClassNames, machineClassRef.Name)
	}
	sort.Strings(machineClassNames)
	return machineClassNames
}

func (c *convertor) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	tab := &metav1.Table{
		ColumnDefinitions: headers,
	}

	if m, err := meta.ListAccessor(obj); err == nil {
		tab.ResourceVersion = m.GetResourceVersion()
		tab.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(obj); err == nil {
			tab.ResourceVersion = m.GetResourceVersion()
		}
	}

	var err error
	tab.Rows, err = table.MetaToTableRow(obj, func(obj runtime.Object, m metav1.Object, name, age string) (cells []interface{}, err error) {
		machinePool := obj.(*compute.MachinePool)

		cells = append(cells, name)
		machineClassNames := machinePoolMachineClassNames(machinePool)
		if len(machineClassNames) == 0 {
			cells = append(cells, "<none>")
		} else {
			cells = append(cells, strings.Join(machineClassNames, ","))
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}

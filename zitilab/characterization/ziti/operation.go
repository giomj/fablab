/*
	Copyright 2020 NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package zitilab_characterization_ziti

import (
	"github.com/netfoundry/fablab/kernel/model"
	operation "github.com/netfoundry/fablab/kernel/runlevel/5_operation"
)

func newOperationFactory() model.Factory {
	return &operationFactory{}
}

func (f *operationFactory) Build(m *model.Model) error {
	c := make(chan struct{})
	m.Operation = model.OperatingBinders{
		func(m *model.Model) model.OperatingStage { return operation.Mesh(c) },
		func(m *model.Model) model.OperatingStage { return operation.Metrics(c) },
		/*
		func(m *model.Model) model.OperatingStage {
			minutes, found := m.GetVariable("sample_minutes")
			if !found {
				minutes = 1
			}
			sampleDuration := time.Duration(minutes.(int)) * time.Minute
			return operation.Iperf(int(sampleDuration.Seconds()))
		},
		*/
		func(m *model.Model) model.OperatingStage { return operation.Closer(c) },
		func(m *model.Model) model.OperatingStage { return operation.Persist() },
	}
	return nil
}

type operationFactory struct{}
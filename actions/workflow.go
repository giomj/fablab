/*
	Copyright 2019 Netfoundry, Inc.

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

package actions

import (
	"bitbucket.org/netfoundry/fablab/kernel"
	"fmt"
)

func Workflow(actions ...kernel.Action) *workflow {
	return &workflow{
		actions: actions,
	}
}

func (workflow *workflow) AddAction(action kernel.Action) {
	workflow.actions = append(workflow.actions, action)
}

func (workflow *workflow) Execute(m *kernel.Model) error {
	for _, action := range workflow.actions {
		if err := action.Execute(m); err != nil {
			return fmt.Errorf("error executing action (%s)", err)
		}
	}
	return nil
}

type workflow struct {
	actions []kernel.Action
}

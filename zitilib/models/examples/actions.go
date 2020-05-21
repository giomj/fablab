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

package zitilib_examples

import (
	"github.com/openziti/fablab/kernel/model"
	"github.com/openziti/fablab/zitilib/actions"
	"github.com/openziti/fablab/zitilib/console"
)
import "github.com/openziti/fablab/zitilib/models/examples/actions"

func newActionsFactory() model.Factory {
	return &actionsFactory{}
}

func (_ *actionsFactory) Build(m *model.Model) error {
	m.Actions = model.ActionBinders{
		"bootstrap": zitilib_examples_actions.NewBootstrapAction(),
		"start":     zitilib_examples_actions.NewStartAction(),
		"stop":      zitilib_examples_actions.NewStopAction(),
		"console":   func(m *model.Model) model.Action { return console.Console() },
		"logs":      func(m *model.Model) model.Action { return zitilib_actions.Logs() },
	}
	return nil
}

type actionsFactory struct{}

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

package subcmd

import (
	"bitbucket.org/netfoundry/fablab/kernel"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(disposeCmd)
}

var disposeCmd = &cobra.Command{
	Use:   "dispose",
	Short: "dispose of all model resources",
	Args:  cobra.ExactArgs(0),
	Run:   dispose,
}

func dispose(_ *cobra.Command, _ []string) {
	if err := kernel.Bootstrap(); err != nil {
		logrus.Fatalf("unable to bootstrap (%s)", err)
	}

	l := kernel.GetLabel()
	if l == nil {
		logrus.Fatalf("no label for instance [%s]", kernel.ActiveInstancePath())
	}

	if l != nil {
		m, found := kernel.GetModel(l.Model)
		if !found {
			logrus.Fatalf("no such model [%s]", l.Model)
		}

		if err := m.Dispose(l); err != nil {
			logrus.Fatalf("error building configuration (%w)", err)
		}

	} else {
		logrus.Fatalf("no label for run")
	}
}

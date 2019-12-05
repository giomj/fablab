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

package host

import (
	"bitbucket.org/netfoundry/fablab/kernel"
	"bitbucket.org/netfoundry/fablab/kernel/lib"
	"fmt"
	"github.com/sirupsen/logrus"
)

func GroupExec(regionSpec, hostSpec, cmd string) kernel.Action {
	return &groupExec{
		regionSpec: regionSpec,
		hostSpec:   hostSpec,
		cmd:        cmd,
	}
}

func (groupExec *groupExec) Execute(m *kernel.Model) error {
	hosts := m.GetHosts(groupExec.regionSpec, groupExec.hostSpec)
	for _, h := range hosts {
		sshUsername := m.MustVariable("credentials", "ssh", "username").(string)
		if o, err := lib.RemoteExec(sshUsername, h.PublicIp, groupExec.cmd); err != nil {
			logrus.Errorf("output [%s]", o)
			return fmt.Errorf("error executing process [%s] on [%s] (%s)", groupExec.cmd, h.PublicIp, err)
		}
	}
	return nil
}

type groupExec struct {
	regionSpec string
	hostSpec   string
	cmd        string
}

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

package lib

import (
	"bitbucket.org/netfoundry/fablab/kernel"
	"text/template"
)

func TemplateFuncMap(m *kernel.Model) template.FuncMap {
	return template.FuncMap{
		"publicIp": func(regionTag, hostTag string) string {
			host := m.GetHostByTags(regionTag, hostTag)
			if host != nil {
				return host.PublicIp
			}
			return ""
		},
	}
}

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

package operation

import (
	"github.com/netfoundry/fablab/kernel"
	"github.com/sirupsen/logrus"
	"time"
)

func Timer(duration time.Duration, closer chan struct{}) kernel.OperatingStage {
	return &timer{duration: duration, close: closer}
}

func (timer *timer) Operate(m *kernel.Model) error {
	logrus.Infof("waiting for %s", timer.duration)
	time.Sleep(timer.duration)
	logrus.Infof("closing")
	close(timer.close)
	return nil
}

type timer struct {
	duration time.Duration
	close chan struct{}
}
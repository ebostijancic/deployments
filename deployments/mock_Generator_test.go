// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package deployments

import "github.com/stretchr/testify/mock"

// MockGenerator is an autogenerated mock type for the Generator type
type MockGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields: deviceID, deployment
func (_m *MockGenerator) Generate(deviceID string, deployment *Deployment) (*DeviceDeployment, error) {
	ret := _m.Called(deviceID, deployment)

	var r0 *DeviceDeployment
	if rf, ok := ret.Get(0).(func(string, *Deployment) *DeviceDeployment); ok {
		r0 = rf(deviceID, deployment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeviceDeployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *Deployment) error); ok {
		r1 = rf(deviceID, deployment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

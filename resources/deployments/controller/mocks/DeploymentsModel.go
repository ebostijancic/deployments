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

package mocks

import (
	"context"
	"github.com/mendersoftware/deployments/resources/deployments"
	"github.com/stretchr/testify/mock"
)

// DeploymentsModel is an autogenerated mock type for the DeploymentsModel type
type DeploymentsModel struct {
	mock.Mock
}

// CreateDeployment provides a mock function with given fields: constructor
func (_m *DeploymentsModel) CreateDeployment(ctx context.Context, constructor *deployments.DeploymentConstructor) (string, error) {
	ret := _m.Called(constructor)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *deployments.DeploymentConstructor) string); ok {
		r0 = rf(ctx, constructor)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *deployments.DeploymentConstructor) error); ok {
		r1 = rf(ctx, constructor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployment provides a mock function with given fields: deploymentID
func (_m *DeploymentsModel) GetDeployment(deploymentID string) (*deployments.Deployment, error) {
	ret := _m.Called(deploymentID)

	var r0 *deployments.Deployment
	if rf, ok := ret.Get(0).(func(string) *deployments.Deployment); ok {
		r0 = rf(deploymentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deployments.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deploymentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeploymentForDevice provides a mock function with given fields: deviceID
func (_m *DeploymentsModel) GetDeploymentForDevice(deviceID string) (*deployments.DeploymentInstructions, error) {
	ret := _m.Called(deviceID)

	var r0 *deployments.DeploymentInstructions
	if rf, ok := ret.Get(0).(func(string) *deployments.DeploymentInstructions); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*deployments.DeploymentInstructions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DeploymentsModel) UpdateDeviceDeploymentStatus(deploymentID string,
	deviceID string, status string) error {

	ret := _m.Called(deploymentID, deviceID, status)
	return ret.Error(0)
}

func (_m *DeploymentsModel) GetDeploymentStats(deploymentID string) (deployments.Stats, error) {
	ret := _m.Called(deploymentID)
	return ret.Get(0).(deployments.Stats), ret.Error(1)
}

func (_m *DeploymentsModel) GetDeviceStatusesForDeployment(deploymentID string) ([]deployments.DeviceDeployment, error) {

	ret := _m.Called(deploymentID)
	return ret.Get(0).([]deployments.DeviceDeployment), ret.Error(1)
}

func (_m *DeploymentsModel) LookupDeployment(query deployments.Query) ([]*deployments.Deployment, error) {

	ret := _m.Called(query)
	return ret.Get(0).([]*deployments.Deployment), ret.Error(1)
}

func (_m *DeploymentsModel) SaveDeviceDeploymentLog(deviceID string,
	deploymentID string, logs []deployments.LogMessage) error {

	ret := _m.Called(deviceID, deploymentID, logs)
	return ret.Error(0)
}

func (_m *DeploymentsModel) HasDeploymentForDevice(deploymentID string, deviceID string) (bool, error) {
	ret := _m.Called(deploymentID, deviceID)
	return ret.Bool(0), ret.Error(1)
}

func (_m *DeploymentsModel) GetDeviceDeploymentLog(deviceID, deploymentID string) (*deployments.DeploymentLog, error) {

	ret := _m.Called(deviceID, deploymentID)
	return ret.Get(0).(*deployments.DeploymentLog), ret.Error(1)
}

func (_m *DeploymentsModel) AbortDeployment(deploymentID string) error {
	ret := _m.Called(deploymentID)
	return ret.Error(0)
}

func (_m *DeploymentsModel) IsDeploymentFinished(deploymentID string) (bool, error) {
	ret := _m.Called(deploymentID)
	return ret.Bool(0), ret.Error(1)
}

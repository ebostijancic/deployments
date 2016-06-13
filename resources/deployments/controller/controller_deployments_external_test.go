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

package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/mendersoftware/deployments/resources/deployments"
	. "github.com/mendersoftware/deployments/resources/deployments/controller"
	"github.com/mendersoftware/deployments/resources/deployments/controller/mocks"
	"github.com/mendersoftware/deployments/resources/deployments/view"
	. "github.com/mendersoftware/deployments/utils/pointers"
	"github.com/stretchr/testify/assert"
)

// Notice: 	Controller tests are not pure unit tests,
// 			they are more of integration test beween controller and view
//			testing actuall HTTP endpoint input/reponse

func ErrorToErrStruct(err error) interface{} {
	return struct{ Error string }{err.Error()}
}

type JSONResponseParams struct {
	OutputStatus     int
	OutputBodyObject interface{}
	OutputHeaders    map[string]string
}

func CheckRecordedResponse(t *testing.T, recorded *test.Recorded, params JSONResponseParams) {

	recorded.CodeIs(params.OutputStatus)
	recorded.ContentTypeIsJson()

	if params.OutputBodyObject != nil {
		assert.NotEmpty(t, recorded.Recorder.Body.String())

		expectedJSON, err := json.Marshal(params.OutputBodyObject)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), recorded.Recorder.Body.String())
	} else {
		assert.Empty(t, recorded.Recorder.Body.String())
	}

	for name, value := range params.OutputHeaders {
		assert.Equal(t, value, recorded.Recorder.HeaderMap.Get(name))
	}
}

func TestControllerGetDeploymentForDevice(t *testing.T) {

	t.Parallel()

	testCases := []struct {
		JSONResponseParams

		InputID string

		InputModelDeploymentInstructions *deployments.DeploymentInstructions
		InputModelError                  error
	}{
		{
			InputID: "broken_id",
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusBadRequest,
				OutputBodyObject: ErrorToErrStruct(ErrIDNotUUIDv4),
			},
		},
		{
			InputID:         "f826484e-1157-4109-af21-304e6d711560",
			InputModelError: errors.New("model error"),
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusInternalServerError,
				OutputBodyObject: ErrorToErrStruct(errors.New("model error")),
			},
		},
		{
			InputID: "f826484e-1157-4109-af21-304e6d711560",
			JSONResponseParams: JSONResponseParams{
				OutputStatus: http.StatusNoContent,
			},
		},
		{
			InputID: "f826484e-1157-4109-af21-304e6d711560",
			InputModelDeploymentInstructions: deployments.NewDeploymentInstructions("", nil, nil),
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusOK,
				OutputBodyObject: deployments.NewDeploymentInstructions("", nil, nil),
			},
		},
	}

	for _, testCase := range testCases {

		deploymentModel := new(mocks.DeploymentsModel)
		deploymentModel.On("GetDeploymentForDevice", testCase.InputID).
			Return(testCase.InputModelDeploymentInstructions, testCase.InputModelError)

		router, err := rest.MakeRouter(
			rest.Get("/r/:id",
				NewDeploymentsController(deploymentModel, new(view.DeploymentsView)).GetDeploymentForDevice))
		assert.NoError(t, err)

		api := rest.NewApi()
		api.SetApp(router)

		recorded := test.RunRequest(t, api.MakeHandler(),
			test.MakeSimpleRequest("GET", "http://localhost/r/"+testCase.InputID, nil))

		CheckRecordedResponse(t, recorded, testCase.JSONResponseParams)
	}
}

func TestControllerGetDeployment(t *testing.T) {

	t.Parallel()

	testCases := []struct {
		JSONResponseParams

		InputID string

		InputModelDeployment *deployments.Deployment
		InputModelError      error
	}{
		{
			InputID: "broken_id",
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusBadRequest,
				OutputBodyObject: ErrorToErrStruct(ErrIDNotUUIDv4),
			},
		},
		{
			InputID:         "f826484e-1157-4109-af21-304e6d711560",
			InputModelError: errors.New("model error"),
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusInternalServerError,
				OutputBodyObject: ErrorToErrStruct(errors.New("model error")),
			},
		},
		{
			InputID: "f826484e-1157-4109-af21-304e6d711560",
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusNotFound,
				OutputBodyObject: ErrorToErrStruct(errors.New("Resource not found")),
			},
		},
		{
			InputID: "f826484e-1157-4109-af21-304e6d711560",

			InputModelDeployment: &deployments.Deployment{Id: StringToPointer("id 123")},

			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusOK,
				OutputBodyObject: &deployments.Deployment{Id: StringToPointer("id 123")},
			},
		},
	}

	for _, testCase := range testCases {

		deploymentModel := new(mocks.DeploymentsModel)
		deploymentModel.On("GetDeployment", testCase.InputID).
			Return(testCase.InputModelDeployment, testCase.InputModelError)

		router, err := rest.MakeRouter(
			rest.Get("/r/:id",
				NewDeploymentsController(deploymentModel, new(view.DeploymentsView)).GetDeployment))
		assert.NoError(t, err)

		api := rest.NewApi()
		api.SetApp(router)

		recorded := test.RunRequest(t, api.MakeHandler(),
			test.MakeSimpleRequest("GET", "http://localhost/r/"+testCase.InputID, nil))

		CheckRecordedResponse(t, recorded, testCase.JSONResponseParams)
	}
}

func TestControllerPostDeployment(t *testing.T) {

	t.Parallel()

	testCases := []struct {
		JSONResponseParams

		InputBodyObject interface{}

		InputModelID    string
		InputModelError error
	}{
		{
			InputBodyObject: nil,
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusBadRequest,
				OutputBodyObject: ErrorToErrStruct(errors.New("Validating request body: JSON payload is empty")),
			},
		},
		{
			InputBodyObject: deployments.NewDeploymentConstructor(),
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusBadRequest,
				OutputBodyObject: ErrorToErrStruct(errors.New(`Validating request body: Name: non zero value required;ArtifactName: non zero value required;Devices: non zero value required;`)),
			},
		},
		{
			InputBodyObject: &deployments.DeploymentConstructor{
				Name:         StringToPointer("NYC Production"),
				ArtifactName: StringToPointer("App 123"),
				Devices:      []string{"f826484e-1157-4109-af21-304e6d711560"},
			},
			InputModelError: errors.New("model error"),
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusInternalServerError,
				OutputBodyObject: ErrorToErrStruct(errors.New("model error")),
			},
		},
		{
			InputBodyObject: &deployments.DeploymentConstructor{
				Name:         StringToPointer("NYC Production"),
				ArtifactName: StringToPointer("App 123"),
				Devices:      []string{"f826484e-1157-4109-af21-304e6d711560"},
			},
			InputModelID: "1234",
			JSONResponseParams: JSONResponseParams{
				OutputStatus:     http.StatusCreated,
				OutputBodyObject: nil,
				OutputHeaders:    map[string]string{"Location": "http://localhost/r/1234"},
			},
		},
	}

	for _, testCase := range testCases {

		deploymentModel := new(mocks.DeploymentsModel)

		deploymentModel.On("CreateDeployment", testCase.InputBodyObject).
			Return(testCase.InputModelID, testCase.InputModelError)

		router, err := rest.MakeRouter(
			rest.Post("/r",
				NewDeploymentsController(deploymentModel, new(view.DeploymentsView)).PostDeployment))
		assert.NoError(t, err)

		api := rest.NewApi()
		api.SetApp(router)

		recorded := test.RunRequest(t, api.MakeHandler(),
			test.MakeSimpleRequest("POST", "http://localhost/r", testCase.InputBodyObject))

		CheckRecordedResponse(t, recorded, testCase.JSONResponseParams)
	}
}

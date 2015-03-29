// Package prediction provides access to the Prediction API.
//
// See https://developers.google.com/prediction/docs/developer-guide
//
// Usage example:
//
//   import "github.com/jfcote87/api/prediction/v1.2"
//   ...
//   predictionService, err := prediction.New(oauthHttpClient)
package prediction

import (
	"errors"
	"fmt"
	"github.com/jfcote87/api/googleapi"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "prediction:v1.2"
const apiName = "prediction"
const apiVersion = "v1.2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/prediction/v1.2/"}

// OAuth2 scopes used by this API.
const (
	// Manage your data and permissions in Google Cloud Storage
	DevstorageFull_controlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageRead_onlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// Manage your data in the Google Prediction API
	PredictionScope = "https://www.googleapis.com/auth/prediction"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Hostedmodels = NewHostedmodelsService(s)
	s.Training = NewTrainingService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Hostedmodels *HostedmodelsService

	Training *TrainingService
}

func NewHostedmodelsService(s *Service) *HostedmodelsService {
	rs := &HostedmodelsService{s: s}
	return rs
}

type HostedmodelsService struct {
	s *Service
}

func NewTrainingService(s *Service) *TrainingService {
	rs := &TrainingService{s: s}
	return rs
}

type TrainingService struct {
	s *Service
}

type Input struct {
	Input *InputInput `json:"input,omitempty"`
}

type InputInput struct {
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

type Output struct {
	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	OutputLabel string `json:"outputLabel,omitempty"`

	OutputMulti []*OutputOutputMulti `json:"outputMulti,omitempty"`

	OutputValue float64 `json:"outputValue,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type OutputOutputMulti struct {
	Label string `json:"label,omitempty"`

	Score float64 `json:"score,omitempty"`
}

type Training struct {
	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	ModelInfo *TrainingModelInfo `json:"modelInfo,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`

	TrainingStatus string `json:"trainingStatus,omitempty"`
}

type TrainingModelInfo struct {
	ClassificationAccuracy float64 `json:"classificationAccuracy,omitempty"`

	MeanSquaredError float64 `json:"meanSquaredError,omitempty"`

	ModelType string `json:"modelType,omitempty"`
}

type Update struct {
	// ClassLabel: The true class label of this instance
	ClassLabel string `json:"classLabel,omitempty"`

	// CsvInstance: The input features for this instance
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

// method id "prediction.predict":

type PredictCall struct {
	s             *Service
	data          string
	input         *Input
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Predict: Submit data and request a prediction

func (s *Service) Predict(data string, input *Input) *PredictCall {
	return &PredictCall{
		s:             s,
		data:          data,
		input:         input,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "training/{data}/predict",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PredictCall) Fields(s ...googleapi.Field) *PredictCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PredictCall) Do() (*Output, error) {
	var returnValue *Output
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"data": c.data,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.input,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Submit data and request a prediction",
	//   "httpMethod": "POST",
	//   "id": "prediction.predict",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket%2Fmydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.hostedmodels.predict":

type HostedmodelsPredictCall struct {
	s               *Service
	hostedModelName string
	input           *Input
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Predict: Submit input and request an output against a hosted model

func (r *HostedmodelsService) Predict(hostedModelName string, input *Input) *HostedmodelsPredictCall {
	return &HostedmodelsPredictCall{
		s:               r.s,
		hostedModelName: hostedModelName,
		input:           input,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "hostedmodels/{hostedModelName}/predict",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HostedmodelsPredictCall) Fields(s ...googleapi.Field) *HostedmodelsPredictCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *HostedmodelsPredictCall) Do() (*Output, error) {
	var returnValue *Output
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"hostedModelName": c.hostedModelName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.input,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Submit input and request an output against a hosted model",
	//   "httpMethod": "POST",
	//   "id": "prediction.hostedmodels.predict",
	//   "parameterOrder": [
	//     "hostedModelName"
	//   ],
	//   "parameters": {
	//     "hostedModelName": {
	//       "description": "The name of a hosted model",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "hostedmodels/{hostedModelName}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.delete":

type TrainingDeleteCall struct {
	s             *Service
	data          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete a trained model

func (r *TrainingService) Delete(data string) *TrainingDeleteCall {
	return &TrainingDeleteCall{
		s:             r.s,
		data:          data,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "training/{data}",
	}
}

func (c *TrainingDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"data": c.data,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a trained model",
	//   "httpMethod": "DELETE",
	//   "id": "prediction.training.delete",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.get":

type TrainingGetCall struct {
	s             *Service
	data          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Check training status of your model

func (r *TrainingService) Get(data string) *TrainingGetCall {
	return &TrainingGetCall{
		s:             r.s,
		data:          data,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "training/{data}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingGetCall) Fields(s ...googleapi.Field) *TrainingGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainingGetCall) Do() (*Training, error) {
	var returnValue *Training
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"data": c.data,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Check training status of your model",
	//   "httpMethod": "GET",
	//   "id": "prediction.training.get",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.insert":

type TrainingInsertCall struct {
	s             *Service
	training      *Training
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Begin training your model

func (r *TrainingService) Insert(training *Training) *TrainingInsertCall {
	return &TrainingInsertCall{
		s:             r.s,
		training:      training,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "training",
	}
}

// Data sets the optional parameter "data": mybucket/mydata resource in
// Google Storage
func (c *TrainingInsertCall) Data(data string) *TrainingInsertCall {
	c.params_.Set("data", fmt.Sprintf("%v", data))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingInsertCall) Fields(s ...googleapi.Field) *TrainingInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainingInsertCall) Do() (*Training, error) {
	var returnValue *Training
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.training,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Begin training your model",
	//   "httpMethod": "POST",
	//   "id": "prediction.training.insert",
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "training",
	//   "request": {
	//     "$ref": "Training"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.update":

type TrainingUpdateCall struct {
	s             *Service
	data          string
	update        *Update
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Add new data to a trained model

func (r *TrainingService) Update(data string, update *Update) *TrainingUpdateCall {
	return &TrainingUpdateCall{
		s:             r.s,
		data:          data,
		update:        update,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "training/{data}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingUpdateCall) Fields(s ...googleapi.Field) *TrainingUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainingUpdateCall) Do() (*Training, error) {
	var returnValue *Training
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"data": c.data,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.update,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add new data to a trained model",
	//   "httpMethod": "PUT",
	//   "id": "prediction.training.update",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "request": {
	//     "$ref": "Update"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

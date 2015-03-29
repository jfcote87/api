// Package prediction provides access to the Prediction API.
//
// See https://developers.google.com/prediction/docs/developer-guide
//
// Usage example:
//
//   import "github.com/jfcote87/api/prediction/v1.6"
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

const apiId = "prediction:v1.6"
const apiName = "prediction"
const apiVersion = "v1.6"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/prediction/v1.6/projects/"}

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
	s.Trainedmodels = NewTrainedmodelsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Hostedmodels *HostedmodelsService

	Trainedmodels *TrainedmodelsService
}

func NewHostedmodelsService(s *Service) *HostedmodelsService {
	rs := &HostedmodelsService{s: s}
	return rs
}

type HostedmodelsService struct {
	s *Service
}

func NewTrainedmodelsService(s *Service) *TrainedmodelsService {
	rs := &TrainedmodelsService{s: s}
	return rs
}

type TrainedmodelsService struct {
	s *Service
}

type Analyze struct {
	// DataDescription: Description of the data the model was trained on.
	DataDescription *AnalyzeDataDescription `json:"dataDescription,omitempty"`

	// Errors: List of errors with the data.
	Errors []map[string]string `json:"errors,omitempty"`

	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// ModelDescription: Description of the model.
	ModelDescription *AnalyzeModelDescription `json:"modelDescription,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type AnalyzeDataDescription struct {
	// Features: Description of the input features in the data set.
	Features []*AnalyzeDataDescriptionFeatures `json:"features,omitempty"`

	// OutputFeature: Description of the output value or label.
	OutputFeature *AnalyzeDataDescriptionOutputFeature `json:"outputFeature,omitempty"`
}

type AnalyzeDataDescriptionFeatures struct {
	// Categorical: Description of the categorical values of this feature.
	Categorical *AnalyzeDataDescriptionFeaturesCategorical `json:"categorical,omitempty"`

	// Index: The feature index.
	Index int64 `json:"index,omitempty,string"`

	// Numeric: Description of the numeric values of this feature.
	Numeric *AnalyzeDataDescriptionFeaturesNumeric `json:"numeric,omitempty"`

	// Text: Description of multiple-word text values of this feature.
	Text *AnalyzeDataDescriptionFeaturesText `json:"text,omitempty"`
}

type AnalyzeDataDescriptionFeaturesCategorical struct {
	// Count: Number of categorical values for this feature in the data.
	Count int64 `json:"count,omitempty,string"`

	// Values: List of all the categories for this feature in the data set.
	Values []*AnalyzeDataDescriptionFeaturesCategoricalValues `json:"values,omitempty"`
}

type AnalyzeDataDescriptionFeaturesCategoricalValues struct {
	// Count: Number of times this feature had this value.
	Count int64 `json:"count,omitempty,string"`

	// Value: The category name.
	Value string `json:"value,omitempty"`
}

type AnalyzeDataDescriptionFeaturesNumeric struct {
	// Count: Number of numeric values for this feature in the data set.
	Count int64 `json:"count,omitempty,string"`

	// Mean: Mean of the numeric values of this feature in the data set.
	Mean string `json:"mean,omitempty"`

	// Variance: Variance of the numeric values of this feature in the data
	// set.
	Variance string `json:"variance,omitempty"`
}

type AnalyzeDataDescriptionFeaturesText struct {
	// Count: Number of multiple-word text values for this feature.
	Count int64 `json:"count,omitempty,string"`
}

type AnalyzeDataDescriptionOutputFeature struct {
	// Numeric: Description of the output values in the data set.
	Numeric *AnalyzeDataDescriptionOutputFeatureNumeric `json:"numeric,omitempty"`

	// Text: Description of the output labels in the data set.
	Text []*AnalyzeDataDescriptionOutputFeatureText `json:"text,omitempty"`
}

type AnalyzeDataDescriptionOutputFeatureNumeric struct {
	// Count: Number of numeric output values in the data set.
	Count int64 `json:"count,omitempty,string"`

	// Mean: Mean of the output values in the data set.
	Mean string `json:"mean,omitempty"`

	// Variance: Variance of the output values in the data set.
	Variance string `json:"variance,omitempty"`
}

type AnalyzeDataDescriptionOutputFeatureText struct {
	// Count: Number of times the output label occurred in the data set.
	Count int64 `json:"count,omitempty,string"`

	// Value: The output label.
	Value string `json:"value,omitempty"`
}

type AnalyzeModelDescription struct {
	// ConfusionMatrix: An output confusion matrix. This shows an estimate
	// for how this model will do in predictions. This is first indexed by
	// the true class label. For each true class label, this provides a pair
	// {predicted_label, count}, where count is the estimated number of
	// times the model will predict the predicted label given the true
	// label. Will not output if more then 100 classes (Categorical models
	// only).
	ConfusionMatrix *AnalyzeModelDescriptionConfusionMatrix `json:"confusionMatrix,omitempty"`

	// ConfusionMatrixRowTotals: A list of the confusion matrix row totals.
	ConfusionMatrixRowTotals map[string]string `json:"confusionMatrixRowTotals,omitempty"`

	// Modelinfo: Basic information about the model.
	Modelinfo *Insert2 `json:"modelinfo,omitempty"`
}

type AnalyzeModelDescriptionConfusionMatrix struct {
}

type Input struct {
	// Input: Input to the model for a prediction.
	Input *InputInput `json:"input,omitempty"`
}

type InputInput struct {
	// CsvInstance: A list of input features, these can be strings or
	// doubles.
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

type Insert struct {
	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// ModelType: Type of predictive model (classification or regression).
	ModelType string `json:"modelType,omitempty"`

	// SourceModel: The Id of the model to be copied over.
	SourceModel string `json:"sourceModel,omitempty"`

	// StorageDataLocation: Google storage location of the training data
	// file.
	StorageDataLocation string `json:"storageDataLocation,omitempty"`

	// StoragePMMLLocation: Google storage location of the preprocessing
	// pmml file.
	StoragePMMLLocation string `json:"storagePMMLLocation,omitempty"`

	// StoragePMMLModelLocation: Google storage location of the pmml model
	// file.
	StoragePMMLModelLocation string `json:"storagePMMLModelLocation,omitempty"`

	// TrainingInstances: Instances to train model on.
	TrainingInstances []*InsertTrainingInstances `json:"trainingInstances,omitempty"`

	// Utility: A class weighting function, which allows the importance
	// weights for class labels to be specified (Categorical models only).
	Utility []*InsertUtility `json:"utility,omitempty"`
}

type InsertTrainingInstances struct {
	// CsvInstance: The input features for this instance.
	CsvInstance []interface{} `json:"csvInstance,omitempty"`

	// Output: The generic output value - could be regression or class
	// label.
	Output string `json:"output,omitempty"`
}

type InsertUtility struct {
}

type Insert2 struct {
	// Created: Insert time of the model (as a RFC 3339 timestamp).
	Created string `json:"created,omitempty"`

	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// ModelInfo: Model metadata.
	ModelInfo *Insert2ModelInfo `json:"modelInfo,omitempty"`

	// ModelType: Type of predictive model (CLASSIFICATION or REGRESSION).
	ModelType string `json:"modelType,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StorageDataLocation: Google storage location of the training data
	// file.
	StorageDataLocation string `json:"storageDataLocation,omitempty"`

	// StoragePMMLLocation: Google storage location of the preprocessing
	// pmml file.
	StoragePMMLLocation string `json:"storagePMMLLocation,omitempty"`

	// StoragePMMLModelLocation: Google storage location of the pmml model
	// file.
	StoragePMMLModelLocation string `json:"storagePMMLModelLocation,omitempty"`

	// TrainingComplete: Training completion time (as a RFC 3339 timestamp).
	TrainingComplete string `json:"trainingComplete,omitempty"`

	// TrainingStatus: The current status of the training job. This can be
	// one of following: RUNNING; DONE; ERROR; ERROR: TRAINING JOB NOT FOUND
	TrainingStatus string `json:"trainingStatus,omitempty"`
}

type Insert2ModelInfo struct {
	// ClassWeightedAccuracy: Estimated accuracy of model taking utility
	// weights into account (Categorical models only).
	ClassWeightedAccuracy string `json:"classWeightedAccuracy,omitempty"`

	// ClassificationAccuracy: A number between 0.0 and 1.0, where 1.0 is
	// 100% accurate. This is an estimate, based on the amount and quality
	// of the training data, of the estimated prediction accuracy. You can
	// use this is a guide to decide whether the results are accurate enough
	// for your needs. This estimate will be more reliable if your real
	// input data is similar to your training data (Categorical models
	// only).
	ClassificationAccuracy string `json:"classificationAccuracy,omitempty"`

	// MeanSquaredError: An estimated mean squared error. The can be used to
	// measure the quality of the predicted model (Regression models only).
	MeanSquaredError string `json:"meanSquaredError,omitempty"`

	// ModelType: Type of predictive model (CLASSIFICATION or REGRESSION).
	ModelType string `json:"modelType,omitempty"`

	// NumberInstances: Number of valid data instances used in the trained
	// model.
	NumberInstances int64 `json:"numberInstances,omitempty,string"`

	// NumberLabels: Number of class labels in the trained model
	// (Categorical models only).
	NumberLabels int64 `json:"numberLabels,omitempty,string"`
}

type List struct {
	// Items: List of models.
	Items []*Insert2 `json:"items,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to fetch the next page, if one
	// exists.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type Output struct {
	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// OutputLabel: The most likely class label (Categorical models only).
	OutputLabel string `json:"outputLabel,omitempty"`

	// OutputMulti: A list of class labels with their estimated
	// probabilities (Categorical models only).
	OutputMulti []*OutputOutputMulti `json:"outputMulti,omitempty"`

	// OutputValue: The estimated regression value (Regression models only).
	OutputValue string `json:"outputValue,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type OutputOutputMulti struct {
	// Label: The class label.
	Label string `json:"label,omitempty"`

	// Score: The probability of the class label.
	Score string `json:"score,omitempty"`
}

type Update struct {
	// CsvInstance: The input features for this instance.
	CsvInstance []interface{} `json:"csvInstance,omitempty"`

	// Output: The generic output value - could be regression or class
	// label.
	Output string `json:"output,omitempty"`
}

// method id "prediction.hostedmodels.predict":

type HostedmodelsPredictCall struct {
	s               *Service
	project         string
	hostedModelName string
	input           *Input
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Predict: Submit input and request an output against a hosted model.

func (r *HostedmodelsService) Predict(project string, hostedModelName string, input *Input) *HostedmodelsPredictCall {
	return &HostedmodelsPredictCall{
		s:               r.s,
		project:         project,
		hostedModelName: hostedModelName,
		input:           input,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "{project}/hostedmodels/{hostedModelName}/predict",
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
		"project":         c.project,
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
	//   "description": "Submit input and request an output against a hosted model.",
	//   "httpMethod": "POST",
	//   "id": "prediction.hostedmodels.predict",
	//   "parameterOrder": [
	//     "project",
	//     "hostedModelName"
	//   ],
	//   "parameters": {
	//     "hostedModelName": {
	//       "description": "The name of a hosted model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/hostedmodels/{hostedModelName}/predict",
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

// method id "prediction.trainedmodels.analyze":

type TrainedmodelsAnalyzeCall struct {
	s             *Service
	project       string
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Analyze: Get analysis of the model and the data the model was trained
// on.

func (r *TrainedmodelsService) Analyze(project string, id string) *TrainedmodelsAnalyzeCall {
	return &TrainedmodelsAnalyzeCall{
		s:             r.s,
		project:       project,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/{id}/analyze",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsAnalyzeCall) Fields(s ...googleapi.Field) *TrainedmodelsAnalyzeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsAnalyzeCall) Do() (*Analyze, error) {
	var returnValue *Analyze
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"id":      c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get analysis of the model and the data the model was trained on.",
	//   "httpMethod": "GET",
	//   "id": "prediction.trainedmodels.analyze",
	//   "parameterOrder": [
	//     "project",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/{id}/analyze",
	//   "response": {
	//     "$ref": "Analyze"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.delete":

type TrainedmodelsDeleteCall struct {
	s             *Service
	project       string
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete a trained model.

func (r *TrainedmodelsService) Delete(project string, id string) *TrainedmodelsDeleteCall {
	return &TrainedmodelsDeleteCall{
		s:             r.s,
		project:       project,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/{id}",
	}
}

func (c *TrainedmodelsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"id":      c.id,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a trained model.",
	//   "httpMethod": "DELETE",
	//   "id": "prediction.trainedmodels.delete",
	//   "parameterOrder": [
	//     "project",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.get":

type TrainedmodelsGetCall struct {
	s             *Service
	project       string
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Check training status of your model.

func (r *TrainedmodelsService) Get(project string, id string) *TrainedmodelsGetCall {
	return &TrainedmodelsGetCall{
		s:             r.s,
		project:       project,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsGetCall) Fields(s ...googleapi.Field) *TrainedmodelsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsGetCall) Do() (*Insert2, error) {
	var returnValue *Insert2
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"id":      c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Check training status of your model.",
	//   "httpMethod": "GET",
	//   "id": "prediction.trainedmodels.get",
	//   "parameterOrder": [
	//     "project",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/{id}",
	//   "response": {
	//     "$ref": "Insert2"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.insert":

type TrainedmodelsInsertCall struct {
	s             *Service
	project       string
	insert        *Insert
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Train a Prediction API model.

func (r *TrainedmodelsService) Insert(project string, insert *Insert) *TrainedmodelsInsertCall {
	return &TrainedmodelsInsertCall{
		s:             r.s,
		project:       project,
		insert:        insert,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsInsertCall) Fields(s ...googleapi.Field) *TrainedmodelsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsInsertCall) Do() (*Insert2, error) {
	var returnValue *Insert2
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.insert,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Train a Prediction API model.",
	//   "httpMethod": "POST",
	//   "id": "prediction.trainedmodels.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels",
	//   "request": {
	//     "$ref": "Insert"
	//   },
	//   "response": {
	//     "$ref": "Insert2"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.list":

type TrainedmodelsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List available models.

func (r *TrainedmodelsService) List(project string) *TrainedmodelsListCall {
	return &TrainedmodelsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/list",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *TrainedmodelsListCall) MaxResults(maxResults int64) *TrainedmodelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Pagination token.
func (c *TrainedmodelsListCall) PageToken(pageToken string) *TrainedmodelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsListCall) Fields(s ...googleapi.Field) *TrainedmodelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsListCall) Do() (*List, error) {
	var returnValue *List
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List available models.",
	//   "httpMethod": "GET",
	//   "id": "prediction.trainedmodels.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Pagination token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/list",
	//   "response": {
	//     "$ref": "List"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.trainedmodels.predict":

type TrainedmodelsPredictCall struct {
	s             *Service
	project       string
	id            string
	input         *Input
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Predict: Submit model id and request a prediction.

func (r *TrainedmodelsService) Predict(project string, id string, input *Input) *TrainedmodelsPredictCall {
	return &TrainedmodelsPredictCall{
		s:             r.s,
		project:       project,
		id:            id,
		input:         input,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/{id}/predict",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsPredictCall) Fields(s ...googleapi.Field) *TrainedmodelsPredictCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsPredictCall) Do() (*Output, error) {
	var returnValue *Output
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"id":      c.id,
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
	//   "description": "Submit model id and request a prediction.",
	//   "httpMethod": "POST",
	//   "id": "prediction.trainedmodels.predict",
	//   "parameterOrder": [
	//     "project",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/{id}/predict",
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

// method id "prediction.trainedmodels.update":

type TrainedmodelsUpdateCall struct {
	s             *Service
	project       string
	id            string
	update        *Update
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Add new data to a trained model.

func (r *TrainedmodelsService) Update(project string, id string, update *Update) *TrainedmodelsUpdateCall {
	return &TrainedmodelsUpdateCall{
		s:             r.s,
		project:       project,
		id:            id,
		update:        update,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/trainedmodels/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainedmodelsUpdateCall) Fields(s ...googleapi.Field) *TrainedmodelsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TrainedmodelsUpdateCall) Do() (*Insert2, error) {
	var returnValue *Insert2
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"id":      c.id,
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
	//   "description": "Add new data to a trained model.",
	//   "httpMethod": "PUT",
	//   "id": "prediction.trainedmodels.update",
	//   "parameterOrder": [
	//     "project",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The unique name for the predictive model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project associated with the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/trainedmodels/{id}",
	//   "request": {
	//     "$ref": "Update"
	//   },
	//   "response": {
	//     "$ref": "Insert2"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// Package autoscaler provides access to the Google Compute Engine Autoscaler API.
//
// See http://developers.google.com/compute/docs/autoscaler
//
// Usage example:
//
//   import "github.com/jfcote87/api/autoscaler/v1beta2"
//   ...
//   autoscalerService, err := autoscaler.New(oauthHttpClient)
package autoscaler

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

const apiId = "autoscaler:v1beta2"
const apiName = "autoscaler"
const apiVersion = "v1beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/autoscaler/v1beta2/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Compute Engine resources
	ComputeScope = "https://www.googleapis.com/auth/compute"

	// View your Google Compute Engine resources
	ComputeReadonlyScope = "https://www.googleapis.com/auth/compute.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Autoscalers = NewAutoscalersService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	s.Zones = NewZonesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Autoscalers *AutoscalersService

	ZoneOperations *ZoneOperationsService

	Zones *ZonesService
}

func NewAutoscalersService(s *Service) *AutoscalersService {
	rs := &AutoscalersService{s: s}
	return rs
}

type AutoscalersService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

func NewZonesService(s *Service) *ZonesService {
	rs := &ZonesService{s: s}
	return rs
}

type ZonesService struct {
	s *Service
}

type Autoscaler struct {
	// AutoscalingPolicy: Configuration parameters for autoscaling
	// algorithm.
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource provided
	// by the client.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the Autoscaler resource. Must be unique per project and
	// zone.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] A self-link to the Autoscaler configuration
	// resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Target: URL to the entity which will be autoscaled. Currently the
	// only supported value is ReplicaPool?s URL. Note: it is illegal to
	// specify multiple Autoscalers for the same target.
	Target string `json:"target,omitempty"`
}

type AutoscalerListResponse struct {
	// Items: Autoscaler resources.
	Items []*Autoscaler `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AutoscalingPolicy struct {
	// CoolDownPeriodSec: The number of seconds that the Autoscaler should
	// wait between two succeeding changes to the number of virtual
	// machines. You should define an interval that is at least as long as
	// the initialization time of a virtual machine and the time it may take
	// for replica pool to create the virtual machine. The default is 60
	// seconds.
	CoolDownPeriodSec int64 `json:"coolDownPeriodSec,omitempty"`

	// CpuUtilization: Exactly one utilization policy should be provided.
	// Configuration parameters of CPU based autoscaling policy.
	CpuUtilization *AutoscalingPolicyCpuUtilization `json:"cpuUtilization,omitempty"`

	// CustomMetricUtilizations: Configuration parameters of autoscaling
	// based on custom metric.
	CustomMetricUtilizations []*AutoscalingPolicyCustomMetricUtilization `json:"customMetricUtilizations,omitempty"`

	// LoadBalancingUtilization: Configuration parameters of autoscaling
	// based on load balancer.
	LoadBalancingUtilization *AutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty"`

	// MaxNumReplicas: The maximum number of replicas that the Autoscaler
	// can scale up to.
	MaxNumReplicas int64 `json:"maxNumReplicas,omitempty"`

	// MinNumReplicas: The minimum number of replicas that the Autoscaler
	// can scale down to.
	MinNumReplicas int64 `json:"minNumReplicas,omitempty"`
}

type AutoscalingPolicyCpuUtilization struct {
	// UtilizationTarget: The target utilization that the Autoscaler should
	// maintain. It is represented as a fraction of used cores. For example:
	// 6 cores used in 8-core VM are represented here as 0.75. Must be a
	// float value between (0, 1]. If not defined, the default is 0.8.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`
}

type AutoscalingPolicyCustomMetricUtilization struct {
	// Metric: Identifier of the metric. It should be a Cloud Monitoring
	// metric. The metric can not have negative values. The metric should be
	// an utilization metric (increasing number of VMs handling requests x
	// times should reduce average value of the metric roughly x times). For
	// example you could use:
	// compute.googleapis.com/instance/network/received_bytes_count.
	Metric string `json:"metric,omitempty"`

	// UtilizationTarget: Target value of the metric which Autoscaler should
	// maintain. Must be a positive value.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`

	// UtilizationTargetType: Defines type in which utilization_target is
	// expressed.
	UtilizationTargetType string `json:"utilizationTargetType,omitempty"`
}

type AutoscalingPolicyLoadBalancingUtilization struct {
	// UtilizationTarget: Fraction of backend capacity utilization (set in
	// HTTP load balancing configuration) that Autoscaler should maintain.
	// Must be a positive float value. If not defined, the default is 0.8.
	// For example if your maxRatePerInstance capacity (in HTTP Load
	// Balancing configuration) is set at 10 and you would like to keep
	// number of instances such that each instance receives 7 QPS on
	// average, set this to 0.7.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`
}

type DeprecationStatus struct {
	Deleted string `json:"deleted,omitempty"`

	Deprecated string `json:"deprecated,omitempty"`

	Obsolete string `json:"obsolete,omitempty"`

	Replacement string `json:"replacement,omitempty"`

	State string `json:"state,omitempty"`
}

type Operation struct {
	ClientOperationId string `json:"clientOperationId,omitempty"`

	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	Error *OperationError `json:"error,omitempty"`

	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	Id uint64 `json:"id,omitempty,string"`

	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always kind#operation for
	// Operation resources.
	Kind string `json:"kind,omitempty"`

	Name string `json:"name,omitempty"`

	OperationType string `json:"operationType,omitempty"`

	Progress int64 `json:"progress,omitempty"`

	Region string `json:"region,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	Status string `json:"status,omitempty"`

	StatusMessage string `json:"statusMessage,omitempty"`

	TargetId uint64 `json:"targetId,omitempty,string"`

	TargetLink string `json:"targetLink,omitempty"`

	User string `json:"user,omitempty"`

	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	Zone string `json:"zone,omitempty"`
}

type OperationError struct {
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	Code string `json:"code,omitempty"`

	Location string `json:"location,omitempty"`

	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	Code string `json:"code,omitempty"`

	Data []*OperationWarningsData `json:"data,omitempty"`

	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type OperationList struct {
	Id string `json:"id,omitempty"`

	Items []*Operation `json:"items,omitempty"`

	// Kind: Type of resource. Always compute#operations for Operations
	// resource.
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type Zone struct {
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	Description string `json:"description,omitempty"`

	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of the resource.
	Kind string `json:"kind,omitempty"`

	MaintenanceWindows []*ZoneMaintenanceWindows `json:"maintenanceWindows,omitempty"`

	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`

	// SelfLink: Server defined URL for the resource (output only).
	SelfLink string `json:"selfLink,omitempty"`

	Status string `json:"status,omitempty"`
}

type ZoneMaintenanceWindows struct {
	BeginTime string `json:"beginTime,omitempty"`

	Description string `json:"description,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	Name string `json:"name,omitempty"`
}

type ZoneList struct {
	Id string `json:"id,omitempty"`

	Items []*Zone `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "autoscaler.autoscalers.delete":

type AutoscalersDeleteCall struct {
	s             *Service
	project       string
	zone          string
	autoscaler    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified Autoscaler resource.

func (r *AutoscalersService) Delete(project string, zone string, autoscaler string) *AutoscalersDeleteCall {
	return &AutoscalersDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		autoscaler:    autoscaler,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersDeleteCall) Fields(s ...googleapi.Field) *AutoscalersDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified Autoscaler resource.",
	//   "httpMethod": "DELETE",
	//   "id": "autoscaler.autoscalers.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.get":

type AutoscalersGetCall struct {
	s             *Service
	project       string
	zone          string
	autoscaler    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets the specified Autoscaler resource.

func (r *AutoscalersService) Get(project string, zone string, autoscaler string) *AutoscalersGetCall {
	return &AutoscalersGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		autoscaler:    autoscaler,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersGetCall) Fields(s ...googleapi.Field) *AutoscalersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersGetCall) Do() (*Autoscaler, error) {
	var returnValue *Autoscaler
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the specified Autoscaler resource.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.autoscalers.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "response": {
	//     "$ref": "Autoscaler"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.insert":

type AutoscalersInsertCall struct {
	s             *Service
	project       string
	zone          string
	autoscaler    *Autoscaler
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Adds new Autoscaler resource.

func (r *AutoscalersService) Insert(project string, zone string, autoscaler *Autoscaler) *AutoscalersInsertCall {
	return &AutoscalersInsertCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		autoscaler:    autoscaler,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersInsertCall) Fields(s ...googleapi.Field) *AutoscalersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.autoscaler,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds new Autoscaler resource.",
	//   "httpMethod": "POST",
	//   "id": "autoscaler.autoscalers.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.list":

type AutoscalersListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all Autoscaler resources in this zone.

func (r *AutoscalersService) List(project string, zone string) *AutoscalersListCall {
	return &AutoscalersListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers",
	}
}

// Filter sets the optional parameter "filter":
func (c *AutoscalersListCall) Filter(filter string) *AutoscalersListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *AutoscalersListCall) MaxResults(maxResults int64) *AutoscalersListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *AutoscalersListCall) PageToken(pageToken string) *AutoscalersListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersListCall) Fields(s ...googleapi.Field) *AutoscalersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersListCall) Do() (*AutoscalerListResponse, error) {
	var returnValue *AutoscalerListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all Autoscaler resources in this zone.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.autoscalers.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers",
	//   "response": {
	//     "$ref": "AutoscalerListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.patch":

type AutoscalersPatchCall struct {
	s             *Service
	project       string
	zone          string
	autoscaler    string
	autoscaler2   *Autoscaler
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update the entire content of the Autoscaler resource. This
// method supports patch semantics.

func (r *AutoscalersService) Patch(project string, zone string, autoscaler string, autoscaler2 *Autoscaler) *AutoscalersPatchCall {
	return &AutoscalersPatchCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		autoscaler:    autoscaler,
		autoscaler2:   autoscaler2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersPatchCall) Fields(s ...googleapi.Field) *AutoscalersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersPatchCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.autoscaler2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the entire content of the Autoscaler resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "autoscaler.autoscalers.patch",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.update":

type AutoscalersUpdateCall struct {
	s             *Service
	project       string
	zone          string
	autoscaler    string
	autoscaler2   *Autoscaler
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update the entire content of the Autoscaler resource.

func (r *AutoscalersService) Update(project string, zone string, autoscaler string, autoscaler2 *Autoscaler) *AutoscalersUpdateCall {
	return &AutoscalersUpdateCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		autoscaler:    autoscaler,
		autoscaler2:   autoscaler2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersUpdateCall) Fields(s ...googleapi.Field) *AutoscalersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AutoscalersUpdateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.autoscaler2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the entire content of the Autoscaler resource.",
	//   "httpMethod": "PUT",
	//   "id": "autoscaler.autoscalers.update",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.delete":

type ZoneOperationsDeleteCall struct {
	s             *Service
	project       string
	zone          string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified zone-specific operation resource.

func (r *ZoneOperationsService) Delete(project string, zone string, operation string) *ZoneOperationsDeleteCall {
	return &ZoneOperationsDeleteCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations/{operation}",
	}
}

func (c *ZoneOperationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified zone-specific operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "autoscaler.zoneOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.get":

type ZoneOperationsGetCall struct {
	s             *Service
	project       string
	zone          string
	operation     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the specified zone-specific operation resource.

func (r *ZoneOperationsService) Get(project string, zone string, operation string) *ZoneOperationsGetCall {
	return &ZoneOperationsGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		operation:     operation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations/{operation}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsGetCall) Fields(s ...googleapi.Field) *ZoneOperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneOperationsGetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the specified zone-specific operation resource.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.list":

type ZoneOperationsListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of operation resources contained within the
// specified zone.

func (r *ZoneOperationsService) List(project string, zone string) *ZoneOperationsListCall {
	return &ZoneOperationsListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/operations",
	}
}

// Filter sets the optional parameter "filter":
func (c *ZoneOperationsListCall) Filter(filter string) *ZoneOperationsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *ZoneOperationsListCall) MaxResults(maxResults int64) *ZoneOperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ZoneOperationsListCall) PageToken(pageToken string) *ZoneOperationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsListCall) Fields(s ...googleapi.Field) *ZoneOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZoneOperationsListCall) Do() (*OperationList, error) {
	var returnValue *OperationList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of operation resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zoneOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.zones.list":

type ZonesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List:

func (r *ZonesService) List(project string) *ZonesListCall {
	return &ZonesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones",
	}
}

// Filter sets the optional parameter "filter":
func (c *ZonesListCall) Filter(filter string) *ZonesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *ZonesListCall) MaxResults(maxResults int64) *ZonesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ZonesListCall) PageToken(pageToken string) *ZonesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZonesListCall) Fields(s ...googleapi.Field) *ZonesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ZonesListCall) Do() (*ZoneList, error) {
	var returnValue *ZoneList
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
	//   "description": "",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zones.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones",
	//   "response": {
	//     "$ref": "ZoneList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// Package replicapool provides access to the Google Compute Engine Instance Group Manager API.
//
// See https://developers.google.com/compute/docs/instance-groups/manager/v1beta2
//
// Usage example:
//
//   import "github.com/jfcote87/api/replicapool/v1beta2"
//   ...
//   replicapoolService, err := replicapool.New(oauthHttpClient)
package replicapool

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

const apiId = "replicapool:v1beta2"
const apiName = "replicapool"
const apiVersion = "v1beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/replicapool/v1beta2/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

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
	s.InstanceGroupManagers = NewInstanceGroupManagersService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	InstanceGroupManagers *InstanceGroupManagersService

	ZoneOperations *ZoneOperationsService
}

func NewInstanceGroupManagersService(s *Service) *InstanceGroupManagersService {
	rs := &InstanceGroupManagersService{s: s}
	return rs
}

type InstanceGroupManagersService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

type InstanceGroupManager struct {
	// BaseInstanceName: The base instance name to use for instances in this
	// group. The value must be a valid RFC1035 name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named
	// by appending a hyphen and a random four-character string to the base
	// instance name.
	BaseInstanceName string `json:"baseInstanceName,omitempty"`

	// CreationTimestamp: [Output only] The time the instance group manager
	// was created, in RFC3339 text format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// CurrentSize: [Output only] The number of instances that currently
	// exist and are a part of this group. This includes instances that are
	// starting but are not yet RUNNING, and instances that are in the
	// process of being deleted or abandoned.
	CurrentSize int64 `json:"currentSize,omitempty"`

	// Description: An optional textual description of the instance group
	// manager.
	Description string `json:"description,omitempty"`

	// Fingerprint: [Output only] Fingerprint of the instance group manager.
	// This field is used for optimistic locking. An up-to-date fingerprint
	// must be provided in order to modify the Instance Group Manager
	// resource.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Group: [Output only] The full URL of the instance group created by
	// the manager. This group contains all of the instances being managed,
	// and cannot contain non-managed instances.
	Group string `json:"group,omitempty"`

	// Id: [Output only] A server-assigned unique identifier for the
	// resource.
	Id uint64 `json:"id,omitempty,string"`

	// InstanceTemplate: The full URL to an instance template from which all
	// new instances will be created.
	InstanceTemplate string `json:"instanceTemplate,omitempty"`

	// Kind: [Output only] The resource type. Always
	// replicapool#instanceGroupManager.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the instance group manager. Must be 1-63 characters
	// long and comply with RFC1035. Supported characters include lowercase
	// letters, numbers, and hyphens.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output only] The fully qualified URL for this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// TargetPools: The full URL of all target pools to which new instances
	// in the group are added. Updating the target pool values does not
	// affect existing instances.
	TargetPools []string `json:"targetPools,omitempty"`

	// TargetSize: [Output only] The number of instances that the manager is
	// attempting to maintain. Deleting or abandoning instances affects this
	// number, as does resizing the group.
	TargetSize int64 `json:"targetSize,omitempty"`
}

type InstanceGroupManagerList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: A list of instance resources.
	Items []*InstanceGroupManager `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

type InstanceGroupManagersAbandonInstancesRequest struct {
	// Instances: The names of one or more instances to abandon. For
	// example:
	// { 'instances': [ 'instance-c3po', 'instance-r2d2' ] }
	Instances []string `json:"instances,omitempty"`
}

type InstanceGroupManagersDeleteInstancesRequest struct {
	// Instances: Names of instances to delete.
	//
	// Example: 'instance-foo',
	// 'instance-bar'
	Instances []string `json:"instances,omitempty"`
}

type InstanceGroupManagersRecreateInstancesRequest struct {
	// Instances: The names of one or more instances to recreate. For
	// example:
	// { 'instances': [ 'instance-c3po', 'instance-r2d2' ] }
	Instances []string `json:"instances,omitempty"`
}

type InstanceGroupManagersSetInstanceTemplateRequest struct {
	// InstanceTemplate: The full URL to an Instance Template from which all
	// new instances will be created.
	InstanceTemplate string `json:"instanceTemplate,omitempty"`
}

type InstanceGroupManagersSetTargetPoolsRequest struct {
	// Fingerprint: The current fingerprint of the Instance Group Manager
	// resource. If this does not match the server-side fingerprint of the
	// resource, then the request will be rejected.
	Fingerprint string `json:"fingerprint,omitempty"`

	// TargetPools: A list of fully-qualified URLs to existing Target Pool
	// resources. New instances in the Instance Group Manager will be added
	// to the specified target pools; existing instances are not affected.
	TargetPools []string `json:"targetPools,omitempty"`
}

type Operation struct {
	// ClientOperationId: [Output only] An optional identifier specified by
	// the client when the mutation was initiated. Must be unique for all
	// operation resources in the project.
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Output Only] The time that this operation was
	// requested, in RFC3339 text format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// EndTime: [Output Only] The time that this operation was completed, in
	// RFC3339 text format.
	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors occurred during processing of this
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	// HttpErrorMessage: [Output only] If operation fails, the HTTP error
	// message returned.
	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	// HttpErrorStatusCode: [Output only] If operation fails, the HTTP error
	// status code returned.
	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] Unique identifier for the resource, generated by
	// the server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested,
	// in RFC3339 text format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	// OperationType: [Output only] Type of the operation. Operations
	// include insert, update, and delete.
	OperationType string `json:"operationType,omitempty"`

	// Progress: [Output only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess at
	// when the operation will be complete. This number should be
	// monotonically increasing as the operation progresses.
	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] URL of the region where the operation resides.
	// Only available when performing regional operations.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server-defined fully-qualified URL for this
	// resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server, in RFC3339 text format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] Status of the operation.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] Unique target ID which identifies a
	// particular incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	// User: [Output Only] User who requested the operation, for example:
	// user@example.com.
	User string `json:"user,omitempty"`

	// Warnings: [Output Only] If there are issues with this operation, a
	// warning is returned.
	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] URL of the zone where the operation resides. Only
	// available when performing per-zone operations.
	Zone string `json:"zone,omitempty"`
}

type OperationError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request which
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	// Code: [Output only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output only] Metadata for this warning in key:value format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: [Output only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	// Key: [Output Only] Metadata key for this warning.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] Metadata value for this warning.
	Value string `json:"value,omitempty"`
}

type OperationList struct {
	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Items: The operation resources.
	Items []*Operation `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request
	// (output only).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "replicapool.instanceGroupManagers.abandonInstances":

type InstanceGroupManagersAbandonInstancesCall struct {
	s                                            *Service
	project                                      string
	zone                                         string
	instanceGroupManager                         string
	instancegroupmanagersabandoninstancesrequest *InstanceGroupManagersAbandonInstancesRequest
	caller_                                      googleapi.Caller
	params_                                      url.Values
	pathTemplate_                                string
}

// AbandonInstances: Removes the specified instances from the managed
// instance group, and from any target pools of which they were members,
// without deleting the instances.

func (r *InstanceGroupManagersService) AbandonInstances(project string, zone string, instanceGroupManager string, instancegroupmanagersabandoninstancesrequest *InstanceGroupManagersAbandonInstancesRequest) *InstanceGroupManagersAbandonInstancesCall {
	return &InstanceGroupManagersAbandonInstancesCall{
		s:                                            r.s,
		project:                                      project,
		zone:                                         zone,
		instanceGroupManager:                         instanceGroupManager,
		instancegroupmanagersabandoninstancesrequest: instancegroupmanagersabandoninstancesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/abandonInstances",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersAbandonInstancesCall) Fields(s ...googleapi.Field) *InstanceGroupManagersAbandonInstancesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersAbandonInstancesCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanagersabandoninstancesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes the specified instances from the managed instance group, and from any target pools of which they were members, without deleting the instances.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.abandonInstances",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/abandonInstances",
	//   "request": {
	//     "$ref": "InstanceGroupManagersAbandonInstancesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.delete":

type InstanceGroupManagersDeleteCall struct {
	s                    *Service
	project              string
	zone                 string
	instanceGroupManager string
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// Delete: Deletes the instance group manager and all instances
// contained within. If you'd like to delete the manager without
// deleting the instances, you must first abandon the instances to
// remove them from the group.

func (r *InstanceGroupManagersService) Delete(project string, zone string, instanceGroupManager string) *InstanceGroupManagersDeleteCall {
	return &InstanceGroupManagersDeleteCall{
		s:                    r.s,
		project:              project,
		zone:                 zone,
		instanceGroupManager: instanceGroupManager,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersDeleteCall) Fields(s ...googleapi.Field) *InstanceGroupManagersDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the instance group manager and all instances contained within. If you'd like to delete the manager without deleting the instances, you must first abandon the instances to remove them from the group.",
	//   "httpMethod": "DELETE",
	//   "id": "replicapool.instanceGroupManagers.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "Name of the Instance Group Manager resource to delete.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.deleteInstances":

type InstanceGroupManagersDeleteInstancesCall struct {
	s                                           *Service
	project                                     string
	zone                                        string
	instanceGroupManager                        string
	instancegroupmanagersdeleteinstancesrequest *InstanceGroupManagersDeleteInstancesRequest
	caller_                                     googleapi.Caller
	params_                                     url.Values
	pathTemplate_                               string
}

// DeleteInstances: Deletes the specified instances. The instances are
// removed from the instance group and any target pools of which they
// are a member, then deleted. The targetSize of the instance group
// manager is reduced by the number of instances deleted.

func (r *InstanceGroupManagersService) DeleteInstances(project string, zone string, instanceGroupManager string, instancegroupmanagersdeleteinstancesrequest *InstanceGroupManagersDeleteInstancesRequest) *InstanceGroupManagersDeleteInstancesCall {
	return &InstanceGroupManagersDeleteInstancesCall{
		s:                                           r.s,
		project:                                     project,
		zone:                                        zone,
		instanceGroupManager:                        instanceGroupManager,
		instancegroupmanagersdeleteinstancesrequest: instancegroupmanagersdeleteinstancesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/deleteInstances",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersDeleteInstancesCall) Fields(s ...googleapi.Field) *InstanceGroupManagersDeleteInstancesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersDeleteInstancesCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanagersdeleteinstancesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified instances. The instances are removed from the instance group and any target pools of which they are a member, then deleted. The targetSize of the instance group manager is reduced by the number of instances deleted.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.deleteInstances",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/deleteInstances",
	//   "request": {
	//     "$ref": "InstanceGroupManagersDeleteInstancesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.get":

type InstanceGroupManagersGetCall struct {
	s                    *Service
	project              string
	zone                 string
	instanceGroupManager string
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// Get: Returns the specified Instance Group Manager resource.

func (r *InstanceGroupManagersService) Get(project string, zone string, instanceGroupManager string) *InstanceGroupManagersGetCall {
	return &InstanceGroupManagersGetCall{
		s:                    r.s,
		project:              project,
		zone:                 zone,
		instanceGroupManager: instanceGroupManager,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersGetCall) Fields(s ...googleapi.Field) *InstanceGroupManagersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersGetCall) Do() (*InstanceGroupManager, error) {
	var returnValue *InstanceGroupManager
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the specified Instance Group Manager resource.",
	//   "httpMethod": "GET",
	//   "id": "replicapool.instanceGroupManagers.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "Name of the instance resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}",
	//   "response": {
	//     "$ref": "InstanceGroupManager"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.insert":

type InstanceGroupManagersInsertCall struct {
	s                    *Service
	project              string
	zone                 string
	size                 int64
	instancegroupmanager *InstanceGroupManager
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// Insert: Creates an instance group manager, as well as the instance
// group and the specified number of instances.

func (r *InstanceGroupManagersService) Insert(project string, zone string, size int64, instancegroupmanager *InstanceGroupManager) *InstanceGroupManagersInsertCall {
	return &InstanceGroupManagersInsertCall{
		s:                    r.s,
		project:              project,
		zone:                 zone,
		size:                 size,
		instancegroupmanager: instancegroupmanager,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "{project}/zones/{zone}/instanceGroupManagers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersInsertCall) Fields(s ...googleapi.Field) *InstanceGroupManagersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("size", fmt.Sprintf("%v", c.size))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanager,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates an instance group manager, as well as the instance group and the specified number of instances.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "size"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "size": {
	//       "description": "Number of instances that should exist.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers",
	//   "request": {
	//     "$ref": "InstanceGroupManager"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.list":

type InstanceGroupManagersListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of Instance Group Manager resources
// contained within the specified zone.

func (r *InstanceGroupManagersService) List(project string, zone string) *InstanceGroupManagersListCall {
	return &InstanceGroupManagersListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers",
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *InstanceGroupManagersListCall) Filter(filter string) *InstanceGroupManagersListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *InstanceGroupManagersListCall) MaxResults(maxResults int64) *InstanceGroupManagersListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *InstanceGroupManagersListCall) PageToken(pageToken string) *InstanceGroupManagersListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersListCall) Fields(s ...googleapi.Field) *InstanceGroupManagersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersListCall) Do() (*InstanceGroupManagerList, error) {
	var returnValue *InstanceGroupManagerList
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
	//   "description": "Retrieves the list of Instance Group Manager resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "replicapool.instanceGroupManagers.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers",
	//   "response": {
	//     "$ref": "InstanceGroupManagerList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.recreateInstances":

type InstanceGroupManagersRecreateInstancesCall struct {
	s                                             *Service
	project                                       string
	zone                                          string
	instanceGroupManager                          string
	instancegroupmanagersrecreateinstancesrequest *InstanceGroupManagersRecreateInstancesRequest
	caller_                                       googleapi.Caller
	params_                                       url.Values
	pathTemplate_                                 string
}

// RecreateInstances: Recreates the specified instances. The instances
// are deleted, then recreated using the instance group manager's
// current instance template.

func (r *InstanceGroupManagersService) RecreateInstances(project string, zone string, instanceGroupManager string, instancegroupmanagersrecreateinstancesrequest *InstanceGroupManagersRecreateInstancesRequest) *InstanceGroupManagersRecreateInstancesCall {
	return &InstanceGroupManagersRecreateInstancesCall{
		s:                                             r.s,
		project:                                       project,
		zone:                                          zone,
		instanceGroupManager:                          instanceGroupManager,
		instancegroupmanagersrecreateinstancesrequest: instancegroupmanagersrecreateinstancesrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/recreateInstances",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersRecreateInstancesCall) Fields(s ...googleapi.Field) *InstanceGroupManagersRecreateInstancesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersRecreateInstancesCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanagersrecreateinstancesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Recreates the specified instances. The instances are deleted, then recreated using the instance group manager's current instance template.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.recreateInstances",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/recreateInstances",
	//   "request": {
	//     "$ref": "InstanceGroupManagersRecreateInstancesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.resize":

type InstanceGroupManagersResizeCall struct {
	s                    *Service
	project              string
	zone                 string
	instanceGroupManager string
	size                 int64
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// Resize: Resizes the managed instance group up or down. If resized up,
// new instances are created using the current instance template. If
// resized down, instances are removed in the order outlined in Resizing
// a managed instance group.

func (r *InstanceGroupManagersService) Resize(project string, zone string, instanceGroupManager string, size int64) *InstanceGroupManagersResizeCall {
	return &InstanceGroupManagersResizeCall{
		s:                    r.s,
		project:              project,
		zone:                 zone,
		instanceGroupManager: instanceGroupManager,
		size:                 size,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/resize",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersResizeCall) Fields(s ...googleapi.Field) *InstanceGroupManagersResizeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersResizeCall) Do() (*Operation, error) {
	var returnValue *Operation
	c.params_.Set("size", fmt.Sprintf("%v", c.size))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resizes the managed instance group up or down. If resized up, new instances are created using the current instance template. If resized down, instances are removed in the order outlined in Resizing a managed instance group.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.resize",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager",
	//     "size"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "size": {
	//       "description": "Number of instances that should exist in this Instance Group Manager.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/resize",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.setInstanceTemplate":

type InstanceGroupManagersSetInstanceTemplateCall struct {
	s                                               *Service
	project                                         string
	zone                                            string
	instanceGroupManager                            string
	instancegroupmanagerssetinstancetemplaterequest *InstanceGroupManagersSetInstanceTemplateRequest
	caller_                                         googleapi.Caller
	params_                                         url.Values
	pathTemplate_                                   string
}

// SetInstanceTemplate: Sets the instance template to use when creating
// new instances in this group. Existing instances are not affected.

func (r *InstanceGroupManagersService) SetInstanceTemplate(project string, zone string, instanceGroupManager string, instancegroupmanagerssetinstancetemplaterequest *InstanceGroupManagersSetInstanceTemplateRequest) *InstanceGroupManagersSetInstanceTemplateCall {
	return &InstanceGroupManagersSetInstanceTemplateCall{
		s:                                               r.s,
		project:                                         project,
		zone:                                            zone,
		instanceGroupManager:                            instanceGroupManager,
		instancegroupmanagerssetinstancetemplaterequest: instancegroupmanagerssetinstancetemplaterequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setInstanceTemplate",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersSetInstanceTemplateCall) Fields(s ...googleapi.Field) *InstanceGroupManagersSetInstanceTemplateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersSetInstanceTemplateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanagerssetinstancetemplaterequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Sets the instance template to use when creating new instances in this group. Existing instances are not affected.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.setInstanceTemplate",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setInstanceTemplate",
	//   "request": {
	//     "$ref": "InstanceGroupManagersSetInstanceTemplateRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.instanceGroupManagers.setTargetPools":

type InstanceGroupManagersSetTargetPoolsCall struct {
	s                                          *Service
	project                                    string
	zone                                       string
	instanceGroupManager                       string
	instancegroupmanagerssettargetpoolsrequest *InstanceGroupManagersSetTargetPoolsRequest
	caller_                                    googleapi.Caller
	params_                                    url.Values
	pathTemplate_                              string
}

// SetTargetPools: Modifies the target pools to which all new instances
// in this group are assigned. Existing instances in the group are not
// affected.

func (r *InstanceGroupManagersService) SetTargetPools(project string, zone string, instanceGroupManager string, instancegroupmanagerssettargetpoolsrequest *InstanceGroupManagersSetTargetPoolsRequest) *InstanceGroupManagersSetTargetPoolsCall {
	return &InstanceGroupManagersSetTargetPoolsCall{
		s:                                          r.s,
		project:                                    project,
		zone:                                       zone,
		instanceGroupManager:                       instanceGroupManager,
		instancegroupmanagerssettargetpoolsrequest: instancegroupmanagerssettargetpoolsrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setTargetPools",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstanceGroupManagersSetTargetPoolsCall) Fields(s ...googleapi.Field) *InstanceGroupManagersSetTargetPoolsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstanceGroupManagersSetTargetPoolsCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":              c.project,
		"zone":                 c.zone,
		"instanceGroupManager": c.instanceGroupManager,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.instancegroupmanagerssettargetpoolsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modifies the target pools to which all new instances in this group are assigned. Existing instances in the group are not affected.",
	//   "httpMethod": "POST",
	//   "id": "replicapool.instanceGroupManagers.setTargetPools",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "instanceGroupManager"
	//   ],
	//   "parameters": {
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the instance group manager resides.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setTargetPools",
	//   "request": {
	//     "$ref": "InstanceGroupManagersSetTargetPoolsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.zoneOperations.get":

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
	//   "id": "replicapool.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "replicapool.zoneOperations.list":

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

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ZoneOperationsListCall) Filter(filter string) *ZoneOperationsListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ZoneOperationsListCall) MaxResults(maxResults int64) *ZoneOperationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
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
	//   "id": "replicapool.zoneOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

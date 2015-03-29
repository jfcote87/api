// Package replicapoolupdater provides access to the Google Compute Engine Instance Group Updater API.
//
// See https://cloud.google.com/compute/docs/instance-groups/manager/#applying_rolling_updates_using_the_updater_service
//
// Usage example:
//
//   import "github.com/jfcote87/api/replicapoolupdater/v1beta1"
//   ...
//   replicapoolupdaterService, err := replicapoolupdater.New(oauthHttpClient)
package replicapoolupdater

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

const apiId = "replicapoolupdater:v1beta1"
const apiName = "replicapoolupdater"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/replicapoolupdater/v1beta1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage replica pools
	ReplicapoolScope = "https://www.googleapis.com/auth/replicapool"

	// View replica pools
	ReplicapoolReadonlyScope = "https://www.googleapis.com/auth/replicapool.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.RollingUpdates = NewRollingUpdatesService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	RollingUpdates *RollingUpdatesService

	ZoneOperations *ZoneOperationsService
}

func NewRollingUpdatesService(s *Service) *RollingUpdatesService {
	rs := &RollingUpdatesService{s: s}
	return rs
}

type RollingUpdatesService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

type InstanceUpdate struct {
	// Error: Errors that occurred during the instance update.
	Error *InstanceUpdateError `json:"error,omitempty"`

	// Instance: URL of the instance being updated.
	Instance string `json:"instance,omitempty"`

	// Status: Status of the instance update. Possible values are:
	// -
	// "PENDING": The instance update is pending execution.
	// -
	// "ROLLING_FORWARD": The instance update is going forward.
	// -
	// "ROLLING_BACK": The instance update is being rolled back.
	// -
	// "PAUSED": The instance update is temporarily paused (inactive).
	// -
	// "ROLLED_OUT": The instance update is finished, the instance is
	// running the new template.
	// - "ROLLED_BACK": The instance update is
	// finished, the instance has been reverted to the previous template.
	// -
	// "CANCELLED": The instance update is paused and no longer can be
	// resumed, undefined in which template the instance is running.
	Status string `json:"status,omitempty"`
}

type InstanceUpdateError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*InstanceUpdateErrorErrors `json:"errors,omitempty"`
}

type InstanceUpdateErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request which
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type InstanceUpdateList struct {
	// Items: Collection of requested instance updates.
	Items []*InstanceUpdate `json:"items,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type Operation struct {
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format (output only).
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors occurred during processing of this
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested.
	// This is in RFC 3339 format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always kind#operation for
	// Operation resources.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource (output only).
	Name string `json:"name,omitempty"`

	OperationType string `json:"operationType,omitempty"`

	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] URL of the region where the operation resides
	// (output only).
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server. This is in RFC 3339 format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] Status of the operation. Can be one of the
	// following: "PENDING", "RUNNING", or "DONE".
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] Unique target id which identifies a
	// particular incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output Only] URL of the resource the operation is
	// mutating (output only).
	TargetLink string `json:"targetLink,omitempty"`

	User string `json:"user,omitempty"`

	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] URL of the zone where the operation resides
	// (output only).
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

type RollingUpdate struct {
	// ActionType: Specifies the action to take for each instance within the
	// instance group. This can be RECREATE which will recreate each
	// instance and is only available for managed instance groups. It can
	// also be REBOOT which performs a soft reboot for each instance and is
	// only available for regular (non-managed) instance groups and explicit
	// lists of instances.
	ActionType string `json:"actionType,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Error: [Output Only] Errors that occurred during rolling update.
	Error *RollingUpdateError `json:"error,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// InstanceGroup: Fully-qualified URL of an instance group being
	// updated. Exactly one of instanceGroupManager, instanceGroup and
	// instance list must be set.
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// InstanceGroupManager: Fully-qualified URL of an instance group
	// manager being updated. Exactly one of instanceGroupManager,
	// instanceGroup and instance list must be set.
	InstanceGroupManager string `json:"instanceGroupManager,omitempty"`

	// InstanceTemplate: Fully-qualified URL of an instance template to
	// apply.
	InstanceTemplate string `json:"instanceTemplate,omitempty"`

	// Instances: List of fully-qualified URLs of instances to be updated.
	// Exactly one of instanceGroupManager, instanceGroup and instance list
	// must be set.
	Instances []string `json:"instances,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// Policy: Parameters of the update process.
	Policy *RollingUpdatePolicy `json:"policy,omitempty"`

	// Progress: [Output Only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess at
	// when the update will be complete. This number should be monotonically
	// increasing as the update progresses.
	Progress int64 `json:"progress,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] Status of the update. Possible values are:
	// -
	// "ROLLING_FORWARD": The update is going forward.
	// - "ROLLING_BACK":
	// The update is being rolled back.
	// - "PAUSED": The update is
	// temporarily paused (inactive).
	// - "ROLLED_OUT": The update is
	// finished, all instances have been updated successfully.
	// -
	// "ROLLED_BACK": The update is finished, all instances have been
	// reverted to the previous template.
	// - "CANCELLED": The update is
	// paused and no longer can be resumed, undefined how many instances are
	// running in which template.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the update.
	StatusMessage string `json:"statusMessage,omitempty"`

	// User: [Output Only] User who requested the update, for example:
	// user@example.com.
	User string `json:"user,omitempty"`
}

type RollingUpdateError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*RollingUpdateErrorErrors `json:"errors,omitempty"`
}

type RollingUpdateErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request which
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`
}

type RollingUpdatePolicy struct {
	// AutoPauseAfterInstances: Number of instances updated before the
	// update gets automatically paused.
	AutoPauseAfterInstances int64 `json:"autoPauseAfterInstances,omitempty"`

	// InstanceStartupTimeoutSec: Maximum amount of time we will wait after
	// finishing all steps until we receive HEALTHY state for instance. If
	// this deadline is exceeded instance update is considered as failed.
	InstanceStartupTimeoutSec int64 `json:"instanceStartupTimeoutSec,omitempty"`

	// MaxNumConcurrentInstances: Maximum number of instances that can be
	// updated simultaneously (concurrently). An update of an instance
	// starts when the instance is about to be restarted and finishes after
	// the instance has been restarted and the sleep period (defined by
	// sleepAfterInstanceRestartSec) has passed.
	MaxNumConcurrentInstances int64 `json:"maxNumConcurrentInstances,omitempty"`

	// MaxNumFailedInstances: Maximum number of instance updates that can
	// fail without failing the group update. Instance update is considered
	// failed if any of it's update actions (e.g. Stop call on Instance
	// resource in Rolling Reboot) failed with permanent failure, or if
	// after finishing all update actions this instance is in UNHEALTHY
	// state.
	MaxNumFailedInstances int64 `json:"maxNumFailedInstances,omitempty"`

	// MinInstanceUpdateTimeSec: Specifies minimum amount of time we will
	// spend on updating single instance, measuring at the start of the
	// first update action (e.g. Recreate call on Instance Group Manager or
	// Stop call on Instance resource). If actual instance update takes less
	// time we will simply sleep before proceeding with next instance.
	MinInstanceUpdateTimeSec int64 `json:"minInstanceUpdateTimeSec,omitempty"`

	// SleepAfterInstanceRestartSec: Time period after the instance has been
	// restarted but before marking the update of this instance as done.
	// This field is deprecated and ignored by Rolling Updater.
	SleepAfterInstanceRestartSec int64 `json:"sleepAfterInstanceRestartSec,omitempty"`
}

type RollingUpdateList struct {
	// Items: Collection of requested updates.
	Items []*RollingUpdate `json:"items,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "replicapoolupdater.rollingUpdates.cancel":

type RollingUpdatesCancelCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Cancel: Cancels an update. The update must be PAUSED before it can be
// cancelled. This has no effect if the update is already CANCELLED.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#cancelrollingupdate
func (r *RollingUpdatesService) Cancel(project string, zone string, rollingUpdate string) *RollingUpdatesCancelCall {
	return &RollingUpdatesCancelCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/cancel",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesCancelCall) Fields(s ...googleapi.Field) *RollingUpdatesCancelCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesCancelCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Cancels an update. The update must be PAUSED before it can be cancelled. This has no effect if the update is already CANCELLED.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.cancel",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/cancel",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.get":

type RollingUpdatesGetCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns information about an update.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) Get(project string, zone string, rollingUpdate string) *RollingUpdatesGetCall {
	return &RollingUpdatesGetCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesGetCall) Fields(s ...googleapi.Field) *RollingUpdatesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesGetCall) Do() (*RollingUpdate, error) {
	var returnValue *RollingUpdate
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns information about an update.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}",
	//   "response": {
	//     "$ref": "RollingUpdate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.insert":

type RollingUpdatesInsertCall struct {
	s             *Service
	project       string
	zone          string
	rollingupdate *RollingUpdate
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts and starts a new update.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#starting_an_update
func (r *RollingUpdatesService) Insert(project string, zone string, rollingupdate *RollingUpdate) *RollingUpdatesInsertCall {
	return &RollingUpdatesInsertCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingupdate: rollingupdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesInsertCall) Fields(s ...googleapi.Field) *RollingUpdatesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesInsertCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.rollingupdate,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts and starts a new update.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates",
	//   "request": {
	//     "$ref": "RollingUpdate"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.list":

type RollingUpdatesListCall struct {
	s             *Service
	project       string
	zone          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists recent updates for a given managed instance group, in
// reverse chronological order and paginated format.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) List(project string, zone string) *RollingUpdatesListCall {
	return &RollingUpdatesListCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates",
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RollingUpdatesListCall) Filter(filter string) *RollingUpdatesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// InstanceGroupManager sets the optional parameter
// "instanceGroupManager": The name of the instance group manager used
// for filtering.
func (c *RollingUpdatesListCall) InstanceGroupManager(instanceGroupManager string) *RollingUpdatesListCall {
	c.params_.Set("instanceGroupManager", fmt.Sprintf("%v", instanceGroupManager))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RollingUpdatesListCall) MaxResults(maxResults int64) *RollingUpdatesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RollingUpdatesListCall) PageToken(pageToken string) *RollingUpdatesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesListCall) Fields(s ...googleapi.Field) *RollingUpdatesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesListCall) Do() (*RollingUpdateList, error) {
	var returnValue *RollingUpdateList
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
	//   "description": "Lists recent updates for a given managed instance group, in reverse chronological order and paginated format.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.list",
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
	//     "instanceGroupManager": {
	//       "description": "The name of the instance group manager used for filtering.",
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
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates",
	//   "response": {
	//     "$ref": "RollingUpdateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.listInstanceUpdates":

type RollingUpdatesListInstanceUpdatesCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ListInstanceUpdates: Lists the current status for each instance
// within a given update.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) ListInstanceUpdates(project string, zone string, rollingUpdate string) *RollingUpdatesListInstanceUpdatesCall {
	return &RollingUpdatesListInstanceUpdatesCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/instanceUpdates",
	}
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RollingUpdatesListInstanceUpdatesCall) Filter(filter string) *RollingUpdatesListInstanceUpdatesCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RollingUpdatesListInstanceUpdatesCall) MaxResults(maxResults int64) *RollingUpdatesListInstanceUpdatesCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RollingUpdatesListInstanceUpdatesCall) PageToken(pageToken string) *RollingUpdatesListInstanceUpdatesCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesListInstanceUpdatesCall) Fields(s ...googleapi.Field) *RollingUpdatesListInstanceUpdatesCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesListInstanceUpdatesCall) Do() (*InstanceUpdateList, error) {
	var returnValue *InstanceUpdateList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists the current status for each instance within a given update.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.listInstanceUpdates",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
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
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/instanceUpdates",
	//   "response": {
	//     "$ref": "InstanceUpdateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.pause":

type RollingUpdatesPauseCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Pause: Pauses the update in state from ROLLING_FORWARD or
// ROLLING_BACK. Has no effect if invoked when the state of the update
// is PAUSED.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#pausing_a_rolling_update
func (r *RollingUpdatesService) Pause(project string, zone string, rollingUpdate string) *RollingUpdatesPauseCall {
	return &RollingUpdatesPauseCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/pause",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesPauseCall) Fields(s ...googleapi.Field) *RollingUpdatesPauseCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesPauseCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Pauses the update in state from ROLLING_FORWARD or ROLLING_BACK. Has no effect if invoked when the state of the update is PAUSED.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.pause",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/pause",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.resume":

type RollingUpdatesResumeCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Resume: Continues an update in PAUSED state. Has no effect if invoked
// when the state of the update is ROLLED_OUT.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#continuerollingupdate
func (r *RollingUpdatesService) Resume(project string, zone string, rollingUpdate string) *RollingUpdatesResumeCall {
	return &RollingUpdatesResumeCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/resume",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesResumeCall) Fields(s ...googleapi.Field) *RollingUpdatesResumeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesResumeCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Continues an update in PAUSED state. Has no effect if invoked when the state of the update is ROLLED_OUT.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.resume",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/resume",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.rollback":

type RollingUpdatesRollbackCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Rollback: Rolls back the update in state from ROLLING_FORWARD or
// PAUSED. Has no effect if invoked when the state of the update is
// ROLLED_BACK.

// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#rollingbackupdate
func (r *RollingUpdatesService) Rollback(project string, zone string, rollingUpdate string) *RollingUpdatesRollbackCall {
	return &RollingUpdatesRollbackCall{
		s:             r.s,
		project:       project,
		zone:          zone,
		rollingUpdate: rollingUpdate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/rollback",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesRollbackCall) Fields(s ...googleapi.Field) *RollingUpdatesRollbackCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RollingUpdatesRollbackCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Rolls back the update in state from ROLLING_FORWARD or PAUSED. Has no effect if invoked when the state of the update is ROLLED_BACK.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.rollback",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/rollback",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.zoneOperations.get":

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
	//   "id": "replicapoolupdater.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
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
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

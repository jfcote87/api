// Package dns provides access to the Google Cloud DNS API.
//
// See https://developers.google.com/cloud-dns
//
// Usage example:
//
//   import "github.com/jfcote87/api/dns/v1beta1"
//   ...
//   dnsService, err := dns.New(oauthHttpClient)
package dns

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

const apiId = "dns:v1beta1"
const apiName = "dns"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/dns/v1beta1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your DNS records hosted by Google Cloud DNS
	NdevClouddnsReadonlyScope = "https://www.googleapis.com/auth/ndev.clouddns.readonly"

	// View and manage your DNS records hosted by Google Cloud DNS
	NdevClouddnsReadwriteScope = "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Changes = NewChangesService(s)
	s.ManagedZones = NewManagedZonesService(s)
	s.Projects = NewProjectsService(s)
	s.ResourceRecordSets = NewResourceRecordSetsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Changes *ChangesService

	ManagedZones *ManagedZonesService

	Projects *ProjectsService

	ResourceRecordSets *ResourceRecordSetsService
}

func NewChangesService(s *Service) *ChangesService {
	rs := &ChangesService{s: s}
	return rs
}

type ChangesService struct {
	s *Service
}

func NewManagedZonesService(s *Service) *ManagedZonesService {
	rs := &ManagedZonesService{s: s}
	return rs
}

type ManagedZonesService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

func NewResourceRecordSetsService(s *Service) *ResourceRecordSetsService {
	rs := &ResourceRecordSetsService{s: s}
	return rs
}

type ResourceRecordSetsService struct {
	s *Service
}

type Change struct {
	// Additions: Which ResourceRecordSets to add?
	Additions []*ResourceRecordSet `json:"additions,omitempty"`

	// Deletions: Which ResourceRecordSets to remove? Must match existing
	// data exactly.
	Deletions []*ResourceRecordSet `json:"deletions,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only).
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dns#change".
	Kind string `json:"kind,omitempty"`

	// StartTime: The time that this operation was started by the server.
	// This is in RFC3339 text format.
	StartTime string `json:"startTime,omitempty"`

	// Status: Status of the operation. Can be one of the following:
	// "PENDING" or "DONE" (output only).
	Status string `json:"status,omitempty"`
}

type ChangesListResponse struct {
	// Changes: The requested changes.
	Changes []*Change `json:"changes,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The presence of this field indicates that there exist
	// more results following your last page of results in pagination order.
	// To fetch them, make another list request using this value as your
	// pagination token.
	//
	// In this way you can retrieve the complete contents
	// of even very large collections one page at a time. However, if the
	// contents of the collection change between the first and last
	// paginated list request, the set of all elements returned will be an
	// inconsistent view of the collection. There is no way to retrieve a
	// "snapshot" of collections larger than the maximum page size.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ManagedZone struct {
	// CreationTime: The time that this resource was created on the server.
	// This is in RFC3339 text format. Output only.
	CreationTime string `json:"creationTime,omitempty"`

	// Description: A mutable string of at most 1024 characters associated
	// with this resource for the user's convenience. Has no effect on the
	// managed zone's function.
	Description string `json:"description,omitempty"`

	// DnsName: The DNS name of this managed zone, for instance
	// "example.com.".
	DnsName string `json:"dnsName,omitempty"`

	// Id: Unique identifier for the resource; defined by the server (output
	// only)
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dns#managedZone".
	Kind string `json:"kind,omitempty"`

	// Name: User assigned name for this resource. Must be unique within the
	// project. The name must be 1-32 characters long, must begin with a
	// letter, end with a letter or digit, and only contain lowercase
	// letters, digits or dashes.
	Name string `json:"name,omitempty"`

	// NameServerSet: Optionally specifies the NameServerSet for this
	// ManagedZone. A NameServerSet is a set of DNS name servers that all
	// host the same ManagedZones. Most users will leave this field unset.
	NameServerSet string `json:"nameServerSet,omitempty"`

	// NameServers: Delegate your managed_zone to these virtual name
	// servers; defined by the server (output only)
	NameServers []string `json:"nameServers,omitempty"`
}

type ManagedZonesListResponse struct {
	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// ManagedZones: The managed zone resources.
	ManagedZones []*ManagedZone `json:"managedZones,omitempty"`

	// NextPageToken: The presence of this field indicates that there exist
	// more results following your last page of results in pagination order.
	// To fetch them, make another list request using this value as your
	// page token.
	//
	// In this way you can retrieve the complete contents of
	// even very large collections one page at a time. However, if the
	// contents of the collection change between the first and last
	// paginated list request, the set of all elements returned will be an
	// inconsistent view of the collection. There is no way to retrieve a
	// consistent snapshot of a collection larger than the maximum page
	// size.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Project struct {
	// Id: User assigned unique identifier for the resource (output only).
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dns#project".
	Kind string `json:"kind,omitempty"`

	// Number: Unique numeric identifier for the resource; defined by the
	// server (output only).
	Number uint64 `json:"number,omitempty,string"`

	// Quota: Quotas assigned to this project (output only).
	Quota *Quota `json:"quota,omitempty"`
}

type Quota struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dns#quota".
	Kind string `json:"kind,omitempty"`

	// ManagedZones: Maximum allowed number of managed zones in the project.
	ManagedZones int64 `json:"managedZones,omitempty"`

	// ResourceRecordsPerRrset: Maximum allowed number of ResourceRecords
	// per ResourceRecordSet.
	ResourceRecordsPerRrset int64 `json:"resourceRecordsPerRrset,omitempty"`

	// RrsetAdditionsPerChange: Maximum allowed number of ResourceRecordSets
	// to add per ChangesCreateRequest.
	RrsetAdditionsPerChange int64 `json:"rrsetAdditionsPerChange,omitempty"`

	// RrsetDeletionsPerChange: Maximum allowed number of ResourceRecordSets
	// to delete per ChangesCreateRequest.
	RrsetDeletionsPerChange int64 `json:"rrsetDeletionsPerChange,omitempty"`

	// RrsetsPerManagedZone: Maximum allowed number of ResourceRecordSets
	// per zone in the project.
	RrsetsPerManagedZone int64 `json:"rrsetsPerManagedZone,omitempty"`

	// TotalRrdataSizePerChange: Maximum allowed size for total rrdata in
	// one ChangesCreateRequest in bytes.
	TotalRrdataSizePerChange int64 `json:"totalRrdataSizePerChange,omitempty"`
}

type ResourceRecordSet struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dns#resourceRecordSet".
	Kind string `json:"kind,omitempty"`

	// Name: For example, www.example.com.
	Name string `json:"name,omitempty"`

	// Rrdatas: As defined in RFC 1035 (section 5) and RFC 1034 (section
	// 3.6.1)
	Rrdatas []string `json:"rrdatas,omitempty"`

	// Ttl: Number of seconds that this ResourceRecordSet can be cached by
	// resolvers.
	Ttl int64 `json:"ttl,omitempty"`

	// Type: One of A, AAAA, SOA, MX, NS, TXT
	Type string `json:"type,omitempty"`
}

type ResourceRecordSetsListResponse struct {
	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The presence of this field indicates that there exist
	// more results following your last page of results in pagination order.
	// To fetch them, make another list request using this value as your
	// pagination token.
	//
	// In this way you can retrieve the complete contents
	// of even very large collections one page at a time. However, if the
	// contents of the collection change between the first and last
	// paginated list request, the set of all elements returned will be an
	// inconsistent view of the collection. There is no way to retrieve a
	// consistent snapshot of a collection larger than the maximum page
	// size.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Rrsets: The resource record set resources.
	Rrsets []*ResourceRecordSet `json:"rrsets,omitempty"`
}

// method id "dns.changes.create":

type ChangesCreateCall struct {
	s             *Service
	project       string
	managedZone   string
	change        *Change
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Atomically update the ResourceRecordSet collection.

func (r *ChangesService) Create(project string, managedZone string, change *Change) *ChangesCreateCall {
	return &ChangesCreateCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		change:        change,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}/changes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChangesCreateCall) Fields(s ...googleapi.Field) *ChangesCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChangesCreateCall) Do() (*Change, error) {
	var returnValue *Change
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.change,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Atomically update the ResourceRecordSet collection.",
	//   "httpMethod": "POST",
	//   "id": "dns.changes.create",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone"
	//   ],
	//   "parameters": {
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}/changes",
	//   "request": {
	//     "$ref": "Change"
	//   },
	//   "response": {
	//     "$ref": "Change"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.changes.get":

type ChangesGetCall struct {
	s             *Service
	project       string
	managedZone   string
	changeId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Fetch the representation of an existing Change.

func (r *ChangesService) Get(project string, managedZone string, changeId string) *ChangesGetCall {
	return &ChangesGetCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		changeId:      changeId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}/changes/{changeId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChangesGetCall) Fields(s ...googleapi.Field) *ChangesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChangesGetCall) Do() (*Change, error) {
	var returnValue *Change
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
		"changeId":    c.changeId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Fetch the representation of an existing Change.",
	//   "httpMethod": "GET",
	//   "id": "dns.changes.get",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone",
	//     "changeId"
	//   ],
	//   "parameters": {
	//     "changeId": {
	//       "description": "The identifier of the requested change, from a previous ResourceRecordSetsChangeResponse.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}/changes/{changeId}",
	//   "response": {
	//     "$ref": "Change"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.changes.list":

type ChangesListCall struct {
	s             *Service
	project       string
	managedZone   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Enumerate Changes to a ResourceRecordSet collection.

func (r *ChangesService) List(project string, managedZone string) *ChangesListCall {
	return &ChangesListCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}/changes",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to be returned. If unspecified, the server will decide how
// many results to return.
func (c *ChangesListCall) MaxResults(maxResults int64) *ChangesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A tag returned by
// a previous list request that was truncated. Use this parameter to
// continue a previous list request.
func (c *ChangesListCall) PageToken(pageToken string) *ChangesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SortBy sets the optional parameter "sortBy": Sorting criterion. The
// only supported value is change sequence.
func (c *ChangesListCall) SortBy(sortBy string) *ChangesListCall {
	c.params_.Set("sortBy", fmt.Sprintf("%v", sortBy))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Sorting order
// direction: 'ascending' or 'descending'.
func (c *ChangesListCall) SortOrder(sortOrder string) *ChangesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChangesListCall) Fields(s ...googleapi.Field) *ChangesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChangesListCall) Do() (*ChangesListResponse, error) {
	var returnValue *ChangesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Enumerate Changes to a ResourceRecordSet collection.",
	//   "httpMethod": "GET",
	//   "id": "dns.changes.list",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone"
	//   ],
	//   "parameters": {
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Optional. Maximum number of results to be returned. If unspecified, the server will decide how many results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A tag returned by a previous list request that was truncated. Use this parameter to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sortBy": {
	//       "default": "changeSequence",
	//       "description": "Sorting criterion. The only supported value is change sequence.",
	//       "enum": [
	//         "changeSequence"
	//       ],
	//       "enumDescriptions": [
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Sorting order direction: 'ascending' or 'descending'.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}/changes",
	//   "response": {
	//     "$ref": "ChangesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.managedZones.create":

type ManagedZonesCreateCall struct {
	s             *Service
	project       string
	managedzone   *ManagedZone
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Create a new ManagedZone.

func (r *ManagedZonesService) Create(project string, managedzone *ManagedZone) *ManagedZonesCreateCall {
	return &ManagedZonesCreateCall{
		s:             r.s,
		project:       project,
		managedzone:   managedzone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManagedZonesCreateCall) Fields(s ...googleapi.Field) *ManagedZonesCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ManagedZonesCreateCall) Do() (*ManagedZone, error) {
	var returnValue *ManagedZone
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.managedzone,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create a new ManagedZone.",
	//   "httpMethod": "POST",
	//   "id": "dns.managedZones.create",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones",
	//   "request": {
	//     "$ref": "ManagedZone"
	//   },
	//   "response": {
	//     "$ref": "ManagedZone"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.managedZones.delete":

type ManagedZonesDeleteCall struct {
	s             *Service
	project       string
	managedZone   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete a previously created ManagedZone.

func (r *ManagedZonesService) Delete(project string, managedZone string) *ManagedZonesDeleteCall {
	return &ManagedZonesDeleteCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}",
	}
}

func (c *ManagedZonesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a previously created ManagedZone.",
	//   "httpMethod": "DELETE",
	//   "id": "dns.managedZones.delete",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone"
	//   ],
	//   "parameters": {
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.managedZones.get":

type ManagedZonesGetCall struct {
	s             *Service
	project       string
	managedZone   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Fetch the representation of an existing ManagedZone.

func (r *ManagedZonesService) Get(project string, managedZone string) *ManagedZonesGetCall {
	return &ManagedZonesGetCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManagedZonesGetCall) Fields(s ...googleapi.Field) *ManagedZonesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ManagedZonesGetCall) Do() (*ManagedZone, error) {
	var returnValue *ManagedZone
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Fetch the representation of an existing ManagedZone.",
	//   "httpMethod": "GET",
	//   "id": "dns.managedZones.get",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone"
	//   ],
	//   "parameters": {
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}",
	//   "response": {
	//     "$ref": "ManagedZone"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.managedZones.list":

type ManagedZonesListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Enumerate ManagedZones that have been created but not yet
// deleted.

func (r *ManagedZonesService) List(project string) *ManagedZonesListCall {
	return &ManagedZonesListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to be returned. If unspecified, the server will decide how
// many results to return.
func (c *ManagedZonesListCall) MaxResults(maxResults int64) *ManagedZonesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A tag returned by
// a previous list request that was truncated. Use this parameter to
// continue a previous list request.
func (c *ManagedZonesListCall) PageToken(pageToken string) *ManagedZonesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ManagedZonesListCall) Fields(s ...googleapi.Field) *ManagedZonesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ManagedZonesListCall) Do() (*ManagedZonesListResponse, error) {
	var returnValue *ManagedZonesListResponse
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
	//   "description": "Enumerate ManagedZones that have been created but not yet deleted.",
	//   "httpMethod": "GET",
	//   "id": "dns.managedZones.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Optional. Maximum number of results to be returned. If unspecified, the server will decide how many results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A tag returned by a previous list request that was truncated. Use this parameter to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones",
	//   "response": {
	//     "$ref": "ManagedZonesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.projects.get":

type ProjectsGetCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Fetch the representation of an existing Project.

func (r *ProjectsService) Get(project string) *ProjectsGetCall {
	return &ProjectsGetCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsGetCall) Do() (*Project, error) {
	var returnValue *Project
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
	//   "description": "Fetch the representation of an existing Project.",
	//   "httpMethod": "GET",
	//   "id": "dns.projects.get",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

// method id "dns.resourceRecordSets.list":

type ResourceRecordSetsListCall struct {
	s             *Service
	project       string
	managedZone   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Enumerate ResourceRecordSets that have been created but not yet
// deleted.

func (r *ResourceRecordSetsService) List(project string, managedZone string) *ResourceRecordSetsListCall {
	return &ResourceRecordSetsListCall{
		s:             r.s,
		project:       project,
		managedZone:   managedZone,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/managedZones/{managedZone}/rrsets",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to be returned. If unspecified, the server will decide how
// many results to return.
func (c *ResourceRecordSetsListCall) MaxResults(maxResults int64) *ResourceRecordSetsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// Name sets the optional parameter "name": Restricts the list to return
// only records with this fully qualified domain name.
func (c *ResourceRecordSetsListCall) Name(name string) *ResourceRecordSetsListCall {
	c.params_.Set("name", fmt.Sprintf("%v", name))
	return c
}

// PageToken sets the optional parameter "pageToken": A tag returned by
// a previous list request that was truncated. Use this parameter to
// continue a previous list request.
func (c *ResourceRecordSetsListCall) PageToken(pageToken string) *ResourceRecordSetsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Type sets the optional parameter "type": Restricts the list to return
// only records of this type. If present, the "name" parameter must also
// be present.
func (c *ResourceRecordSetsListCall) Type(type_ string) *ResourceRecordSetsListCall {
	c.params_.Set("type", fmt.Sprintf("%v", type_))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourceRecordSetsListCall) Fields(s ...googleapi.Field) *ResourceRecordSetsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ResourceRecordSetsListCall) Do() (*ResourceRecordSetsListResponse, error) {
	var returnValue *ResourceRecordSetsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project":     c.project,
		"managedZone": c.managedZone,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Enumerate ResourceRecordSets that have been created but not yet deleted.",
	//   "httpMethod": "GET",
	//   "id": "dns.resourceRecordSets.list",
	//   "parameterOrder": [
	//     "project",
	//     "managedZone"
	//   ],
	//   "parameters": {
	//     "managedZone": {
	//       "description": "Identifies the managed zone addressed by this request. Can be the managed zone name or id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Optional. Maximum number of results to be returned. If unspecified, the server will decide how many results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "name": {
	//       "description": "Restricts the list to return only records with this fully qualified domain name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A tag returned by a previous list request that was truncated. Use this parameter to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Identifies the project addressed by this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Restricts the list to return only records of this type. If present, the \"name\" parameter must also be present.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/managedZones/{managedZone}/rrsets",
	//   "response": {
	//     "$ref": "ResourceRecordSetsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readonly",
	//     "https://www.googleapis.com/auth/ndev.clouddns.readwrite"
	//   ]
	// }

}

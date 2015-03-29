// Package container provides access to the Google Container Engine API.
//
// See https://cloud.google.com/container-engine/docs/v1beta1/
//
// Usage example:
//
//   import "github.com/jfcote87/api/container/v1beta1"
//   ...
//   containerService, err := container.New(oauthHttpClient)
package container

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

const apiId = "container:v1beta1"
const apiName = "container"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/container/v1beta1/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Clusters = NewProjectsClustersService(s)
	rs.Operations = NewProjectsOperationsService(s)
	rs.Zones = NewProjectsZonesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Clusters *ProjectsClustersService

	Operations *ProjectsOperationsService

	Zones *ProjectsZonesService
}

func NewProjectsClustersService(s *Service) *ProjectsClustersService {
	rs := &ProjectsClustersService{s: s}
	return rs
}

type ProjectsClustersService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

func NewProjectsZonesService(s *Service) *ProjectsZonesService {
	rs := &ProjectsZonesService{s: s}
	rs.Clusters = NewProjectsZonesClustersService(s)
	rs.Operations = NewProjectsZonesOperationsService(s)
	return rs
}

type ProjectsZonesService struct {
	s *Service

	Clusters *ProjectsZonesClustersService

	Operations *ProjectsZonesOperationsService
}

func NewProjectsZonesClustersService(s *Service) *ProjectsZonesClustersService {
	rs := &ProjectsZonesClustersService{s: s}
	return rs
}

type ProjectsZonesClustersService struct {
	s *Service
}

func NewProjectsZonesOperationsService(s *Service) *ProjectsZonesOperationsService {
	rs := &ProjectsZonesOperationsService{s: s}
	return rs
}

type ProjectsZonesOperationsService struct {
	s *Service
}

type Cluster struct {
	// ClusterApiVersion: The API version of the Kubernetes master and
	// kubelets running in this cluster. Leave blank to pick up the latest
	// stable release, or specify a version of the form "x.y.z". The Google
	// Container Engine release notes lists the currently supported
	// versions. If an incorrect version is specified, the server returns an
	// error listing the currently supported versions.
	ClusterApiVersion string `json:"clusterApiVersion,omitempty"`

	// ContainerIpv4Cidr: [Output only] The IP addresses of the container
	// pods in this cluster, in  CIDR notation (e.g. 1.2.3.4/29).
	ContainerIpv4Cidr string `json:"containerIpv4Cidr,omitempty"`

	// CreationTimestamp: [Output only] The time the cluster was created, in
	// RFC3339 text format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional description of this cluster.
	Description string `json:"description,omitempty"`

	// EnableCloudLogging: Whether logs from the cluster should be made
	// available via the Google Cloud Logging service. This includes both
	// logs from your applications running in the cluster as well as logs
	// from the Kubernetes components themselves.
	EnableCloudLogging bool `json:"enableCloudLogging,omitempty"`

	// Endpoint: [Output only] The IP address of this cluster's Kubernetes
	// master. The endpoint can be accessed from the internet at
	// https://username:password@endpoint/.
	//
	// See the masterAuth property of
	// this resource for username and password information.
	Endpoint string `json:"endpoint,omitempty"`

	// MasterAuth: The HTTP basic authentication information for accessing
	// the master. Because the master endpoint is open to the internet, you
	// should create a strong password.
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`

	// Name: The name of this cluster. The name must be unique within this
	// project and zone, and can be up to 40 characters with the following
	// restrictions:
	// - Lowercase letters, numbers, and hyphens only.
	// -
	// Must start with a letter.
	// - Must end with a number or a letter.
	Name string `json:"name,omitempty"`

	// Network: The name of the Google Compute Engine network to which the
	// cluster is connected.
	Network string `json:"network,omitempty"`

	// NodeConfig: The machine type and image to use for all nodes in this
	// cluster. See the descriptions of the child properties of nodeConfig.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// NodeRoutingPrefixSize: [Output only] The size of the address space on
	// each node for hosting containers.
	NodeRoutingPrefixSize int64 `json:"nodeRoutingPrefixSize,omitempty"`

	// NumNodes: The number of nodes to create in this cluster. You must
	// ensure that your Compute Engine resource quota is sufficient for this
	// number of instances plus one (to include the master). You must also
	// have available firewall and routes quota.
	NumNodes int64 `json:"numNodes,omitempty"`

	// SelfLink: [Output only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServicesIpv4Cidr: [Output only] The IP addresses of the Kubernetes
	// services in this cluster, in  CIDR notation (e.g. 1.2.3.4/29).
	// Service addresses are always in the 10.0.0.0/16 range.
	ServicesIpv4Cidr string `json:"servicesIpv4Cidr,omitempty"`

	// Status: [Output only] The current status of this cluster.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output only] Additional information about the current
	// status of this cluster, if available.
	StatusMessage string `json:"statusMessage,omitempty"`

	// Zone: [Output only] The name of the Google Compute Engine zone in
	// which the cluster resides.
	Zone string `json:"zone,omitempty"`
}

type CreateClusterRequest struct {
	// Cluster: A cluster resource.
	Cluster *Cluster `json:"cluster,omitempty"`
}

type ListAggregatedClustersResponse struct {
	// Clusters: A list of clusters in the project, across all zones.
	Clusters []*Cluster `json:"clusters,omitempty"`
}

type ListAggregatedOperationsResponse struct {
	// Operations: A list of operations in the project, across all zones.
	Operations []*Operation `json:"operations,omitempty"`
}

type ListClustersResponse struct {
	// Clusters: A list of clusters in the project in the specified zone.
	Clusters []*Cluster `json:"clusters,omitempty"`
}

type ListOperationsResponse struct {
	// Operations: A list of operations in the project in the specified
	// zone.
	Operations []*Operation `json:"operations,omitempty"`
}

type MasterAuth struct {
	// Password: The password to use when accessing the Kubernetes master
	// endpoint.
	Password string `json:"password,omitempty"`

	// User: The username to use when accessing the Kubernetes master
	// endpoint.
	User string `json:"user,omitempty"`
}

type NodeConfig struct {
	// MachineType: The name of a Google Compute Engine machine type (e.g.
	// n1-standard-1).
	//
	// If unspecified, the default machine type is
	// n1-standard-1.
	MachineType string `json:"machineType,omitempty"`

	// ServiceAccounts: The optional list of ServiceAccounts, each with
	// their specified scopes, to be made available on all of the node VMs.
	// In addition to the service accounts and scopes specified, the
	// "default" account will always be created with the following scopes to
	// ensure the correct functioning of the cluster:
	// -
	// https://www.googleapis.com/auth/compute,
	// -
	// https://www.googleapis.com/auth/devstorage.read_only
	ServiceAccounts []*ServiceAccount `json:"serviceAccounts,omitempty"`

	// SourceImage: The fully-specified name of a Google Compute Engine
	// image. For example:
	// https://www.googleapis.com/compute/v1/projects/debian-cloud/global/ima
	// ges/backports-debian-7-wheezy-vYYYYMMDD (where YYYMMDD is the version
	// date).
	//
	// If specifying an image, you are responsible for ensuring its
	// compatibility with the Debian 7 backports image. We recommend leaving
	// this field blank to accept the default backports-debian-7-wheezy
	// value.
	SourceImage string `json:"sourceImage,omitempty"`
}

type Operation struct {
	// ErrorMessage: If an error has occurred, a textual description of the
	// error.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Name: The server-assigned ID for this operation. If the operation is
	// fulfilled upfront, it may not have a resource name.
	Name string `json:"name,omitempty"`

	// OperationType: The operation type.
	OperationType string `json:"operationType,omitempty"`

	// SelfLink: Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The current status of the operation.
	Status string `json:"status,omitempty"`

	// Target: [Optional] The URL of the cluster resource that this
	// operation is associated with.
	Target string `json:"target,omitempty"`

	// TargetLink: Server-defined URL for the target of the operation.
	TargetLink string `json:"targetLink,omitempty"`

	// Zone: The name of the Google Compute Engine zone in which the
	// operation is taking place.
	Zone string `json:"zone,omitempty"`
}

type ServiceAccount struct {
	// Email: Email address of the service account.
	Email string `json:"email,omitempty"`

	// Scopes: The list of scopes to be made available for this service
	// account.
	Scopes []string `json:"scopes,omitempty"`
}

// method id "container.projects.clusters.list":

type ProjectsClustersListCall struct {
	s             *Service
	projectId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all clusters owned by a project across all zones.

func (r *ProjectsClustersService) List(projectId string) *ProjectsClustersListCall {
	return &ProjectsClustersListCall{
		s:             r.s,
		projectId:     projectId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/clusters",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersListCall) Fields(s ...googleapi.Field) *ProjectsClustersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsClustersListCall) Do() (*ListAggregatedClustersResponse, error) {
	var returnValue *ListAggregatedClustersResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all clusters owned by a project across all zones.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.clusters.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/clusters",
	//   "response": {
	//     "$ref": "ListAggregatedClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.operations.list":

type ProjectsOperationsListCall struct {
	s             *Service
	projectId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all operations in a project, across all zones.

func (r *ProjectsOperationsService) List(projectId string) *ProjectsOperationsListCall {
	return &ProjectsOperationsListCall{
		s:             r.s,
		projectId:     projectId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/operations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsOperationsListCall) Do() (*ListAggregatedOperationsResponse, error) {
	var returnValue *ListAggregatedOperationsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all operations in a project, across all zones.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.operations.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/operations",
	//   "response": {
	//     "$ref": "ListAggregatedOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.create":

type ProjectsZonesClustersCreateCall struct {
	s                    *Service
	projectId            string
	zoneId               string
	createclusterrequest *CreateClusterRequest
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// Create: Creates a cluster, consisting of the specified number and
// type of Google Compute Engine instances, plus a Kubernetes master
// instance.
//
// The cluster is created in the project's default
// network.
//
// A firewall is added that allows traffic into port 443 on
// the master, which enables HTTPS. A firewall and a route is added for
// each node to allow the containers on that node to communicate with
// all other instances in the cluster.
//
// Finally, a route named
// k8s-iproute-10-xx-0-0 is created to track that the cluster's
// 10.xx.0.0/16 CIDR has been assigned.

func (r *ProjectsZonesClustersService) Create(projectId string, zoneId string, createclusterrequest *CreateClusterRequest) *ProjectsZonesClustersCreateCall {
	return &ProjectsZonesClustersCreateCall{
		s:                    r.s,
		projectId:            projectId,
		zoneId:               zoneId,
		createclusterrequest: createclusterrequest,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "{projectId}/zones/{zoneId}/clusters",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesClustersCreateCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"zoneId":    c.zoneId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.createclusterrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a cluster, consisting of the specified number and type of Google Compute Engine instances, plus a Kubernetes master instance.\n\nThe cluster is created in the project's default network.\n\nA firewall is added that allows traffic into port 443 on the master, which enables HTTPS. A firewall and a route is added for each node to allow the containers on that node to communicate with all other instances in the cluster.\n\nFinally, a route named k8s-iproute-10-xx-0-0 is created to track that the cluster's 10.xx.0.0/16 CIDR has been assigned.",
	//   "httpMethod": "POST",
	//   "id": "container.projects.zones.clusters.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone in which the cluster resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/clusters",
	//   "request": {
	//     "$ref": "CreateClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.delete":

type ProjectsZonesClustersDeleteCall struct {
	s             *Service
	projectId     string
	zoneId        string
	clusterId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the cluster, including the Kubernetes master and all
// worker nodes.
//
// Firewalls and routes that were configured at cluster
// creation are also deleted.

func (r *ProjectsZonesClustersService) Delete(projectId string, zoneId string, clusterId string) *ProjectsZonesClustersDeleteCall {
	return &ProjectsZonesClustersDeleteCall{
		s:             r.s,
		projectId:     projectId,
		zoneId:        zoneId,
		clusterId:     clusterId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/zones/{zoneId}/clusters/{clusterId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesClustersDeleteCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"zoneId":    c.zoneId,
		"clusterId": c.clusterId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the cluster, including the Kubernetes master and all worker nodes.\n\nFirewalls and routes that were configured at cluster creation are also deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "container.projects.zones.clusters.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone in which the cluster resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/clusters/{clusterId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.get":

type ProjectsZonesClustersGetCall struct {
	s             *Service
	projectId     string
	zoneId        string
	clusterId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a specific cluster.

func (r *ProjectsZonesClustersService) Get(projectId string, zoneId string, clusterId string) *ProjectsZonesClustersGetCall {
	return &ProjectsZonesClustersGetCall{
		s:             r.s,
		projectId:     projectId,
		zoneId:        zoneId,
		clusterId:     clusterId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/zones/{zoneId}/clusters/{clusterId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersGetCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesClustersGetCall) Do() (*Cluster, error) {
	var returnValue *Cluster
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"zoneId":    c.zoneId,
		"clusterId": c.clusterId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a specific cluster.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone in which the cluster resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/clusters/{clusterId}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.list":

type ProjectsZonesClustersListCall struct {
	s             *Service
	projectId     string
	zoneId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all clusters owned by a project in the specified zone.

func (r *ProjectsZonesClustersService) List(projectId string, zoneId string) *ProjectsZonesClustersListCall {
	return &ProjectsZonesClustersListCall{
		s:             r.s,
		projectId:     projectId,
		zoneId:        zoneId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/zones/{zoneId}/clusters",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersListCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesClustersListCall) Do() (*ListClustersResponse, error) {
	var returnValue *ListClustersResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"zoneId":    c.zoneId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all clusters owned by a project in the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone in which the cluster resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.operations.get":

type ProjectsZonesOperationsGetCall struct {
	s             *Service
	projectId     string
	zoneId        string
	operationId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets the specified operation.

func (r *ProjectsZonesOperationsService) Get(projectId string, zoneId string, operationId string) *ProjectsZonesOperationsGetCall {
	return &ProjectsZonesOperationsGetCall{
		s:             r.s,
		projectId:     projectId,
		zoneId:        zoneId,
		operationId:   operationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/zones/{zoneId}/operations/{operationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesOperationsGetCall) Do() (*Operation, error) {
	var returnValue *Operation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId":   c.projectId,
		"zoneId":      c.zoneId,
		"operationId": c.operationId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the specified operation.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.operations.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId",
	//     "operationId"
	//   ],
	//   "parameters": {
	//     "operationId": {
	//       "description": "The server-assigned name of the operation.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone in which the operation resides. This is always the same zone as the cluster with which the operation is associated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/operations/{operationId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.operations.list":

type ProjectsZonesOperationsListCall struct {
	s             *Service
	projectId     string
	zoneId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all operations in a project in a specific zone.

func (r *ProjectsZonesOperationsService) List(projectId string, zoneId string) *ProjectsZonesOperationsListCall {
	return &ProjectsZonesOperationsListCall{
		s:             r.s,
		projectId:     projectId,
		zoneId:        zoneId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{projectId}/zones/{zoneId}/operations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsListCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsZonesOperationsListCall) Do() (*ListOperationsResponse, error) {
	var returnValue *ListOperationsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"zoneId":    c.zoneId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all operations in a project in a specific zone.",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.operations.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zoneId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console project ID or  project number.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zoneId": {
	//       "description": "The name of the Google Compute Engine zone to return operations for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/zones/{zoneId}/operations",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

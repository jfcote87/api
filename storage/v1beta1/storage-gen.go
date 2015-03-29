// Package storage provides access to the Cloud Storage API.
//
// See https://developers.google.com/storage/docs/json_api/
//
// Usage example:
//
//   import "github.com/jfcote87/api/storage/v1beta1"
//   ...
//   storageService, err := storage.New(oauthHttpClient)
package storage

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

const apiId = "storage:v1beta1"
const apiName = "storage"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/storage/v1beta1/"}

// OAuth2 scopes used by this API.
const (
	// Manage your data and permissions in Google Cloud Storage
	DevstorageFull_controlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageRead_onlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.BucketAccessControls = NewBucketAccessControlsService(s)
	s.Buckets = NewBucketsService(s)
	s.ObjectAccessControls = NewObjectAccessControlsService(s)
	s.Objects = NewObjectsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	BucketAccessControls *BucketAccessControlsService

	Buckets *BucketsService

	ObjectAccessControls *ObjectAccessControlsService

	Objects *ObjectsService
}

func NewBucketAccessControlsService(s *Service) *BucketAccessControlsService {
	rs := &BucketAccessControlsService{s: s}
	return rs
}

type BucketAccessControlsService struct {
	s *Service
}

func NewBucketsService(s *Service) *BucketsService {
	rs := &BucketsService{s: s}
	return rs
}

type BucketsService struct {
	s *Service
}

func NewObjectAccessControlsService(s *Service) *ObjectAccessControlsService {
	rs := &ObjectAccessControlsService{s: s}
	return rs
}

type ObjectAccessControlsService struct {
	s *Service
}

func NewObjectsService(s *Service) *ObjectsService {
	rs := &ObjectsService{s: s}
	return rs
}

type ObjectsService struct {
	s *Service
}

type Bucket struct {
	// Acl: Access controls on the bucket.
	Acl []*BucketAccessControl `json:"acl,omitempty"`

	// DefaultObjectAcl: Default access controls to apply to new objects
	// when no ACL is provided.
	DefaultObjectAcl []*ObjectAccessControl `json:"defaultObjectAcl,omitempty"`

	// Id: The name of the bucket.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For buckets, this is always
	// storage#bucket.
	Kind string `json:"kind,omitempty"`

	// Location: The location of the bucket. Object data for objects in the
	// bucket resides in physical storage in this location. Can be US or EU.
	// Defaults to US.
	Location string `json:"location,omitempty"`

	// Owner: The owner of the bucket. This will always be the project
	// team's owner group.
	Owner *BucketOwner `json:"owner,omitempty"`

	// ProjectId: The project the bucket belongs to.
	ProjectId uint64 `json:"projectId,omitempty,string"`

	// SelfLink: The URI of this bucket.
	SelfLink string `json:"selfLink,omitempty"`

	// TimeCreated: Creation time of the bucket in RFC 3339 format.
	TimeCreated string `json:"timeCreated,omitempty"`

	// Website: The bucket's website configuration.
	Website *BucketWebsite `json:"website,omitempty"`
}

type BucketOwner struct {
	// Entity: The entity, in the form group-groupId.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity.
	EntityId string `json:"entityId,omitempty"`
}

type BucketWebsite struct {
	// MainPageSuffix: Behaves as the bucket's directory index where missing
	// objects are treated as potential directories.
	MainPageSuffix string `json:"mainPageSuffix,omitempty"`

	// NotFoundPage: The custom object to return when a requested resource
	// is not found.
	NotFoundPage string `json:"notFoundPage,omitempty"`
}

type BucketAccessControl struct {
	// Bucket: The name of the bucket.
	Bucket string `json:"bucket,omitempty"`

	// Domain: The domain associated with the entity, if any.
	Domain string `json:"domain,omitempty"`

	// Email: The email address associated with the entity, if any.
	Email string `json:"email,omitempty"`

	// Entity: The entity holding the permission, in one of the following
	// forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	//
	// - domain-domain
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// -
	// The user liz@example.com would be user-liz@example.com.
	// - The group
	// example@googlegroups.com would be group-example@googlegroups.com.
	// -
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity, if any.
	EntityId string `json:"entityId,omitempty"`

	// Id: The ID of the access-control entry.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For bucket access control entries,
	// this is always storage#bucketAccessControl.
	Kind string `json:"kind,omitempty"`

	// Role: The access permission for the entity. Can be READER, WRITER, or
	// OWNER.
	Role string `json:"role,omitempty"`

	// SelfLink: The link to this access-control entry.
	SelfLink string `json:"selfLink,omitempty"`
}

type BucketAccessControls struct {
	// Items: The list of items.
	Items []*BucketAccessControl `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of bucket access control
	// entries, this is always storage#bucketAccessControls.
	Kind string `json:"kind,omitempty"`
}

type Buckets struct {
	// Items: The list of items.
	Items []*Bucket `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of buckets, this is always
	// storage#buckets.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Object struct {
	// Acl: Access controls on the object.
	Acl []*ObjectAccessControl `json:"acl,omitempty"`

	// Bucket: The bucket containing this object.
	Bucket string `json:"bucket,omitempty"`

	// CacheControl: Cache-Control directive for the object data.
	CacheControl string `json:"cacheControl,omitempty"`

	// ContentDisposition: Content-Disposition of the object data.
	ContentDisposition string `json:"contentDisposition,omitempty"`

	// ContentEncoding: Content-Encoding of the object data.
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// ContentLanguage: Content-Language of the object data.
	ContentLanguage string `json:"contentLanguage,omitempty"`

	// Id: The ID of the object.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For objects, this is always
	// storage#object.
	Kind string `json:"kind,omitempty"`

	// Media: Object media data. Provided on your behalf when uploading raw
	// media or multipart/related with an auxiliary media part.
	Media *ObjectMedia `json:"media,omitempty"`

	// Metadata: User-provided metadata, in key/value pairs.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name: The name of this object. Required if not specified by URL
	// parameter.
	Name string `json:"name,omitempty"`

	// Owner: The owner of the object. This will always be the uploader of
	// the object.
	Owner *ObjectOwner `json:"owner,omitempty"`

	// SelfLink: The link to this object.
	SelfLink string `json:"selfLink,omitempty"`
}

type ObjectMedia struct {
	// Algorithm: Hash algorithm used. Currently only MD5 is supported.
	// Required if a hash is provided.
	Algorithm string `json:"algorithm,omitempty"`

	// ContentType: Content-Type of the object data.
	ContentType string `json:"contentType,omitempty"`

	// Data: URL-safe Base64-encoded data. This property can be used to
	// insert objects under 64KB in size, and will only be returned in
	// response to the get method for objects so created. When this resource
	// is returned in response to the list method, this property is omitted.
	Data string `json:"data,omitempty"`

	// Hash: Hash of the data. Required if a hash algorithm is provided.
	Hash string `json:"hash,omitempty"`

	// Length: Content-Length of the data in bytes.
	Length uint64 `json:"length,omitempty,string"`

	// Link: Media download link.
	Link string `json:"link,omitempty"`

	// TimeCreated: Creation time of the data in RFC 3339 format.
	TimeCreated string `json:"timeCreated,omitempty"`
}

type ObjectOwner struct {
	// Entity: The entity, in the form user-userId.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity.
	EntityId string `json:"entityId,omitempty"`
}

type ObjectAccessControl struct {
	// Bucket: The name of the bucket.
	Bucket string `json:"bucket,omitempty"`

	// Domain: The domain associated with the entity, if any.
	Domain string `json:"domain,omitempty"`

	// Email: The email address associated with the entity, if any.
	Email string `json:"email,omitempty"`

	// Entity: The entity holding the permission, in one of the following
	// forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	//
	// - domain-domain
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// -
	// The user liz@example.com would be user-liz@example.com.
	// - The group
	// example@googlegroups.com would be group-example@googlegroups.com.
	// -
	// To refer to all members of the Google Apps for Business domain
	// example.com, the entity would be domain-example.com.
	Entity string `json:"entity,omitempty"`

	// EntityId: The ID for the entity, if any.
	EntityId string `json:"entityId,omitempty"`

	// Id: The ID of the access-control entry.
	Id string `json:"id,omitempty"`

	// Kind: The kind of item this is. For object access control entries,
	// this is always storage#objectAccessControl.
	Kind string `json:"kind,omitempty"`

	// Object: The name of the object.
	Object string `json:"object,omitempty"`

	// Role: The access permission for the entity. Can be READER or OWNER.
	Role string `json:"role,omitempty"`

	// SelfLink: The link to this access-control entry.
	SelfLink string `json:"selfLink,omitempty"`
}

type ObjectAccessControls struct {
	// Items: The list of items.
	Items []*ObjectAccessControl `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of object access control
	// entries, this is always storage#objectAccessControls.
	Kind string `json:"kind,omitempty"`
}

type Objects struct {
	// Items: The list of items.
	Items []*Object `json:"items,omitempty"`

	// Kind: The kind of item this is. For lists of objects, this is always
	// storage#objects.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Prefixes: The list of prefixes of objects matching-but-not-listed up
	// to and including the requested delimiter.
	Prefixes []string `json:"prefixes,omitempty"`
}

// method id "storage.bucketAccessControls.delete":

type BucketAccessControlsDeleteCall struct {
	s             *Service
	bucket        string
	entity        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the ACL entry for the specified entity on the
// specified bucket.

func (r *BucketAccessControlsService) Delete(bucket string, entity string) *BucketAccessControlsDeleteCall {
	return &BucketAccessControlsDeleteCall{
		s:             r.s,
		bucket:        bucket,
		entity:        entity,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/acl/{entity}",
	}
}

func (c *BucketAccessControlsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.bucketAccessControls.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.get":

type BucketAccessControlsGetCall struct {
	s             *Service
	bucket        string
	entity        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns the ACL entry for the specified entity on the specified
// bucket.

func (r *BucketAccessControlsService) Get(bucket string, entity string) *BucketAccessControlsGetCall {
	return &BucketAccessControlsGetCall{
		s:             r.s,
		bucket:        bucket,
		entity:        entity,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsGetCall) Fields(s ...googleapi.Field) *BucketAccessControlsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketAccessControlsGetCall) Do() (*BucketAccessControl, error) {
	var returnValue *BucketAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the ACL entry for the specified entity on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.bucketAccessControls.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.insert":

type BucketAccessControlsInsertCall struct {
	s                   *Service
	bucket              string
	bucketaccesscontrol *BucketAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Insert: Creates a new ACL entry on the specified bucket.

func (r *BucketAccessControlsService) Insert(bucket string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsInsertCall {
	return &BucketAccessControlsInsertCall{
		s:                   r.s,
		bucket:              bucket,
		bucketaccesscontrol: bucketaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/acl",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsInsertCall) Fields(s ...googleapi.Field) *BucketAccessControlsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketAccessControlsInsertCall) Do() (*BucketAccessControl, error) {
	var returnValue *BucketAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucketaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new ACL entry on the specified bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.bucketAccessControls.insert",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.list":

type BucketAccessControlsListCall struct {
	s             *Service
	bucket        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves ACL entries on the specified bucket.

func (r *BucketAccessControlsService) List(bucket string) *BucketAccessControlsListCall {
	return &BucketAccessControlsListCall{
		s:             r.s,
		bucket:        bucket,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/acl",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsListCall) Fields(s ...googleapi.Field) *BucketAccessControlsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketAccessControlsListCall) Do() (*BucketAccessControls, error) {
	var returnValue *BucketAccessControls
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves ACL entries on the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.bucketAccessControls.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl",
	//   "response": {
	//     "$ref": "BucketAccessControls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.patch":

type BucketAccessControlsPatchCall struct {
	s                   *Service
	bucket              string
	entity              string
	bucketaccesscontrol *BucketAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Patch: Updates an ACL entry on the specified bucket. This method
// supports patch semantics.

func (r *BucketAccessControlsService) Patch(bucket string, entity string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsPatchCall {
	return &BucketAccessControlsPatchCall{
		s:                   r.s,
		bucket:              bucket,
		entity:              entity,
		bucketaccesscontrol: bucketaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsPatchCall) Fields(s ...googleapi.Field) *BucketAccessControlsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketAccessControlsPatchCall) Do() (*BucketAccessControl, error) {
	var returnValue *BucketAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucketaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an ACL entry on the specified bucket. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.bucketAccessControls.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.bucketAccessControls.update":

type BucketAccessControlsUpdateCall struct {
	s                   *Service
	bucket              string
	entity              string
	bucketaccesscontrol *BucketAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Update: Updates an ACL entry on the specified bucket.

func (r *BucketAccessControlsService) Update(bucket string, entity string, bucketaccesscontrol *BucketAccessControl) *BucketAccessControlsUpdateCall {
	return &BucketAccessControlsUpdateCall{
		s:                   r.s,
		bucket:              bucket,
		entity:              entity,
		bucketaccesscontrol: bucketaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketAccessControlsUpdateCall) Fields(s ...googleapi.Field) *BucketAccessControlsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketAccessControlsUpdateCall) Do() (*BucketAccessControl, error) {
	var returnValue *BucketAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucketaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an ACL entry on the specified bucket.",
	//   "httpMethod": "PUT",
	//   "id": "storage.bucketAccessControls.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/acl/{entity}",
	//   "request": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "response": {
	//     "$ref": "BucketAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.buckets.delete":

type BucketsDeleteCall struct {
	s             *Service
	bucket        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an empty bucket.

func (r *BucketsService) Delete(bucket string) *BucketsDeleteCall {
	return &BucketsDeleteCall{
		s:             r.s,
		bucket:        bucket,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}",
	}
}

func (c *BucketsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an empty bucket.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.buckets.delete",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.get":

type BucketsGetCall struct {
	s             *Service
	bucket        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns metadata for the specified bucket.

func (r *BucketsService) Get(bucket string) *BucketsGetCall {
	return &BucketsGetCall{
		s:             r.s,
		bucket:        bucket,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl.
func (c *BucketsGetCall) Projection(projection string) *BucketsGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsGetCall) Fields(s ...googleapi.Field) *BucketsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketsGetCall) Do() (*Bucket, error) {
	var returnValue *Bucket
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns metadata for the specified bucket.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.get",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.insert":

type BucketsInsertCall struct {
	s             *Service
	bucket        *Bucket
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new bucket.

func (r *BucketsService) Insert(bucket *Bucket) *BucketsInsertCall {
	return &BucketsInsertCall{
		s:             r.s,
		bucket:        bucket,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl, unless the bucket resource
// specifies acl or defaultObjectAcl properties, when it defaults to
// full.
func (c *BucketsInsertCall) Projection(projection string) *BucketsInsertCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsInsertCall) Fields(s ...googleapi.Field) *BucketsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketsInsertCall) Do() (*Bucket, error) {
	var returnValue *Bucket
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucket,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new bucket.",
	//   "httpMethod": "POST",
	//   "id": "storage.buckets.insert",
	//   "parameters": {
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl, unless the bucket resource specifies acl or defaultObjectAcl properties, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.list":

type BucketsListCall struct {
	s             *Service
	projectId     uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of buckets for a given project.

func (r *BucketsService) List(projectId uint64) *BucketsListCall {
	return &BucketsListCall{
		s:             r.s,
		projectId:     projectId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b",
	}
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of buckets to return.
func (c *BucketsListCall) MaxResults(maxResults int64) *BucketsListCall {
	c.params_.Set("max-results", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *BucketsListCall) PageToken(pageToken string) *BucketsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl.
func (c *BucketsListCall) Projection(projection string) *BucketsListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsListCall) Fields(s ...googleapi.Field) *BucketsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketsListCall) Do() (*Buckets, error) {
	var returnValue *Buckets
	c.params_.Set("projectId", fmt.Sprintf("%v", c.projectId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of buckets for a given project.",
	//   "httpMethod": "GET",
	//   "id": "storage.buckets.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "max-results": {
	//       "description": "Maximum number of buckets to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A valid API project identifier.",
	//       "format": "uint64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b",
	//   "response": {
	//     "$ref": "Buckets"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.patch":

type BucketsPatchCall struct {
	s             *Service
	bucket        string
	bucket2       *Bucket
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a bucket. This method supports patch semantics.

func (r *BucketsService) Patch(bucket string, bucket2 *Bucket) *BucketsPatchCall {
	return &BucketsPatchCall{
		s:             r.s,
		bucket:        bucket,
		bucket2:       bucket2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
func (c *BucketsPatchCall) Projection(projection string) *BucketsPatchCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsPatchCall) Fields(s ...googleapi.Field) *BucketsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketsPatchCall) Do() (*Bucket, error) {
	var returnValue *Bucket
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucket2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a bucket. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.buckets.patch",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.buckets.update":

type BucketsUpdateCall struct {
	s             *Service
	bucket        string
	bucket2       *Bucket
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a bucket.

func (r *BucketsService) Update(bucket string, bucket2 *Bucket) *BucketsUpdateCall {
	return &BucketsUpdateCall{
		s:             r.s,
		bucket:        bucket,
		bucket2:       bucket2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
func (c *BucketsUpdateCall) Projection(projection string) *BucketsUpdateCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BucketsUpdateCall) Fields(s ...googleapi.Field) *BucketsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BucketsUpdateCall) Do() (*Bucket, error) {
	var returnValue *Bucket
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.bucket2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a bucket.",
	//   "httpMethod": "PUT",
	//   "id": "storage.buckets.update",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit acl and defaultObjectAcl properties."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}",
	//   "request": {
	//     "$ref": "Bucket"
	//   },
	//   "response": {
	//     "$ref": "Bucket"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objectAccessControls.delete":

type ObjectAccessControlsDeleteCall struct {
	s             *Service
	bucket        string
	object        string
	entity        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the ACL entry for the specified entity on the
// specified object.

func (r *ObjectAccessControlsService) Delete(bucket string, object string, entity string) *ObjectAccessControlsDeleteCall {
	return &ObjectAccessControlsDeleteCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		entity:        entity,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}/acl/{entity}",
	}
}

func (c *ObjectAccessControlsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the ACL entry for the specified entity on the specified object.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.objectAccessControls.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.get":

type ObjectAccessControlsGetCall struct {
	s             *Service
	bucket        string
	object        string
	entity        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns the ACL entry for the specified entity on the specified
// object.

func (r *ObjectAccessControlsService) Get(bucket string, object string, entity string) *ObjectAccessControlsGetCall {
	return &ObjectAccessControlsGetCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		entity:        entity,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsGetCall) Fields(s ...googleapi.Field) *ObjectAccessControlsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectAccessControlsGetCall) Do() (*ObjectAccessControl, error) {
	var returnValue *ObjectAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the ACL entry for the specified entity on the specified object.",
	//   "httpMethod": "GET",
	//   "id": "storage.objectAccessControls.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.insert":

type ObjectAccessControlsInsertCall struct {
	s                   *Service
	bucket              string
	object              string
	objectaccesscontrol *ObjectAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Insert: Creates a new ACL entry on the specified object.

func (r *ObjectAccessControlsService) Insert(bucket string, object string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsInsertCall {
	return &ObjectAccessControlsInsertCall{
		s:                   r.s,
		bucket:              bucket,
		object:              object,
		objectaccesscontrol: objectaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/o/{object}/acl",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsInsertCall) Fields(s ...googleapi.Field) *ObjectAccessControlsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectAccessControlsInsertCall) Do() (*ObjectAccessControl, error) {
	var returnValue *ObjectAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.objectaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new ACL entry on the specified object.",
	//   "httpMethod": "POST",
	//   "id": "storage.objectAccessControls.insert",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.list":

type ObjectAccessControlsListCall struct {
	s             *Service
	bucket        string
	object        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves ACL entries on the specified object.

func (r *ObjectAccessControlsService) List(bucket string, object string) *ObjectAccessControlsListCall {
	return &ObjectAccessControlsListCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}/acl",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsListCall) Fields(s ...googleapi.Field) *ObjectAccessControlsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectAccessControlsListCall) Do() (*ObjectAccessControls, error) {
	var returnValue *ObjectAccessControls
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves ACL entries on the specified object.",
	//   "httpMethod": "GET",
	//   "id": "storage.objectAccessControls.list",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl",
	//   "response": {
	//     "$ref": "ObjectAccessControls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.patch":

type ObjectAccessControlsPatchCall struct {
	s                   *Service
	bucket              string
	object              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Patch: Updates an ACL entry on the specified object. This method
// supports patch semantics.

func (r *ObjectAccessControlsService) Patch(bucket string, object string, entity string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsPatchCall {
	return &ObjectAccessControlsPatchCall{
		s:                   r.s,
		bucket:              bucket,
		object:              object,
		entity:              entity,
		objectaccesscontrol: objectaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/o/{object}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsPatchCall) Fields(s ...googleapi.Field) *ObjectAccessControlsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectAccessControlsPatchCall) Do() (*ObjectAccessControl, error) {
	var returnValue *ObjectAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.objectaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an ACL entry on the specified object. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.objectAccessControls.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objectAccessControls.update":

type ObjectAccessControlsUpdateCall struct {
	s                   *Service
	bucket              string
	object              string
	entity              string
	objectaccesscontrol *ObjectAccessControl
	caller_             googleapi.Caller
	params_             url.Values
	pathTemplate_       string
}

// Update: Updates an ACL entry on the specified object.

func (r *ObjectAccessControlsService) Update(bucket string, object string, entity string, objectaccesscontrol *ObjectAccessControl) *ObjectAccessControlsUpdateCall {
	return &ObjectAccessControlsUpdateCall{
		s:                   r.s,
		bucket:              bucket,
		object:              object,
		entity:              entity,
		objectaccesscontrol: objectaccesscontrol,
		caller_:             googleapi.JSONCall{},
		params_:             make(map[string][]string),
		pathTemplate_:       "b/{bucket}/o/{object}/acl/{entity}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectAccessControlsUpdateCall) Fields(s ...googleapi.Field) *ObjectAccessControlsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectAccessControlsUpdateCall) Do() (*ObjectAccessControl, error) {
	var returnValue *ObjectAccessControl
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
		"entity": c.entity,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.objectaccesscontrol,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an ACL entry on the specified object.",
	//   "httpMethod": "PUT",
	//   "id": "storage.objectAccessControls.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "object",
	//     "entity"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of a bucket.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entity": {
	//       "description": "The entity holding the permission. Can be user-userId, user-emailAddress, group-groupId, group-emailAddress, allUsers, or allAuthenticatedUsers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}/acl/{entity}",
	//   "request": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "response": {
	//     "$ref": "ObjectAccessControl"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control"
	//   ]
	// }

}

// method id "storage.objects.delete":

type ObjectsDeleteCall struct {
	s             *Service
	bucket        string
	object        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes data blobs and associated metadata.

func (r *ObjectsService) Delete(bucket string, object string) *ObjectsDeleteCall {
	return &ObjectsDeleteCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}",
	}
}

func (c *ObjectsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes data blobs and associated metadata.",
	//   "httpMethod": "DELETE",
	//   "id": "storage.objects.delete",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.get":

type ObjectsGetCall struct {
	s             *Service
	bucket        string
	object        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves objects or their associated metadata.

func (r *ObjectsService) Get(bucket string, object string) *ObjectsGetCall {
	return &ObjectsGetCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl.
func (c *ObjectsGetCall) Projection(projection string) *ObjectsGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsGetCall) Fields(s ...googleapi.Field) *ObjectsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectsGetCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *ObjectsGetCall) Do() (*Object, error) {
	var returnValue *Object
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves objects or their associated metadata.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.get",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "storage.objects.insert":

type ObjectsInsertCall struct {
	s             *Service
	bucket        string
	object        *Object
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Stores new data blobs and associated metadata.

func (r *ObjectsService) Insert(bucket string, object *Object) *ObjectsInsertCall {
	return &ObjectsInsertCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o",
		context_:      context.TODO(),
	}
}

// Name sets the optional parameter "name": Name of the object. Required
// when the object metadata is not otherwise provided. Overrides the
// object metadata's name value, if any.
func (c *ObjectsInsertCall) Name(name string) *ObjectsInsertCall {
	c.params_.Set("name", fmt.Sprintf("%v", name))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl, unless the object resource
// specifies the acl property, when it defaults to full.
func (c *ObjectsInsertCall) Projection(projection string) *ObjectsInsertCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// MediaUpload takes a context and UploadCaller interface
func (c *ObjectsInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *ObjectsInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/storage/v1beta1/b/{bucket}/o"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/storage/v1beta1/b/{bucket}/o"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *ObjectsInsertCall) Media(r io.Reader) *ObjectsInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/storage/v1beta1/b/{bucket}/o"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *ObjectsInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ObjectsInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/storage/v1beta1/b/{bucket}/o"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *ObjectsInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ObjectsInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsInsertCall) Fields(s ...googleapi.Field) *ObjectsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectsInsertCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.object,
		Result:  &res,
	}

	return res, c.caller_.Do(c.context_, c.s.client, call)
}

func (c *ObjectsInsertCall) Do() (*Object, error) {
	var returnValue *Object
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.object,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Stores new data blobs and associated metadata.",
	//   "httpMethod": "POST",
	//   "id": "storage.objects.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/storage/v1beta1/b/{bucket}/o"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/storage/v1beta1/b/{bucket}/o"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which to store the new object. Overrides the provided object metadata's bucket value, if any.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Name of the object. Required when the object metadata is not otherwise provided. Overrides the object metadata's name value, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl, unless the object resource specifies the acl property, when it defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaDownload": true,
	//   "supportsMediaUpload": true
	// }

}

// method id "storage.objects.list":

type ObjectsListCall struct {
	s             *Service
	bucket        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of objects matching the criteria.

func (r *ObjectsService) List(bucket string) *ObjectsListCall {
	return &ObjectsListCall{
		s:             r.s,
		bucket:        bucket,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o",
	}
}

// Delimiter sets the optional parameter "delimiter": Returns results in
// a directory-like mode. items will contain only objects whose names,
// aside from the prefix, do not contain delimiter. Objects whose names,
// aside from the prefix, contain delimiter will have their name,
// truncated after the delimiter, returned in prefixes. Duplicate
// prefixes are omitted.
func (c *ObjectsListCall) Delimiter(delimiter string) *ObjectsListCall {
	c.params_.Set("delimiter", fmt.Sprintf("%v", delimiter))
	return c
}

// MaxResults sets the optional parameter "max-results": Maximum number
// of items plus prefixes to return. As duplicate prefixes are omitted,
// fewer total results may be returned than requested.
func (c *ObjectsListCall) MaxResults(maxResults int64) *ObjectsListCall {
	c.params_.Set("max-results", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *ObjectsListCall) PageToken(pageToken string) *ObjectsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Prefix sets the optional parameter "prefix": Filter results to
// objects whose names begin with this prefix.
func (c *ObjectsListCall) Prefix(prefix string) *ObjectsListCall {
	c.params_.Set("prefix", fmt.Sprintf("%v", prefix))
	return c
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to no_acl.
func (c *ObjectsListCall) Projection(projection string) *ObjectsListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsListCall) Fields(s ...googleapi.Field) *ObjectsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectsListCall) Do() (*Objects, error) {
	var returnValue *Objects
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of objects matching the criteria.",
	//   "httpMethod": "GET",
	//   "id": "storage.objects.list",
	//   "parameterOrder": [
	//     "bucket"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which to look for objects.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "delimiter": {
	//       "description": "Returns results in a directory-like mode. items will contain only objects whose names, aside from the prefix, do not contain delimiter. Objects whose names, aside from the prefix, contain delimiter will have their name, truncated after the delimiter, returned in prefixes. Duplicate prefixes are omitted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "Maximum number of items plus prefixes to return. As duplicate prefixes are omitted, fewer total results may be returned than requested.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "prefix": {
	//       "description": "Filter results to objects whose names begin with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to no_acl.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o",
	//   "response": {
	//     "$ref": "Objects"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "storage.objects.patch":

type ObjectsPatchCall struct {
	s             *Service
	bucket        string
	object        string
	object2       *Object
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a data blob's associated metadata. This method
// supports patch semantics.

func (r *ObjectsService) Patch(bucket string, object string, object2 *Object) *ObjectsPatchCall {
	return &ObjectsPatchCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		object2:       object2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
func (c *ObjectsPatchCall) Projection(projection string) *ObjectsPatchCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsPatchCall) Fields(s ...googleapi.Field) *ObjectsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectsPatchCall) Do() (*Object, error) {
	var returnValue *Object
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.object2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a data blob's associated metadata. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "storage.objects.patch",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ]
	// }

}

// method id "storage.objects.update":

type ObjectsUpdateCall struct {
	s             *Service
	bucket        string
	object        string
	object2       *Object
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a data blob's associated metadata.

func (r *ObjectsService) Update(bucket string, object string, object2 *Object) *ObjectsUpdateCall {
	return &ObjectsUpdateCall{
		s:             r.s,
		bucket:        bucket,
		object:        object,
		object2:       object2,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "b/{bucket}/o/{object}",
	}
}

// Projection sets the optional parameter "projection": Set of
// properties to return. Defaults to full.
func (c *ObjectsUpdateCall) Projection(projection string) *ObjectsUpdateCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ObjectsUpdateCall) Fields(s ...googleapi.Field) *ObjectsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ObjectsUpdateCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.object2,
		Result:  &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *ObjectsUpdateCall) Do() (*Object, error) {
	var returnValue *Object
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"bucket": c.bucket,
		"object": c.object,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.object2,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a data blob's associated metadata.",
	//   "httpMethod": "PUT",
	//   "id": "storage.objects.update",
	//   "parameterOrder": [
	//     "bucket",
	//     "object"
	//   ],
	//   "parameters": {
	//     "bucket": {
	//       "description": "Name of the bucket in which the object resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "object": {
	//       "description": "Name of the object.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Set of properties to return. Defaults to full.",
	//       "enum": [
	//         "full",
	//         "no_acl"
	//       ],
	//       "enumDescriptions": [
	//         "Include all properties.",
	//         "Omit the acl property."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "b/{bucket}/o/{object}",
	//   "request": {
	//     "$ref": "Object"
	//   },
	//   "response": {
	//     "$ref": "Object"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

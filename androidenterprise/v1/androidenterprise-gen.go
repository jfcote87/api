// Package androidenterprise provides access to the Google Play MDM API.
//
// Usage example:
//
//   import "github.com/jfcote87/api/androidenterprise/v1"
//   ...
//   androidenterpriseService, err := androidenterprise.New(oauthHttpClient)
package androidenterprise

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

const apiId = "androidenterprise:v1"
const apiName = "androidenterprise"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/androidenterprise/v1/"}

// OAuth2 scopes used by this API.
const (
	// Manage corporate Android devices
	AndroidenterpriseScope = "https://www.googleapis.com/auth/androidenterprise"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Collections = NewCollectionsService(s)
	s.Collectionviewers = NewCollectionviewersService(s)
	s.Devices = NewDevicesService(s)
	s.Enterprises = NewEnterprisesService(s)
	s.Entitlements = NewEntitlementsService(s)
	s.Grouplicenses = NewGrouplicensesService(s)
	s.Grouplicenseusers = NewGrouplicenseusersService(s)
	s.Installs = NewInstallsService(s)
	s.Permissions = NewPermissionsService(s)
	s.Products = NewProductsService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Collections *CollectionsService

	Collectionviewers *CollectionviewersService

	Devices *DevicesService

	Enterprises *EnterprisesService

	Entitlements *EntitlementsService

	Grouplicenses *GrouplicensesService

	Grouplicenseusers *GrouplicenseusersService

	Installs *InstallsService

	Permissions *PermissionsService

	Products *ProductsService

	Users *UsersService
}

func NewCollectionsService(s *Service) *CollectionsService {
	rs := &CollectionsService{s: s}
	return rs
}

type CollectionsService struct {
	s *Service
}

func NewCollectionviewersService(s *Service) *CollectionviewersService {
	rs := &CollectionviewersService{s: s}
	return rs
}

type CollectionviewersService struct {
	s *Service
}

func NewDevicesService(s *Service) *DevicesService {
	rs := &DevicesService{s: s}
	return rs
}

type DevicesService struct {
	s *Service
}

func NewEnterprisesService(s *Service) *EnterprisesService {
	rs := &EnterprisesService{s: s}
	return rs
}

type EnterprisesService struct {
	s *Service
}

func NewEntitlementsService(s *Service) *EntitlementsService {
	rs := &EntitlementsService{s: s}
	return rs
}

type EntitlementsService struct {
	s *Service
}

func NewGrouplicensesService(s *Service) *GrouplicensesService {
	rs := &GrouplicensesService{s: s}
	return rs
}

type GrouplicensesService struct {
	s *Service
}

func NewGrouplicenseusersService(s *Service) *GrouplicenseusersService {
	rs := &GrouplicenseusersService{s: s}
	return rs
}

type GrouplicenseusersService struct {
	s *Service
}

func NewInstallsService(s *Service) *InstallsService {
	rs := &InstallsService{s: s}
	return rs
}

type InstallsService struct {
	s *Service
}

func NewPermissionsService(s *Service) *PermissionsService {
	rs := &PermissionsService{s: s}
	return rs
}

type PermissionsService struct {
	s *Service
}

func NewProductsService(s *Service) *ProductsService {
	rs := &ProductsService{s: s}
	return rs
}

type ProductsService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type AppRestrictionsSchema struct {
	// Restrictions: The set of restrictions that make up this schema.
	Restrictions []*AppRestrictionsSchemaRestriction `json:"restrictions,omitempty"`
}

type AppRestrictionsSchemaRestriction struct {
	// DefaultValue: The default value of the restriction.
	DefaultValue *AppRestrictionsSchemaRestrictionRestrictionValue `json:"defaultValue,omitempty"`

	// Description: A longer description of the restriction, giving more
	// detail of what it affects.
	Description string `json:"description,omitempty"`

	// Entry: For choice or multiselect restrictions, the list of possible
	// entries' human-readable names.
	Entry []string `json:"entry,omitempty"`

	// EntryValue: For choice or multiselect restrictions, the list of
	// possible entries' machine-readable values.
	EntryValue []string `json:"entryValue,omitempty"`

	// Key: The unique key that the product uses to identify the
	// restriction, e.g. "com.google.android.gm.fieldname".
	Key string `json:"key,omitempty"`

	// RestrictionType: The type of the restriction.
	RestrictionType string `json:"restrictionType,omitempty"`

	// Title: The name of the restriction.
	Title string `json:"title,omitempty"`
}

type AppRestrictionsSchemaRestrictionRestrictionValue struct {
	// Type: The type of the value being provided.
	Type string `json:"type,omitempty"`

	// ValueBool: The boolean value - this will only be present if type is
	// bool.
	ValueBool bool `json:"valueBool,omitempty"`

	// ValueInteger: The integer value - this will only be present if type
	// is integer.
	ValueInteger int64 `json:"valueInteger,omitempty"`

	// ValueMultiselect: The list of string values - this will only be
	// present if type is multiselect.
	ValueMultiselect []string `json:"valueMultiselect,omitempty"`

	// ValueString: The string value - this will be present for types
	// string, choice and hidden.
	ValueString string `json:"valueString,omitempty"`
}

type Collection struct {
	// CollectionId: Arbitrary unique ID, allocated by the API on creation.
	CollectionId string `json:"collectionId,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collection".
	Kind string `json:"kind,omitempty"`

	// Name: A user-friendly name for the collection (should be unique),
	// e.g. "Accounting apps".
	Name string `json:"name,omitempty"`

	// ProductId: The IDs of the products in the collection, in the order in
	// which they should be displayed.
	ProductId []string `json:"productId,omitempty"`

	// Visibility: Whether this collection is visible to all users, or only
	// to the users that have been granted access through the
	// collection_viewers api. Even if a collection is visible to allUsers,
	// it is possible to add and remove viewers, but this will have no
	// effect until the collection's visibility changes to viewersOnly.
	Visibility string `json:"visibility,omitempty"`
}

type CollectionViewersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionViewersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

type CollectionsListResponse struct {
	// Collection: An ordered collection of products which can be made
	// visible on the Google Play Store app to a selected group of users.
	Collection []*Collection `json:"collection,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#collectionsListResponse".
	Kind string `json:"kind,omitempty"`
}

type Device struct {
	// AndroidId: The Google Play Services Android ID for the device encoded
	// as a lowercase hex string, e.g. "123456789abcdef0".
	AndroidId string `json:"androidId,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#device".
	Kind string `json:"kind,omitempty"`
}

type DeviceState struct {
	// AccountState: The state of the Google account on the device.
	// "enabled" indicates that the Google account on the device can be used
	// to access Google services (including Google Play), while "disabled"
	// means that it cannot. A new device is initially in the "disabled"
	// state.
	AccountState string `json:"accountState,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#deviceState".
	Kind string `json:"kind,omitempty"`
}

type DevicesListResponse struct {
	// Device: A managed device.
	Device []*Device `json:"device,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#devicesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Enterprise struct {
	// Id: The unique ID for the enterprise.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterprise".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the enterprise, e.g. "Example Inc".
	Name string `json:"name,omitempty"`

	// PrimaryDomain: The enterprise's primary domain, e.g. "example.com".
	PrimaryDomain string `json:"primaryDomain,omitempty"`
}

type EnterpriseAccount struct {
	// AccountEmail: The email address of the service account.
	AccountEmail string `json:"accountEmail,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterpriseAccount".
	Kind string `json:"kind,omitempty"`
}

type EnterprisesListResponse struct {
	// Enterprise: An enterprise.
	Enterprise []*Enterprise `json:"enterprise,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#enterprisesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Entitlement struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#entitlement".
	Kind string `json:"kind,omitempty"`

	// ProductId: The ID of the product that the entitlement is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`

	// Reason: The reason for the entitlement, e.g. "free" for free apps.
	// This is temporary, it will be replaced by the acquisition kind field
	// of group licenses.
	Reason string `json:"reason,omitempty"`
}

type EntitlementsListResponse struct {
	// Entitlement: An entitlement of a user to a product (e.g. an app). For
	// example, a free app that they have installed, or a paid app that they
	// have been allocated a license to.
	Entitlement []*Entitlement `json:"entitlement,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#entitlementsListResponse".
	Kind string `json:"kind,omitempty"`
}

type GroupLicense struct {
	// AcquisitionKind: How this group license was acquired. "bulkPurchase"
	// means that this group license object was created because the
	// enterprise purchased licenses for this product; this is "free"
	// otherwise (for free products).
	AcquisitionKind string `json:"acquisitionKind,omitempty"`

	// Approval: Whether the product to which this group license relates is
	// currently approved by the enterprise, as either "approved" or
	// "unapproved". Products are approved when a group license is first
	// created, but this approval may be revoked by an enterprise admin via
	// Google Play. Unapproved products will not be visible to end users in
	// collections and new entitlements to them should not normally be
	// created.
	Approval string `json:"approval,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicense".
	Kind string `json:"kind,omitempty"`

	// NumProvisioned: The total number of provisioned licenses for this
	// product. Returned by read operations, but ignored in write
	// operations.
	NumProvisioned int64 `json:"numProvisioned,omitempty"`

	// NumPurchased: The number of purchased licenses (possibly in multiple
	// purchases). If this field is omitted then there is no limit on the
	// number of licenses that can be provisioned (e.g. if the acquisition
	// kind is "free").
	NumPurchased int64 `json:"numPurchased,omitempty"`

	// ProductId: The ID of the product that the license is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`
}

type GroupLicenseUsersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicenseUsersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

type GroupLicensesListResponse struct {
	// GroupLicense: A group license for a product approved for use in the
	// enterprise.
	GroupLicense []*GroupLicense `json:"groupLicense,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#groupLicensesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Install struct {
	// InstallState: Install state. The state "installPending" means that an
	// install request has recently been made and download to the device is
	// in progress. The state "installed" means that the app has been
	// installed. This field is read-only.
	InstallState string `json:"installState,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#install".
	Kind string `json:"kind,omitempty"`

	// ProductId: The ID of the product that the install is for, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`

	// VersionCode: The version of the installed product. Guaranteed to be
	// set only if the install state is "installed".
	VersionCode int64 `json:"versionCode,omitempty"`
}

type InstallsListResponse struct {
	// Install: An installation of an app for a user on a specific device.
	// The existence of an install implies that the user must have an
	// entitlement to the app.
	Install []*Install `json:"install,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#installsListResponse".
	Kind string `json:"kind,omitempty"`
}

type Permission struct {
	// Description: A longer description of the permissions giving more
	// details of what it affects.
	Description string `json:"description,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#permission".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the permission.
	Name string `json:"name,omitempty"`

	// PermissionId: An opaque string uniquely identifying the permission.
	PermissionId string `json:"permissionId,omitempty"`
}

type Product struct {
	// AuthorName: The name of the author of the product (e.g. the app
	// developer).
	AuthorName string `json:"authorName,omitempty"`

	// DetailsUrl: A link to the (consumer) Google Play details page for the
	// product.
	DetailsUrl string `json:"detailsUrl,omitempty"`

	// IconUrl: A link to an image that can be used as an icon for the
	// product.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#product".
	Kind string `json:"kind,omitempty"`

	// ProductId: A string of the form "app:
	// " - e.g.
	// "app:com.google.android.gm" represents the GMail app.
	ProductId string `json:"productId,omitempty"`

	// Title: The name of the product.
	Title string `json:"title,omitempty"`

	// WorkDetailsUrl: A link to the Google Play for Work details page for
	// the product, for use by an Enterprise administrator.
	WorkDetailsUrl string `json:"workDetailsUrl,omitempty"`
}

type ProductPermission struct {
	// PermissionId: An opaque string uniquely identifying the permission.
	PermissionId string `json:"permissionId,omitempty"`

	// State: Whether the permission has been accepted or not.
	State string `json:"state,omitempty"`
}

type ProductPermissions struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#productPermissions".
	Kind string `json:"kind,omitempty"`

	// Permission: The permissions required by the app.
	Permission []*ProductPermission `json:"permission,omitempty"`

	// ProductId: The ID of the app that the permissions relate to, e.g.
	// "app:com.google.android.gm".
	ProductId string `json:"productId,omitempty"`
}

type User struct {
	// Id: The unique ID for the user.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#user".
	Kind string `json:"kind,omitempty"`

	// PrimaryEmail: The user's primary email, e.g. "jsmith@example.com".
	PrimaryEmail string `json:"primaryEmail,omitempty"`
}

type UserToken struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#userToken".
	Kind string `json:"kind,omitempty"`

	// Token: The token (activation code) to be entered by the user. This
	// consists of a sequence of decimal digits. Note that the leading digit
	// may be 0.
	Token string `json:"token,omitempty"`

	// UserId: The unique ID for the user.
	UserId string `json:"userId,omitempty"`
}

type UsersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "androidenterprise#usersListResponse".
	Kind string `json:"kind,omitempty"`

	// User: A user of an enterprise.
	User []*User `json:"user,omitempty"`
}

// method id "androidenterprise.collections.delete":

type CollectionsDeleteCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a collection.

func (r *CollectionsService) Delete(enterpriseId string, collectionId string) *CollectionsDeleteCall {
	return &CollectionsDeleteCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}",
	}
}

func (c *CollectionsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a collection.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.collections.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.get":

type CollectionsGetCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the details of a collection.

func (r *CollectionsService) Get(enterpriseId string, collectionId string) *CollectionsGetCall {
	return &CollectionsGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsGetCall) Fields(s ...googleapi.Field) *CollectionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionsGetCall) Do() (*Collection, error) {
	var returnValue *Collection
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the details of a collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collections.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.insert":

type CollectionsInsertCall struct {
	s             *Service
	enterpriseId  string
	collection    *Collection
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new collection.

func (r *CollectionsService) Insert(enterpriseId string, collection *Collection) *CollectionsInsertCall {
	return &CollectionsInsertCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collection:    collection,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsInsertCall) Fields(s ...googleapi.Field) *CollectionsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionsInsertCall) Do() (*Collection, error) {
	var returnValue *Collection
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.collection,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new collection.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.collections.insert",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.list":

type CollectionsListCall struct {
	s             *Service
	enterpriseId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the IDs of all the collections for an enterprise.

func (r *CollectionsService) List(enterpriseId string) *CollectionsListCall {
	return &CollectionsListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsListCall) Fields(s ...googleapi.Field) *CollectionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionsListCall) Do() (*CollectionsListResponse, error) {
	var returnValue *CollectionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the IDs of all the collections for an enterprise.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collections.list",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections",
	//   "response": {
	//     "$ref": "CollectionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.patch":

type CollectionsPatchCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	collection    *Collection
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a collection. This method supports patch semantics.

func (r *CollectionsService) Patch(enterpriseId string, collectionId string, collection *Collection) *CollectionsPatchCall {
	return &CollectionsPatchCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		collection:    collection,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsPatchCall) Fields(s ...googleapi.Field) *CollectionsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionsPatchCall) Do() (*Collection, error) {
	var returnValue *Collection
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.collection,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a collection. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.collections.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collections.update":

type CollectionsUpdateCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	collection    *Collection
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a collection.

func (r *CollectionsService) Update(enterpriseId string, collectionId string, collection *Collection) *CollectionsUpdateCall {
	return &CollectionsUpdateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		collection:    collection,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionsUpdateCall) Fields(s ...googleapi.Field) *CollectionsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionsUpdateCall) Do() (*Collection, error) {
	var returnValue *Collection
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.collection,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a collection.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.collections.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}",
	//   "request": {
	//     "$ref": "Collection"
	//   },
	//   "response": {
	//     "$ref": "Collection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.delete":

type CollectionviewersDeleteCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Removes the user from the list of those specifically allowed
// to see the collection. If the collection's visibility is set to
// viewersOnly then only such users will see the collection.

func (r *CollectionviewersService) Delete(enterpriseId string, collectionId string, userId string) *CollectionviewersDeleteCall {
	return &CollectionviewersDeleteCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	}
}

func (c *CollectionviewersDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes the user from the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.collectionviewers.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.get":

type CollectionviewersGetCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the ID of the user if they have been specifically
// allowed to see the collection. If the collection's visibility is set
// to viewersOnly then only these users will see the collection.

func (r *CollectionviewersService) Get(enterpriseId string, collectionId string, userId string) *CollectionviewersGetCall {
	return &CollectionviewersGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersGetCall) Fields(s ...googleapi.Field) *CollectionviewersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionviewersGetCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the ID of the user if they have been specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only these users will see the collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collectionviewers.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.list":

type CollectionviewersListCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the IDs of the users who have been specifically
// allowed to see the collection. If the collection's visibility is set
// to viewersOnly then only these users will see the collection.

func (r *CollectionviewersService) List(enterpriseId string, collectionId string) *CollectionviewersListCall {
	return &CollectionviewersListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}/users",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersListCall) Fields(s ...googleapi.Field) *CollectionviewersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionviewersListCall) Do() (*CollectionViewersListResponse, error) {
	var returnValue *CollectionViewersListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the IDs of the users who have been specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only these users will see the collection.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.collectionviewers.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users",
	//   "response": {
	//     "$ref": "CollectionViewersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.patch":

type CollectionviewersPatchCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	userId        string
	user          *User
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Adds the user to the list of those specifically allowed to see
// the collection. If the collection's visibility is set to viewersOnly
// then only such users will see the collection. This method supports
// patch semantics.

func (r *CollectionviewersService) Patch(enterpriseId string, collectionId string, userId string, user *User) *CollectionviewersPatchCall {
	return &CollectionviewersPatchCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		userId:        userId,
		user:          user,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersPatchCall) Fields(s ...googleapi.Field) *CollectionviewersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionviewersPatchCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.user,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds the user to the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.collectionviewers.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.collectionviewers.update":

type CollectionviewersUpdateCall struct {
	s             *Service
	enterpriseId  string
	collectionId  string
	userId        string
	user          *User
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Adds the user to the list of those specifically allowed to
// see the collection. If the collection's visibility is set to
// viewersOnly then only such users will see the collection.

func (r *CollectionviewersService) Update(enterpriseId string, collectionId string, userId string, user *User) *CollectionviewersUpdateCall {
	return &CollectionviewersUpdateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		collectionId:  collectionId,
		userId:        userId,
		user:          user,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CollectionviewersUpdateCall) Fields(s ...googleapi.Field) *CollectionviewersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CollectionviewersUpdateCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"collectionId": c.collectionId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.user,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds the user to the list of those specifically allowed to see the collection. If the collection's visibility is set to viewersOnly then only such users will see the collection.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.collectionviewers.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "collectionId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The ID of the collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/collections/{collectionId}/users/{userId}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.get":

type DevicesGetCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the details of a device.

func (r *DevicesService) Get(enterpriseId string, userId string, deviceId string) *DevicesGetCall {
	return &DevicesGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesGetCall) Fields(s ...googleapi.Field) *DevicesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DevicesGetCall) Do() (*Device, error) {
	var returnValue *Device
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the details of a device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}",
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.getState":

type DevicesGetStateCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetState: Retrieves whether a device is enabled or disabled for
// access by the user to Google services. The device state takes effect
// only if enforcing EMM policies on Android devices is enabled in the
// Google Admin Console. Otherwise, the device state is ignored and all
// devices are allowed access to Google services.

func (r *DevicesService) GetState(enterpriseId string, userId string, deviceId string) *DevicesGetStateCall {
	return &DevicesGetStateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesGetStateCall) Fields(s ...googleapi.Field) *DevicesGetStateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DevicesGetStateCall) Do() (*DeviceState, error) {
	var returnValue *DeviceState
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves whether a device is enabled or disabled for access by the user to Google services. The device state takes effect only if enforcing EMM policies on Android devices is enabled in the Google Admin Console. Otherwise, the device state is ignored and all devices are allowed access to Google services.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.getState",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	//   "response": {
	//     "$ref": "DeviceState"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.list":

type DevicesListCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the IDs of all of a user's devices.

func (r *DevicesService) List(enterpriseId string, userId string) *DevicesListCall {
	return &DevicesListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesListCall) Fields(s ...googleapi.Field) *DevicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DevicesListCall) Do() (*DevicesListResponse, error) {
	var returnValue *DevicesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the IDs of all of a user's devices.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.devices.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices",
	//   "response": {
	//     "$ref": "DevicesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.devices.setState":

type DevicesSetStateCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	devicestate   *DeviceState
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// SetState: Sets whether a device is enabled or disabled for access by
// the user to Google services. The device state takes effect only if
// enforcing EMM policies on Android devices is enabled in the Google
// Admin Console. Otherwise, the device state is ignored and all devices
// are allowed access to Google services.

func (r *DevicesService) SetState(enterpriseId string, userId string, deviceId string, devicestate *DeviceState) *DevicesSetStateCall {
	return &DevicesSetStateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		devicestate:   devicestate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DevicesSetStateCall) Fields(s ...googleapi.Field) *DevicesSetStateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DevicesSetStateCall) Do() (*DeviceState, error) {
	var returnValue *DeviceState
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.devicestate,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Sets whether a device is enabled or disabled for access by the user to Google services. The device state takes effect only if enforcing EMM policies on Android devices is enabled in the Google Admin Console. Otherwise, the device state is ignored and all devices are allowed access to Google services.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.devices.setState",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/state",
	//   "request": {
	//     "$ref": "DeviceState"
	//   },
	//   "response": {
	//     "$ref": "DeviceState"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.delete":

type EnterprisesDeleteCall struct {
	s             *Service
	enterpriseId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the binding between the MDM and enterprise. This is
// now deprecated; use this to unenroll customers that were previously
// enrolled with the 'insert' call, then enroll them again with the
// 'enroll' call.

func (r *EnterprisesService) Delete(enterpriseId string) *EnterprisesDeleteCall {
	return &EnterprisesDeleteCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}",
	}
}

func (c *EnterprisesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the binding between the MDM and enterprise. This is now deprecated; use this to unenroll customers that were previously enrolled with the 'insert' call, then enroll them again with the 'enroll' call.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.enterprises.delete",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.enroll":

type EnterprisesEnrollCall struct {
	s             *Service
	token         string
	enterprise    *Enterprise
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Enroll: Enrolls an enterprise with the calling MDM.

func (r *EnterprisesService) Enroll(token string, enterprise *Enterprise) *EnterprisesEnrollCall {
	return &EnterprisesEnrollCall{
		s:             r.s,
		token:         token,
		enterprise:    enterprise,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/enroll",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesEnrollCall) Fields(s ...googleapi.Field) *EnterprisesEnrollCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EnterprisesEnrollCall) Do() (*Enterprise, error) {
	var returnValue *Enterprise
	c.params_.Set("token", fmt.Sprintf("%v", c.token))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.enterprise,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Enrolls an enterprise with the calling MDM.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.enroll",
	//   "parameterOrder": [
	//     "token"
	//   ],
	//   "parameters": {
	//     "token": {
	//       "description": "The token provided by the enterprise to register the MDM.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/enroll",
	//   "request": {
	//     "$ref": "Enterprise"
	//   },
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.get":

type EnterprisesGetCall struct {
	s             *Service
	enterpriseId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the name and domain of an enterprise.

func (r *EnterprisesService) Get(enterpriseId string) *EnterprisesGetCall {
	return &EnterprisesGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesGetCall) Fields(s ...googleapi.Field) *EnterprisesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EnterprisesGetCall) Do() (*Enterprise, error) {
	var returnValue *Enterprise
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the name and domain of an enterprise.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.enterprises.get",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}",
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.insert":

type EnterprisesInsertCall struct {
	s             *Service
	token         string
	enterprise    *Enterprise
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Establishes the binding between the MDM and an enterprise.
// This is now deprecated; use enroll instead.

func (r *EnterprisesService) Insert(token string, enterprise *Enterprise) *EnterprisesInsertCall {
	return &EnterprisesInsertCall{
		s:             r.s,
		token:         token,
		enterprise:    enterprise,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesInsertCall) Fields(s ...googleapi.Field) *EnterprisesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EnterprisesInsertCall) Do() (*Enterprise, error) {
	var returnValue *Enterprise
	c.params_.Set("token", fmt.Sprintf("%v", c.token))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.enterprise,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Establishes the binding between the MDM and an enterprise. This is now deprecated; use enroll instead.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.insert",
	//   "parameterOrder": [
	//     "token"
	//   ],
	//   "parameters": {
	//     "token": {
	//       "description": "The token provided by the enterprise to register the MDM.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises",
	//   "request": {
	//     "$ref": "Enterprise"
	//   },
	//   "response": {
	//     "$ref": "Enterprise"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.list":

type EnterprisesListCall struct {
	s             *Service
	domain        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Looks up an enterprise by domain name.

func (r *EnterprisesService) List(domain string) *EnterprisesListCall {
	return &EnterprisesListCall{
		s:             r.s,
		domain:        domain,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesListCall) Fields(s ...googleapi.Field) *EnterprisesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EnterprisesListCall) Do() (*EnterprisesListResponse, error) {
	var returnValue *EnterprisesListResponse
	c.params_.Set("domain", fmt.Sprintf("%v", c.domain))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Looks up an enterprise by domain name.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.enterprises.list",
	//   "parameterOrder": [
	//     "domain"
	//   ],
	//   "parameters": {
	//     "domain": {
	//       "description": "The exact primary domain name of the enterprise to look up.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises",
	//   "response": {
	//     "$ref": "EnterprisesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.setAccount":

type EnterprisesSetAccountCall struct {
	s                 *Service
	enterpriseId      string
	enterpriseaccount *EnterpriseAccount
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// SetAccount: Set the account that will be used to authenticate to the
// API as the enterprise.

func (r *EnterprisesService) SetAccount(enterpriseId string, enterpriseaccount *EnterpriseAccount) *EnterprisesSetAccountCall {
	return &EnterprisesSetAccountCall{
		s:                 r.s,
		enterpriseId:      enterpriseId,
		enterpriseaccount: enterpriseaccount,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "enterprises/{enterpriseId}/account",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EnterprisesSetAccountCall) Fields(s ...googleapi.Field) *EnterprisesSetAccountCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EnterprisesSetAccountCall) Do() (*EnterpriseAccount, error) {
	var returnValue *EnterpriseAccount
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.enterpriseaccount,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Set the account that will be used to authenticate to the API as the enterprise.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.enterprises.setAccount",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/account",
	//   "request": {
	//     "$ref": "EnterpriseAccount"
	//   },
	//   "response": {
	//     "$ref": "EnterpriseAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.enterprises.unenroll":

type EnterprisesUnenrollCall struct {
	s             *Service
	enterpriseId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Unenroll: Unenrolls an enterprise from the calling MDM.

func (r *EnterprisesService) Unenroll(enterpriseId string) *EnterprisesUnenrollCall {
	return &EnterprisesUnenrollCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/unenroll",
	}
}

func (c *EnterprisesUnenrollCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Unenrolls an enterprise from the calling MDM.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.enterprises.unenroll",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/unenroll",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.delete":

type EntitlementsDeleteCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Removes an entitlement to an app for a user and uninstalls
// it.

func (r *EntitlementsService) Delete(enterpriseId string, userId string, entitlementId string) *EntitlementsDeleteCall {
	return &EntitlementsDeleteCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		entitlementId: entitlementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	}
}

func (c *EntitlementsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes an entitlement to an app for a user and uninstalls it.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.entitlements.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.get":

type EntitlementsGetCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves details of an entitlement.

func (r *EntitlementsService) Get(enterpriseId string, userId string, entitlementId string) *EntitlementsGetCall {
	return &EntitlementsGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		entitlementId: entitlementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsGetCall) Fields(s ...googleapi.Field) *EntitlementsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EntitlementsGetCall) Do() (*Entitlement, error) {
	var returnValue *Entitlement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details of an entitlement.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.entitlements.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.list":

type EntitlementsListCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List of all entitlements for the specified user. Only the ID is
// set.

func (r *EntitlementsService) List(enterpriseId string, userId string) *EntitlementsListCall {
	return &EntitlementsListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/entitlements",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsListCall) Fields(s ...googleapi.Field) *EntitlementsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EntitlementsListCall) Do() (*EntitlementsListResponse, error) {
	var returnValue *EntitlementsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List of all entitlements for the specified user. Only the ID is set.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.entitlements.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements",
	//   "response": {
	//     "$ref": "EntitlementsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.patch":

type EntitlementsPatchCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	entitlement   *Entitlement
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Adds or updates an entitlement to an app for a user. This
// method supports patch semantics.

func (r *EntitlementsService) Patch(enterpriseId string, userId string, entitlementId string, entitlement *Entitlement) *EntitlementsPatchCall {
	return &EntitlementsPatchCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		entitlementId: entitlementId,
		entitlement:   entitlement,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	}
}

// Install sets the optional parameter "install": Set to true to also
// install the product on all the user's devices where possible. Failure
// to install on one or more devices will not prevent this operation
// from returning successfully, as long as the entitlement was
// successfully assigned to the user.
func (c *EntitlementsPatchCall) Install(install bool) *EntitlementsPatchCall {
	c.params_.Set("install", fmt.Sprintf("%v", install))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsPatchCall) Fields(s ...googleapi.Field) *EntitlementsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EntitlementsPatchCall) Do() (*Entitlement, error) {
	var returnValue *Entitlement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.entitlement,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds or updates an entitlement to an app for a user. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.entitlements.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "install": {
	//       "description": "Set to true to also install the product on all the user's devices where possible. Failure to install on one or more devices will not prevent this operation from returning successfully, as long as the entitlement was successfully assigned to the user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "request": {
	//     "$ref": "Entitlement"
	//   },
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.entitlements.update":

type EntitlementsUpdateCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	entitlementId string
	entitlement   *Entitlement
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Adds or updates an entitlement to an app for a user.

func (r *EntitlementsService) Update(enterpriseId string, userId string, entitlementId string, entitlement *Entitlement) *EntitlementsUpdateCall {
	return &EntitlementsUpdateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		entitlementId: entitlementId,
		entitlement:   entitlement,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	}
}

// Install sets the optional parameter "install": Set to true to also
// install the product on all the user's devices where possible. Failure
// to install on one or more devices will not prevent this operation
// from returning successfully, as long as the entitlement was
// successfully assigned to the user.
func (c *EntitlementsUpdateCall) Install(install bool) *EntitlementsUpdateCall {
	c.params_.Set("install", fmt.Sprintf("%v", install))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntitlementsUpdateCall) Fields(s ...googleapi.Field) *EntitlementsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EntitlementsUpdateCall) Do() (*Entitlement, error) {
	var returnValue *Entitlement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":  c.enterpriseId,
		"userId":        c.userId,
		"entitlementId": c.entitlementId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.entitlement,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds or updates an entitlement to an app for a user.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.entitlements.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "entitlementId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "entitlementId": {
	//       "description": "The ID of the entitlement, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "install": {
	//       "description": "Set to true to also install the product on all the user's devices where possible. Failure to install on one or more devices will not prevent this operation from returning successfully, as long as the entitlement was successfully assigned to the user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/entitlements/{entitlementId}",
	//   "request": {
	//     "$ref": "Entitlement"
	//   },
	//   "response": {
	//     "$ref": "Entitlement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenses.get":

type GrouplicensesGetCall struct {
	s              *Service
	enterpriseId   string
	groupLicenseId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Get: Retrieves details of an enterprise's group license for a
// product.

func (r *GrouplicensesService) Get(enterpriseId string, groupLicenseId string) *GrouplicensesGetCall {
	return &GrouplicensesGetCall{
		s:              r.s,
		enterpriseId:   enterpriseId,
		groupLicenseId: groupLicenseId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicensesGetCall) Fields(s ...googleapi.Field) *GrouplicensesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GrouplicensesGetCall) Do() (*GroupLicense, error) {
	var returnValue *GroupLicense
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":   c.enterpriseId,
		"groupLicenseId": c.groupLicenseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details of an enterprise's group license for a product.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenses.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "groupLicenseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupLicenseId": {
	//       "description": "The ID of the product the group license is for, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}",
	//   "response": {
	//     "$ref": "GroupLicense"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenses.list":

type GrouplicensesListCall struct {
	s             *Service
	enterpriseId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves IDs of all products for which the enterprise has a
// group license.

func (r *GrouplicensesService) List(enterpriseId string) *GrouplicensesListCall {
	return &GrouplicensesListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/groupLicenses",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicensesListCall) Fields(s ...googleapi.Field) *GrouplicensesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GrouplicensesListCall) Do() (*GroupLicensesListResponse, error) {
	var returnValue *GroupLicensesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves IDs of all products for which the enterprise has a group license.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenses.list",
	//   "parameterOrder": [
	//     "enterpriseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses",
	//   "response": {
	//     "$ref": "GroupLicensesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.grouplicenseusers.list":

type GrouplicenseusersListCall struct {
	s              *Service
	enterpriseId   string
	groupLicenseId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// List: Retrieves the IDs of the users who have been granted
// entitlements under the license.

func (r *GrouplicenseusersService) List(enterpriseId string, groupLicenseId string) *GrouplicenseusersListCall {
	return &GrouplicenseusersListCall{
		s:              r.s,
		enterpriseId:   enterpriseId,
		groupLicenseId: groupLicenseId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}/users",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GrouplicenseusersListCall) Fields(s ...googleapi.Field) *GrouplicenseusersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GrouplicenseusersListCall) Do() (*GroupLicenseUsersListResponse, error) {
	var returnValue *GroupLicenseUsersListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId":   c.enterpriseId,
		"groupLicenseId": c.groupLicenseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the IDs of the users who have been granted entitlements under the license.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.grouplicenseusers.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "groupLicenseId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupLicenseId": {
	//       "description": "The ID of the product the group license is for, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/groupLicenses/{groupLicenseId}/users",
	//   "response": {
	//     "$ref": "GroupLicenseUsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.delete":

type InstallsDeleteCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	installId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Requests to remove an app from a device. A call to get or
// list will still show the app as installed on the device until it is
// actually removed.

func (r *InstallsService) Delete(enterpriseId string, userId string, deviceId string, installId string) *InstallsDeleteCall {
	return &InstallsDeleteCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		installId:     installId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	}
}

func (c *InstallsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Requests to remove an app from a device. A call to get or list will still show the app as installed on the device until it is actually removed.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.installs.delete",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.get":

type InstallsGetCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	installId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves details of an installation of an app on a device.

func (r *InstallsService) Get(enterpriseId string, userId string, deviceId string, installId string) *InstallsGetCall {
	return &InstallsGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		installId:     installId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsGetCall) Fields(s ...googleapi.Field) *InstallsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstallsGetCall) Do() (*Install, error) {
	var returnValue *Install
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details of an installation of an app on a device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.installs.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.list":

type InstallsListCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the details of all apps installed on the specified
// device.

func (r *InstallsService) List(enterpriseId string, userId string, deviceId string) *InstallsListCall {
	return &InstallsListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsListCall) Fields(s ...googleapi.Field) *InstallsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstallsListCall) Do() (*InstallsListResponse, error) {
	var returnValue *InstallsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the details of all apps installed on the specified device.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.installs.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs",
	//   "response": {
	//     "$ref": "InstallsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.patch":

type InstallsPatchCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	installId     string
	install       *Install
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Requests to install the latest version of an app to a device.
// If the app is already installed then it is updated to the latest
// version if necessary. This method supports patch semantics.

func (r *InstallsService) Patch(enterpriseId string, userId string, deviceId string, installId string, install *Install) *InstallsPatchCall {
	return &InstallsPatchCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		installId:     installId,
		install:       install,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsPatchCall) Fields(s ...googleapi.Field) *InstallsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstallsPatchCall) Do() (*Install, error) {
	var returnValue *Install
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.install,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Requests to install the latest version of an app to a device. If the app is already installed then it is updated to the latest version if necessary. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "androidenterprise.installs.patch",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "request": {
	//     "$ref": "Install"
	//   },
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.installs.update":

type InstallsUpdateCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	deviceId      string
	installId     string
	install       *Install
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Requests to install the latest version of an app to a device.
// If the app is already installed then it is updated to the latest
// version if necessary.

func (r *InstallsService) Update(enterpriseId string, userId string, deviceId string, installId string, install *Install) *InstallsUpdateCall {
	return &InstallsUpdateCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		deviceId:      deviceId,
		installId:     installId,
		install:       install,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstallsUpdateCall) Fields(s ...googleapi.Field) *InstallsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *InstallsUpdateCall) Do() (*Install, error) {
	var returnValue *Install
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
		"deviceId":     c.deviceId,
		"installId":    c.installId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.install,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Requests to install the latest version of an app to a device. If the app is already installed then it is updated to the latest version if necessary.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.installs.update",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId",
	//     "deviceId",
	//     "installId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The Android ID of the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "installId": {
	//       "description": "The ID of the product represented by the install, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/devices/{deviceId}/installs/{installId}",
	//   "request": {
	//     "$ref": "Install"
	//   },
	//   "response": {
	//     "$ref": "Install"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.permissions.get":

type PermissionsGetCall struct {
	s             *Service
	permissionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves details of an Android app permission for display to an
// enterprise admin.

func (r *PermissionsService) Get(permissionId string) *PermissionsGetCall {
	return &PermissionsGetCall{
		s:             r.s,
		permissionId:  permissionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "permissions/{permissionId}",
	}
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de")
func (c *PermissionsGetCall) Language(language string) *PermissionsGetCall {
	c.params_.Set("language", fmt.Sprintf("%v", language))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PermissionsGetCall) Fields(s ...googleapi.Field) *PermissionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PermissionsGetCall) Do() (*Permission, error) {
	var returnValue *Permission
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"permissionId": c.permissionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details of an Android app permission for display to an enterprise admin.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.permissions.get",
	//   "parameterOrder": [
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\")",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The ID of the permission.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "permissions/{permissionId}",
	//   "response": {
	//     "$ref": "Permission"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.get":

type ProductsGetCall struct {
	s             *Service
	enterpriseId  string
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves details of a product for display to an enterprise
// admin.

func (r *ProductsService) Get(enterpriseId string, productId string) *ProductsGetCall {
	return &ProductsGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/products/{productId}",
	}
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de").
func (c *ProductsGetCall) Language(language string) *ProductsGetCall {
	c.params_.Set("language", fmt.Sprintf("%v", language))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetCall) Fields(s ...googleapi.Field) *ProductsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsGetCall) Do() (*Product, error) {
	var returnValue *Product
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details of a product for display to an enterprise admin.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\").",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product, e.g. \"app:com.google.android.gm\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}",
	//   "response": {
	//     "$ref": "Product"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.getAppRestrictionsSchema":

type ProductsGetAppRestrictionsSchemaCall struct {
	s             *Service
	enterpriseId  string
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetAppRestrictionsSchema: Retrieves the schema defining app
// restrictions configurable for this product. All products have a
// schema, but this may be empty if no app restrictions are defined.

func (r *ProductsService) GetAppRestrictionsSchema(enterpriseId string, productId string) *ProductsGetAppRestrictionsSchemaCall {
	return &ProductsGetAppRestrictionsSchemaCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/products/{productId}/appRestrictionsSchema",
	}
}

// Language sets the optional parameter "language": The BCP47 tag for
// the user's preferred language (e.g. "en-US", "de").
func (c *ProductsGetAppRestrictionsSchemaCall) Language(language string) *ProductsGetAppRestrictionsSchemaCall {
	c.params_.Set("language", fmt.Sprintf("%v", language))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetAppRestrictionsSchemaCall) Fields(s ...googleapi.Field) *ProductsGetAppRestrictionsSchemaCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsGetAppRestrictionsSchemaCall) Do() (*AppRestrictionsSchema, error) {
	var returnValue *AppRestrictionsSchema
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the schema defining app restrictions configurable for this product. All products have a schema, but this may be empty if no app restrictions are defined.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.getAppRestrictionsSchema",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The BCP47 tag for the user's preferred language (e.g. \"en-US\", \"de\").",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/appRestrictionsSchema",
	//   "response": {
	//     "$ref": "AppRestrictionsSchema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.getPermissions":

type ProductsGetPermissionsCall struct {
	s             *Service
	enterpriseId  string
	productId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetPermissions: Retrieves the Android app permissions required by
// this app.

func (r *ProductsService) GetPermissions(enterpriseId string, productId string) *ProductsGetPermissionsCall {
	return &ProductsGetPermissionsCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		productId:     productId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/products/{productId}/permissions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsGetPermissionsCall) Fields(s ...googleapi.Field) *ProductsGetPermissionsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsGetPermissionsCall) Do() (*ProductPermissions, error) {
	var returnValue *ProductPermissions
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the Android app permissions required by this app.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.products.getPermissions",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/permissions",
	//   "response": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.products.updatePermissions":

type ProductsUpdatePermissionsCall struct {
	s                  *Service
	enterpriseId       string
	productId          string
	productpermissions *ProductPermissions
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// UpdatePermissions: Updates the set of Android app permissions for
// this app that have been accepted by the enterprise.

func (r *ProductsService) UpdatePermissions(enterpriseId string, productId string, productpermissions *ProductPermissions) *ProductsUpdatePermissionsCall {
	return &ProductsUpdatePermissionsCall{
		s:                  r.s,
		enterpriseId:       enterpriseId,
		productId:          productId,
		productpermissions: productpermissions,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "enterprises/{enterpriseId}/products/{productId}/permissions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProductsUpdatePermissionsCall) Fields(s ...googleapi.Field) *ProductsUpdatePermissionsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProductsUpdatePermissionsCall) Do() (*ProductPermissions, error) {
	var returnValue *ProductPermissions
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"productId":    c.productId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.productpermissions,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates the set of Android app permissions for this app that have been accepted by the enterprise.",
	//   "httpMethod": "PUT",
	//   "id": "androidenterprise.products.updatePermissions",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "productId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The ID of the product.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/products/{productId}/permissions",
	//   "request": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "response": {
	//     "$ref": "ProductPermissions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.generateToken":

type UsersGenerateTokenCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GenerateToken: Generates a token (activation code) to allow this user
// to configure their work account in the Android Setup Wizard. Revokes
// any previously generated token.

func (r *UsersService) GenerateToken(enterpriseId string, userId string) *UsersGenerateTokenCall {
	return &UsersGenerateTokenCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/token",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGenerateTokenCall) Fields(s ...googleapi.Field) *UsersGenerateTokenCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersGenerateTokenCall) Do() (*UserToken, error) {
	var returnValue *UserToken
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generates a token (activation code) to allow this user to configure their work account in the Android Setup Wizard. Revokes any previously generated token.",
	//   "httpMethod": "POST",
	//   "id": "androidenterprise.users.generateToken",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/token",
	//   "response": {
	//     "$ref": "UserToken"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.get":

type UsersGetCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a user's details.

func (r *UsersService) Get(enterpriseId string, userId string) *UsersGetCall {
	return &UsersGetCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersGetCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a user's details.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.users.get",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.list":

type UsersListCall struct {
	s             *Service
	enterpriseId  string
	email         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Looks up a user by email address.

func (r *UsersService) List(enterpriseId string, email string) *UsersListCall {
	return &UsersListCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		email:         email,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersListCall) Do() (*UsersListResponse, error) {
	var returnValue *UsersListResponse
	c.params_.Set("email", fmt.Sprintf("%v", c.email))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Looks up a user by email address.",
	//   "httpMethod": "GET",
	//   "id": "androidenterprise.users.list",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "email"
	//   ],
	//   "parameters": {
	//     "email": {
	//       "description": "The exact primary email address of the user to look up.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users",
	//   "response": {
	//     "$ref": "UsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

// method id "androidenterprise.users.revokeToken":

type UsersRevokeTokenCall struct {
	s             *Service
	enterpriseId  string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// RevokeToken: Revokes a previously generated token (activation code)
// for the user.

func (r *UsersService) RevokeToken(enterpriseId string, userId string) *UsersRevokeTokenCall {
	return &UsersRevokeTokenCall{
		s:             r.s,
		enterpriseId:  enterpriseId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "enterprises/{enterpriseId}/users/{userId}/token",
	}
}

func (c *UsersRevokeTokenCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"enterpriseId": c.enterpriseId,
		"userId":       c.userId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Revokes a previously generated token (activation code) for the user.",
	//   "httpMethod": "DELETE",
	//   "id": "androidenterprise.users.revokeToken",
	//   "parameterOrder": [
	//     "enterpriseId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "enterpriseId": {
	//       "description": "The ID of the enterprise.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "enterprises/{enterpriseId}/users/{userId}/token",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidenterprise"
	//   ]
	// }

}

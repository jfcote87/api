// Package reseller provides access to the Enterprise Apps Reseller API.
//
// See https://developers.google.com/google-apps/reseller/
//
// Usage example:
//
//   import "github.com/jfcote87/api/reseller/v1"
//   ...
//   resellerService, err := reseller.New(oauthHttpClient)
package reseller

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

const apiId = "reseller:v1"
const apiName = "reseller"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/apps/reseller/v1/"}

// OAuth2 scopes used by this API.
const (
	// Manage users on your domain
	AppsOrderScope = "https://www.googleapis.com/auth/apps.order"

	// Manage users on your domain
	AppsOrderReadonlyScope = "https://www.googleapis.com/auth/apps.order.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Customers = NewCustomersService(s)
	s.Subscriptions = NewSubscriptionsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Customers *CustomersService

	Subscriptions *SubscriptionsService
}

func NewCustomersService(s *Service) *CustomersService {
	rs := &CustomersService{s: s}
	return rs
}

type CustomersService struct {
	s *Service
}

func NewSubscriptionsService(s *Service) *SubscriptionsService {
	rs := &SubscriptionsService{s: s}
	return rs
}

type SubscriptionsService struct {
	s *Service
}

type Address struct {
	// AddressLine1: Address line 1 of the address.
	AddressLine1 string `json:"addressLine1,omitempty"`

	// AddressLine2: Address line 2 of the address.
	AddressLine2 string `json:"addressLine2,omitempty"`

	// AddressLine3: Address line 3 of the address.
	AddressLine3 string `json:"addressLine3,omitempty"`

	// ContactName: Name of the contact person.
	ContactName string `json:"contactName,omitempty"`

	// CountryCode: ISO 3166 country code.
	CountryCode string `json:"countryCode,omitempty"`

	// Kind: Identifies the resource as a customer address.
	Kind string `json:"kind,omitempty"`

	// Locality: Name of the locality. This is in accordance with -
	// http://portablecontacts.net/draft-spec.html#address_element.
	Locality string `json:"locality,omitempty"`

	// OrganizationName: Name of the organization.
	OrganizationName string `json:"organizationName,omitempty"`

	// PostalCode: The postal code. This is in accordance with -
	// http://portablecontacts.net/draft-spec.html#address_element.
	PostalCode string `json:"postalCode,omitempty"`

	// Region: Name of the region. This is in accordance with -
	// http://portablecontacts.net/draft-spec.html#address_element.
	Region string `json:"region,omitempty"`
}

type ChangePlanRequest struct {
	// Kind: Identifies the resource as a subscription change plan request.
	Kind string `json:"kind,omitempty"`

	// PlanName: Name of the plan to change to.
	PlanName string `json:"planName,omitempty"`

	// PurchaseOrderId: Purchase order id for your order tracking purposes.
	PurchaseOrderId string `json:"purchaseOrderId,omitempty"`

	// Seats: Number/Limit of seats in the new plan.
	Seats *Seats `json:"seats,omitempty"`
}

type Customer struct {
	// AlternateEmail: The alternate email of the customer.
	AlternateEmail string `json:"alternateEmail,omitempty"`

	// CustomerDomain: The domain name of the customer.
	CustomerDomain string `json:"customerDomain,omitempty"`

	// CustomerId: The id of the customer.
	CustomerId string `json:"customerId,omitempty"`

	// Kind: Identifies the resource as a customer.
	Kind string `json:"kind,omitempty"`

	// PhoneNumber: The phone number of the customer.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// PostalAddress: The postal address of the customer.
	PostalAddress *Address `json:"postalAddress,omitempty"`

	// ResourceUiUrl: Ui url for customer resource.
	ResourceUiUrl string `json:"resourceUiUrl,omitempty"`
}

type RenewalSettings struct {
	// Kind: Identifies the resource as a subscription renewal setting.
	Kind string `json:"kind,omitempty"`

	// RenewalType: Subscription renewal type.
	RenewalType string `json:"renewalType,omitempty"`
}

type Seats struct {
	// Kind: Identifies the resource as a subscription change plan request.
	Kind string `json:"kind,omitempty"`

	// LicensedNumberOfSeats: Read-only field containing the current number
	// of licensed seats for FLEXIBLE Google-Apps subscriptions and
	// secondary subscriptions such as Google-Vault and Drive-storage.
	LicensedNumberOfSeats int64 `json:"licensedNumberOfSeats,omitempty"`

	// MaximumNumberOfSeats: Maximum number of seats that can be purchased.
	// This needs to be provided only for a non-commitment plan. For a
	// commitment plan it is decided by the contract.
	MaximumNumberOfSeats int64 `json:"maximumNumberOfSeats,omitempty"`

	// NumberOfSeats: Number of seats to purchase. This is applicable only
	// for a commitment plan.
	NumberOfSeats int64 `json:"numberOfSeats,omitempty"`
}

type Subscription struct {
	// BillingMethod: Billing method of this subscription.
	BillingMethod string `json:"billingMethod,omitempty"`

	// CreationTime: Creation time of this subscription in milliseconds
	// since Unix epoch.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// CustomerId: The id of the customer to whom the subscription belongs.
	CustomerId string `json:"customerId,omitempty"`

	// Kind: Identifies the resource as a Subscription.
	Kind string `json:"kind,omitempty"`

	// Plan: Plan details of the subscription
	Plan *SubscriptionPlan `json:"plan,omitempty"`

	// PurchaseOrderId: Purchase order id for your order tracking purposes.
	PurchaseOrderId string `json:"purchaseOrderId,omitempty"`

	// RenewalSettings: Renewal settings of the subscription.
	RenewalSettings *RenewalSettings `json:"renewalSettings,omitempty"`

	// ResourceUiUrl: Ui url for subscription resource.
	ResourceUiUrl string `json:"resourceUiUrl,omitempty"`

	// Seats: Number/Limit of seats in the new plan.
	Seats *Seats `json:"seats,omitempty"`

	// SkuId: Name of the sku for which this subscription is purchased.
	SkuId string `json:"skuId,omitempty"`

	// Status: Status of the subscription.
	Status string `json:"status,omitempty"`

	// SubscriptionId: The id of the subscription.
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// TransferInfo: Transfer related information for the subscription.
	TransferInfo *SubscriptionTransferInfo `json:"transferInfo,omitempty"`

	// TrialSettings: Trial Settings of the subscription.
	TrialSettings *SubscriptionTrialSettings `json:"trialSettings,omitempty"`
}

type SubscriptionPlan struct {
	// CommitmentInterval: Interval of the commitment if it is a commitment
	// plan.
	CommitmentInterval *SubscriptionPlanCommitmentInterval `json:"commitmentInterval,omitempty"`

	// IsCommitmentPlan: Whether the plan is a commitment plan or not.
	IsCommitmentPlan bool `json:"isCommitmentPlan,omitempty"`

	// PlanName: The plan name of this subscription's plan.
	PlanName string `json:"planName,omitempty"`
}

type SubscriptionPlanCommitmentInterval struct {
	// EndTime: End time of the commitment interval in milliseconds since
	// Unix epoch.
	EndTime int64 `json:"endTime,omitempty,string"`

	// StartTime: Start time of the commitment interval in milliseconds
	// since Unix epoch.
	StartTime int64 `json:"startTime,omitempty,string"`
}

type SubscriptionTransferInfo struct {
	MinimumTransferableSeats int64 `json:"minimumTransferableSeats,omitempty"`

	// TransferabilityExpirationTime: Time when transfer token or intent to
	// transfer will expire.
	TransferabilityExpirationTime int64 `json:"transferabilityExpirationTime,omitempty,string"`
}

type SubscriptionTrialSettings struct {
	// IsInTrial: Whether the subscription is in trial.
	IsInTrial bool `json:"isInTrial,omitempty"`

	// TrialEndTime: End time of the trial in milliseconds since Unix epoch.
	TrialEndTime int64 `json:"trialEndTime,omitempty,string"`
}

type Subscriptions struct {
	// Kind: Identifies the resource as a collection of subscriptions.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Subscriptions: The subscriptions in this page of results.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
}

// method id "reseller.customers.get":

type CustomersGetCall struct {
	s             *Service
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a customer resource if one exists and is owned by the
// reseller.

func (r *CustomersService) Get(customerId string) *CustomersGetCall {
	return &CustomersGetCall{
		s:             r.s,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customers/{customerId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersGetCall) Fields(s ...googleapi.Field) *CustomersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomersGetCall) Do() (*Customer, error) {
	var returnValue *Customer
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a customer resource if one exists and is owned by the reseller.",
	//   "httpMethod": "GET",
	//   "id": "reseller.customers.get",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.customers.insert":

type CustomersInsertCall struct {
	s             *Service
	customer      *Customer
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a customer resource if one does not already exist.

func (r *CustomersService) Insert(customer *Customer) *CustomersInsertCall {
	return &CustomersInsertCall{
		s:             r.s,
		customer:      customer,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customers",
	}
}

// CustomerAuthToken sets the optional parameter "customerAuthToken": An
// auth token needed for inserting a customer for which domain already
// exists. Can be generated at
// https://www.google.com/a/cpanel//TransferToken.
func (c *CustomersInsertCall) CustomerAuthToken(customerAuthToken string) *CustomersInsertCall {
	c.params_.Set("customerAuthToken", fmt.Sprintf("%v", customerAuthToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersInsertCall) Fields(s ...googleapi.Field) *CustomersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomersInsertCall) Do() (*Customer, error) {
	var returnValue *Customer
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.customer,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a customer resource if one does not already exist.",
	//   "httpMethod": "POST",
	//   "id": "reseller.customers.insert",
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "An auth token needed for inserting a customer for which domain already exists. Can be generated at https://www.google.com/a/cpanel//TransferToken. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.customers.patch":

type CustomersPatchCall struct {
	s             *Service
	customerId    string
	customer      *Customer
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update a customer resource if one it exists and is owned by
// the reseller. This method supports patch semantics.

func (r *CustomersService) Patch(customerId string, customer *Customer) *CustomersPatchCall {
	return &CustomersPatchCall{
		s:             r.s,
		customerId:    customerId,
		customer:      customer,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customers/{customerId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersPatchCall) Fields(s ...googleapi.Field) *CustomersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomersPatchCall) Do() (*Customer, error) {
	var returnValue *Customer
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.customer,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update a customer resource if one it exists and is owned by the reseller. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "reseller.customers.patch",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.customers.update":

type CustomersUpdateCall struct {
	s             *Service
	customerId    string
	customer      *Customer
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update a customer resource if one it exists and is owned by
// the reseller.

func (r *CustomersService) Update(customerId string, customer *Customer) *CustomersUpdateCall {
	return &CustomersUpdateCall{
		s:             r.s,
		customerId:    customerId,
		customer:      customer,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customers/{customerId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersUpdateCall) Fields(s ...googleapi.Field) *CustomersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomersUpdateCall) Do() (*Customer, error) {
	var returnValue *Customer
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.customer,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update a customer resource if one it exists and is owned by the reseller.",
	//   "httpMethod": "PUT",
	//   "id": "reseller.customers.update",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.activate":

type SubscriptionsActivateCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Activate: Activates a subscription previously suspended by the
// reseller

func (r *SubscriptionsService) Activate(customerId string, subscriptionId string) *SubscriptionsActivateCall {
	return &SubscriptionsActivateCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}/activate",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsActivateCall) Fields(s ...googleapi.Field) *SubscriptionsActivateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsActivateCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Activates a subscription previously suspended by the reseller",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.activate",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/activate",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changePlan":

type SubscriptionsChangePlanCall struct {
	s                 *Service
	customerId        string
	subscriptionId    string
	changeplanrequest *ChangePlanRequest
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// ChangePlan: Changes the plan of a subscription

func (r *SubscriptionsService) ChangePlan(customerId string, subscriptionId string, changeplanrequest *ChangePlanRequest) *SubscriptionsChangePlanCall {
	return &SubscriptionsChangePlanCall{
		s:                 r.s,
		customerId:        customerId,
		subscriptionId:    subscriptionId,
		changeplanrequest: changeplanrequest,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "customers/{customerId}/subscriptions/{subscriptionId}/changePlan",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangePlanCall) Fields(s ...googleapi.Field) *SubscriptionsChangePlanCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsChangePlanCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.changeplanrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Changes the plan of a subscription",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changePlan",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changePlan",
	//   "request": {
	//     "$ref": "ChangePlanRequest"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changeRenewalSettings":

type SubscriptionsChangeRenewalSettingsCall struct {
	s               *Service
	customerId      string
	subscriptionId  string
	renewalsettings *RenewalSettings
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// ChangeRenewalSettings: Changes the renewal settings of a subscription

func (r *SubscriptionsService) ChangeRenewalSettings(customerId string, subscriptionId string, renewalsettings *RenewalSettings) *SubscriptionsChangeRenewalSettingsCall {
	return &SubscriptionsChangeRenewalSettingsCall{
		s:               r.s,
		customerId:      customerId,
		subscriptionId:  subscriptionId,
		renewalsettings: renewalsettings,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "customers/{customerId}/subscriptions/{subscriptionId}/changeRenewalSettings",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangeRenewalSettingsCall) Fields(s ...googleapi.Field) *SubscriptionsChangeRenewalSettingsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsChangeRenewalSettingsCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.renewalsettings,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Changes the renewal settings of a subscription",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changeRenewalSettings",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changeRenewalSettings",
	//   "request": {
	//     "$ref": "RenewalSettings"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changeSeats":

type SubscriptionsChangeSeatsCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	seats          *Seats
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// ChangeSeats: Changes the seats configuration of a subscription

func (r *SubscriptionsService) ChangeSeats(customerId string, subscriptionId string, seats *Seats) *SubscriptionsChangeSeatsCall {
	return &SubscriptionsChangeSeatsCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		seats:          seats,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}/changeSeats",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangeSeatsCall) Fields(s ...googleapi.Field) *SubscriptionsChangeSeatsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsChangeSeatsCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.seats,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Changes the seats configuration of a subscription",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changeSeats",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changeSeats",
	//   "request": {
	//     "$ref": "Seats"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.delete":

type SubscriptionsDeleteCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	deletionType   string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Delete: Cancels/Downgrades a subscription.

func (r *SubscriptionsService) Delete(customerId string, subscriptionId string, deletionType string) *SubscriptionsDeleteCall {
	return &SubscriptionsDeleteCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		deletionType:   deletionType,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}",
	}
}

func (c *SubscriptionsDeleteCall) Do() error {
	c.params_.Set("deletionType", fmt.Sprintf("%v", c.deletionType))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Cancels/Downgrades a subscription.",
	//   "httpMethod": "DELETE",
	//   "id": "reseller.subscriptions.delete",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId",
	//     "deletionType"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deletionType": {
	//       "description": "Whether the subscription is to be fully cancelled or downgraded",
	//       "enum": [
	//         "cancel",
	//         "downgrade",
	//         "suspend",
	//         "transfer_to_direct"
	//       ],
	//       "enumDescriptions": [
	//         "Cancels the subscription immediately",
	//         "Downgrades a Google Apps for Business subscription to Google Apps",
	//         "Suspends the subscriptions for 4 days before cancelling it",
	//         "Transfers a subscription directly to Google"
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.get":

type SubscriptionsGetCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Get: Gets a subscription of the customer.

func (r *SubscriptionsService) Get(customerId string, subscriptionId string) *SubscriptionsGetCall {
	return &SubscriptionsGetCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsGetCall) Fields(s ...googleapi.Field) *SubscriptionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsGetCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a subscription of the customer.",
	//   "httpMethod": "GET",
	//   "id": "reseller.subscriptions.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.subscriptions.insert":

type SubscriptionsInsertCall struct {
	s             *Service
	customerId    string
	subscription  *Subscription
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates/Transfers a subscription for the customer.

func (r *SubscriptionsService) Insert(customerId string, subscription *Subscription) *SubscriptionsInsertCall {
	return &SubscriptionsInsertCall{
		s:             r.s,
		customerId:    customerId,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customers/{customerId}/subscriptions",
	}
}

// CustomerAuthToken sets the optional parameter "customerAuthToken": An
// auth token needed for transferring a subscription. Can be generated
// at https://www.google.com/a/cpanel/customer-domain/TransferToken.
func (c *SubscriptionsInsertCall) CustomerAuthToken(customerAuthToken string) *SubscriptionsInsertCall {
	c.params_.Set("customerAuthToken", fmt.Sprintf("%v", customerAuthToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsInsertCall) Fields(s ...googleapi.Field) *SubscriptionsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsInsertCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.subscription,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates/Transfers a subscription for the customer.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.insert",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "An auth token needed for transferring a subscription. Can be generated at https://www.google.com/a/cpanel/customer-domain/TransferToken. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.list":

type SubscriptionsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists subscriptions of a reseller, optionally filtered by a
// customer name prefix.

func (r *SubscriptionsService) List() *SubscriptionsListCall {
	return &SubscriptionsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "subscriptions",
	}
}

// CustomerAuthToken sets the optional parameter "customerAuthToken": An
// auth token needed if the customer is not a resold customer of this
// reseller. Can be generated at
// https://www.google.com/a/cpanel/customer-domain/TransferToken.
func (c *SubscriptionsListCall) CustomerAuthToken(customerAuthToken string) *SubscriptionsListCall {
	c.params_.Set("customerAuthToken", fmt.Sprintf("%v", customerAuthToken))
	return c
}

// CustomerId sets the optional parameter "customerId": Id of the
// Customer
func (c *SubscriptionsListCall) CustomerId(customerId string) *SubscriptionsListCall {
	c.params_.Set("customerId", fmt.Sprintf("%v", customerId))
	return c
}

// CustomerNamePrefix sets the optional parameter "customerNamePrefix":
// Prefix of the customer's domain name by which the subscriptions
// should be filtered. Optional
func (c *SubscriptionsListCall) CustomerNamePrefix(customerNamePrefix string) *SubscriptionsListCall {
	c.params_.Set("customerNamePrefix", fmt.Sprintf("%v", customerNamePrefix))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *SubscriptionsListCall) MaxResults(maxResults int64) *SubscriptionsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *SubscriptionsListCall) PageToken(pageToken string) *SubscriptionsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsListCall) Fields(s ...googleapi.Field) *SubscriptionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsListCall) Do() (*Subscriptions, error) {
	var returnValue *Subscriptions
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists subscriptions of a reseller, optionally filtered by a customer name prefix.",
	//   "httpMethod": "GET",
	//   "id": "reseller.subscriptions.list",
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "An auth token needed if the customer is not a resold customer of this reseller. Can be generated at https://www.google.com/a/cpanel/customer-domain/TransferToken.Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerNamePrefix": {
	//       "description": "Prefix of the customer's domain name by which the subscriptions should be filtered. Optional",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions",
	//   "response": {
	//     "$ref": "Subscriptions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.subscriptions.startPaidService":

type SubscriptionsStartPaidServiceCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// StartPaidService: Starts paid service of a trial subscription

func (r *SubscriptionsService) StartPaidService(customerId string, subscriptionId string) *SubscriptionsStartPaidServiceCall {
	return &SubscriptionsStartPaidServiceCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}/startPaidService",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsStartPaidServiceCall) Fields(s ...googleapi.Field) *SubscriptionsStartPaidServiceCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsStartPaidServiceCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Starts paid service of a trial subscription",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.startPaidService",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/startPaidService",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.suspend":

type SubscriptionsSuspendCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Suspend: Suspends an active subscription

func (r *SubscriptionsService) Suspend(customerId string, subscriptionId string) *SubscriptionsSuspendCall {
	return &SubscriptionsSuspendCall{
		s:              r.s,
		customerId:     customerId,
		subscriptionId: subscriptionId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customers/{customerId}/subscriptions/{subscriptionId}/suspend",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsSuspendCall) Fields(s ...googleapi.Field) *SubscriptionsSuspendCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsSuspendCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Suspends an active subscription",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.suspend",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Id of the Customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "Id of the subscription, which is unique for a customer",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/suspend",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

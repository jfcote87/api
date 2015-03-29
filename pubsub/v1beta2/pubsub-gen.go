// Package pubsub provides access to the Google Cloud Pub/Sub API.
//
// Usage example:
//
//   import "github.com/jfcote87/api/pubsub/v1beta2"
//   ...
//   pubsubService, err := pubsub.New(oauthHttpClient)
package pubsub

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

const apiId = "pubsub:v1beta2"
const apiName = "pubsub"
const apiVersion = "v1beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "pubsub.googleapis.com", Path: "/v1beta2/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage Pub/Sub topics and subscriptions
	PubsubScope = "https://www.googleapis.com/auth/pubsub"
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
	rs.Subscriptions = NewProjectsSubscriptionsService(s)
	rs.Topics = NewProjectsTopicsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Subscriptions *ProjectsSubscriptionsService

	Topics *ProjectsTopicsService
}

func NewProjectsSubscriptionsService(s *Service) *ProjectsSubscriptionsService {
	rs := &ProjectsSubscriptionsService{s: s}
	return rs
}

type ProjectsSubscriptionsService struct {
	s *Service
}

func NewProjectsTopicsService(s *Service) *ProjectsTopicsService {
	rs := &ProjectsTopicsService{s: s}
	rs.Subscriptions = NewProjectsTopicsSubscriptionsService(s)
	return rs
}

type ProjectsTopicsService struct {
	s *Service

	Subscriptions *ProjectsTopicsSubscriptionsService
}

func NewProjectsTopicsSubscriptionsService(s *Service) *ProjectsTopicsSubscriptionsService {
	rs := &ProjectsTopicsSubscriptionsService{s: s}
	return rs
}

type ProjectsTopicsSubscriptionsService struct {
	s *Service
}

type AcknowledgeRequest struct {
	AckIds []string `json:"ackIds,omitempty"`
}

type Empty struct {
}

type ListSubscriptionsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
}

type ListTopicSubscriptionsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Subscriptions []string `json:"subscriptions,omitempty"`
}

type ListTopicsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Topics []*Topic `json:"topics,omitempty"`
}

type ModifyAckDeadlineRequest struct {
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	AckId string `json:"ackId,omitempty"`
}

type ModifyPushConfigRequest struct {
	PushConfig *PushConfig `json:"pushConfig,omitempty"`
}

type PublishRequest struct {
	Messages []*PubsubMessage `json:"messages,omitempty"`
}

type PublishResponse struct {
	MessageIds []string `json:"messageIds,omitempty"`
}

type PubsubMessage struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	Data string `json:"data,omitempty"`

	MessageId string `json:"messageId,omitempty"`
}

type PullRequest struct {
	MaxMessages int64 `json:"maxMessages,omitempty"`

	ReturnImmediately bool `json:"returnImmediately,omitempty"`
}

type PullResponse struct {
	ReceivedMessages []*ReceivedMessage `json:"receivedMessages,omitempty"`
}

type PushConfig struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	PushEndpoint string `json:"pushEndpoint,omitempty"`
}

type ReceivedMessage struct {
	AckId string `json:"ackId,omitempty"`

	Message *PubsubMessage `json:"message,omitempty"`
}

type Subscription struct {
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	Name string `json:"name,omitempty"`

	PushConfig *PushConfig `json:"pushConfig,omitempty"`

	Topic string `json:"topic,omitempty"`
}

type Topic struct {
	Name string `json:"name,omitempty"`
}

// method id "pubsub.projects.subscriptions.acknowledge":

type ProjectsSubscriptionsAcknowledgeCall struct {
	s                  *Service
	subscription       string
	acknowledgerequest *AcknowledgeRequest
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Acknowledge: Acknowledges the messages associated with the ack tokens
// in the AcknowledgeRequest. The Pub/Sub system can remove the relevant
// messages from the subscription. Acknowledging a message whose ack
// deadline has expired may succeed, but such a message may be
// redelivered later. Acknowledging a message more than once will not
// result in an error.

func (r *ProjectsSubscriptionsService) Acknowledge(subscription string, acknowledgerequest *AcknowledgeRequest) *ProjectsSubscriptionsAcknowledgeCall {
	return &ProjectsSubscriptionsAcknowledgeCall{
		s:                  r.s,
		subscription:       subscription,
		acknowledgerequest: acknowledgerequest,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "{+subscription}:acknowledge",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsAcknowledgeCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsAcknowledgeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsAcknowledgeCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.acknowledgerequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Acknowledges the messages associated with the ack tokens in the AcknowledgeRequest. The Pub/Sub system can remove the relevant messages from the subscription. Acknowledging a message whose ack deadline has expired may succeed, but such a message may be redelivered later. Acknowledging a message more than once will not result in an error.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.acknowledge",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:acknowledge",
	//   "request": {
	//     "$ref": "AcknowledgeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.create":

type ProjectsSubscriptionsCreateCall struct {
	s             *Service
	name          string
	subscription  *Subscription
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a subscription to a given topic for a given
// subscriber. If the subscription already exists, returns
// ALREADY_EXISTS. If the corresponding topic doesn't exist, returns
// NOT_FOUND. If the name is not provided in the request, the server
// will assign a random name for this subscription on the same project
// as the topic.

func (r *ProjectsSubscriptionsService) Create(name string, subscription *Subscription) *ProjectsSubscriptionsCreateCall {
	return &ProjectsSubscriptionsCreateCall{
		s:             r.s,
		name:          name,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+name}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsCreateCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsCreateCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"name": c.name,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.subscription,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a subscription to a given topic for a given subscriber. If the subscription already exists, returns ALREADY_EXISTS. If the corresponding topic doesn't exist, returns NOT_FOUND. If the name is not provided in the request, the server will assign a random name for this subscription on the same project as the topic.",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.subscriptions.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.delete":

type ProjectsSubscriptionsDeleteCall struct {
	s             *Service
	subscription  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing subscription. All pending messages in the
// subscription are immediately dropped. Calls to Pull after deletion
// will return NOT_FOUND. After a subscription is deleted, a new one may
// be created with the same name, but the new one has no association
// with the old subscription, or its topic unless the same topic is
// specified.

func (r *ProjectsSubscriptionsService) Delete(subscription string) *ProjectsSubscriptionsDeleteCall {
	return &ProjectsSubscriptionsDeleteCall{
		s:             r.s,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+subscription}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing subscription. All pending messages in the subscription are immediately dropped. Calls to Pull after deletion will return NOT_FOUND. After a subscription is deleted, a new one may be created with the same name, but the new one has no association with the old subscription, or its topic unless the same topic is specified.",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.subscriptions.delete",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.get":

type ProjectsSubscriptionsGetCall struct {
	s             *Service
	subscription  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets the configuration details of a subscription.

func (r *ProjectsSubscriptionsService) Get(subscription string) *ProjectsSubscriptionsGetCall {
	return &ProjectsSubscriptionsGetCall{
		s:             r.s,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+subscription}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsGetCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsGetCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the configuration details of a subscription.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.get",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.list":

type ProjectsSubscriptionsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists matching subscriptions.

func (r *ProjectsSubscriptionsService) List(project string) *ProjectsSubscriptionsListCall {
	return &ProjectsSubscriptionsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+project}/subscriptions",
	}
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsSubscriptionsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsSubscriptionsListCall) PageToken(pageToken string) *ProjectsSubscriptionsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsListCall) Do() (*ListSubscriptionsResponse, error) {
	var returnValue *ListSubscriptionsResponse
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
	//   "description": "Lists matching subscriptions.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+project}/subscriptions",
	//   "response": {
	//     "$ref": "ListSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyAckDeadline":

type ProjectsSubscriptionsModifyAckDeadlineCall struct {
	s                        *Service
	subscription             string
	modifyackdeadlinerequest *ModifyAckDeadlineRequest
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// ModifyAckDeadline: Modifies the ack deadline for a specific message.
// This method is useful to indicate that more time is needed to process
// a message by the subscriber, or to make the message available for
// redelivery if the processing was interrupted.

func (r *ProjectsSubscriptionsService) ModifyAckDeadline(subscription string, modifyackdeadlinerequest *ModifyAckDeadlineRequest) *ProjectsSubscriptionsModifyAckDeadlineCall {
	return &ProjectsSubscriptionsModifyAckDeadlineCall{
		s:                        r.s,
		subscription:             subscription,
		modifyackdeadlinerequest: modifyackdeadlinerequest,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "{+subscription}:modifyAckDeadline",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyAckDeadlineCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.modifyackdeadlinerequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modifies the ack deadline for a specific message. This method is useful to indicate that more time is needed to process a message by the subscriber, or to make the message available for redelivery if the processing was interrupted.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyAckDeadline",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:modifyAckDeadline",
	//   "request": {
	//     "$ref": "ModifyAckDeadlineRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyPushConfig":

type ProjectsSubscriptionsModifyPushConfigCall struct {
	s                       *Service
	subscription            string
	modifypushconfigrequest *ModifyPushConfigRequest
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// ModifyPushConfig: Modifies the PushConfig for a specified
// subscription. This may be used to change a push subscription to a
// pull one (signified by an empty PushConfig) or vice versa, or change
// the endpoint URL and other attributes of a push subscription.
// Messages will accumulate for delivery continuously through the call
// regardless of changes to the PushConfig.

func (r *ProjectsSubscriptionsService) ModifyPushConfig(subscription string, modifypushconfigrequest *ModifyPushConfigRequest) *ProjectsSubscriptionsModifyPushConfigCall {
	return &ProjectsSubscriptionsModifyPushConfigCall{
		s:                       r.s,
		subscription:            subscription,
		modifypushconfigrequest: modifypushconfigrequest,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "{+subscription}:modifyPushConfig",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyPushConfigCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyPushConfigCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsModifyPushConfigCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.modifypushconfigrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modifies the PushConfig for a specified subscription. This may be used to change a push subscription to a pull one (signified by an empty PushConfig) or vice versa, or change the endpoint URL and other attributes of a push subscription. Messages will accumulate for delivery continuously through the call regardless of changes to the PushConfig.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyPushConfig",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:modifyPushConfig",
	//   "request": {
	//     "$ref": "ModifyPushConfigRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.pull":

type ProjectsSubscriptionsPullCall struct {
	s             *Service
	subscription  string
	pullrequest   *PullRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Pull: Pulls messages from the server. Returns an empty list if there
// are no messages available in the backlog. The server may return
// UNAVAILABLE if there are too many concurrent pull requests pending
// for the given subscription.

func (r *ProjectsSubscriptionsService) Pull(subscription string, pullrequest *PullRequest) *ProjectsSubscriptionsPullCall {
	return &ProjectsSubscriptionsPullCall{
		s:             r.s,
		subscription:  subscription,
		pullrequest:   pullrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+subscription}:pull",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsPullCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsPullCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsSubscriptionsPullCall) Do() (*PullResponse, error) {
	var returnValue *PullResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"subscription": c.subscription,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.pullrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Pulls messages from the server. Returns an empty list if there are no messages available in the backlog. The server may return UNAVAILABLE if there are too many concurrent pull requests pending for the given subscription.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.pull",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+subscription}:pull",
	//   "request": {
	//     "$ref": "PullRequest"
	//   },
	//   "response": {
	//     "$ref": "PullResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.create":

type ProjectsTopicsCreateCall struct {
	s             *Service
	name          string
	topic         *Topic
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates the given topic with the given name.

func (r *ProjectsTopicsService) Create(name string, topic *Topic) *ProjectsTopicsCreateCall {
	return &ProjectsTopicsCreateCall{
		s:             r.s,
		name:          name,
		topic:         topic,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+name}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsCreateCall) Fields(s ...googleapi.Field) *ProjectsTopicsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsCreateCall) Do() (*Topic, error) {
	var returnValue *Topic
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"name": c.name,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.topic,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates the given topic with the given name.",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.topics.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Topic"
	//   },
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.delete":

type ProjectsTopicsDeleteCall struct {
	s             *Service
	topic         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the topic with the given name. All subscriptions to
// this topic are detached from the topic. Returns NOT_FOUND if the
// topic does not exist. After a topic is deleted, a new topic may be
// created with the same name; this is an entirely new topic with none
// of the old configuration or subscriptions.

func (r *ProjectsTopicsService) Delete(topic string) *ProjectsTopicsDeleteCall {
	return &ProjectsTopicsDeleteCall{
		s:             r.s,
		topic:         topic,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+topic}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsDeleteCall) Fields(s ...googleapi.Field) *ProjectsTopicsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"topic": c.topic,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the topic with the given name. All subscriptions to this topic are detached from the topic. Returns NOT_FOUND if the topic does not exist. After a topic is deleted, a new topic may be created with the same name; this is an entirely new topic with none of the old configuration or subscriptions.",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.topics.delete",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.get":

type ProjectsTopicsGetCall struct {
	s             *Service
	topic         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets the configuration of a topic.

func (r *ProjectsTopicsService) Get(topic string) *ProjectsTopicsGetCall {
	return &ProjectsTopicsGetCall{
		s:             r.s,
		topic:         topic,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+topic}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsGetCall) Fields(s ...googleapi.Field) *ProjectsTopicsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsGetCall) Do() (*Topic, error) {
	var returnValue *Topic
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"topic": c.topic,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the configuration of a topic.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.get",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}",
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.list":

type ProjectsTopicsListCall struct {
	s             *Service
	project       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists matching topics.

func (r *ProjectsTopicsService) List(project string) *ProjectsTopicsListCall {
	return &ProjectsTopicsListCall{
		s:             r.s,
		project:       project,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+project}/topics",
	}
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsTopicsListCall) PageSize(pageSize int64) *ProjectsTopicsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsTopicsListCall) PageToken(pageToken string) *ProjectsTopicsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsListCall) Do() (*ListTopicsResponse, error) {
	var returnValue *ListTopicsResponse
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
	//   "description": "Lists matching topics.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+project}/topics",
	//   "response": {
	//     "$ref": "ListTopicsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.publish":

type ProjectsTopicsPublishCall struct {
	s              *Service
	topic          string
	publishrequest *PublishRequest
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Publish: Adds one or more messages to the topic. Returns NOT_FOUND if
// the topic does not exist.

func (r *ProjectsTopicsService) Publish(topic string, publishrequest *PublishRequest) *ProjectsTopicsPublishCall {
	return &ProjectsTopicsPublishCall{
		s:              r.s,
		topic:          topic,
		publishrequest: publishrequest,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "{+topic}:publish",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsPublishCall) Fields(s ...googleapi.Field) *ProjectsTopicsPublishCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsPublishCall) Do() (*PublishResponse, error) {
	var returnValue *PublishResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"topic": c.topic,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.publishrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds one or more messages to the topic. Returns NOT_FOUND if the topic does not exist.",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.topics.publish",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}:publish",
	//   "request": {
	//     "$ref": "PublishRequest"
	//   },
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.subscriptions.list":

type ProjectsTopicsSubscriptionsListCall struct {
	s             *Service
	topic         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists the name of the subscriptions for this topic.

func (r *ProjectsTopicsSubscriptionsService) List(topic string) *ProjectsTopicsSubscriptionsListCall {
	return &ProjectsTopicsSubscriptionsListCall{
		s:             r.s,
		topic:         topic,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{+topic}/subscriptions",
	}
}

// PageSize sets the optional parameter "pageSize":
func (c *ProjectsTopicsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsTopicsSubscriptionsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ProjectsTopicsSubscriptionsListCall) PageToken(pageToken string) *ProjectsTopicsSubscriptionsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsSubscriptionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsTopicsSubscriptionsListCall) Do() (*ListTopicSubscriptionsResponse, error) {
	var returnValue *ListTopicSubscriptionsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"topic": c.topic,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists the name of the subscriptions for this topic.",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.subscriptions.list",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "topic": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+topic}/subscriptions",
	//   "response": {
	//     "$ref": "ListTopicSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// Package admin provides access to the Email Migration API v2.
//
// See https://developers.google.com/admin-sdk/email-migration/v2/
//
// Usage example:
//
//   import "github.com/jfcote87/api/admin/email_migration_v2"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin

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

const apiId = "admin:email_migration_v2"
const apiName = "admin"
const apiVersion = "email_migration_v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/email/v2/users/"}

// OAuth2 scopes used by this API.
const (
	// Manage email messages of users on your domain
	EmailMigrationScope = "https://www.googleapis.com/auth/email.migration"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Mail = NewMailService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Mail *MailService
}

func NewMailService(s *Service) *MailService {
	rs := &MailService{s: s}
	return rs
}

type MailService struct {
	s *Service
}

type MailItem struct {
	// IsDeleted: Boolean indicating if the mail is deleted (used in Vault)
	IsDeleted bool `json:"isDeleted,omitempty"`

	// IsDraft: Boolean indicating if the mail is draft
	IsDraft bool `json:"isDraft,omitempty"`

	// IsInbox: Boolean indicating if the mail is in inbox
	IsInbox bool `json:"isInbox,omitempty"`

	// IsSent: Boolean indicating if the mail is in 'sent mails'
	IsSent bool `json:"isSent,omitempty"`

	// IsStarred: Boolean indicating if the mail is starred
	IsStarred bool `json:"isStarred,omitempty"`

	// IsTrash: Boolean indicating if the mail is in trash
	IsTrash bool `json:"isTrash,omitempty"`

	// IsUnread: Boolean indicating if the mail is unread
	IsUnread bool `json:"isUnread,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Labels: List of labels (strings)
	Labels []string `json:"labels,omitempty"`
}

// method id "emailMigration.mail.insert":

type MailInsertCall struct {
	s             *Service
	userKey       string
	mailitem      *MailItem
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Insert Mail into Google's Gmail backends

func (r *MailService) Insert(userKey string, mailitem *MailItem) *MailInsertCall {
	return &MailInsertCall{
		s:             r.s,
		userKey:       userKey,
		mailitem:      mailitem,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userKey}/mail",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *MailInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *MailInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/email/v2/users/{userKey}/mail"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/email/v2/users/{userKey}/mail"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *MailInsertCall) Media(r io.Reader) *MailInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/email/v2/users/{userKey}/mail"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *MailInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *MailInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/email/v2/users/{userKey}/mail"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *MailInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *MailInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

func (c *MailInsertCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.mailitem,
	}

	return c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Insert Mail into Google's Gmail backends",
	//   "httpMethod": "POST",
	//   "id": "emailMigration.mail.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "message/rfc822"
	//     ],
	//     "maxSize": "35MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/email/v2/users/{userKey}/mail"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/email/v2/users/{userKey}/mail"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "The email or immutable id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userKey}/mail",
	//   "request": {
	//     "$ref": "MailItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/email.migration"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// Package groupsmigration provides access to the Groups Migration API.
//
// See https://developers.google.com/google-apps/groups-migration/
//
// Usage example:
//
//   import "github.com/jfcote87/api/groupsmigration/v1"
//   ...
//   groupsmigrationService, err := groupsmigration.New(oauthHttpClient)
package groupsmigration

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

const apiId = "groupsmigration:v1"
const apiName = "groupsmigration"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/groups/v1/groups/"}

// OAuth2 scopes used by this API.
const (
	// Manage messages in groups on your domain
	AppsGroupsMigrationScope = "https://www.googleapis.com/auth/apps.groups.migration"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Archive = NewArchiveService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Archive *ArchiveService
}

func NewArchiveService(s *Service) *ArchiveService {
	rs := &ArchiveService{s: s}
	return rs
}

type ArchiveService struct {
	s *Service
}

type Groups struct {
	// Kind: The kind of insert resource this is.
	Kind string `json:"kind,omitempty"`

	// ResponseCode: The status of the insert request.
	ResponseCode string `json:"responseCode,omitempty"`
}

// method id "groupsmigration.archive.insert":

type ArchiveInsertCall struct {
	s             *Service
	groupId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Inserts a new mail into the archive of the Google group.

func (r *ArchiveService) Insert(groupId string) *ArchiveInsertCall {
	return &ArchiveInsertCall{
		s:             r.s,
		groupId:       groupId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{groupId}/archive",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *ArchiveInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *ArchiveInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/groups/v1/groups/{groupId}/archive"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/groups/v1/groups/{groupId}/archive"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *ArchiveInsertCall) Media(r io.Reader) *ArchiveInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/groups/v1/groups/{groupId}/archive"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *ArchiveInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ArchiveInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/groups/v1/groups/{groupId}/archive"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *ArchiveInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ArchiveInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ArchiveInsertCall) Fields(s ...googleapi.Field) *ArchiveInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ArchiveInsertCall) Do() (*Groups, error) {
	var returnValue *Groups
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupId": c.groupId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Inserts a new mail into the archive of the Google group.",
	//   "httpMethod": "POST",
	//   "id": "groupsmigration.archive.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "message/rfc822"
	//     ],
	//     "maxSize": "16MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/groups/v1/groups/{groupId}/archive"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/groups/v1/groups/{groupId}/archive"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The group ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupId}/archive",
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.migration"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

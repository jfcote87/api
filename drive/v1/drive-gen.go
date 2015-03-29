// Package drive provides access to the Drive API.
//
// See https://developers.google.com/drive/
//
// Usage example:
//
//   import "github.com/jfcote87/api/drive/v1"
//   ...
//   driveService, err := drive.New(oauthHttpClient)
package drive

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

const apiId = "drive:v1"
const apiName = "drive"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/drive/v1/"}

// OAuth2 scopes used by this API.
const (
	// View and manage Google Drive files that you have opened or created
	// with this app
	DriveFileScope = "https://www.googleapis.com/auth/drive.file"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Files = NewFilesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Files *FilesService
}

func NewFilesService(s *Service) *FilesService {
	rs := &FilesService{s: s}
	return rs
}

type FilesService struct {
	s *Service
}

type File struct {
	// CreatedDate: Create time for this file (formatted RFC 3339
	// timestamp).
	CreatedDate string `json:"createdDate,omitempty"`

	// Description: A short description of the file
	Description string `json:"description,omitempty"`

	// DownloadUrl: Short term download URL for the file. This will only be
	// populated on files with content stored in Drive.
	DownloadUrl string `json:"downloadUrl,omitempty"`

	// Etag: ETag of the file.
	Etag string `json:"etag,omitempty"`

	// FileExtension: The file extension used when downloading this file.
	// This field is read only. To set the extension, include it on title
	// when creating the file. This will only be populated on files with
	// content stored in Drive.
	FileExtension string `json:"fileExtension,omitempty"`

	// FileSize: The size of the file in bytes. This will only be populated
	// on files with content stored in Drive.
	FileSize int64 `json:"fileSize,omitempty,string"`

	// Id: The id of the file.
	Id string `json:"id,omitempty"`

	// IndexableText: Indexable text attributes for the file (can only be
	// written)
	IndexableText *FileIndexableText `json:"indexableText,omitempty"`

	// Kind: The type of file. This is always drive#file
	Kind string `json:"kind,omitempty"`

	// Labels: Labels for the file.
	Labels *FileLabels `json:"labels,omitempty"`

	// LastViewedDate: Last time this file was viewed by the user (formatted
	// RFC 3339 timestamp).
	LastViewedDate string `json:"lastViewedDate,omitempty"`

	// Md5Checksum: An MD5 checksum for the content of this file. This will
	// only be populated on files with content stored in Drive.
	Md5Checksum string `json:"md5Checksum,omitempty"`

	// MimeType: The mimetype of the file
	MimeType string `json:"mimeType,omitempty"`

	// ModifiedByMeDate: Last time this file was modified by the user
	// (formatted RFC 3339 timestamp).
	ModifiedByMeDate string `json:"modifiedByMeDate,omitempty"`

	// ModifiedDate: Last time this file was modified by anyone (formatted
	// RFC 3339 timestamp).
	ModifiedDate string `json:"modifiedDate,omitempty"`

	// ParentsCollection: Collection of parent folders which contain this
	// file.
	// On insert, setting this field will put the file in all of the
	// provided folders. If no folders are provided, the file will be placed
	// in the default root folder. On update, this field is ignored.
	ParentsCollection []*FileParentsCollection `json:"parentsCollection,omitempty"`

	// SelfLink: A link back to this file.
	SelfLink string `json:"selfLink,omitempty"`

	// Title: The title of this file.
	Title string `json:"title,omitempty"`

	// UserPermission: The permissions for the authenticated user on this
	// file.
	UserPermission *Permission `json:"userPermission,omitempty"`
}

type FileIndexableText struct {
	// Text: The text to be indexed for this file
	Text string `json:"text,omitempty"`
}

type FileLabels struct {
	// Hidden: Whether this file is hidden from the user
	Hidden bool `json:"hidden,omitempty"`

	// Starred: Whether this file is starred by the user.
	Starred bool `json:"starred,omitempty"`

	// Trashed: Whether this file has been trashed.
	Trashed bool `json:"trashed,omitempty"`
}

type FileParentsCollection struct {
	// Id: The id of this parent
	Id string `json:"id,omitempty"`

	// ParentLink: A link to get the metadata for this parent
	ParentLink string `json:"parentLink,omitempty"`
}

type Permission struct {
	// AdditionalRoles: Any additional roles that this permission describes.
	AdditionalRoles []string `json:"additionalRoles,omitempty"`

	// Etag: An etag for this permission.
	Etag string `json:"etag,omitempty"`

	// Kind: The kind of this permission. This is always drive#permission
	Kind string `json:"kind,omitempty"`

	// Role: The role that this permission describes. (For example: reader,
	// writer, owner)
	Role string `json:"role,omitempty"`

	// Type: The type of permission (For example: user, group etc).
	Type string `json:"type,omitempty"`
}

// method id "drive.files.get":

type FilesGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a file's metadata by id.

func (r *FilesService) Get(id string) *FilesGetCall {
	return &FilesGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "files/{id}",
	}
}

// Projection sets the optional parameter "projection": This parameter
// is deprecated and has no function.
func (c *FilesGetCall) Projection(projection string) *FilesGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// UpdateViewedDate sets the optional parameter "updateViewedDate":
// Whether to update the view date after successfully retrieving the
// file.
func (c *FilesGetCall) UpdateViewedDate(updateViewedDate bool) *FilesGetCall {
	c.params_.Set("updateViewedDate", fmt.Sprintf("%v", updateViewedDate))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesGetCall) Fields(s ...googleapi.Field) *FilesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesGetCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a file's metadata by id.",
	//   "httpMethod": "GET",
	//   "id": "drive.files.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id for the file in question.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "This parameter is deprecated and has no function.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Deprecated",
	//         "Deprecated"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updateViewedDate": {
	//       "default": "true",
	//       "description": "Whether to update the view date after successfully retrieving the file.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "files/{id}",
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/drive.file"
	//   ]
	// }

}

// method id "drive.files.insert":

type FilesInsertCall struct {
	s             *Service
	file          *File
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Inserts a file, and any settable metadata or blob content
// sent with the request.

func (r *FilesService) Insert(file *File) *FilesInsertCall {
	return &FilesInsertCall{
		s:             r.s,
		file:          file,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "files",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *FilesInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *FilesInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/drive/v1/files"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/drive/v1/files"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *FilesInsertCall) Media(r io.Reader) *FilesInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/drive/v1/files"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *FilesInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *FilesInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/drive/v1/files"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *FilesInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *FilesInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesInsertCall) Fields(s ...googleapi.Field) *FilesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesInsertCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.file,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Inserts a file, and any settable metadata or blob content sent with the request.",
	//   "httpMethod": "POST",
	//   "id": "drive.files.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "5120GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/drive/v1/files"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/drive/v1/files"
	//       }
	//     }
	//   },
	//   "path": "files",
	//   "request": {
	//     "$ref": "File"
	//   },
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/drive.file"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "drive.files.patch":

type FilesPatchCall struct {
	s             *Service
	id            string
	file          *File
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates file metadata and/or content. This method supports
// patch semantics.

func (r *FilesService) Patch(id string, file *File) *FilesPatchCall {
	return &FilesPatchCall{
		s:             r.s,
		id:            id,
		file:          file,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "files/{id}",
	}
}

// NewRevision sets the optional parameter "newRevision": Whether a blob
// upload should create a new revision. If false, the blob data in the
// current head revision is replaced. If true or not set, a new blob is
// created as head revision, and previous revisions are preserved
// (causing increased use of the user's data storage quota).
func (c *FilesPatchCall) NewRevision(newRevision bool) *FilesPatchCall {
	c.params_.Set("newRevision", fmt.Sprintf("%v", newRevision))
	return c
}

// UpdateModifiedDate sets the optional parameter "updateModifiedDate":
// Controls updating the modified date of the file. If true, the
// modified date will be updated to the current time, regardless of
// whether other changes are being made. If false, the modified date
// will only be updated to the current time if other changes are also
// being made (changing the title, for example).
func (c *FilesPatchCall) UpdateModifiedDate(updateModifiedDate bool) *FilesPatchCall {
	c.params_.Set("updateModifiedDate", fmt.Sprintf("%v", updateModifiedDate))
	return c
}

// UpdateViewedDate sets the optional parameter "updateViewedDate":
// Whether to update the view date after successfully updating the file.
func (c *FilesPatchCall) UpdateViewedDate(updateViewedDate bool) *FilesPatchCall {
	c.params_.Set("updateViewedDate", fmt.Sprintf("%v", updateViewedDate))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesPatchCall) Fields(s ...googleapi.Field) *FilesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesPatchCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.file,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates file metadata and/or content. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "drive.files.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id for the file in question.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "newRevision": {
	//       "default": "true",
	//       "description": "Whether a blob upload should create a new revision. If false, the blob data in the current head revision is replaced. If true or not set, a new blob is created as head revision, and previous revisions are preserved (causing increased use of the user's data storage quota).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updateModifiedDate": {
	//       "default": "false",
	//       "description": "Controls updating the modified date of the file. If true, the modified date will be updated to the current time, regardless of whether other changes are being made. If false, the modified date will only be updated to the current time if other changes are also being made (changing the title, for example).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updateViewedDate": {
	//       "default": "true",
	//       "description": "Whether to update the view date after successfully updating the file.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "files/{id}",
	//   "request": {
	//     "$ref": "File"
	//   },
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/drive.file"
	//   ]
	// }

}

// method id "drive.files.update":

type FilesUpdateCall struct {
	s             *Service
	id            string
	file          *File
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Update: Updates file metadata and/or content

func (r *FilesService) Update(id string, file *File) *FilesUpdateCall {
	return &FilesUpdateCall{
		s:             r.s,
		id:            id,
		file:          file,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "files/{id}",
		context_:      context.TODO(),
	}
}

// NewRevision sets the optional parameter "newRevision": Whether a blob
// upload should create a new revision. If false, the blob data in the
// current head revision is replaced. If true or not set, a new blob is
// created as head revision, and previous revisions are preserved
// (causing increased use of the user's data storage quota).
func (c *FilesUpdateCall) NewRevision(newRevision bool) *FilesUpdateCall {
	c.params_.Set("newRevision", fmt.Sprintf("%v", newRevision))
	return c
}

// UpdateModifiedDate sets the optional parameter "updateModifiedDate":
// Controls updating the modified date of the file. If true, the
// modified date will be updated to the current time, regardless of
// whether other changes are being made. If false, the modified date
// will only be updated to the current time if other changes are also
// being made (changing the title, for example).
func (c *FilesUpdateCall) UpdateModifiedDate(updateModifiedDate bool) *FilesUpdateCall {
	c.params_.Set("updateModifiedDate", fmt.Sprintf("%v", updateModifiedDate))
	return c
}

// UpdateViewedDate sets the optional parameter "updateViewedDate":
// Whether to update the view date after successfully updating the file.
func (c *FilesUpdateCall) UpdateViewedDate(updateViewedDate bool) *FilesUpdateCall {
	c.params_.Set("updateViewedDate", fmt.Sprintf("%v", updateViewedDate))
	return c
}

// MediaUpload takes a context and UploadCaller interface
func (c *FilesUpdateCall) Upload(ctx context.Context, u googleapi.UploadCaller) *FilesUpdateCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/drive/v1/files/{id}"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/drive/v1/files/{id}"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *FilesUpdateCall) Media(r io.Reader) *FilesUpdateCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/drive/v1/files/{id}"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *FilesUpdateCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *FilesUpdateCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/drive/v1/files/{id}"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *FilesUpdateCall) ProgressUpdater(pu googleapi.ProgressUpdater) *FilesUpdateCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesUpdateCall) Fields(s ...googleapi.Field) *FilesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesUpdateCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.file,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates file metadata and/or content",
	//   "httpMethod": "PUT",
	//   "id": "drive.files.update",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "5120GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/drive/v1/files/{id}"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/drive/v1/files/{id}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id for the file in question.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "newRevision": {
	//       "default": "true",
	//       "description": "Whether a blob upload should create a new revision. If false, the blob data in the current head revision is replaced. If true or not set, a new blob is created as head revision, and previous revisions are preserved (causing increased use of the user's data storage quota).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updateModifiedDate": {
	//       "default": "false",
	//       "description": "Controls updating the modified date of the file. If true, the modified date will be updated to the current time, regardless of whether other changes are being made. If false, the modified date will only be updated to the current time if other changes are also being made (changing the title, for example).",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updateViewedDate": {
	//       "default": "true",
	//       "description": "Whether to update the view date after successfully updating the file.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "files/{id}",
	//   "request": {
	//     "$ref": "File"
	//   },
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/drive.file"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

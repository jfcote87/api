// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package googleapi contains the common code shared by all Google API
// libraries.
package googleapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/api/googleapi/internal/uritemplates"
)

// ContentTyper is an interface for Readers which know (or would like
// to override) their Content-Type. If a media body doesn't implement
// ContentTyper, the type is sniffed from the content using
// http.DetectContentType.
type ContentTyper interface {
	ContentType() string
}

// A SizeReaderAt is a ReaderAt with a Size method.
// An io.SectionReader implements SizeReaderAt.
type SizeReaderAt interface {
	io.Reader
	io.ReaderAt
	Size() int64
}

const (
	Version = "0.5"

	// statusResumeIncomplete is the code returned by the Google uploader when the transfer is not yet complete.
	statusResumeIncomplete = 308

	// userAgent is the header string used to identify itself to the Google uploader.
	userAgent = "google-api-go-client/" + Version

	// uploadPause determines the delay between failed upload attempts
	uploadPause = 1 * time.Second
)

var (
	NoContext = context.TODO()

	batchClient = &http.Client{Transport: batchTransport{}}
)

type batchTransport struct{}

func (b batchTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		req.Body.Close()
	}
	return nil, errors.New("googleapi: Batch Client does not process calls with media uploads.")
}

func BatchClient() *http.Client {
	return batchClient
}

// Error contains an error response from the server.
type Error struct {
	// Code is the HTTP response status code and will always be populated.
	Code int `json:"code"`
	// Message is the server response message and is only populated when
	// explicitly referenced by the JSON server response.
	Message string `json:"message"`
	// Body is the raw response returned by the server.
	// It is often but not always JSON, depending on how the request fails.
	Body string

	Errors []ErrorItem
}

// ErrorItem is a detailed error code & message from the Google API frontend.
type ErrorItem struct {
	// Reason is the typed error code. For example: "some_example".
	Reason string `json:"reason"`
	// Message is the human-readable description of the error.
	Message string `json:"message"`
}

func (e *Error) Error() string {
	if len(e.Errors) == 0 && e.Message == "" {
		return fmt.Sprintf("googleapi: got HTTP response code %d with body: %v", e.Code, e.Body)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "googleapi: Error %d: ", e.Code)
	if e.Message != "" {
		fmt.Fprintf(&buf, "%s", e.Message)
	}
	if len(e.Errors) == 0 {
		return strings.TrimSpace(buf.String())
	}
	if len(e.Errors) == 1 && e.Errors[0].Message == e.Message {
		fmt.Fprintf(&buf, ", %s", e.Errors[0].Reason)
		return buf.String()
	}
	fmt.Fprintln(&buf, "\nMore details:")
	for _, v := range e.Errors {
		fmt.Fprintf(&buf, "Reason: %s, Message: %s\n", v.Reason, v.Message)
	}
	return buf.String()

}

type errorReply struct {
	Error *Error `json:"error"`
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer CloseBody(res)
	slurp, err := ioutil.ReadAll(res.Body)
	if err == nil {
		jerr := new(errorReply)
		err = json.Unmarshal(slurp, jerr)
		if err == nil && jerr.Error != nil {
			if jerr.Error.Code == 0 {
				jerr.Error.Code = res.StatusCode
			}
			jerr.Error.Body = string(slurp)
			return jerr.Error
		}
	}
	return &Error{
		Code: res.StatusCode,
		Body: string(slurp),
	}
}

// Wrap returns a struct with passed interface
// as a field named Data.
func Wrap(d interface{}) interface{} {
	var dataWrapper struct {
		Data interface{} `json:"data"`
	}
	dataWrapper.Data = d
	return dataWrapper
}

func typeHeader(contentType string) textproto.MIMEHeader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", contentType)
	return h
}

var errAborted = errors.New("googleapi: upload aborted")

// ProgressUpdater is a function that is called upon every progress update of a resumable upload.
// This is the only part of a resumable upload (from googleapi) that is usable by the developer.
// The remaining usable pieces of resumable uploads is exposed in each auto-generated API.
type ProgressUpdater func(current, total int64, err error)

// ResumableUpload is used by the generated APIs to provide resumable uploads.
// It is not used by developers directly.
type ResumableUpload struct {
	// URL is the resumable resource destination provided by the server after specifying "&uploadType=resumable".
	URL string
	// Media is the object being uploaded.
	Media io.ReaderAt
	// MediaType defines the media type, e.g. "image/jpeg".
	MediaType string
	// ContentLength is the full size of the object being uploaded.
	ContentLength int64
	// ChunkSize is the size of the chunks created during a resumable upload and should be a power of two.
	// 1<<18 is the minimum size supported by the Google uploader, and there is no maximum.
	// If ChunkSize is zero, chunking is ignored.
	ChunkSize int64
	// MaxRetries is the number of times the upload will retry after 5XX errors.
	MaxRetries int

	//	mu       sync.Mutex // guards progress
	progress int64 // number of bytes uploaded so far

	// Callback is an optional function that will be called upon every progress update.
	Callback ProgressUpdater
	// Error causing the upload to stop.  If set to
	Err error
}

func NewResumableUpload(r io.ReaderAt, mediaType string, size int64, opts ...option) (*ResumableUpload, error) {
	c := &ResumableUpload{
		Media:         r, // io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		MaxRetries:    5,
		//Callback:      updater,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (rx *ResumableUpload) Error() string {
	if rx.Err != nil {
		return fmt.Sprintf("googleapi: ResumableUpload - %v", rx.Err)
	}
	return "googleapi: ResumableUpload - no error"
}

type option func(*ResumableUpload) error

// ChunkSize sets the chunk size of the
func ChunkSize(n int64) option {
	return func(rx *ResumableUpload) error {
		if n > 0 && n%0x40000 == 0 {
			rx.ChunkSize = n
			return nil
		}
		return errors.New("invalid chunk size. Must be at least 256Kb and divisible by 256Kb")
	}
}

// MediaSize set the content length
func MediaSize(n int64) option {
	return func(rx *ResumableUpload) error {
		rx.ContentLength = n
		return nil
	}
}

// Retries sets the number of resumable max
func MaxRetries(n int) option {
	return func(rx *ResumableUpload) error {
		rx.MaxRetries = n
		return nil
	}
}

// ResumableUrl sets the url
func ResumableUrl(url string) option {
	return func(rx *ResumableUpload) error {
		rx.URL = url
		return nil
	}
}

func (rx *ResumableUpload) parseRange(res *http.Response) (err error) {
	rangeString := res.Header.Get("Range")
	// Range header not returned if no bytes have been sent.
	if rangeString == "" {
		rx.progress = 0
		return nil
	}
	if ranges := strings.Split(rangeString, "-"); len(ranges) == 2 {
		if rx.progress, err = strconv.ParseInt(ranges[1], 10, 64); err == nil {
			rx.progress += 1
			return
		}
	}
	err = fmt.Errorf("unable to parse range header: %s", rangeString)
	return
}

// isValidChunk checks that the chunk size is > (1<<18) and is a multiple of 256kb
func isValidChunk(n int64) bool {
	return n > 0 && n%0x40000 == 0
}

func (rx *ResumableUpload) transferChunk(client *http.Client) (*http.Response, error) {
	var rangeStr string
	var body io.Reader
	var reqSize int64
	if rx.progress >= 0 {
		reqSize = rx.ContentLength - rx.progress
		if rx.ChunkSize > 0 && reqSize > rx.ChunkSize {
			reqSize = rx.ChunkSize
		}
		rangeStr = fmt.Sprintf("bytes %d-%d/%d", rx.progress, rx.progress+reqSize-1, rx.ContentLength)
		body = io.NewSectionReader(rx.Media, rx.progress, reqSize)
	} else {
		rangeStr = fmt.Sprintf("bytes */%d", rx.ContentLength) // Status request
	}
	req, _ := http.NewRequest("PUT", rx.URL, body)
	req.ContentLength = reqSize
	req.Header.Set("Content-Range", rangeStr)
	req.Header.Set("Content-Type", rx.MediaType)
	req.Header.Set("User-Agent", userAgent)
	return client.Do(req)
}

// Resume restarts an existing upload
func (rx *ResumableUpload) Resume(ctx context.Context, client *http.Client, v interface{}) error {
	rx.progress = -1
	return rx.Upload(ctx, client, v)
}

// Upload starts the process of a resumable upload with a cancellable context.
func (rx *ResumableUpload) Upload(ctx context.Context, client *http.Client, v interface{}) error {
	maxRetries := uint(rx.MaxRetries)
	retries := uint(0)
	var res *http.Response
	var err error

	rx.Err = nil
	for err == nil {
		// if progress is unknown then check transfer status
		res, err = rx.transferChunk(client)
		if err != nil { // probably no communication with server. retry.
			retries++
		} else if err = CheckResponse(res); err != nil {
			// Res.StatusCode != 200 or 201 and res.Body is closed
			switch res.StatusCode {
			case statusResumeIncomplete:
				err = rx.parseRange(res)
				retries = 0
			case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
				retries++
			default:
				retries = 0 // not retry error, so exit
			}
		} else {
			break // Res.StatusCode = 200 or 201
		}

		if retries > 0 && retries <= maxRetries {
			if rx.Callback != nil {
				rx.Callback(rx.progress, rx.ContentLength, err)
			}
			select {
			case <-ctx.Done():
				err = ctx.Err()
			case <-time.After(time.Second * (1 << (retries - 1))):
				rx.progress = -1
				err = nil
				continue
			}
		} else {
			select { // Check for cancellation
			case <-ctx.Done():
				err = ctx.Err()
			default:
			}
		}
		if rx.Callback != nil {
			rx.Callback(rx.progress, rx.ContentLength, err)
		}
	}

	if err != nil {
		rx.Err = err
		return rx
	}
	defer CloseBody(res)
	return json.NewDecoder(res.Body).Decode(v)

}

// has4860Fix is whether this Go environment contains the fix for
// http://golang.org/issue/4860
var has4860Fix bool

// init initializes has4860Fix by checking the behavior of the net/http package.
func init() {
	r := http.Request{
		URL: &url.URL{
			Scheme: "http",
			Opaque: "//opaque",
		},
	}
	b := &bytes.Buffer{}
	r.Write(b)
	has4860Fix = bytes.HasPrefix(b.Bytes(), []byte("GET http"))
}

// SetOpaque sets u.Opaque from u.Path such that HTTP requests to it
// don't alter any hex-escaped characters in u.Path.
func SetOpaque(u *url.URL) {
	u.Opaque = "//" + u.Host + u.Path
	if !has4860Fix {
		u.Opaque = u.Scheme + ":" + u.Opaque
	}
}

// Expand subsitutes any {encoded} strings in the URL passed in using
// the map supplied.
//
// This calls SetOpaque to avoid encoding of the parameters in the URL path.
func Expand(baseURL *url.URL, path string, expansions map[string]string) (u *url.URL) {
	if expansions != nil {
		if px, err := uritemplates.Expand(path, expansions); err == nil {
			path = px
		}
	}
	pathURL := &url.URL{Path: path}
	u = baseURL.ResolveReference(pathURL)
	SetOpaque(u)
	return
}

// CloseBody is used to close res.Body.
// Prior to calling Close, it also tries to Read a small amount to see an EOF.
// Not seeing an EOF can prevent HTTP Transports from reusing connections.
func CloseBody(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}
	// Justification for 3 byte reads: two for up to "\r\n" after
	// a JSON/XML document, and then 1 to see EOF if we haven't yet.
	// TODO(bradfitz): detect Go 1.3+ and skip these reads.
	// See https://codereview.appspot.com/58240043
	// and https://codereview.appspot.com/49570044
	buf := make([]byte, 1)
	for i := 0; i < 3; i++ {
		_, err := res.Body.Read(buf)
		if err != nil {
			break
		}
	}
	res.Body.Close()

}

// VariantType returns the type name of the given variant.
// If the map doesn't contain the named key or the value is not a []interface{}, "" is returned.
// This is used to support "variant" APIs that can return one of a number of different types.
func VariantType(t map[string]interface{}) string {
	s, _ := t["type"].(string)
	return s
}

// ConvertVariant uses the JSON encoder/decoder to fill in the struct 'dst' with the fields found in variant 'v'.
// This is used to support "variant" APIs that can return one of a number of different types.
// It reports whether the conversion was successful.
func ConvertVariant(v map[string]interface{}, dst interface{}) bool {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return false
	}
	return json.Unmarshal(buf.Bytes(), dst) == nil
}

// A Field names a field to be retrieved with a partial response.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
//
// Partial responses can dramatically reduce the amount of data that must be sent to your application.
// In order to request partial responses, you can specify the full list of fields
// that your application needs by adding the Fields option to your request.
//
// Field strings use camelCase with leading lower-case characters to identify fields within the response.
//
// For example, if your response has a "NextPageToken" and a slice of "Items" with "Id" fields,
// you could request just those fields like this:
//
//     svc.Events.List().Fields("nextPageToken", "items/id").Do()
//
// or if you were also interested in each Item's "Updated" field, you can combine them like this:
//
//     svc.Events.List().Fields("nextPageToken", "items(id,updated)").Do()
//
// More information about field formatting can be found here:
// https://developers.google.com/+/api/#fields-syntax
//
// Another way to find field names is through the Google API explorer:
// https://developers.google.com/apis-explorer/#p/
type Field string

// CombineFields combines fields into a single string.
func CombineFields(s []Field) string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = string(v)
	}
	return strings.Join(r, ",")
}

// Caller is the interface used to execute api calls
//
// *http.Client parameter should perform oauth2.0 authorization
// *Call contains url and payload data
// inteface{} parameter should be a ptr to a struct that the function
// will unmarshal the http.Response into.  If the value is an *http.Response,
// the function will set the value to the response.
type Caller interface {
	Do(context.Context, *http.Client, *Call) error
}

type UploadCaller interface {
	Caller
	getMediaType() (io.Reader, string)
}

type JSONCall struct{}

func (p JSONCall) Do(ctx context.Context, cl *http.Client, c *Call) error {
	// if batchClient just return the call
	if cl == batchClient {
		return c
	}
	body, err := c.PayloadReader()
	if err != nil {
		return err
	}
	params := c.GetParams()

	hdrs := make(http.Header)
	if body != nil {
		hdrs.Add("Content-Type", "application/json")
	}

	return do(cl, c.Method, c.URL, params, hdrs, body, c.Result)
	/*	res, err := do(cl, c.Method, c.URL, params, hdrs, body)
		if err != nil {
			return err
		}
		return handleResponse(res, c.Result) */
}

// MediaUpload is a caller that executes simple and multipart uploads
type MediaUpload struct {
	Media     io.Reader
	MediaType string
}

func (m *MediaUpload) getMediaType() (io.Reader, string) {
	if m.MediaType != "" {
		return m.Media, m.MediaType
	}
	if typer, ok := m.Media.(ContentTyper); ok {
		return m.Media, typer.ContentType()
	}

	typ := "application/octet-stream"
	buf := make([]byte, 1024)
	n, err := m.Media.Read(buf)
	buf = buf[:n]
	if err == nil {
		typ = http.DetectContentType(buf)
	}
	return io.MultiReader(bytes.NewBuffer(buf), m.Media), typ
}

func multipartMedia(jsonBody, mediaBody io.Reader, ct string) (io.ReadCloser, string) {
	pr, pw := io.Pipe()
	mpw := multipart.NewWriter(pw)
	go func() {
		var err error
		var w io.Writer
		defer func() {
			if err != nil {
				pr.CloseWithError(err)
			}
			mpw.Close()
			pw.Close()
		}()
		if w, err = mpw.CreatePart(textproto.MIMEHeader{
			"Content-Type": []string{"application/json"},
		}); err != nil {
			return
		}
		if _, err = io.Copy(w, jsonBody); err != nil {
			return
		}
		if w, err = mpw.CreatePart(textproto.MIMEHeader{
			"Content-Type": []string{ct},
		}); err != nil {
			return
		}
		if _, err = io.Copy(w, mediaBody); err != nil {
			return
		}
		return
	}()
	return pr, "multipart/related; boundary=" + mpw.Boundary()
}

func (m *MediaUpload) Do(ctx context.Context, cl *http.Client, c *Call) error {
	defer func() {
		if closer, ok := m.Media.(io.Closer); ok {
			closer.Close()
		}
	}()

	uploadType := "media"
	params := c.GetParams()

	body, ct := m.getMediaType()
	jsonBody, err := c.PayloadReader()
	if err != nil {
		return err
	}
	if jsonBody != nil {
		uploadType = "multipart"
		pr, tct := multipartMedia(jsonBody, body, ct)
		defer pr.Close()
		body = pr
		ct = tct
	}
	params.Set("uploadType", uploadType)
	return do(cl, c.Method, c.URL, params, http.Header{"content-type": []string{ct}}, body, c.Result)
}

func (rx *ResumableUpload) getMediaType() (io.Reader, string) {
	if rx.MediaType != "" {
		return nil, rx.MediaType
	}
	if typer, ok := rx.Media.(ContentTyper); ok {
		return nil, typer.ContentType()
	}

	typ := "application/octet-stream"
	buf := make([]byte, 1024)
	n, err := rx.Media.ReadAt(buf, 0)
	buf = buf[:n]
	if err == nil || err == io.EOF {
		typ = http.DetectContentType(buf)
	}
	return nil, typ
}

func (rx *ResumableUpload) Do(ctx context.Context, cl *http.Client, c *Call) error {
	var jsonBody io.Reader
	//var err error
	params := c.GetParams()
	params.Set("uploadType", "resumable")
	_, ct := rx.getMediaType()
	hdrs := http.Header{"X-Upload-Content-Type": []string{ct}}
	if c.Payload != nil {
		var err error
		if jsonBody, err = c.PayloadReader(); err != nil {
			return err
		} //payloadReader(c.Payload)
		hdrs.Add("Content-Type", "application/json")
	}

	var res *http.Response
	if err := do(cl, c.Method, c.URL, params, hdrs, jsonBody, &res); err != nil {
		return err
	}
	CloseBody(res) // no need to read response after this point

	rx.URL = res.Header.Get("Location")
	return rx.Upload(ctx, cl, c.Result)

}

func do(cl *http.Client, method string, url *url.URL, params url.Values, hdrs http.Header, body io.Reader, v interface{}) error {
	resPtr, isResponse := v.(**http.Response)
	if isResponse {
		params.Set("alt", "media")
	} else if params.Get("alt") != "" {
		params.Set("alt", "json")
	}

	req, err := http.NewRequest(method, "", body)
	if err != nil {
		return err
	}
	req.URL = url
	if hdrs != nil {
		req.Header = hdrs
	}
	req.Header.Set("User-Agent", userAgent)
	req.URL.RawQuery = params.Encode()
	res, err := cl.Do(req)
	if err != nil {
		return err
	}
	if err = CheckResponse(res); err != nil {
		return err
	}
	if isResponse {
		*resPtr = res
		return nil
	}
	defer CloseBody(res)
	if v != nil {
		err = json.NewDecoder(res.Body).Decode(v)
	}
	return err

}

/*
func handleResponse(res *http.Response, v interface{}) error {
	closeBody := true
	defer func() {
		if closeBody {
			CloseBody(res)
		}
	}()
	var err error
	switch ret := v.(type) {
	case **http.Response:
		ret = &res
		closeBody = false
	case interface{}:
		err = json.NewDecoder(res.Body).Decode(ret)
	}
	return err
}
*/
// Call object is
type Call struct {
	// Payload is the value to be marshalled into
	// the request body.
	Payload interface{}

	// http method for request
	Method string

	// URL for request, not including querystring
	URL *url.URL

	// Params contain the values for calls querystring
	Params url.Values

	// Result is used to return data back to the calling function.
	// It may contain a pointer to a struct for unmarshalling the
	// response, or an *http.Response which
	Result interface{}
}

// GetParams checks for nil value in params
// field.  Returns empty map if nil.
func (c Call) GetParams() url.Values {
	if c.Params == nil {
		return make(url.Values)
	}
	return c.Params
}

// Error so that Call fulfills the error interface.
func (c Call) Error() string {
	return "googleapi: batch call"
}

// JSONBody returns an io reader containing
// the results of a marshaled payload.  If
// payload is nil then a nil reader is returned.
func (c Call) PayloadReader() (io.Reader, error) {
	if c.Payload != nil && !reflect.ValueOf(c.Payload).IsNil() {
		b, err := json.Marshal(c.Payload)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), err
	}
	return nil, nil
}

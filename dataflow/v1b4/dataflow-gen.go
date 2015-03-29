// Package dataflow provides access to the Google Dataflow API.
//
// Usage example:
//
//   import "github.com/jfcote87/api/dataflow/v1b4"
//   ...
//   dataflowService, err := dataflow.New(oauthHttpClient)
package dataflow

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

const apiId = "dataflow:v1b4"
const apiName = "dataflow"
const apiVersion = "v1b4"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "dataflow.googleapis.com", Path: "/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	return s, nil
}

type Service struct {
	client *http.Client
}

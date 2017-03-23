package client

import (
	"habor-template-demo/api/errors"
	"habor-template-demo/api/v1"

	"github.com/golang/glog"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	apiScheme = "https"
	apiHost   = "api.whispir.com"

	// Add resource type here
	ResourceTemplate  = "templates"
	ResourceWorkspace = "workspaces"
	ResourceMessage   = "messages"
)

var allowedType map[string]string

func init() {
	// Add mappings of resource type to content type here
	allowedType = map[string]string{
		ResourceTemplate:  whispirContentType("template", "v1"),
		ResourceWorkspace: whispirContentType("workspace", "v1"),
		ResourceMessage:   whispirContentType("message", "v1"),
	}
}

func whispirContentType(resource, version string) string {
	return fmt.Sprintf("application/vnd.whispir.%s-%s+json", resource, version)
}

// Wrap the client to query whispir API
type Client struct {
	ApiKey   string
	User     string
	Password string
}

func NewClient(apiKey, user, password string) *Client {
	return &Client{
		ApiKey:   apiKey,
		User:     user,
		Password: password,
	}
}

// Set 'GET' method
func (c *Client) Get() *Request {
	return c.Verb(http.MethodGet)
}

// Set 'POST' method
func (c *Client) Post() *Request {
	return c.Verb(http.MethodPost)
}

func (c *Client) Put() *Request {
	return c.Verb(http.MethodPut)
}

// Set 'DELETE' method
func (c *Client) Delete() *Request {
	return c.Verb(http.MethodDelete)
}

// Set method of request
func (c *Client) Verb(method string) *Request {
	return &Request{
		client: c,
		method: method,
	}
}

// Wrap http request
type Request struct {
	// http.Request object to send request
	request *http.Request
	client  *Client
	err     error
	// Request method
	method          string
	workspace       string
	resource        string
	id              string
	body            io.Reader
	acceptFlag      bool
	contentTypeFlag bool
	listOpts        *v1.PaginationOptions
}

// Set workspace
func (r *Request) Workspace(workspace string) *Request {
	if nil != r.err {
		return r
	}
	r.workspace = workspace
	return r
}

// Set resource to be requested
func (r *Request) Resource(resource string) *Request {
	if nil != r.err {
		return r
	}
	r.resource = resource
	return r
}

// Set resource id
func (r *Request) ID(id string) *Request {
	if nil != r.err {
		return r
	}
	r.id = id
	return r
}

// Set body from file
func (r *Request) BodyFromFile(fpath string) *Request {
	if nil != r.err {
		return r
	}
	data, err := ioutil.ReadFile(fpath)
	if nil != err {
		r.err = fmt.Errorf("read file error: %v\n", err)
		return r
	}
	r.body = bytes.NewReader(data)
	return r
}

// Set body
func (r *Request) Body(data []byte) *Request {
	if nil != r.err {
		return r
	}
	r.body = bytes.NewReader(data)
	return r
}

// Set accept header of request
func (r *Request) SetAccept() *Request {
	if nil != r.err {
		return r
	}
	r.acceptFlag = true
	return r
}

// Set content type header of request
func (r *Request) SetContentType() *Request {
	if nil != r.err {
		return r
	}
	r.contentTypeFlag = true
	return r
}

func (r *Request) Pagination(opts *v1.PaginationOptions) *Request {
	if nil != r.err {
		return r
	}
	r.listOpts = opts
	return r
}

// Send request and return a Result
func (r *Request) Do() *Result {
	result := &Result{}
	if nil != r.err {
		result.err = r.err
		return result
	}
	var err error
	// Get http.Request object
	r.request, err = http.NewRequest(r.method, r.getUrl(), r.body)
	if nil != err {
		result.err = internalError("new request error: %v", err)
		return result
	}

	// Set 'Accept' and 'Content-Type' headers
	err = r.setHeaders()
	if nil != err {
		result.err = internalError("set request headers error: %v", err)
		return result
	}

	// Do request and set Result
	resp, err := http.DefaultClient.Do(r.request)
	if nil != err {
		result.err = internalError("send request error: %v", err)
		return result
	}
	result.header = resp.Header

	// Read response body
	if nil != resp.Body {
		defer resp.Body.Close()
		result.body, err = ioutil.ReadAll(resp.Body)
		if glog.V(10) {
			glog.Infof("response %d: %s\n", resp.StatusCode, string(result.body))
		}
		if nil != err {
			result.err = internalError("read response body error: %v", err)
			return result
		}
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusNotModified {
		result.err = errors.New(resp.StatusCode, string(result.body))
	}
	return result
}

func internalError(format string, args ...interface{}) error {
	return errors.NewInternalServerError(fmt.Sprintf(format, args...))
}

// Get request URL
func (r *Request) getUrl() string {
	// Set scheme and host
	u := url.URL{
		Scheme: apiScheme,
		Host:   apiHost,
	}

	// Set path
	p := ""
	if "" != r.workspace {
		p = path.Join(p, ResourceWorkspace, r.workspace)
	}
	p = path.Join(p, r.resource)
	if "" != r.id {
		p = path.Join(p, r.id)
	}
	u.Path = p

	// Set query
	query := url.Values{}
	query.Set("apikey", r.client.ApiKey)
	if nil != r.listOpts {
		query.Set("offset", strconv.Itoa(r.listOpts.Offset))
		query.Set("limit", strconv.Itoa(r.listOpts.Limit))
	}
	u.RawQuery = query.Encode()

	if glog.V(8) {
		glog.Infoln("request url: ", u.String())
	}
	return u.String()
}

// Set request headers
func (r *Request) setHeaders() error {
	// Set basic auth headers
	if "" == r.client.User || "" == r.client.Password {
		return fmt.Errorf("user name or password is not specified")
	}
	r.request.SetBasicAuth(r.client.User, r.client.Password)

	// Set headers if needed
	if "" != r.resource && (r.acceptFlag || r.contentTypeFlag) {
		t, allowed := allowedType[r.resource]
		if !allowed {
			return fmt.Errorf("disallowed resource type '%s'", r.resource)
		}
		if r.acceptFlag {
			r.request.Header.Add("Accept", t)
		}
		if r.contentTypeFlag {
			r.request.Header.Add("Content-Type", t)
		}
	}
	return nil
}

// Wrap http response
type Result struct {
	// response body
	body []byte
	// response status
	statusCode int

	header http.Header

	err error
}

// Decode response into the specified object
func (res *Result) Into(obj interface{}) error {
	if nil != res.err {
		return res.err
	}
	if len(res.body) > 0 {
		if err := json.Unmarshal(res.body, obj); nil != err {
			return internalError("decode response error: %v", err)
		}
	}
	return nil
}

// Return response error
func (res *Result) Err() error {
	return res.err
}

func (res *Result) ExtractLocation(loc *string) *Result {
	if nil != res.err {
		return res
	}
	if len(res.header["Location"]) > 0 {
		*loc = res.header["Location"][0]
	}
	return res
}

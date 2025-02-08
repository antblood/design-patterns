// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: subject.proto

package v1

import context "context"
import fmt "fmt"
import http "net/http"
import io "io"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// =================
// Subject Interface
// =================

// The Subject service definition.
type Subject interface {
	// Add an observer to the subject.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// Remove an observer from the subject.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// Get the latest state of the subject.
	LatestState(context.Context, *LatestStateRequest) (*LatestStateResponse, error)
}

// =======================
// Subject Protobuf Client
// =======================

type subjectProtobufClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewSubjectProtobufClient creates a Protobuf client that implements the Subject interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewSubjectProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Subject {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "subject", "Subject")
	urls := [3]string{
		serviceURL + "Add",
		serviceURL + "Remove",
		serviceURL + "LatestState",
	}

	return &subjectProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *subjectProtobufClient) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "Add")
	caller := c.callAdd
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *AddRequest) (*AddResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*AddRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*AddRequest) when calling interceptor")
					}
					return c.callAdd(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*AddResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*AddResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectProtobufClient) callAdd(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	out := new(AddResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *subjectProtobufClient) Remove(ctx context.Context, in *RemoveRequest) (*RemoveResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "Remove")
	caller := c.callRemove
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*RemoveRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*RemoveRequest) when calling interceptor")
					}
					return c.callRemove(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*RemoveResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*RemoveResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectProtobufClient) callRemove(ctx context.Context, in *RemoveRequest) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *subjectProtobufClient) LatestState(ctx context.Context, in *LatestStateRequest) (*LatestStateResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "LatestState")
	caller := c.callLatestState
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *LatestStateRequest) (*LatestStateResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*LatestStateRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*LatestStateRequest) when calling interceptor")
					}
					return c.callLatestState(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*LatestStateResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*LatestStateResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectProtobufClient) callLatestState(ctx context.Context, in *LatestStateRequest) (*LatestStateResponse, error) {
	out := new(LatestStateResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===================
// Subject JSON Client
// ===================

type subjectJSONClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewSubjectJSONClient creates a JSON client that implements the Subject interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewSubjectJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Subject {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "subject", "Subject")
	urls := [3]string{
		serviceURL + "Add",
		serviceURL + "Remove",
		serviceURL + "LatestState",
	}

	return &subjectJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *subjectJSONClient) Add(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "Add")
	caller := c.callAdd
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *AddRequest) (*AddResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*AddRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*AddRequest) when calling interceptor")
					}
					return c.callAdd(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*AddResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*AddResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectJSONClient) callAdd(ctx context.Context, in *AddRequest) (*AddResponse, error) {
	out := new(AddResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *subjectJSONClient) Remove(ctx context.Context, in *RemoveRequest) (*RemoveResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "Remove")
	caller := c.callRemove
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*RemoveRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*RemoveRequest) when calling interceptor")
					}
					return c.callRemove(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*RemoveResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*RemoveResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectJSONClient) callRemove(ctx context.Context, in *RemoveRequest) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *subjectJSONClient) LatestState(ctx context.Context, in *LatestStateRequest) (*LatestStateResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithMethodName(ctx, "LatestState")
	caller := c.callLatestState
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *LatestStateRequest) (*LatestStateResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*LatestStateRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*LatestStateRequest) when calling interceptor")
					}
					return c.callLatestState(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*LatestStateResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*LatestStateResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *subjectJSONClient) callLatestState(ctx context.Context, in *LatestStateRequest) (*LatestStateResponse, error) {
	out := new(LatestStateResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ======================
// Subject Server Handler
// ======================

type subjectServer struct {
	Subject
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewSubjectServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewSubjectServer(svc Subject, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &subjectServer{
		Subject:          svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *subjectServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *subjectServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// SubjectPathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const SubjectPathPrefix = "/twirp/subject.Subject/"

func (s *subjectServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "subject")
	ctx = ctxsetters.WithServiceName(ctx, "Subject")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "subject.Subject" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "Add":
		s.serveAdd(ctx, resp, req)
		return
	case "Remove":
		s.serveRemove(ctx, resp, req)
		return
	case "LatestState":
		s.serveLatestState(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *subjectServer) serveAdd(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveAddJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveAddProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *subjectServer) serveAddJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Add")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(AddRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Subject.Add
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *AddRequest) (*AddResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*AddRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*AddRequest) when calling interceptor")
					}
					return s.Subject.Add(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*AddResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*AddResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *AddResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *AddResponse and nil error while calling Add. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) serveAddProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Add")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(AddRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Subject.Add
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *AddRequest) (*AddResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*AddRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*AddRequest) when calling interceptor")
					}
					return s.Subject.Add(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*AddResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*AddResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *AddResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *AddResponse and nil error while calling Add. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) serveRemove(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveRemoveJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveRemoveProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *subjectServer) serveRemoveJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Remove")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(RemoveRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Subject.Remove
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*RemoveRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*RemoveRequest) when calling interceptor")
					}
					return s.Subject.Remove(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*RemoveResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*RemoveResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *RemoveResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *RemoveResponse and nil error while calling Remove. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) serveRemoveProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Remove")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(RemoveRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Subject.Remove
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*RemoveRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*RemoveRequest) when calling interceptor")
					}
					return s.Subject.Remove(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*RemoveResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*RemoveResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *RemoveResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *RemoveResponse and nil error while calling Remove. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) serveLatestState(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveLatestStateJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveLatestStateProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *subjectServer) serveLatestStateJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "LatestState")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(LatestStateRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Subject.LatestState
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *LatestStateRequest) (*LatestStateResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*LatestStateRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*LatestStateRequest) when calling interceptor")
					}
					return s.Subject.LatestState(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*LatestStateResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*LatestStateResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *LatestStateResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *LatestStateResponse and nil error while calling LatestState. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) serveLatestStateProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "LatestState")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(LatestStateRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Subject.LatestState
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *LatestStateRequest) (*LatestStateResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*LatestStateRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*LatestStateRequest) when calling interceptor")
					}
					return s.Subject.LatestState(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*LatestStateResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*LatestStateResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *LatestStateResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *LatestStateResponse and nil error while calling LatestState. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *subjectServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *subjectServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *subjectServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "subject", "Subject")
}

var twirpFileDescriptor1 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3d, 0x0f, 0x82, 0x30,
	0x10, 0x86, 0x63, 0x4c, 0x20, 0x39, 0x44, 0x4d, 0x01, 0x4d, 0xd0, 0x89, 0x5d, 0xfc, 0x9a, 0x1c,
	0x71, 0x72, 0x70, 0x82, 0xcd, 0x0d, 0xec, 0x2d, 0x26, 0x52, 0x6c, 0x0f, 0x7e, 0x9f, 0x3f, 0xcd,
	0x84, 0x02, 0x4a, 0x70, 0xbc, 0xf7, 0x7d, 0xae, 0xf7, 0xa4, 0x60, 0xab, 0x32, 0x7b, 0xe0, 0x9d,
	0xc2, 0x42, 0x0a, 0x12, 0xcc, 0x6c, 0xc6, 0x60, 0x02, 0x10, 0x71, 0x1e, 0xe3, 0xab, 0x44, 0x45,
	0x81, 0x0d, 0x56, 0x3d, 0xa9, 0x42, 0xe4, 0x0a, 0x83, 0x19, 0xd8, 0x31, 0x3e, 0x45, 0x85, 0x6d,
	0x3f, 0x87, 0x69, 0x1b, 0x34, 0x88, 0x0b, 0xec, 0x9a, 0x12, 0x2a, 0x4a, 0x28, 0xa5, 0x8e, 0xf3,
	0xc0, 0xe9, 0xa5, 0x1a, 0x3e, 0xbc, 0x47, 0x60, 0x26, 0xfa, 0x30, 0xdb, 0xc1, 0x38, 0xe2, 0x9c,
	0x39, 0x61, 0x2b, 0xf6, 0xd5, 0xf0, 0xdd, 0x7e, 0xa8, 0xb7, 0xd9, 0x09, 0x0c, 0x7d, 0x9c, 0x2d,
	0xba, 0xbe, 0xa7, 0xe7, 0x2f, 0x07, 0x79, 0xb3, 0x7a, 0x01, 0xeb, 0xc7, 0x87, 0xad, 0x3a, 0x6e,
	0xe8, 0xee, 0xaf, 0xff, 0x97, 0xfa, 0xa5, 0xb3, 0x77, 0x73, 0x44, 0xa6, 0x50, 0x56, 0x28, 0x37,
	0x45, 0x4a, 0x84, 0x32, 0xdf, 0x56, 0xfb, 0xcc, 0xa8, 0xbf, 0xf5, 0xf8, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xba, 0xfd, 0xa6, 0x67, 0x01, 0x00, 0x00,
}

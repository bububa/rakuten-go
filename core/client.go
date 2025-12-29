package core

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bububa/rakuten-go/internal/debug"
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHTTPClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	client  *http.Client
	appID   string
	secret  string
	baseURL string
	preReqs []PreRequest
	debug   bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		appID:  appID,
		secret: secret,
		client: defaultHTTPClient(),
	}
}

func (c *SDKClient) AppID() string {
	return c.appID
}

func (c *SDKClient) Secret() string {
	return c.secret
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHTTPClient 设置http.Client
func (c *SDKClient) SetHTTPClient(client *http.Client) {
	c.client = client
}

func (c *SDKClient) WithPreRequests(reqs ...PreRequest) {
	c.preReqs = reqs
}

func (c *SDKClient) AddPreRequests(reqs ...PreRequest) {
	c.preReqs = append(c.preReqs, reqs...)
}

func (c *SDKClient) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		appID:   c.appID,
		secret:  c.secret,
		debug:   c.debug,
		client:  c.client,
		baseURL: c.baseURL,
		preReqs: c.preReqs,
	}
}

// Post post api
func (c *SDKClient) Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.doBodyRequest(ctx, http.MethodPost, c.baseURL, gw, req, resp, accessToken)
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	return c.get(ctx, c.baseURL, gw, req, resp, accessToken)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(ctx context.Context, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	return c.upload(ctx, c.baseURL, gw, req, resp, accessToken)
}

// Put post api
func (c *SDKClient) Put(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.doBodyRequest(ctx, http.MethodPut, c.baseURL, gw, req, resp, accessToken)
}

// Patch post api
func (c *SDKClient) Patch(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.doBodyRequest(ctx, http.MethodPatch, c.baseURL, gw, req, resp, accessToken)
}

// Delete post api
func (c *SDKClient) Delete(ctx context.Context, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	return c.doBodyRequest(ctx, http.MethodDelete, c.baseURL, gw, req, resp, accessToken)
}

func (c *SDKClient) get(ctx context.Context, base string, gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	reqURL := util.StringsJoin(base, gw)
	if req != nil {
		reqURL = util.StringsJoin(reqURL, "?", req.Encode())
	}
	debug.PrintGetRequest(reqURL, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return err
	}
	if r, ok := req.(model.RequestWithContentType); ok {
		httpReq.Header.Add("Content-Type", r.ContentType())
	}
	if accessToken != "" {
		httpReq.Header.Add("Authorization", accessToken)
	}
	return c.fetch(httpReq, resp)
}

func (c *SDKClient) doBodyRequest(ctx context.Context, method string, base string, gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqURL := util.StringsJoin(base, gw)
	httpReq, err := http.NewRequestWithContext(ctx, method, reqURL, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", req.ContentType())
	if accessToken != "" {
		httpReq.Header.Add("Authorization", accessToken)
	}
	debug.PrintJSONRequest(method, reqURL, httpReq.Header, reqBytes, c.debug)
	return c.fetch(httpReq, resp)
}

func (c *SDKClient) upload(ctx context.Context, base string, gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	mw := multipart.NewWriter(buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	for _, v := range params {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			builder := util.NewStringsBuilder()
			builder.WriteString("@")
			builder.WriteString(v.Value)
			mp[v.Key] = builder.String()
			util.ReleaseStringsBuilder(builder)
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	_ = mw.Close()
	reqURL := util.StringsJoin(base, gw)
	debug.PrintPostMultipartRequest(reqURL, mp, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Authorization", accessToken)
	}

	return c.fetch(httpReq, resp)
}

type PreRequest func(httpReq *http.Request) error

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	if len(c.preReqs) > 0 {
		for _, req := range c.preReqs {
			if err := req(httpReq); err != nil {
				return err
			}
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	if unmarshaler, ok := resp.(model.Unmarshaler); ok {
		if err := unmarshaler.Unmarshal(httpResp.Body); err != nil {
			return err
		}
		return nil
	}
	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	contentType := strings.ToLower(httpResp.Header.Get("Content-Type"))
	if strings.Contains(contentType, "xml") {
		err = debug.DecodeXMLHttpResponse(httpResp.Body, buf, resp, c.debug)
	} else {
		err = debug.DecodeJSONHttpResponse(httpResp.Body, buf, resp, c.debug)
	}
	if resp.IsError() {
		return resp
	} else if err != nil {
		debug.PrintError(err, c.debug)
		return errors.Join(err, model.ResponseError{
			ErrorCode:        strconv.Itoa(httpResp.StatusCode),
			ErrorDescription: buf.String(),
		})
	} else if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return errors.Join(err, model.ResponseError{
			ErrorCode:        strconv.Itoa(httpResp.StatusCode),
			ErrorDescription: buf.String(),
		})
	}
	if httpResp.ContentLength <= 0 {
		httpResp.ContentLength = int64(buf.Len())
	}
	return nil
}

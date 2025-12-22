package ichiba

import (
	"github.com/bububa/rakuten-go/core"
)

type Client struct {
	core.SDKClient
}

// NewClient returns a new SDKClient for Rakuten Webservice API
func NewClient(appID string, secret string) *Client {
	clt := core.NewSDKClient(appID, secret)
	clt.SetBaseURL(BaseURL)
	return &Client{*clt}
}

package advertising

import (
	"encoding/base64"

	"github.com/bububa/rakuten-go/core"
	"github.com/bububa/rakuten-go/util"
)

type Client struct {
	core.SDKClient
}

func (c *Client) TokenKey() string {
	buf := util.NewBuffer()
	defer util.ReleaseBuffer(buf)
	buf.WriteString(c.AppID())
	buf.WriteByte(':')
	buf.WriteString(c.Secret())
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// NewClient returns a new SDKClient for Rakuten Advertisering
func NewClient(appID string, secret string) *Client {
	clt := core.NewSDKClient(appID, secret)
	clt.SetBaseURL(BaseURL)
	return &Client{*clt}
}

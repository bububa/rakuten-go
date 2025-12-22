package rms

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
	buf.WriteString(c.Secret())
	buf.WriteByte(':')
	buf.WriteString(c.AppID())
	return util.StringsJoin("ESA ", base64.StdEncoding.EncodeToString(buf.Bytes()))
}

// NewClient returns a new SDKClient for Rakuten Webservice API
func NewClient(appID string, secret string) *Client {
	clt := core.NewSDKClient(appID, secret)
	clt.SetBaseURL(BaseURL)
	return &Client{*clt}
}

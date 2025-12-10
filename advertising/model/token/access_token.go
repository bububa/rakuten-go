package token

import (
	"strconv"

	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

type AccessTokenRequest struct {
	SID          uint64 `json:"sid,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Encode implements Post Request
func (r AccessTokenRequest) Encode() []byte {
	values := util.NewURLValues()
	values.Set("scope", strconv.FormatUint(r.SID, 10))
	if r.RefreshToken != "" {
		values.Set("refresh_token", r.RefreshToken)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return []byte(ret)
}

// ContentType implements Request interface
func (r AccessTokenRequest) ContentType() string {
	return "application/x-www-form-urlencoded"
}

type AccessTokenResponse struct {
	model.BaseResponse
	AccessToken
}

type AccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}

package token

import (
	"context"

	"github.com/bububa/rakuten-go/core"
	"github.com/bububa/rakuten-go/model/token"
)

// AccessToken get a new API access token or with/without refresh token
func AccessToken(ctx context.Context, clt *core.SDKClient, req *token.AccessTokenRequest) (*token.AccessToken, error) {
	var resp token.AccessTokenResponse
	if err := clt.Post(ctx, "/token", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.AccessToken, nil
}

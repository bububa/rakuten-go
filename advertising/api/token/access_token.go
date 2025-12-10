package token

import (
	"context"

	"github.com/bububa/rakuten-go/advertising"
	"github.com/bububa/rakuten-go/advertising/model/token"
)

// AccessToken get a new API access token or with/without refresh token
func AccessToken(ctx context.Context, clt *advertising.Client, req *token.AccessTokenRequest) (*token.AccessToken, error) {
	var resp token.AccessTokenResponse
	if err := clt.Post(ctx, "/token", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.AccessToken, nil
}

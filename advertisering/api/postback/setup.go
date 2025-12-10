package postback

import (
	"context"

	"github.com/bububa/rakuten-go/advertisering/core"
	"github.com/bububa/rakuten-go/advertisering/model/postback"
)

// Setup Set up a postback with a custom URL
func Setup(ctx context.Context, clt *core.SDKClient, req *postback.Postback, accessToken string) error {
	var (
		r = postback.PostbackRequest{
			Postback: *req,
		}
		resp postback.PostbackResponse
	)
	if err := clt.Post(ctx, "/v1/postback", &r, &resp, accessToken); err != nil {
		return err
	}
	*req = *resp.Data
	return nil
}

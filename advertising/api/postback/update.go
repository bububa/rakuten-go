package postback

import (
	"context"
	"strconv"

	"github.com/bububa/rakuten-go/advertising"
	"github.com/bububa/rakuten-go/advertising/model/postback"
	"github.com/bububa/rakuten-go/util"
)

// Update Update an existing postback
func Update(ctx context.Context, clt *advertising.Client, req *postback.Postback, accessToken string) error {
	var (
		r = postback.PostbackRequest{
			Postback: postback.Postback{
				IsActive: req.IsActive,
				URL:      req.URL,
			},
		}
		resp postback.PostbackResponse
	)
	if err := clt.Put(ctx, util.StringsJoin("/v1/postback", strconv.FormatUint(req.SID, 10)), &r, &resp, accessToken); err != nil {
		return err
	}
	*req = *resp.Data
	return nil
}

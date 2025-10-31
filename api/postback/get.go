package postback

import (
	"context"
	"strconv"

	"github.com/bububa/rakuten-go/core"
	"github.com/bububa/rakuten-go/model/postback"
	"github.com/bububa/rakuten-go/util"
)

// Get View current postback settings
func Get(ctx context.Context, clt *core.SDKClient, req *postback.Postback, accessToken string) error {
	var resp postback.PostbackResponse
	if err := clt.Get(ctx, util.StringsJoin("/v1/postback", strconv.FormatUint(req.SID, 10)), nil, &resp, accessToken); err != nil {
		return err
	}
	*req = *resp.Data
	return nil
}

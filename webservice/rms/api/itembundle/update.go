package itembundle

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/itembundle"
)

// Update 更新组合
func Update(ctx context.Context, clt *rms.Client, req *itembundle.CreateRequest) (*itembundle.Bundle, error) {
	var resp itembundle.CreateResponse
	if err := clt.Put(ctx, util.StringsJoin("/es/1.0/bto/bundle", req.BundleManageNumber), req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.Bundle, nil
}

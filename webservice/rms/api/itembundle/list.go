package itembundle

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/itembundle"
)

// List 组合列表响应包装
func List(ctx context.Context, clt *rms.Client, req *itembundle.ListRequest) (*itembundle.ListResponse, error) {
	var resp itembundle.ListResponse
	if err := clt.Get(ctx, "/es/1.0/bto/bundles", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp, nil
}

package itembundle

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/itembundle"
)

// Create 创建组合
func Create(ctx context.Context, clt *rms.Client, req *itembundle.CreateRequest) (*itembundle.Bundle, error) {
	var resp itembundle.CreateResponse
	if err := clt.Post(ctx, "/es/1.0/bto/bundle", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.Bundle, nil
}

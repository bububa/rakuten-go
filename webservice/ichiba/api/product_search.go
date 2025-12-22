package api

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/ichiba"
	"github.com/bububa/rakuten-go/webservice/ichiba/model"
)

// ProductSearch Ichiba Product Search API
func ProductSearch(ctx context.Context, clt *ichiba.Client, req *model.ProductSearchRequest) (*model.ProductSearchResponse, error) {
	req.ApplicationID = clt.AppID()
	req.FormatVersion = 2
	var resp model.ProductSearchResponse
	if err := clt.Get(ctx, "/Product/Search/20170426", req, &resp, ""); err != nil {
		return nil, err
	}
	return &resp, nil
}

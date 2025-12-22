package order

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/order"
)

// Search search order
func Search(ctx context.Context, clt *rms.Client, req *order.SearchRequest) (*order.SearchResponse, error) {
	var resp order.SearchResponse
	if err := clt.Post(ctx, "/es/2.0/order/searchOrder/", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp, nil
}

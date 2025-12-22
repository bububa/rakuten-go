package order

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/order"
)

// Get GetOrder
func Get(ctx context.Context, clt *rms.Client, req *order.GetRequest) (*order.GetResponse, error) {
	var resp order.GetResponse
	if err := clt.Post(ctx, "/es/2.0/order/getOrder/", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp, nil
}

package item20

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/item20"
)

// BulkGet 商品管理番号を指定し、商品情報を取得
func BulkGet(ctx context.Context, clt *rms.Client, req *item20.BulkGetRequest) ([]item20.Item, error) {
	var resp item20.BulkGetResponse
	if err := clt.Post(ctx, "/es/2.0/items/bulk_get", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return resp.Result, nil
}

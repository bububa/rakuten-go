package item20

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/item20"
)

// Search 商品管理番号を指定し、商品情報を取得
func Search(ctx context.Context, clt *rms.Client, req *item20.SearchRequest) (*item20.SearchResponse, error) {
	var resp item20.SearchResponse
	if err := clt.Get(ctx, "/es/2.0/items/search", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp, nil
}

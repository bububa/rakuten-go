package item20

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/item20"
)

// Get 商品管理番号を指定し、商品情報を取得
func Get(ctx context.Context, clt *rms.Client, itemManageNumber string) (*item20.Item, error) {
	var resp item20.GetResponse
	if err := clt.Get(ctx, util.StringsJoin("/es/2.0/items/manage-numbers/", itemManageNumber), nil, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.Item, nil
}

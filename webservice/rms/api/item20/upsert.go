package item20

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/item20"
)

// Upsert 商品管理番号を指定し、商品情報の登録・全項目の更新
func Upsert(ctx context.Context, clt *rms.Client, req *item20.UpsertRequest) error {
	var resp rmsModel.BaseResponse
	return clt.Put(ctx, util.StringsJoin("/es/2.0/items/manage-numbers/", req.ManageNumber), req, &resp, clt.TokenKey())
}

package item20

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/item20"
)

// Patch 商品管理番号を指定し、商品情報の部分更新
func Patch(ctx context.Context, clt *rms.Client, req *item20.PatchRequest) error {
	var resp rmsModel.BaseResponse
	return clt.Patch(ctx, util.StringsJoin("/es/2.0/items/manage-numbers/", req.ManageNumber), req, &resp, clt.TokenKey())
}

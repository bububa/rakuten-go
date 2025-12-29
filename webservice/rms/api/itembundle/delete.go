package itembundle

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// Delete delete ItemBundle
func Delete(ctx context.Context, clt *rms.Client, bundleManageNumber string) error {
	var resp rmsModel.BaseResponse
	return clt.Delete(ctx, util.StringsJoin("/es/1.0/bto/bundle", bundleManageNumber), nil, &resp, clt.TokenKey())
}

package coupon

import (
	"context"
	"strconv"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// GetThanksCoupon get thanks coupon
func GetThanksCoupon(ctx context.Context, clt *rms.Client, couponID int64) (*coupon.ThanksCoupon, error) {
	var resp rmsModel.XMLResponse[*coupon.ThanksCoupon]
	if err := clt.Get(ctx, util.StringsJoin("/es/1.0/thankscoupon/", strconv.FormatInt(couponID, 10)), nil, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return resp.Content, nil
}

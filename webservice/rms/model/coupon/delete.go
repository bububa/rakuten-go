package coupon

import (
	"encoding/xml"

	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// DeleteRequest delete Coupon Request
type DeleteRequest struct {
	XMLName xml.Name        `xml:"couponDeleteRequest" json:"-"`
	Coupon  *CouponToDelete `xml:"coupon" json:"coupon"`
}

type CouponToDelete struct {
	CouponCode string `xml:"couponCode"`
}

func (r *DeleteRequest) ContentType() string {
	return "application/xml"
}

func (r *DeleteRequest) Encode() []byte {
	req := rmsModel.NewXMLRequest(r)
	return req.Encode()
}

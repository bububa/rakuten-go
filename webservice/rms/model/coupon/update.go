package coupon

import (
	"encoding/xml"

	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// UpdateRequest Update Coupon Request
type UpdateRequest struct {
	XMLName xml.Name `xml:"couponUpdateRequest" json:"-"`
	Coupon  *Coupon  `json:"coupon,omitempty" xml:"coupon"`
}

func (r *UpdateRequest) ContentType() string {
	return "application/xml"
}

func (r *UpdateRequest) Encode() []byte {
	req := rmsModel.NewXMLRequest(r)
	return req.Encode()
}

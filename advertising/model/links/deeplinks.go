package links

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

type DeepLinksRequest struct {
	model.JSONRequest
	// URL product link to use for deep-linking
	URL string `json:"url,omitempty"`
	// AdvertiserID Advertiser ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// U1 u1 value
	U1 string `json:"u1,omitempty"`
}

// Encode implements Post Request
func (r DeepLinksRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type DeepLinksResponse struct {
	model.BaseResponse
	Advertiser *AdvertiserDeepLink `json:"advertiser,omitempty"`
}

type AdvertiserDeepLink struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	URL         string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	DeepLink    *DeepLink `json:"deep_link,omitempty"`
}

type DeepLink struct {
	URL         string `json:"url,omitempty"`
	DeepLinkURL string `json:"deep_link_url,omitempty"`
	U1          string `json:"u1,omitempty"`
}

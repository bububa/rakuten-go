package webservice

type BaseRequest struct {
	// ApplicationID アプリケーションID
	ApplicationID string `json:"applicationId,omitempty"`
	// AffiliateId アフェリエイトID
	AffiliateID string `json:"affiliateId,omitempty"`
	// FormatVersion
	FormatVersion int `json:"formatVersion,omitempty"`
}

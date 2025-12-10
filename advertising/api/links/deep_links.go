package links

import (
	"context"

	"github.com/bububa/rakuten-go/advertising"
	"github.com/bububa/rakuten-go/advertising/model/links"
)

// DeepLinks generate deep-links from partner-advertiser's product links
func DeepLinks(ctx context.Context, clt *advertising.Client, req *links.DeepLinksRequest, accessToken string) (*links.AdvertiserDeepLink, error) {
	var resp links.DeepLinksResponse
	if err := clt.Post(ctx, "/v1/links/deep_links", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Advertiser, nil
}

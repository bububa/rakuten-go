package product

import (
	"context"

	"github.com/bububa/rakuten-go/advertisering/core"
	"github.com/bububa/rakuten-go/advertisering/model/product"
)

// Search Use the Product Search API to search partner-advertisers’ product data. Users are able to search for products that match keywords in the following categories
// Product name
// Primary or Secondary category
// Short product description
// Long product description
// Keyword(s)
func Search(ctx context.Context, clt *core.SDKClient, req *product.SearchRequest, accessToken string) (*product.SearchResult, error) {
	var resp product.SearchResponse
	if err := clt.Get(ctx, "/productsearch/1.0", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return &resp.SearchResult, nil
}

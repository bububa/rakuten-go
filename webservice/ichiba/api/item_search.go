package api

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/ichiba"
	"github.com/bububa/rakuten-go/webservice/ichiba/model"
)

// ItemSearch Ichiba Item Search API
func ItemSearch(ctx context.Context, clt *ichiba.Client, req *model.ItemSearchRequest) (*model.ItemSearchResponse, error) {
	req.ApplicationID = clt.AppID()
	req.FormatVersion = 2
	var resp model.ItemSearchResponse
	if err := clt.Get(ctx, "/IchibaItem/Search/20220601", req, &resp, ""); err != nil {
		return nil, err
	}
	return &resp, nil
}

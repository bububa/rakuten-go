package api

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/ichiba"
	"github.com/bububa/rakuten-go/webservice/ichiba/model"
)

// TagSearch Ichiba Tag Search API
func TagSearch(ctx context.Context, clt *ichiba.Client, req *model.TagSearchRequest) (*model.TagSearchResponse, error) {
	req.ApplicationID = clt.AppID()
	req.FormatVersion = 2
	var resp model.TagSearchResponse
	if err := clt.Get(ctx, "/IchibaTag/Search/20140222", req, &resp, ""); err != nil {
		return nil, err
	}
	return &resp, nil
}

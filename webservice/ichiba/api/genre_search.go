package api

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/ichiba"
	"github.com/bububa/rakuten-go/webservice/ichiba/model"
)

// GenreSearch Ichiba Genre Search API
func GenreSearch(ctx context.Context, clt *ichiba.Client, req *model.GenreSearchRequest) (*model.GenreSearchResponse, error) {
	req.ApplicationID = clt.AppID()
	req.FormatVersion = 2
	var resp model.GenreSearchResponse
	if err := clt.Get(ctx, "/IchibaGenre/Search/20120723", req, &resp, ""); err != nil {
		return nil, err
	}
	return &resp, nil
}

package itembundle

import (
	"context"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/itembundle"
)

// Get get ItemBundle
func Get(ctx context.Context, clt *rms.Client, bundleManageNumber string) (*itembundle.Bundle, error) {
	var resp itembundle.GetResponse
	if err := clt.Get(ctx, util.StringsJoin("/es/1.0/bto/bundle", bundleManageNumber), nil, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return &resp.Bundle, nil
}

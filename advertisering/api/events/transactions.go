package events

import (
	"context"

	"github.com/bububa/rakuten-go/advertisering/core"
	"github.com/bububa/rakuten-go/advertisering/model/events"
)

// Transactions Retrieve transaction confirmations of events completed on a partner-advertiser's website.
func Transactions(ctx context.Context, clt *core.SDKClient, req *events.TransactionsRequest, accessToken string) ([]events.Transaction, error) {
	var resp events.TransactionsResponse
	if err := clt.Get(ctx, "/events/1.0/transactions", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Transactions, nil
}

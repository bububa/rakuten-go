package events

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/bububa/rakuten-go/advertisering/enum"
	"github.com/bububa/rakuten-go/advertisering/model"
	"github.com/bububa/rakuten-go/util"
)

type TransactionsRequest struct {
	model.JSONRequest
	// ProcessDateStart start of date range. Returns transactions processed on or after this date. (YYYY-MM-DD HH:mm:ss)
	ProcessDateStart string `json:"process_date_start,omitempty"`
	// ProcessDateEnd end of date range. Returns transactions processed on or before this date. (YYYY-MM-DD HH:mm:ss)
	ProcessDateEnd string `json:"process_date_end,omitempty"`
	// TransactionDateStart start of date range. Returns transactions that occurred on or after this date. (YYYY-MM-DD HH:mm:ss)
	TransactionDateStart string `json:"transaction_date_start,omitempty"`
	// TransactionDateEnd end of date range. Returns transactions that occurres on or before this date. (YYYY-MM-DD HH:mm:ss)
	TransactionDateEnd string `json:"transaction_date_end,omitempty"`
	// Limit maximum number of results (default: 100)
	Limit int `json:"limit,omitempty"`
	// Page page number of transactions to retrieve (Default 1)
	Page int `json:"page,omitempty"`
	// Currency currency of the sale amount for transactions
	Currency enum.Currency `json:"currency,omitempty"`
	// Type specify if the transaction type is a real time or batched
	Type enum.TransactionType `json:"type,omitempty"`
}

// Encode implements Get Request
func (r TransactionsRequest) Encode() string {
	values := util.NewURLValues()
	if r.ProcessDateStart != "" {
		values.Set("process_date_start", r.ProcessDateStart)
	}
	if r.ProcessDateEnd != "" {
		values.Set("process_date_end", r.ProcessDateEnd)
	}
	if r.TransactionDateStart != "" {
		values.Set("transaction_date_start", r.TransactionDateStart)
	}
	if r.TransactionDateEnd != "" {
		values.Set("transaction_date_end", r.TransactionDateEnd)
	}
	if r.Limit > 0 {
		values.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.Currency != "" {
		values.Set("currency", string(r.Currency))
	}
	if r.Type != "" {
		values.Set("type", string(r.Type))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

type TransactionsResponse struct {
	model.BaseResponse
	Transactions []Transaction `json:"transactions,omitempty"`
}

// Unmarshal implements model.Unmarshaler
func (resp *TransactionsResponse) Unmarshal(r io.Reader) error {
	bs, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bs, &resp.BaseResponse); err == nil {
		if resp.IsError() {
			return resp
		}
	}
	return json.Unmarshal(bs, &resp.Transactions)
}

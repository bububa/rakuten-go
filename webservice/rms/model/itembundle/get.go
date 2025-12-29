package itembundle

import (
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// GetResponse は楽天ペイ受注APIの注文情報取得時の検索条件です API Response
type GetResponse struct {
	rmsModel.BaseResponse
	Bundle
}

package order

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// GetRequest は楽天ペイ受注APIの注文情報取得時の検索条件です API Request
type GetRequest struct {
	model.JSONRequest
	// OrderNumberList は注文番号リストです。最大100件まで指定可能で、過去2年分の注文が対象です。
	// この項目は必須です。
	OrderNumberList []string `json:"orderNumberList"`

	// Version はバージョン番号です。現在は必須ではありませんが、今後は必須化される予定です。
	// 次のいずれかが入力されます。
	// 3: 消費税増税対応
	// 4: 共通の送料込みライン対応
	// 5: 領収書、前払い期限版
	// 6: 顧客・配送対応注意表示詳細対応
	// 7: SKU対応
	Version int `json:"version"`
}

// Encode implements PostRequest interface
func (r *GetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// GetResponse は楽天ペイ受注APIの注文情報取得時の検索条件です API Response
type GetResponse struct {
	rmsModel.BaseResponse
	// GetOrderMessageModelList はメッセージモデルリストです。
	MessageModelList []OrderMessage `json:"MessageModelList"`

	// OrderModelList は受注モデルリストです。
	OrderModelList *[]Order `json:"OrderModelList,omitempty"`

	// version バージョン番号
	Version int `json:"version"`
}

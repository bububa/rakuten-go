package model

// Message common message response
type Message struct {
	// MessgeType はメッセージ種別です。以下のいずれかが入力されます。
	// ・INFO
	// ・ERROR
	// ・WARNING
	MessageType string `json:"messageType"`

	// MessageCode はメッセージコードです。
	// メッセージコードは https://webservice.rms.rakuten.co.jp/merchant-portal/view?contents=/ja/common/1-1_service_index/rakutenpayorderapi/rakutenpaymsgcodereference を参照してください(要ログイン)
	MessageCode string `json:"messageCode"`

	// Message はメッセージです。
	Message string `json:"message"`
}

// Sort は楽天ペイ受注APIで注文検索の検索条件のうち、ソートに関する条件です。
type Sort struct {
	// SortColumn は並び替え項目です。以下のいずれかを指定することができます。
	// 1: 注文日時
	SortColumn int `json:"sortColumn,omitempty"`

	// SortDirection は並び替え方法です。以下のいずれかを指定することができます。
	// 1: 昇順
	// 2: 降順
	SortDirection int `json:"sortDirection,omitempty"`
}

// Pagination は楽天ペイ受注APIで注文検索の検索条件のうち、ページングに関する条件です。
type Pagination struct {
	// RequestRecordsAmount は1ページあたりの取得結果数です。最大1,000件まで取得可能です。
	RequestRecordsAmount int `json:"requestRecordsAmount,omitempty"`

	// RequestPage はリクエストページ番号です。
	RequestPage int `json:"requestPage,omitempty"`

	// SortModelList はソート条件です。
	SortModelList []Sort `json:"SortModelList,omitempty"`
}

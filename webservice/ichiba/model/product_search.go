package model

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util/query"
	"github.com/bububa/rakuten-go/webservice/ichiba"
)

// ProductSearchRequest 商品価格ナビ製品検索 API Request
type ProductSearchRequest struct {
	ichiba.BaseRequest
	// Keyword 	検索キーワード
	// UTF-8でURLエンコードした文字列
	// (*1)検索キーワード、ジャンルIDのいずれかが指定されていることが必須です。
	Keyword string `json:"keyword,omitempty"`
	// GenreID ジャンルID
	// 楽天市場におけるジャンルを特定するためのID
	// ジャンル名、ジャンルの親子関係を調べたい場合は、「楽天ジャンル検索API(GenreSearch)」をご利用ください
	// (*1)検索キーワード、ジャンルIDのいずれかが指定されていることが必須です。
	GenreID int64 `json:"genreId,omitempty"`
	// ProductID 楽天プロダクト製品ID
	// 楽天プロダクト用の製品ID 他のAPI入力パラメーターと併用することはできません。
	// ※共通パラメーターを除く
	// (*1)検索キーワード、ジャンルID、楽天プロダクト製品IDのいずれかが指定されていることが必須です。
	ProductID string `json:"productId,omitempty"`
	// Hits How many results to display on each page
	// An integer between 1 and 30
	Hits int `json:"hits,omitempty"`
	// Page Result page
	// An integer between 1 and 100
	Page int `json:"page,omitempty"`
	// Sort 	ソート
	// standard：
	// 楽天標準ソート順
	// -releaseDate：
	// 発売日順（降順）
	// -seller：
	// 売上順（降順）
	// -satisfied：
	// 満足順（降順）
	// ※UTF-8でURLエンコードされている必要があります。
	Sort string `json:"sort,omitempty"`
	// MinPrice Minimum price
	// An integer greater than 0 and less than 999,999,999
	MinPrice int64 `json:"minPrice,omitempty"`
	// MaxPrice Maximum price
	// An integer greater than 0 and less than 999,999,999
	// maxPrice must be larger than minPrice.
	MaxPrice int `json:"maxPrice,omitempty"`
	// OrFlag OR search flag
	// Choose between AND searches and OR searches when there are multiple keywords.
	// 0: AND
	// 1: OR
	// *It isn't possible to use a complex search condition like "(A and B) or C".
	OrFlag int `json:"orFlag,omitempty"`
	// GenreInformationFlag Genre information flag
	// 0: Do not get number of item in each genre.
	// 1: Get number of item in each genre.
	GenreInformationFlag int `json:"genreInformationFlag,omitempty"`
}

// Encode implements GetRequest interface
func (r *ProductSearchRequest) Encode() string {
	values, err := query.Values(r)
	if err != nil {
		return ""
	}
	return values.Encode()
}

// ProductSearchResponse 商品価格ナビ製品検索 API Response
type ProductSearchResponse struct {
	model.BaseResponse
	// Products product list
	Products []Product `json:"Products,omitempty"`
	// Count Total number of items in the search results
	Count int `json:"count"`
	// First First page Result number to display from
	First int `json:"first"`
	// Hits Search hits Number of results to return
	Hits int `json:"hits"`
	// Last Last page Result number to display until
	Last int `json:"last"`
	// Page Current page number
	Page int `json:"page"`
	// PageCount Total page count Maximum of 100
	PageCount int `json:"pageCount"`
	// GenreInformation
	GenreInformation []GenreInformation `json:"GenreInformation,omitempty"`
}

type Product struct {
	// ProductID 楽天プロダクト製品ID
	ProductID string `json:"productId,omitempty"`
	// ProductName 製品名
	ProductName string `json:"productName,omitempty"`
	// ProductNo 型番
	ProductNo string `json:"productNo,omitempty"`
	// BrandName ブランド名
	BrandName string `json:"brandName,omitempty"`
	// ProductURLPC 製品ページURL(PC)
	ProductURLPC string `json:"productUrlPC,omitempty"`
	// ProductURLMobile 製品ページURL(モバイル)
	ProductURLMobile string `json:"productUrlMobile,omitempty"`
	// AffiliateURL アフィリエイトURL
	AffiliateURL string `json:"affiliateUrl,omitempty"`
	// SmallImageURL 製品画像64x64URL
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	// MediumImageURL 製品画像128x128URL
	MediumImageURL string `json:"mediumImageUrl,omitempty"`
	// ProductCaption 製品説明文
	ProductCaption string `json:"productCaption,omitempty"`
	// ReleaseDate 発売年月日
	ReleaseDate string `json:"releaseDate,omitempty"`
	// MakerCode 楽天プロダクトメーカーコード
	MakerCode string `json:"makerCode,omitempty"`
	// MakerName 会社名
	MakerName string `json:"makerName,omitempty"`
	// MakerNameKana 会社名カナ
	MakerNameKana string `json:"makerNameKana,omitempty"`
	// MakerNameFormal 会社正式名称
	MakerNameFormal string `json:"makerNameFormal,omitempty"`
	// MakerPageURLPC メーカーページURL（PC）
	MakerPageURLPC string `json:"makerPageUrlPC,omitempty"`
	// MakerPageURLMobile メーカーページURL（モバイル）
	MakerPageURLMobile string `json:"makerPageUrlMobile,omitempty"`
	// ItemCount 商品数
	ItemCount int64 `json:"itemCount,omitempty"`
	// SalesItemCount 購入可能な商品数
	SalesItemCount int64 `json:"salesItemCount,omitempty"`
	// UsedExcludeCount 中古を除く商品数
	UsedExcludeCount int64 `json:"usedExcludeCount,omitempty"`
	// UsedExcludeSalesCount 中古を除く購買可能な商品数
	UsedExcludeSalesCount int64 `json:"usedExcludeSalesCount,omitempty"`
	// MaxPrice 最高価格
	MaxPrice int64 `json:"maxPrice,omitempty"`
	// SalesMaxPrice 購入可能な最高価格
	SalesMaxPrice int64 `json:"salesMaxPrice,omitempty"`
	// UsedExcludeMaxPrice 中古を除く最高価格
	UsedExcludeMaxPrice int64 `json:"usedExcludeMaxPrice,omitempty"`
	// UsedExcludeSalesMaxPrice 中古を除く購入可能な最高価格
	UsedExcludeSalesMaxPrice int64 `json:"usedExcludeSalesMaxPrice,omitempty"`
	// MinPrice 最低価格
	MinPrice int64 `json:"minPrice,omitempty"`
	// SalesMinPrice 購入可能な最低価格
	SalesMinPrice int64 `json:"salesMinPrice,omitempty"`
	// UsedExcludeMinPrice 中古を除く最低価格
	UsedExcludeMinPrice int64 `json:"usedExcludeMinPrice,omitempty"`
	// UsedExcludeSalesMinPrice 中古を除く購入可能な最低価格
	UsedExcludeSalesMinPrice int64 `json:"usedExcludeSalesMinPrice,omitempty"`
	// AveragePrice 平均価格
	AveragePrice int64 `json:"averagePrice,omitempty"`
	// ReviewCount レビュー数
	ReviewCount int64 `json:"reviewCount,omitempty"`
	// ReviewAverage レビュー平均
	ReviewAverage float64 `json:"reviewAverage,omitempty"`
	// ReviewURLPC レビューページURL(PC)
	ReviewURLPC string `json:"reviewUrlPC,omitempty"`
	// ReviewURLMobile レビューページURL(モバイル)
	ReviewURLMobile string `json:"reviewUrlMobile,omitempty"`
	// Rank ランキング順位
	Rank int64 `json:"rank,omitempty"`
	// RankTargetGenreID ランキングの対象ジャンルID
	RankTargetGenreID model.Int64 `json:"rankTargetGenreId,omitempty"`
	// RankTargetProductCount ランキングの対象製品数
	RankTargetProductCount int64 `json:"rankTargetProductCount,omitempty"`
	// GenreID ジャンルID
	GenreID model.Int64 `json:"genreId,omitempty"`
	// GenreName ジャンル名/td>
	GenreName string `json:"genreName,omitempty"`
	// ProductDetails 製品詳細
	ProductDetails []ProductDetail `json:"ProductDetails,omitempty"`
}

// ProductDetail 製品詳細
type ProductDetail struct {
	// Name 項目名
	Name string `json:"name,omitempty"`
	// Value 項目値
	Value string `json:"value,omitempty"`
}

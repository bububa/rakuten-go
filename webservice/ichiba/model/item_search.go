package model

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util/query"
	"github.com/bububa/rakuten-go/webservice/ichiba"
)

// ItemSearchRequest Ichiba Item Search API Request
type ItemSearchRequest struct {
	ichiba.BaseRequest
	// Keyword UTF-8 URL encoded string
	// The keyword parameter can have a maximum length of 128 single byte characters
	// The keyword parameter is delimited with single byte space characters. This defaults to an AND operation including all the keywords. To use OR instead set the orFlag to 1.
	// Each search keyword must be at least two single byte characters or one double byte character.
	// An exception is a minimum of two characters if the search keywords are using hiragana, katakana, or symbols.
	// (*1) One of the following is required: search keyword, genre ID, item code, or shop code.
	Keyword string `json:"keyword,omitempty"`
	// ShopCode Shop Code The portion of the shop URL listed as "xyz": http://www.rakuten.co.jp/[xyz]
	// (*1) One of the following is required: search keyword, genre ID, item code, or shop code.
	ShopCode string `json:"shopCode,omitempty"`
	// ItemCode Item Code Included in the output parameters of the Item Search, Item Ranking APIs.
	// A value in the form of "shop:1234"
	// (*1) A search request must contain either a keyword, a genre ID, or an item code.
	ItemCode string `json:"itemCode,omitempty"`
	// GenreID Genre ID to specify a genre in Rakuten Ichiba
	// Please use the Rakuten Ichiba Genre Search API 2 to look up genre names and genre relations.
	// (*1) One of the following is required: search keyword, genre ID, item code, or shop code.
	GenreID int64 `json:"genreId,omitempty"`
	// TagID Tag ID Comma separated (maximum 10 IDs)
	// ID to specify a tag in Rakuten Ichiba.
	TagID string `json:"tagId,omitempty"`
	// Hits How many results to display on each page
	// An integer between 1 and 30
	Hits int `json:"hits,omitempty"`
	// Page Result page
	// An integer between 1 and 100
	Page int `json:"page,omitempty"`
	// Sort +affiliateRate：
	// Affiliate Rate ( Ascending order )
	// -affiliateRate：
	// Affiliate Rate ( Descending order )
	// +reviewCount：
	// Number of reviews ( Ascending order )
	// -reviewCount：
	// Number of reviews ( Descending order )
	// +reviewAverage：
	// Average review rating ( Ascending order )
	// -reviewAverage：
	// Average review rating ( Descending order )
	// +itemPrice：
	// Price ( Ascending order )
	// -itemPrice：
	// Price ( Descending order )
	// +updateTimestamp：
	// Time of item update ( Ascending order )
	// -updateTimestamp：
	// Time of item update ( Descending order )
	// standard：
	// Rakuten standard sort
	// *UTF-8 URL encoding is required.
	Sort string `json:"sort,omitempty"`
	// MinPrice Minimum price
	// An integer greater than 0 and less than 999,999,999
	MinPrice int64 `json:"minPrice,omitempty"`
	// MaxPrice Maximum price
	// An integer greater than 0 and less than 999,999,999
	// maxPrice must be larger than minPrice.
	MaxPrice int `json:"maxPrice,omitempty"`
	// Availability
	// 0: All products
	// 1: Only available products
	Availability int `json:"availability,omitempty"`
	// Field Search field
	// 0: Broad search (prefer more matches with the same keyword)
	// 1: Restricted search (prefer fewer matches with the same keyword)
	Field string `json:"field,omitempty"`
	// Carrier Platform
	// Choose the target platform for type of information to return: PC, mobile, or smartphone
	// PC: 0
	// mobile: 1 (*Japan only)
	// smartphone: 2
	// ImageFlag Has image flag
	// 0: Search all products
	// 1: Search only products with images
	ImageFlag int `json:"imageFlag,omitempty"`
	// OrFlag OR search flag
	// Choose between AND searches and OR searches when there are multiple keywords.
	// 0: AND
	// 1: OR
	// *It isn't possible to use a complex search condition like "(A and B) or C".
	OrFlag int `json:"orFlag,omitempty"`
	// NGKeyword Excluded keywords
	// Words to exclude from search results
	// Strings encoded with UTF-8 URL encoding
	// Same format as keyword
	NGKeyword string `json:"NGKeyword,omitempty"`
	// PurchaseType Narrow search by purchase type
	// 0: Normal purchase
	// 1: Periodic purchase (a service that allows customers to buy goods they choose on a cycle they specify).
	// 2: Distribution group purchase (a service that allows the shop to choose items and also choose the number of times they are delivered).
	PurchaseType int `json:"purchaseType,omitempty"`
	// ShipOverseasFlag Overseas shipping flag
	// 0: All products
	// 1: Only products that can be shipped overseas
	ShipOverseasFlag int `json:"shipOverseasFlag,omitempty"`
	// ShipOverseasArea Overseas shipping area
	// It is possible to restrict searches based on delivery areas.
	// Check the overseas delivery area codes
	// *shipOverseasFlag must be 1
	ShipOverseasArea string `json:"shipOverseasArea,omitempty"`
	// AsurakuFlag Asuraku flag
	// 0: All products
	// 1: Only items that can use the Asuraku service
	// *Starting July 1st onwards, the parameter "asurakuFlag" will always be set to the value of "0".
	AsurakuFlag int `json:"asurakuFlag,omitempty"`
	// AsurakuArea Asuraku area
	// It is possible to restrict searches based on delivery areas.
	// Check the Asuraku delivery area codes
	// *asurakuFlag must be 1
	AsurakuArea string `json:"asurakuArea,omitempty"`
	// PointRateFlag Point multiplier flag
	// 0: All products
	// 1: Only items with point multipliers
	PointRateFlag int `json:"pointRateFlag,omitempty"`
	// PointRate Point multiplier
	// An integer from 2 to 10 (e.g. 5 means points will be multiplied by five)
	// Read about point multipliers.
	// *pointRateFlag must be 1
	PointRate int `json:"pointRate,omitempty"`
	// PostageFlag
	// 0: All products
	// 1: Only items with postage included/free shipping
	PostageFlag int `json:"postageFlag,omitempty"`
	// CreditCardFlag Credit card flag
	// 0: All products
	// 1: Only products purchaseable with credit cards
	CreditCardFlag int `json:"creditCardFlag,omitempty"`
	// GiftFlag Giftwrap flag
	// 0: All products
	// 1: Only products with giftwrapping available
	GiftFlag int `json:"giftFlag,omitempty"`
	// HasReviewFlag Has review flag
	// 0: All products
	// 1: Only products with reviews
	HasReviewFlag int `json:"hasReviewFlag,omitempty"`
	// MaxAffiliateRate Maximum affiliate rate
	// A number from 1.0 to 99.9 (e.g. 5.0 for products with an affiliate rate up to 5%).
	// Don't show if the affiliate rate is higher than specified.
	// Up to one decimal place can be specified.
	MaxAffiliateRate float64 `json:"maxAffiliateRate,omitempty"`
	// MinAffiliateRate Minimum affiliate rate
	// A number from 1.0 to 99.9 (e.g. 5.0 for products with an affiliate rate above 5%).
	// Don't show if the affiliate rate is lower than specified.
	// Up to one decimal place can be specified.
	// Specify an affiliate rate lower than the maximum rate.
	MinAffiliateRate float64 `json:"minAffiliateRate,omitempty"`
	// HasMovieFlag Has movie flag
	// 0: All products
	// 1: Only products with movies (the link to the movies will be returned)
	HasMovieFlag int `json:"hasMovieFlag,omitempty"`
	// PamphletFlag Pamphlet flag
	// 0: All products
	// 1: Only products with pamphlets
	PamphletFlag int `json:"pamphletFlag,omitempty"`
	// AppointDeliveryDateFlag 	Specify delivery date flag
	// 0: All products
	// 1: Only products for which the delivery date can be specified.
	AppointDeliveryDateFlag int `json:"appointDeliveryDateFlag,omitempty"`
	// Elements Output elements
	// Comma separated.
	// When required output parameters are specified, only those parameters will be returned.
	// (e.g elements=reviewCount,reviewAverage)
	Elements string `json:"elements,omitempty"`
	// GenreInformationFlag Genre information flag
	// 0: Do not get number of item in each genre.
	// 1: Get number of item in each genre.
	GenreInformationFlag int `json:"genreInformationFlag,omitempty"`
	// TagInformationFlag Tag information flag
	// 0: Do not get number of item in each tag.
	// 1: Get number of item in each tag.
	TagInformationFlag int `json:"tagInformationFlag,omitempty"`
}

// Encode implements GetRequest interface
func (r *ItemSearchRequest) Encode() string {
	values, err := query.Values(r)
	if err != nil {
		return ""
	}
	return values.Encode()
}

// ItemSearchResponse Ichiba Item Search API Response
type ItemSearchResponse struct {
	model.BaseResponse
	// Items Item information
	Items []Item `json:"items,omitempty"`
	// Carrier Platform
	// PC: 0
	// mobile: 1 (*Japan only)
	// smartphone: 2
	Carrier int `json:"carrier"`
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
	// TagInformation
	TagInformation []TagGroup `json:"TagInformation,omitempty"`
}

// Item Item Information
type Item struct {
	// AffiliateURL affiliate URL
	// (Only when the Affiliate ID has been included as an input parameter)
	// *Returns a URL regardless of the value of carrier
	// Affiliate URL, starting with https
	AffiliateURL string `json:"affiliateUrl,omitempty"`
	// Catchcopy Sales message
	Catchcopy string `json:"catchcopy,omitempty"`
	// ItemCaption item caption
	// *Depending on the setting for carrier the results may vary.
	ItemCaption string `json:"itemCaption,omitempty"`
	// ItemCode Item Code
	ItemCode string `json:"itemCode,omitempty"`
	// ItemName Item Name
	// To display the usual name please use catchcopy+itemname.
	// *Depending on the setting for carrier the results may vary.
	ItemName string `json:"itemName,omitempty"`
	// ItemPrice item price
	ItemPrice int64 `json:"itemPrice,omitempty"`
	// ItemPriceBaseField item price base field
	// Contains one of the following strings
	// "item_price_min1", "item_price_min2" or "item_price_min3"
	ItemPriceBaseField string `json:"itemPriceBaseField,omitempty"`
	// ItemPriceMax1 item price max 1
	// All item price max
	ItemPriceMax1 int64 `json:"itemPriceMax1,omitempty"`
	// ItemPriceMax2 item price max 2
	// Searchable item price max
	ItemPriceMax2 int64 `json:"itemPriceMax2,omitempty"`
	// ItemPriceMax3 item price max 3
	// Purchasable item price max
	ItemPriceMax3 int64 `json:"itemPriceMax3,omitempty"`
	// ItemPriceMin1 item price min 1
	// All item price min
	ItemPriceMin1 int64 `json:"itemPriceMin1,omitempty"`
	// ItemPriceMin2 item price min 2
	// Searchable item price min
	ItemPriceMin2 int64 `json:"itemPriceMin2,omitempty"`
	// ItemPriceMin3 item price min 3
	// Purchasable item price min
	ItemPriceMin3 int64 `json:"itemPriceMin3,omitempty"`
	// ItemURL item URL
	// *Depending on the setting for carrier the results may vary.
	// *Depending on the setting for carrier the results may vary. If you include affiliateId in the input parameter, itemUrl will be same as that of affiliateUrl.(From 2015/7/1)
	// URL for each item, starting with https.
	ItemURL string `json:"itemUrl,omitempty"`
	// SmallImageURLs 64 pixel square item image URL
	// Returns an array of up to three image URLs. Images size is 64 pixels square.
	// Small Image URL (64x64), starting with https.
	SmallImageURLs []string `json:"smallImageUrls,omitempty"`
	// MediumImageURLs 128 pixel square item image URL
	// Returns an array of up to three image URLs. Images size is 128 pixels square.
	// Medium Image URL (128x128), starting with https.
	MediumImageURLs []string `json:"mediumImageUrls,omitempty"`
	// ShopCode The portion marked "xyz" in each shop's URL: http://
	// www.rakuten.co.jp/[xyz]
	ShopCode string `json:"shopCode,omitempty"`
	// ShopName shop name
	ShopName string `json:"shopName,omitempty"`
	// ShopURL URL for each shop, starting with https
	// *If you include affiliateId in the input parameter, shopUrl will be same as that of shopAffiliateUrl.(From 2015/7/1)
	ShopURL string `json:"shopUrl,omitempty"`
	// ShopAffiliateURL shop affiliate URL
	// (Only when the Affiliate ID has been included as an input parameter)
	// *Returns a URL regardless of the value of carrier
	// Shop Affiliate URL, starting with https
	ShopAffiliateURL string `json:"shopAffiliateUrl,omitempty"`
	// TagIDs Multiple tag IDs are returned as an array
	TagIds []int `json:"tagIds,omitempty"`
	// GenreID genre id
	GenreID model.Int64 `json:"genreId,omitempty"`
	// Availability availability flag
	// 0: Out of stock
	// 1: Available
	Availability int `json:"availability,omitempty"`
	// ImageFlag has image flag
	// 0: Items without images
	// 1: Items with images
	ImageFlag int `json:"imageFlag,omitempty"`
	// TaxFlag  tax flag
	// 0: Tax included
	// 1: Tax not included
	TaxFlag int `json:"taxFlag,omitempty"`
	// PostageFlag postage flag
	// 0: Postage included
	// 1: Postage not included
	PostageFlag int `json:"postageFlag,omitempty"`
	// CreditCardFlag credit card flag
	// 0: Credit cards not accepted
	// 1: Credit cards accepted
	CreditCardFlag int `json:"creditCardFlag,omitempty"`
	// ShopOfTheYearFlag Shop of the Year flag
	// 0: Shops that have not won Shop of the Year
	// 1: Shops that have won Shop of the Year
	ShopOfTheYearFlag int `json:"shopOfTheYearFlag,omitempty"`
	// ShipOverseasFlag Overseas shipping flag
	// 0: Products that cannot be shipped overseas
	// 1: Products that can be shipped overseas
	ShipOverseasFlag int `json:"shipOverseasFlag,omitempty"`
	// ShipOverseasArea Overseas shipping area
	// Supported countries are shown, separated by slashes. (/)
	ShipOverseasArea string `json:"shipOverseasArea,omitempty"`
	// AsurakuFlag Asuraku flag
	// 0: No next day delivery
	// 1: Next day delivery possible
	// *Starting July 1st onwards, the parameter "asurakuFlag" will always be set to the value of "0".
	// *Read more about Asaraku
	AsurakuFlag int `json:"asurakuFlag,omitempty"`
	// AsurakuClosingTime Asaraku closing time
	// Returned value takes this form: "HH:MM"
	// *Is only returned when asarakuFlag is set to 1
	AsurakuClosingTime string `json:"asurakuClosingTime,omitempty"`
	// AsurakuArea Asuraku delivery area
	// Supported areas are shown, separated by slashes. (/)
	// *asurakuFlag must be 1 for this to be returned.
	AsurakuArea string `json:"asurakuArea,omitempty"`
	// AffiliateRate Affiliate rate
	AffiliateRate float64 `json:"affiliateRate,omitempty"`
	// StartTime sale start time
	// Only when a time limited sale has been set.
	// (format is "YYYY-MM-DD HH:MM")
	StartTime string `json:"startTime,omitempty"`
	// EndTime sale end time
	// Only when a time limited sale has been set.
	// (format is "YYYY-MM-DD HH:MM")
	EndTime string `json:"endTime,omitempty"`
	// ReviewCount review count
	ReviewCount int64 `json:"reviewCount,omitempty"`
	// ReviewAverage average review count
	ReviewAverage float64 `json:"reviewAverage,omitempty"`
	// PointRate Item point multiplier
	// (e.g. 5 for a point multiplier of "5 times points."
	// Read more about item-based point multipliers. *Only shows information for point multipliers that will not end within 24 hours.
	PointRate int `json:"pointRate,omitempty"`
	// PointRateStartTime Start of item point multiplier
	// Date and time of the start of the item-based point multiplier.
	// *Only shows information for point multipliers that will not end within 24 hours.
	PointRateStartTime string `json:"pointRateStartTime,omitempty"`
	// PointRateEndTime End of item point multiplier
	// Date and time of the end of the item-based point multiplier.
	// *Only shows information for point multipliers that will not end within 24 hours.
	PointRateEndTime string `json:"pointRateEndTime,omitempty"`
	// GiftFlag Gift wrap flag
	// 0: Gift wrap service unavailable
	// 1: Gift wrap service available
	GiftFlag int `json:"giftFlag"`
}

type GenreInformation struct {
	// Parents Parent genre
	// Inputted genre Id's parent genres
	// When genrePath=1 is set, the output will include ancestor genres. Ancester genres are displayed as "<parent>~ </parent>" and are listed inside "<parents>~ </parents>" tag.
	Parents []Genre `json:"parents,omitempty"`
	// Current Current genre
	// Genre ID from the input paramter
	Current []Genre `json:"current,omitempty"`
	// Children Child genre
	// Inputted genre Id's child genres
	// Child genres are displayed as "<child>~ </child>" and are listed inside "<children>~ </children>" tag
	// If genreId=0 is set, all genres that have genreLevel=1 will be listed as "<children>~ </children>"
	Children []Genre `json:"children,omitempty"`
}

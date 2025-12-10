package product

import (
	"encoding/xml"
	"strconv"

	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

// SearchRequest Product Search API request
type SearchRequest struct {
	model.XMLRequest
	// Keyword Use the parameter keyword to include all search terms in the result.
	Keyword string `json:"keyword,omitempty"`
	// Exact Use the parameter exact to return results containing the search terms in the exact order.
	Exact string `json:"exact,omitempty"`
	// One Use the one parameter to return results containing at least one of the search terms.
	One string `json:"one,omitempty"`
	// None Use the none parameter to return results containing none of the search terms.
	None string `json:"none,omitempty"`
	// Cat Filter results by primary or secondary category
	Cat string `json:"cat,omitempty"`
	// Language Filter results by language
	// Available values : en_US, fr_FR, de_DE, pt_BR
	Language string `json:"language,omitempty"`
	// Max Maximum results per page (default: 100)
	Max int `json:"max,omitempty"`
	// PageNumber Page number of results to retrieve (default: 1)
	PageNumber int `json:"pagenumber,omitempty"`
	// MID Advertiser ID
	MID uint64 `json:"mid,omitempty"`
	// Sort Sort results by one of the allowed fields
	// Available values : retailprice, productname, categoryname, mid
	Sort string `json:"sort,omitempty"`
	// SortType Sort results in ascending (asc) or descending (dsc) order
	// Available values : asc, dsc
	SortType string `json:"sorttype,omitempty"`
}

// Encode implements GetRequest
func (r SearchRequest) Encode() string {
	values := util.NewURLValues()
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
	}
	if r.Exact != "" {
		values.Set("exact", r.Exact)
	}
	if r.One != "" {
		values.Set("one", r.One)
	}
	if r.None != "" {
		values.Set("none", r.None)
	}
	if r.Cat != "" {
		values.Set("cat", r.Cat)
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.Max > 0 {
		values.Set("max", strconv.Itoa(r.Max))
	}
	if r.PageNumber > 0 {
		values.Set("pagenumber", strconv.Itoa(r.PageNumber))
	}
	if r.MID > 0 {
		values.Set("mid", strconv.FormatUint(r.MID, 10))
	}
	if r.Sort != "" {
		values.Set("sort", r.Sort)
	}
	if r.SortType != "" {
		values.Set("sorttype", r.SortType)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

type SearchResponse struct {
	XMLName xml.Name `xml:"result"`
	model.BaseResponse
	SearchResult
}

type SearchResult struct {
	TotalMatches int       `json:"total_matches" xml:"TotalMatches"`
	TotalPages   int       `json:"total_pages,omitempty" xml:"TotalPages"`
	PageNumber   int       `json:"page_number,omitempty" xml:"PageNumber"`
	Items        []Product `json:"items,omitempty" xml:"item"`
}

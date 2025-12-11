package model

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util/query"
	"github.com/bububa/rakuten-go/webservice"
)

// GenreSearchRequest Ichiba Genre Search API Request
type GenreSearchRequest struct {
	webservice.BaseRequest
	// GenreID Use genreId=0 to search from the genre root
	GenreID int64 `json:"genreId,omitempty"`
	// GenrePath genre path
	// 0: Exclude ancester genres
	// 1: Display ancester genres
	// Ancestor genres are a set of genres in the upper levels (parent, grandparent, etc.)
	GenrePath int `json:"genrePath,omitempty"`
}

// Encode implements GetRequest interface
func (r *GenreSearchRequest) Encode() string {
	values, err := query.Values(r)
	if err != nil {
		return ""
	}
	return values.Encode()
}

// GenreSearchResponse Ichiba Genre Search API Response
type GenreSearchResponse struct {
	model.BaseResponse
	// Parents Parent genre
	// Inputted genre Id's parent genres
	// When genrePath=1 is set, the output will include ancestor genres. Ancester genres are displayed as "<parent>~ </parent>" and are listed inside "<parents>~ </parents>" tag.
	Parents []Genre `json:"parents,omitempty"`
	// Current Current genre
	// Genre ID from the input paramter
	Current *Genre `json:"current,omitempty"`
	// Children Child genre
	// Inputted genre Id's child genres
	// Child genres are displayed as "<child>~ </child>" and are listed inside "<children>~ </children>" tag
	// If genreId=0 is set, all genres that have genreLevel=1 will be listed as "<children>~ </children>"
	Children []Genre `json:"children,omitempty"`
}

// Genre Genre information
type Genre struct {
	// GenreID Genre ID
	GenreID int64 `json:"genreId,omitempty"`
	// GenreName Genre name
	GenreName string `json:"genreName,omitempty"`
	// GenreLevel Genre level , use 0 to search at the genre root
	GenreLevel int `json:"genre_level,omitempty"`
}

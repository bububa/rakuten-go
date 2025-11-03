package product

// Product product details
type Product struct {
	MID          uint64       `json:"mid,omitempty" xml:"mid"`
	MerchantName string       `json:"merchant_name,omitempty" xml:"merchantname"`
	LinkID       uint64       `json:"link_id,omitempty" xml:"linkid"`
	CreatedOn    string       `json:"created_on,omitempty" xml:"createdon"`
	SKU          uint64       `json:"sku,omitempty" xml:"sku"`
	ProductName  string       `json:"product_name,omitempty" xml:"productname"`
	Category     *Category    `json:"category,omitempty" xml:"category"`
	Price        *Price       `json:"price,omitempty" xml:"price"`
	SalePrice    *Price       `json:"sale_price,omitempty" xml:"saleprice"`
	UPCCode      string       `json:"upcode,omitempty" xml:"upccode"`
	Description  *Description `json:"description,omitempty" xml:"description"`
	Keywords     string       `json:"keywords,omitempty" xml:"keywords"`
	LinkURL      string       `json:"link_url,omitempty" xml:"linkurl"`
	ImageURL     string       `json:"image_url,omitempty" xml:"imageurl"`
}

// Category is product category
type Category struct {
	Primary   string `json:"primary,omitempty" xml:"primary"`
	Secondary string `json:"secondary,omitempty" xml:"secondary"`
}

// Price is product price
type Price struct {
	Currency string  `json:"currency,omitempty" xml:"currency,attr"`
	Value    float64 `json:"value,omitempty" xml:",chardata"`
}

// Description is product description
type Description struct {
	Short string `json:"short,omitempty" xml:"short"`
	Long  string `json:"long,omitempty" xml:"long"`
}

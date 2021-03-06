package rakuten

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// IchibaItemSearch search ichiba item
func (c *Client) IchibaItemSearch(param *IchibaItemSearchParam) (*IchibaItemResult, error) {
	ichibaSearchURL := "/Search/20170706"

	param.ApplicationID = c.AppID
	format := "json" // default setting in rakuten API

	v := url.Values{}
	rv := reflect.ValueOf(param).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		name := f.Name
		key := f.Tag.Get("rws")
		value := rv.FieldByName(name).Interface().(string)
		if value != "" {
			v.Set(key, value)
			if name == "format" {
				format = value
			}
		}
	}

	url := fmt.Sprintf("%s%s?%s", c.BaseURL, ichibaSearchURL, v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data IchibaItemResult
	if format == "xml" {
		err = xml.NewDecoder(resp.Body).Decode(&data)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
	}
	if err != nil {
		return nil, err
	}

	return &data, err
}

// IchibaItemRanking ranking ichiba items
func (c *Client) IchibaItemRanking(param *IchibaRankingParam) (*IchibaRankingResult, error) {
	ichibaRankingURL := "/Ranking/20170628"
	param.ApplicationID = c.AppID
	format := "json" // default setting in rakuten API

	v := url.Values{}
	rv := reflect.ValueOf(param).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		name := f.Name
		key := f.Tag.Get("rws")
		value := rv.FieldByName(name).Interface().(string)
		if value != "" {
			v.Set(key, value)
			if name == "format" {
				format = value
			}
		}
	}

	url := fmt.Sprintf("%s%s?%s", c.BaseURL, ichibaRankingURL, v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data IchibaRankingResult
	if format == "xml" {
		err = xml.NewDecoder(resp.Body).Decode(&data)
	} else {
		err = json.NewDecoder(resp.Body).Decode(&data)
	}
	if err != nil {
		return nil, err
	}

	return &data, err
}

// IchibaItemSearchParam param for ichiba item search
type IchibaItemSearchParam struct {
	ApplicationID string `rws:"applicationId"`
	AffiliateID   string `rws:"affiliateId"`
	Format        string `rws:"format"`
	Callback      string `rws:"callback"`
	Elements      string `rws:"elements"`
	FormatVersion string `rws:"formatVersion"`

	NGKeyword               string `rws:"NGKeyword"`
	AppointDeliveryDateFlag string `rws:"appointDeliveryDateFlag"`
	AsurakuArea             string `rws:"asurakuArea"`
	AsurakuFlag             string `rws:"asurakuFlag"`
	Availability            string `rws:"availability"`
	Carrier                 string `rws:"carrier"`
	CreditCardFlag          string `rws:"creditCardFlag"`
	Field                   string `rws:"field"`
	GenreID                 string `rws:"genreId"`
	GenreInformationFlag    string `rws:"genreInformationFlag"`
	GiftFlag                string `rws:"giftFlag"`
	HasMovieFlag            string `rws:"hasMovieFlag"`
	HasReviewFlag           string `rws:"hasReviewFlag"`
	Hits                    string `rws:"hits"`
	ImageFlag               string `rws:"imageFlag"`
	ItemCode                string `rws:"itemCode"`
	Keyword                 string `rws:"keyword"`
	MaxAffiliateRate        string `rws:"maxAffiliateRate"`
	MaxPrice                string `rws:"maxPrice"`
	MinAffiliateRate        string `rws:"minAffiliateRate"`
	MinPrice                string `rws:"minPrice"`
	OrFlag                  string `rws:"orFlag"`
	Page                    string `rws:"page"`
	PamphletFlag            string `rws:"pamphletFlag"`
	PostringRate            string `rws:"postringRate"`
	PostringRateFlag        string `rws:"postringRateFlag"`
	PostageFlag             string `rws:"postageFlag"`
	PurchaseType            string `rws:"purchaseType"`
	ShipOverseasArea        string `rws:"shipOverseasArea"`
	ShipOverseasFlag        string `rws:"shipOverseasFlag"`
	ShopCode                string `rws:"shopCode"`
	Sort                    string `rws:"sort"`
	TagID                   string `rws:"tagId"`
	TagInformationFlag      string `rws:"tagInformationFlag"`
}

// IchibaItemResult is api response result for ichiba item result
type IchibaItemResult struct {
	Count            int           `json:"count"`
	Page             int           `json:"page"`
	First            int           `json:"first"`
	Last             int           `json:"last"`
	Hits             int           `json:"hits"`
	Carrier          int           `json:"carrier"`
	PageCount        int           `json:"pageCount"`
	Items            []Item        `json:"Items"`
	GenreInformation []interface{} `json:"GenreInformation"`
	TagInformation   []interface{} `json:"TagInformation"`
}

// Item is inner struct in ichiba item result
type Item struct {
	Item struct {
		ItemName       string `json:"itemName"`
		Catchcopy      string `json:"catchcopy"`
		ItemCode       string `json:"itemCode"`
		ItemPrice      int    `json:"itemPrice"`
		ItemCaption    string `json:"itemCaption"`
		ItemURL        string `json:"itemUrl"`
		ShopURL        string `json:"shopUrl"`
		SmallImageUrls []struct {
			ImageURL string `json:"imageUrl"`
		} `json:"smallImageUrls"`
		MediumImageUrls []struct {
			ImageURL string `json:"imageUrl"`
		} `json:"mediumImageUrls"`
		AffiliateURL       string  `json:"affiliateUrl"`
		ShopAffiliateURL   string  `json:"shopAffiliateUrl"`
		ImageFlag          int     `json:"imageFlag"`
		Availability       int     `json:"availability"`
		TaxFlag            int     `json:"taxFlag"`
		PostageFlag        int     `json:"postageFlag"`
		CreditCardFlag     int     `json:"creditCardFlag"`
		ShopOfTheYearFlag  int     `json:"shopOfTheYearFlag"`
		ShipOverseasFlag   int     `json:"shipOverseasFlag"`
		ShipOverseasArea   string  `json:"shipOverseasArea"`
		AsurakuFlag        int     `json:"asurakuFlag"`
		AsurakuClosingTime string  `json:"asurakuClosingTime"`
		AsurakuArea        string  `json:"asurakuArea"`
		AffiliateRate      float64 `json:"affiliateRate"`
		StartTime          string  `json:"startTime"`
		EndTime            string  `json:"endTime"`
		ReviewCount        int     `json:"reviewCount"`
		ReviewAverage      float64 `json:"reviewAverage"`
		PointRate          int     `json:"pointRate"`
		PointRateStartTime string  `json:"pointRateStartTime"`
		PointRateEndTime   string  `json:"pointRateEndTime"`
		GiftFlag           int     `json:"giftFlag"`
		ShopName           string  `json:"shopName"`
		ShopCode           string  `json:"shopCode"`
		GenreID            string  `json:"genreId"`
		TagIds             []int   `json:"tagIds"`
	} `json:"Item"`
}

// IchibaRankingParam param for ichiba ranking
type IchibaRankingParam struct {
	ApplicationID string `rws:"applicationId"`
	AffiliateID   string `rws:"affiliateId"`
	Format        string `rws:"format"`
	Callback      string `rws:"callback"`
	Elements      string `rws:"elements"`
	FormatVersion string `rws:"formatVersion"`

	GenreID     string `rws:"genreId"`
	AgeBased    string `rws:"age"`
	GenderBased string `rws:"sex"`
	Carrier     string `rws:"carrier"`
	Page        string `rws:"page"`
	Period      string `rws:"period"`
}

// IchibaRankingResult is api response result for ichiba ranking result
type IchibaRankingResult struct {
	Title         string        `json:"title"`
	LastBuildDate string        `json:"lastBuildDate"`
	Items         []RankingItem `json:"Items"`
}

// RankingItem is inner struct in ichiba item result
type RankingItem struct {
	Item struct {
		Rank           int    `json:"rank"`
		Carrier        int    `json:"carrier"`
		ItemName       string `json:"itemName"`
		Catchcopy      string `json:"catchcopy"`
		ItemCode       string `json:"itemCode"`
		ItemPrice      string `json:"itemPrice"`
		ItemCaption    string `json:"itemCaption"`
		ItemURL        string `json:"itemUrl"`
		AffiliateURL   string `json:"affiliateUrl"`
		ImageFlag      int    `json:"imageFlag"`
		SmallImageUrls []struct {
			ImageURL string `json:"imageUrl"`
		} `json:"smallImageUrls"`
		MediumImageUrls []struct {
			ImageURL string `json:"imageUrl"`
		} `json:"mediumImageUrls"`
		Availability       int    `json:"availability"`
		TaxFlag            int    `json:"taxFlag"`
		PostageFlag        int    `json:"postageFlag"`
		CreditCardFlag     int    `json:"creditCardFlag"`
		ShopOfTheYearFlag  int    `json:"shopOfTheYearFlag"`
		ShipOverseasFlag   int    `json:"shipOverseasFlag"`
		ShipOverseasArea   string `json:"shipOverseasArea"`
		AsurakuFlag        int    `json:"asurakuFlag"`
		AsurakuClosingTime string `json:"asurakuClosingTime"`
		AsurakuArea        string `json:"asurakuArea"`
		AffiliateRate      string `json:"affiliateRate"`
		StartTime          string `json:"startTime"`
		EndTime            string `json:"endTime"`
		ReviewCount        int    `json:"reviewCount"`
		ReviewAverage      string `json:"reviewAverage"`
		PointRate          int    `json:"pointRate"`
		PointRateStartTime string `json:"pointRateStartTime"`
		PointRateEndTime   string `json:"pointRateEndTime"`
		ShopName           string `json:"shopName"`
		ShopCode           string `json:"shopCode"`
		ShopURL            string `json:"shopUrl"`
		GenreID            string `json:"genreId"`

		// ShopAffiliateURL   string  `json:"shopAffiliateUrl"`
		// GiftFlag           int     `json:"giftFlag"`
		// TagIds             []int   `json:"tagIds"`
	} `json:"Item"`
}

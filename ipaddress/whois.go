package ipaddress

import (
	"net/url"
	"sync"
	"time"
)

const (
	whoisAPIURLBase = "https://ipwho.is"
)

var whoIsApiURL = sync.OnceValue(func() *url.URL {
	parsedURL, err := url.Parse(whoisAPIURLBase)
	if err != nil {
		panic("failed to parse WHOIS API URL: " + err.Error())
	}
	return parsedURL
})

type whoisResponse struct {
	Ip            string  `json:"ip"`
	Success       bool    `json:"success"`
	Type          string  `json:"type"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	IsEu          bool    `json:"is_eu"`
	Postal        string  `json:"postal"`
	CallingCode   string  `json:"calling_code"`
	Capital       string  `json:"capital"`
	Borders       string  `json:"borders"`
	Flag          struct {
		Img          string `json:"img"`
		Emoji        string `json:"emoji"`
		EmojiUnicode string `json:"emoji_unicode"`
	} `json:"flag"`
	Connection struct {
		Asn    int    `json:"asn"`
		Org    string `json:"org"`
		Isp    string `json:"isp"`
		Domain string `json:"domain"`
	} `json:"connection"`
	Timezone struct {
		Id          string    `json:"id"`
		Abbr        string    `json:"abbr"`
		IsDst       bool      `json:"is_dst"`
		Offset      int       `json:"offset"`
		Utc         string    `json:"utc"`
		CurrentTime time.Time `json:"current_time"`
	} `json:"timezone"`
}

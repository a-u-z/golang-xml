package main

import "encoding/xml"

type MarketDescriptions struct {
	XMLName      xml.Name `xml:"market_descriptions"`
	Market       []Market `xml:"market"`
	ResponseCode string   `xml:"response_code,attr"`
}

// attr 表示為 xml 標籤內的屬性，而非下一層級物件
type Market struct {
	XMLName  xml.Name `xml:"market"`
	ID       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	Groups   string   `xml:"groups,attr"`
	Outcomes Outcomes `xml:"outcomes"`
}

type Outcomes struct {
	XMLName xml.Name  `xml:"outcomes"`
	Outcome []Outcome `xml:"outcome"`
}

type Outcome struct {
	XMLName xml.Name `xml:"outcome"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

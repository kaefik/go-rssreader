package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
)

// http://www.w3schools.com/rss/rss_syntax.asp
// http://www.w3schools.com/rss/rss_channel.asp
type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}

type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

func main() {
	r := Rss2{}
	xmlContent, _ := ioutil.ReadFile("rss-data.xml")
	err := xml.Unmarshal(xmlContent, &r)
	if err != nil {
		panic(err)
	}
	for _, item := range r.ItemList {
		fmt.Println(item)
	}
}

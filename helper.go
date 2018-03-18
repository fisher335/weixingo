package main

import (
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"time"
	"strings"
	"strconv"
)

type Herf struct {
	Type    string
	Content string
	Img     string
}

func GetJock() Herf {
	c := make([]Herf, 0)
	rand.Seed(time.Now().Unix())
	url := "https://www.qiushibaike.com/text/"+strconv.Itoa(rand.Intn(12)+1)+"/"
	doc, _ := goquery.NewDocument(url)
	doc.Find("div[id^='qiushi']").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("div.content").Find("span").Text()
		//fmt.Printf(s)
		if s != "" {
			var d = Herf{}
			d.Content = s
			d.Type = "content"
			src, _ := selection.Find("img.illustration").Attr("src")
			if src != "" {
				d.Img = "http:" + src
				d.Type = "img"
				//fmt.Println(strings.Replace(src, "//", "http://", -1))
			}
			c = append(c, d)
		}
	})
	rand.Seed(time.Now().Unix())
	res := c[rand.Intn(len(c))]
	if strings.Contains(res.Content, "查看全文") {
		GetJock()
	}
	return res
}

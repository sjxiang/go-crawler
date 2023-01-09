package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/sjxiang/go-crawler/collect"
)


func main() {
	url := "https://book.douban.com/annual/2022?fullscreen=1&source=navigation"
	
	var f collect.Fetcher = collect.BaseFetch{}
	body, err := f.Get(url)

	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}


	// 加载 HTML 文档
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("read content failed:%v", err)
	}

	doc.Find("div.index-leftside h2").Each(func(i int, s *goquery.Selection) {
		// 获取匹配标签中的文本
		title := s.Text()
		fmt.Printf("%d, %s\n", i, title)
	})
	
}



func TestFeature(body string) {
	// 统计小 feature
	numLinks := strings.Count(body, "<a")
	fmt.Printf("当前页面链接总数为：%d 条\n", numLinks)

	exist := strings.Contains(body, "疫情")
	fmt.Printf("是否存在疫情相关：%v\n", exist)	
}
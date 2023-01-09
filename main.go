package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/sjxiang/go-crawler/util"
)


func main() {
	url := "https://www.thepaper.cn/"
	
	body, err := Fetch(url)

	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}


	// 统计小 feature
	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("当前页面链接总数为：%d 条\n", numLinks)

	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情相关：%v\n", exist)
	
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


func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		util.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())  // 将 HTML 文本从特殊编码转换成 UTF-8 编码

	return ioutil.ReadAll(utf8Reader)  
}



// 采样，确定编码属性
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	// 样本取 1024 字节
	// HTML 文本如果少于 1024 字节，请求有问题
	bytes, err := r.Peek(1024)

	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
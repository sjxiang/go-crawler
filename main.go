package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)  // []byte{}
	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("当前页面链接总数为：%d 条\n", numLinks)

	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情相关：%v\n", exist)
}
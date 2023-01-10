package collect

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/sjxiang/go-crawler/util"
)


type BaseFetch struct {}


func (BaseFetch) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		util.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v\n", resp.StatusCode)
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
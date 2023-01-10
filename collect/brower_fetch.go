package collect

import (
	"fmt"
	"net/http"
	"bufio"
	"io/ioutil"

	"golang.org/x/text/transform"

)

// 模拟浏览器访问
type BrowerFetch struct {}


func (BrowerFetch) Get(url string) ([]byte, error) {

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("get url failed:%v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())  // 将 HTML 文本从特殊编码转换成 UTF-8 编码

	return ioutil.ReadAll(utf8Reader)  
}


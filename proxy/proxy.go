package main

import (
	"io"
	"log"
	"net/http"

	"github.com/sjxiang/go-crawler/util"
)

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.Fatal("test")
	nextHTTP(w, r)
}

func main() {
	srv := &http.Server{
		Addr: ":8888",
		Handler: Router{},  // 需要实现 ServeHTTP 接口的实例
	}

	log.Fatal(srv.ListenAndServe())
}


func nextHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Transport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, r.Body)	
}


func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
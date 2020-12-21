package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var urlCount = 0

const (
	VANILA_SERVER = "http://127.0.0.1:1385"
	ECHO_SERVER   = "http://127.0.0.1:1387"
	GIN_SERVER    = "http://127.0.0.1:1389"
	PROXY_PORT    = "1335"
)

func main() {
	log.Printf("start at 127.0.0.1:%s", PROXY_PORT)
	http.HandleFunc("/api", toyLoadBalancer)

	log.Fatal(http.ListenAndServe(":"+PROXY_PORT, nil))
}

func toyLoadBalancer(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL()

	logRequestPayload(url)

	serveReverseUrl(url, res, req)
}

func getProxyURL() string {
	var urls = []string{VANILA_SERVER, ECHO_SERVER, GIN_SERVER}
	url := urls[urlCount]
	urlCount++

	if urlCount >= len(urls) {
		urlCount = 0
	}
	return url
}

func logRequestPayload(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}

func serveReverseUrl(target string, res http.ResponseWriter, req *http.Request) {
	url, err := url.Parse(target)

	if err != nil {
		panic(err)
	}

	director := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = url.Path + req.URL.Path
	}

	reverseProxy := &httputil.ReverseProxy{Director: director}
	reverseProxy.ServeHTTP(res, req)
}

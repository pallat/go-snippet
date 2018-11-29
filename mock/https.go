package mock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Client struct{}

func (Client) Do(*http.Request) (*http.Response, error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
  "args": {},
  "headers": {
    "Accept": "application/json",
    "Accept-Encoding": "gzip, deflate, br",
    "Accept-Language": "en-US,en;q=0.9,th-TH;q=0.8,th;q=0.7",
    "Connection": "close",
    "Cookie": "_gauges_unique_month=1; _gauges_unique_year=1; _gauges_unique=1; _gauges_unique_hour=1; _gauges_unique_day=1",
    "Host": "httpbin.org",
    "Referer": "https://httpbin.org/",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36"
  },
  "origin": "110.168.200.35",
  "url": "https://httpbin.org/get"
}`)
	}))
	defer ts.Close()

	return http.Get(ts.URL)
}

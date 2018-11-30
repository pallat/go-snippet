package example

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo"

	"github.com/pallat/gosnippet/mock"
)

func TestHttpBinOK(t *testing.T) {
	h := &Handler{C: mock.Client{}}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h.HTTPBin(c)

	expected := information{Args: struct{}{}, Headers: struct {
		Accept         string "json:\"Accept\""
		AcceptEncoding string "json:\"Accept-Encoding\""
		AcceptLanguage string "json:\"Accept-Language\""
		Connection     string "json:\"Connection\""
		Cookie         string "json:\"Cookie\""
		Host           string "json:\"Host\""
		Referer        string "json:\"Referer\""
		UserAgent      string "json:\"User-Agent\""
	}{Accept: "application/json", AcceptEncoding: "gzip, deflate, br", AcceptLanguage: "en-US,en;q=0.9,th-TH;q=0.8,th;q=0.7", Connection: "close", Cookie: "_gauges_unique_month=1; _gauges_unique_year=1; _gauges_unique=1; _gauges_unique_hour=1; _gauges_unique_day=1", Host: "httpbin.org", Referer: "https://httpbin.org/", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36"}, Origin: "110.168.200.35", URL: "https://httpbin.org/get"}

	var in information
	err := json.Unmarshal(rec.Body.Bytes(), &in)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(in, expected) {
		t.Errorf("%v\n is expected but get\n%v\n", expected, in)
	}
}

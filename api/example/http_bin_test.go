package example_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"

	"github.com/pallat/gosnippet/api/example"
	"github.com/pallat/gosnippet/mock"
)

func TestHttpBinOK(t *testing.T) {
	h := &example.Handler{C: mock.Client{}}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h.HTTPBin(c)

	fmt.Println(rec.Body.String())
}

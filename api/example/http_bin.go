package example

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pallat/gosnippet/pkg/https"
	"github.com/pallat/gosnippet/pkg/logs"
)

var url = "https://httpbin.org/get"

type Handler struct {
	C https.IClient
}

func (h *Handler) HTTPBin(c echo.Context) error {
	c.Logger().Infof(logs.Info(c, "HTTPBin is starting"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Logger().Errorf(logs.Error(c, err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := h.C.Do(req)
	if err != nil {
		c.Logger().Errorf(logs.Error(c, err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	var in information

	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&in)
	if err != nil {
		c.Logger().Errorf(logs.Error(c, err))
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer res.Body.Close()

	defer c.Logger().Infof(logs.Success(c))
	return c.JSON(http.StatusOK, information{})
}

type information struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
		Connection     string `json:"Connection"`
		Cookie         string `json:"Cookie"`
		Host           string `json:"Host"`
		Referer        string `json:"Referer"`
		UserAgent      string `json:"User-Agent"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func HTTPBinBusinessFail(c echo.Context) error {
	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "request",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

	res, err := https.Client.Get(url)
	if err != nil {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "T",
			"message": err.Error(),
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(http.StatusInternalServerError, err)
	}

	var in information

	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&in)
	if err != nil {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "T",
			"message": err.Error(),
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "response",
		"status":  "E",
		"message": "balance is not enough",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

	return c.JSON(http.StatusOK, information{})
}

func HTTPBinTechnicalFail(c echo.Context) error {
	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "request",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

	res, err := https.Client.Get("https://httpbin.org/status/404")
	if err != nil {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "T",
			"message": err.Error(),
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(http.StatusInternalServerError, err)
	}

	if res.StatusCode != http.StatusOK {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "T",
			"message": res.Status,
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(res.StatusCode, map[string]string{"message": res.Status})
	}

	var in information

	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&in)
	if err != nil {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "T",
			"message": err.Error(),
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "response",
		"status":  "E",
		"message": "balance is not enough",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

	return c.JSON(http.StatusOK, information{})
}

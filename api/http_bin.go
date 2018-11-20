package api

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

var url = "https://httpbin.org/get"

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	},
}

func HTTPBin(c echo.Context) error {
	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "request",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

	res, err := client.Get(url)
	if err != nil {
		c.Logger().Infof("json", map[string]string{
			"service": "HTTPBin",
			"state":   "request",
			"status":  "E",
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
			"status":  "E",
			"message": err.Error(),
			"id":      c.Response().Header().Get(echo.HeaderXRequestID),
		})
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Infof("json", map[string]string{
		"service": "HTTPBin",
		"state":   "response",
		"status":  "S",
		"id":      c.Response().Header().Get(echo.HeaderXRequestID),
	})

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

	res, err := client.Get(url)
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

	res, err := client.Get("https://httpbin.org/status/404")
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

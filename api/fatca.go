package api

import (
	"encoding/xml"
	"net/http"

	"ktb.co.th/api/prototype/pkg/errs"
	"ktb.co.th/api/prototype/pkg/handler"
)

type GetFatcaRequest struct{}
type GetFatcaResponse struct{}

func Fatca(c *handler.Context) {
	var req GetFatcaRequest

	c.Logger.Debug(`{"key": "value"}`)
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.New(err, "400", "Bad Request"))
		return
	}

	fatcaInfoRes, err := getPartyFATCAInfo(reformatToEnvelopeFATCAInfo(req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.New(err, "400", "Bad Request"))
		return
	}

	c.JSON(http.StatusOK, reformatToResponse(fatcaInfoRes))
}

func reformatToEnvelopeFATCAInfo(req GetFatcaRequest) EnvelopeFATCAInfo {
	return EnvelopeFATCAInfo{}
}

func reformatToResponse(EnvelopeFATCAInfoResponse) GetFatcaResponse {
	return GetFatcaResponse{}
}

func getPartyFATCAInfo(req EnvelopeFATCAInfo) (EnvelopeFATCAInfoResponse, error) {
	return EnvelopeFATCAInfoResponse{}, nil
}

type EnvelopeFATCAInfoResponse struct{}

type EnvelopeFATCAInfo struct {
	XMLName xml.Name      `xml:"soapenv:Envelope"`
	Soapenv string        `xml:"soapenv,attr"`
	Ejbs    string        `xml:"ejbs,attr"`
	Body    BodyFATCAInfo `xml:"soapenv:Body"`
}

type BodyFATCAInfo struct {
	GetPartyFATCAInfo GetPartyFATCAInfo `xml:"ejbs:getPartyFATCAInfo"`
}

type GetPartyFATCAInfo struct {
	Request RequestFATCAInfo `xml:"request"`
}

type RequestFATCAInfo struct {
	Control        ControlFATCAInfo `xml:"control"`
	CustomerID     string           `xml:"customerId"`
	CustomerSource string           `xml:"customerSource"`
}

type ControlFATCAInfo struct {
	Branch        string `xml:"branch"`
	Channel       string `xml:"channel"`
	RequestID     string `xml:"requestId"`
	RequesterName string `xml:"requesterName"`
	User          string `xml:"user"`
}

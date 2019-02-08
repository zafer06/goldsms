package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL string = "http://www.goldmesaj.com.tr/api/v1"

// GoldSMS struct
type GoldSMS struct {
	//
}

// APIMultiSend function
func (g *GoldSMS) APIMultiSend(req string) MultiSendRes {
	var res, _ = g.Connect("/multisend", req)

	var m MultiSendRes
	json.Unmarshal(res, &m)
	return m
}

// APIKredi function
func (g *GoldSMS) APIKredi(req string) KrediRes {
	var res, _ = g.Connect("/kredi", req)

	var k KrediRes
	json.Unmarshal(res, &k)
	return k
}

// APIReportByID function
func (g *GoldSMS) APIReportByID(req string) ReportByIDRes {
	var res, _ = g.Connect("/reportbyid", req)

	var r ReportByIDRes
	json.Unmarshal(res, &r)
	return r
}

// APIReport function
func (g *GoldSMS) APIReport(req string) ReportRes {
	var res, _ = g.Connect("/report", req)

	var r ReportRes
	json.Unmarshal(res, &r)
	return r
}

// APISendSms function
func (g *GoldSMS) APISendSms(req string) SendSmsRes {
	var res, _ = g.Connect("/sendsms", req)

	var s SendSmsRes
	json.Unmarshal(res, &s)
	return s
}

// APIAlfanumerik function
func (g *GoldSMS) APIAlfanumerik(req string) AlfanumerikRes {
	var res, _ = g.Connect("/alfanumerik", req)

	var a AlfanumerikRes
	json.Unmarshal(res, &a)
	return a
}

// APILogin function
func (g *GoldSMS) APILogin(req string) LoginRes {
	var res, _ = g.Connect("/login", req)

	var l LoginRes
	json.Unmarshal(res, &l)
	return l
}

// Connect function
func (g *GoldSMS) Connect(url string, request string) ([]byte, error) {
	json := strings.NewReader(request)

	req, err := http.NewRequest("POST", baseURL+url, json)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

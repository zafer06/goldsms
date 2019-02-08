package api

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var smsUser string
var smsPass string

func init() {
	smsUser = os.Getenv("GOLDSMS_USER")
	smsPass = os.Getenv("GOLDSMS_PASS")
}

func TestAPIMultiSend(t *testing.T) {
	var req = fmt.Sprintf(`{
	    "username": "%s",
	    "password": "%s",
	    "sdate": "",
	    "vperiod": "48",
	    "message": [
	        {
	            "sender": "TEST",
	            "text": "Bu bir test!",
	            "utf8": "1",
	            "gsm": ["905*****", "905*****"]
	        },
	        {
	            "sender": "TEST",
	            "text": "Baska bir test!",
	            "utf8": "1",
	            "gsm": ["905*****"]
	        }
	    ]
	}`, smsUser, smsPass)

	sms := GoldSMS{}
	multi := sms.APIMultiSend(req)

	if multi.Status != "ok" {
		t.Errorf("API kredi was incorrect. got: %s, want: %s", multi.Error, "ok")
	}
}

func TestAPIKredi(t *testing.T) {
	var req = fmt.Sprintf(`{ "username": "%s", "password": "%s" }`,
		smsUser, smsPass)

	sms := GoldSMS{}
	kredi := sms.APIKredi(req)

	t.Log("--> ", kredi)

	if kredi.Status != "ok" {
		t.Errorf("API kredi was incorrect. got: %s, want: %s", kredi.Error, "ok")
	}
}

func TestAPIReportByID(t *testing.T) {
	var req = fmt.Sprintf(`{
        "username": "%s",
        "password": "%s",
        "reportid": [
                "m12k0s45",
                "m12k0s46"
            ]
        }`, smsUser, smsPass)

	sms := GoldSMS{}
	report := sms.APIReportByID(req)

	if report.Status != "ok" {
		t.Errorf("ReportByID was incorrect. got: %s, want: %s", report.Error, "ok")
	}
}

func TestAPIReport(t *testing.T) {
	var detail = 0 // 0 -> simple, 1 -> detail

	var req = fmt.Sprintf(`{
        "username": "%s",
        "password": "%s",
        "messageid": "77-ASZ-YQF-JMW",
        "detailed": "%d"
        }`, smsUser, smsPass, detail)

	sms := GoldSMS{}
	report := sms.APIReport(req)

	if report.Status != "ok" {
		t.Errorf("Report was incorrect. got: %s, want: %s", report.Error, "ok")
	}
}

func TestAPISendSms(t *testing.T) {
	var sendDate = time.Now().Format("20060102150405")
	var sender = "TEST"
	var message = "Bu bir test mesajıdır"
	var utf8 = 1 // utf flag for Turkish char 0 -> Turkish char off or 1 -> Turkish char on
	var gsmNumbers string
	for _, v := range []string{"905*******", "905******"} {
		gsmNumbers += "\"" + v + "\","
	}
	gsmNumbers = gsmNumbers[:len(gsmNumbers)-1]

	var req = fmt.Sprintf(`{
    	"username": "%s",
    	"password": "%s",
    	"sdate": "%s",
    	"vperiod": "48",
    	"message": {
    		"sender": "%s",
    		"text": "%s",
    		"utf8": "%d",
    		"gsm": [%s]
    	}
    }`, smsUser, smsPass, sendDate, sender, message, utf8, gsmNumbers)

	fmt.Println("--> ", req)
	/*
		sms := GoldSMS{}
		send := sms.APISendSms(req)

		t.Log("--> ", send)

		if send.Status != "ok" {
			t.Errorf("SendSms was incorrect. got: %s, want: %s", send.Status, "ok")
		}
	*/
}

func TestAPIAlfanumerik(t *testing.T) {
	var req = fmt.Sprintf(`{ "username": "%s", "password": "%s" }`,
		smsUser, smsPass)

	sms := GoldSMS{}
	alfanumerik := sms.APIAlfanumerik(req)

	if alfanumerik.Status != "ok" || alfanumerik.Result[0] == "" {
		t.Errorf("Alfanumerik was incorrect. got: %s, want: %s", alfanumerik.Error, "ok")
	}
}

func TestAPILogin(t *testing.T) {
	var req = fmt.Sprintf(`{ "username": "%s", "password": "%s" }`,
		smsUser, smsPass)

	sms := GoldSMS{}
	login := sms.APILogin(req)

	if login.Status != "ok" {
		t.Errorf("API login was incorrect. got: %s, want: %s", login.Error, "ok")
	}
}

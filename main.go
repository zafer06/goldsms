package main

import (
	"fmt"
	"goldsms/api"
	"os"
)

func main() {
	var smsUser = os.Getenv("GOLDSMS_USER")
	var smsPass = os.Getenv("GOLDSMS_PASS")

	var sms = api.GoldSMS{}
	var req = fmt.Sprintf(`{ "username": "%s", "password": "%s" }`,
		smsUser, smsPass)

	var login = sms.APILogin(req)

	fmt.Println("--> Login Status :", login.Status)
}

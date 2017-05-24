package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"model"
	_ "opcode-server/routers"
)

const (
	base_url = "http://localhost:8080/v1/email"
)

func Test_EmailService_Send(t *testing.T) {
	// request message
	var emails []*model.Email
	var e1 model.Email
	e1.To = "suengguan@126.com"
	e1.Subject = "登录提醒"
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	e1.Body = "【科思世纪官网登录提醒】欢迎您于 " + currentTime + " 登录PME系统。特此提醒。【科思世纪】"
	emails = append(emails, &e1)

	requestData, err := json.Marshal(&emails)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// post
	res, err := http.Post(base_url+"/send/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type TranslateModel struct {
	Q     string
	From  string
	To    string
	Appid int
	Salt  int
	Sign  string
}

func NewTranslateModeler(appID int, password string, q, from, to string) TranslateModel {
	tran := TranslateModel{
		Q:    q,
		From: from,
		To:   to,
	}
	tran.Appid = appID
	tran.Salt = time.Now().Second()
	content := strconv.Itoa(appID) + q + strconv.Itoa(tran.Salt) + password
	sign := SumString(content) //计算sign值
	tran.Sign = sign
	return tran
}

func (tran TranslateModel) ToValues() url.Values {
	values := url.Values{
		"q":     {tran.Q},
		"from":  {tran.From},
		"to":    {tran.To},
		"appid": {strconv.Itoa(tran.Appid)},
		"salt":  {strconv.Itoa(tran.Salt)},
		"sign":  {tran.Sign},
	}
	return values
}

func test2() {
	tran := NewTranslateModeler(3344, "", "世界第一223", "zh", "cht")
	values := tran.ToValues()
	resp, err := http.PostForm("Url", values)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	txt := string(body)
	fmt.Println(txt)
}

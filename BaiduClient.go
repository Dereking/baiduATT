package baiduATT

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BaiduClient struct {
	AppID     string
	APIKey    string
	SecretKey string

	Token *BaiduToken
}

type BaiduToken struct {
	Refresh_token  string `json:"refresh_token"`
	Expires_in     uint64 `json:"expires_in"`
	Access_token   string `json:"access_token"`
	Scope          string `json:"scope"`
	Session_key    string `json:"session_key"`
	Session_secret string `json:"session_secret"`
}

func NewBaiduClient(AppID, APIKey, SecretKey string) *BaiduClient {
	return &BaiduClient{
		AppID:     AppID,
		APIKey:    APIKey,
		SecretKey: SecretKey,
	}
}

func (c *BaiduClient) GetToken() error {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		c.APIKey, c.SecretKey)
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	tk := &BaiduToken{}
	err = json.Unmarshal(body, tk)
	if err != nil {
		return err
	}
	c.Token = tk

	//fmt.Println(tk.Access_token)

	return nil
}

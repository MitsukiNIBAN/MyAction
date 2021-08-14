package WPush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Content struct {
	Msg string `json:"content"`
}

type Data struct {
	Touser  string  `json:"touser"`
	Msgtype string  `json:"msgtype"`
	Agentid string  `json:"agentid"`
	Text    Content `json:"text"`
}

func praseJson(jsonStr string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(string(jsonStr)), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func IsZero(v interface{}) bool {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		panic(fmt.Sprintf("type is not comparable: %v", t))
	}
	return v == reflect.Zero(t).Interface()
}

func GetToken(corpid string, appsecret string) string {
	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + appsecret)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	result, err := praseJson(string(respByte))
	if err != nil {
		panic(err)
	}

	return result["access_token"].(string)
}

func PushMsg(msg string, corpid string, appsecret string, agentid string) {
	token := GetToken(corpid, appsecret)
	paramJson, err := json.Marshal(Data{Touser: "@all", Msgtype: "text", Agentid: agentid, Text: Content{Msg: msg}})
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="+token, bytes.NewBuffer(paramJson))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	signResp, err := (&http.Client{}).Do(request)
	if err != nil {
		panic(err)
	}
	defer signResp.Body.Close()

	respByte, err := ioutil.ReadAll(signResp.Body)
	if err != nil {
		panic(err)
	}

	result, err := praseJson(string(respByte))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	if IsZero(result["errcode"]) {
		fmt.Println("推送成功")
	} else {
		panic(result)
	}
}

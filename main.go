package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	appVer := os.Args[1]
	uid := os.Args[2]
	cookie := os.Args[3]
	actID := "e202009291139501"
	region := "cn_gf01"

	userAgent := fmt.Sprintf("Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) miHoYoBBS/%v", appVer)
	referer := fmt.Sprintf("https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=%v&act_id=%v&utm_source=%v&utm_medium=%v&utm_campaign=%v",
		"true", actID, "bbs", "mys", "icon")

	infoUrl := fmt.Sprintf("https://api-takumi.mihoyo.com/event/bbs_sign_reward/info?region=%v&act_id=%v&uid=%v", region, actID, uid)
	// signUrl := "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign"

	infoReq, _ := http.NewRequest("GET", infoUrl, nil)
	infoReq.Header.Set("User-Agent", userAgent)
	infoReq.Header.Set("Referer", referer)
	infoReq.Header.Set("Accept-Encoding", "gzip, deflate, br")
	infoReq.Header.Set("Cookie", cookie)
	infoResp, err := (&http.Client{}).Do(infoReq)
	if err != nil {
		fmt.Println(err)
	}
	defer infoResp.Body.Close()
	respByte, _ := ioutil.ReadAll(infoResp.Body)
	fmt.Println(string(respByte))
}

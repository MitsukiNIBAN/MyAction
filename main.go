package main

import (
	"fmt"
	"os"
)

func main()  {
	appVer := os.Args[1]
	uid:= os.Args[2]
	cookie := os.Args[3]
	actID:="e202009291139501"
	region := "cn_gf01"

	infoUrl := fmt.Sprintf("https://api-takumi.mihoyo.com/event/bbs_sign_reward/info?region=%v&act_id=%v&uid=%v", region, actID, uid)
	signUrl := "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign"

	fmt.Println(infoUrl)
	fmt.Println(appVer)
	fmt.Println(cookie)
}
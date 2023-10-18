package main

import (
	"BilibiliAutoLikeRoport/utils"
	"fmt"
	login "github.com/boxie123/GoBilibiliLogin"
	"log"
	"net/http"
)

func main() {
	_, _, filePath := login.Login()
	_, cookie, roomId := utils.ReaderSetting(filePath)

	client := &http.Client{}
	err := utils.SendLikeReport(client, cookie, roomId)
	if err != nil {
		log.Printf("点赞失败: %v", err)
	}
	fmt.Scanln()
}

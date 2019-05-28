package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"time"
	"golang.org/x/net/websocket"
	"strconv"
	"log"
	"./dto"
	"encoding/json"
)

var (
	appId string
	key   string
	call  Callback
)

func md5String(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

func start(callback Callback) {

	call = callback

	var data = "{\"roomId\":518512}"
	var appId = ""
	var key = "33db6b3b"
	var timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	var signStr = "data=" + data + "&key=" + key + "&timestamp=" + timestamp

	sign := md5String(signStr)

	fmt.Println(timestamp)

	websocketUrl := "ws://openapi.huya.com/index.html?do=getMessageNotice&data=" + data + "&appId=" + appId + "&timestamp=" + timestamp + "&sign=" + sign

	fmt.Println(websocketUrl)

	ws, err := websocket.Dial(websocketUrl, "", "http://openapi.huya.com/")
	if err != nil {
		panic(err)
	}

	go keep(ws)

	for ; ; {
		recvMessage(ws)
	}

}

func recvMessage(ws *websocket.Conn) {

	var str string
	err := websocket.Message.Receive(ws, &str)
	if err != nil {
		log.Println(err)
	}

	var res = &dto.ResponseData{}
	json.Unmarshal([]byte(str), res)

	call.MessageNotice(res.Data)
}

func main() {

	start(&CallbackImpl{})

}

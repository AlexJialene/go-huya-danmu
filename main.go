package main

import (
	md52 "crypto/md5"
	"fmt"
	"encoding/hex"
	"time"
	"golang.org/x/net/websocket"
	"strconv"
	"./dto"
	"log"
	"encoding/json"
)

var (
	appId string
	key string
)

//callback

func start(url string)  {

}

func md5String(str string) string {
	md5:=md52.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}


func main() {

	//start(MESSAGE_NOTICE)

	//
	var data = "{\"roomId\":518512}"
	var appId = "155868690397009450"
	var key = "33db6b3b"
	var timestamp = strconv.FormatInt(time.Now().Unix() , 10)


	var signStr = "data=" + data + "&key=" + key + "&timestamp=" + timestamp

	sign:=md5String(signStr)

	fmt.Println(timestamp)

	websocketUrl := "ws://openapi.huya.com/index.html?do=getMessageNotice&data="+data+"&appId="+appId+"&timestamp="+timestamp+"&sign="+sign

	fmt.Println(websocketUrl)


	test:=dto.Test


	ws, err := websocket.Dial(websocketUrl, "", "http://openapi.huya.com/")
	if err != nil {
		panic(err)
	}

	for ; ;  {
		recvMessage(ws)
	}



}

func recvMessage(ws *websocket.Conn) {
	//data := dto.MessageNotice{}
	//res  := dto.ResponseData{ 0 , "" , data}
	var res dto.ResponseData
	//err :=websocket.Message.Receive(ws , res)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(res.StatusCode)
	//fmt.Println(res.Data.SendNick)
	//fmt.Println(res.Data.FansLevel)

	data := make([]byte, 2048)
	_, err := ws.Read(data)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(data , res)

	fmt.Println(string(data))
	fmt.Println(res.StatusCode)
}

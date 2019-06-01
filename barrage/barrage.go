package barrage

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"strconv"
	"strings"
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"encoding/json"
)

const (
	MESSAGE_NOTICE = "ws://openapi.huya.com/index.html?do=getMessageNotice&data={}&appId={}&timestamp={}&sign={}"
	ITEM_NOTICE    = "ws://openapi.huya.com/index.html?do=getSendItemNotice&data={}&appId={}&timestamp={}&sign={}"
	VIP_NOTICE     = "ws://openapi.huya.com/index.html?do=getVipEnterBannerNotice&data={}&appId={}&timestamp={}&sign={}"
)

var C = make(chan int)

func Md5String(s string) string {
	md := md5.New()
	md.Write([]byte(s))
	return hex.EncodeToString(md.Sum(nil))
}

func GetTimestampUnixString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func urlAssembly(url string, s ... string) string {
	for _, v := range s {
		url = strings.Replace(url, "{}", v, 1)
	}
	return url
}

func keep(ws *websocket.Conn) {
	if nil != ws {
		for ; ; {
			fmt.Println("[info] | send keep-alive ping")
			ws.Write([]byte("ping"))
			time.Sleep(time.Duration(15) * time.Second)
		}
	}
}

func ToStruct(call, obj interface{}) error {
	m := call.(map[string]interface{})
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, obj)
}

func (c Client) recvMessage(ws *websocket.Conn, call func(dto interface{})) {
	defer ws.Close()
	for {
		var str string
		err := websocket.Message.Receive(ws, &str)
		if err != nil {
			log.Println("[Error]", err)
		}
		var res = &ResponseData{}
		json.Unmarshal([]byte(str), res)
		call(res.Data)
	}

}

func (c Client) Connect(url string, call func(call interface{})) {
	ws, err := websocket.Dial(urlAssembly(
		url,
		c.getRoomStr(),
		c.AppId,
		GetTimestampUnixString(),
		c.genSign()),
		"",
		"http://openapi.huya.com/")

	if err != nil {
		panic(err)
	}

	go keep(ws)
	go c.recvMessage(ws, call)

}

func (c Client) genSign() string {

	signStr := "data=" + c.getRoomStr() + "&key=" + c.Key + "&timestamp=" + GetTimestampUnixString()
	return Md5String(signStr)
}

func (c Client) getRoomStr() string {
	roomStr := "{\"roomId\":" + c.RoomId + "}"
	return roomStr
}

func Sync() {
	if C == nil {
		C = make(chan int)
	}
	<-C
}

func ConnectMessage(c Client, callback func(message *MessageNotice)) {
	c.Connect(MESSAGE_NOTICE, func(call interface{}) {
		b := &MessageNotice{}
		err := ToStruct(call, b)
		if err != nil {
			log.Println("[Error]", err)
		}
		callback(b)
	})
}

func ConnectItemMessage(c Client, callback func(message *ItemNotice)) {
	c.Connect(ITEM_NOTICE, func(call interface{}) {
		b := &ItemNotice{}
		err := ToStruct(call, b)
		if err != nil {
			log.Println("[Error]", err)
		}
		callback(b)
	})
}

func ConnectVipMessage(c Client, callback func(message *VipNotice)) {
	c.Connect(VIP_NOTICE, func(call interface{}) {
		b := &VipNotice{}
		err := ToStruct(call, b)
		if err != nil {
			log.Println("[Error]", err)
		}
		callback(b)
	})
}

type Client struct {
	RoomId string
	AppId  string
	Key    string
}

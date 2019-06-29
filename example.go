package main

import (
	"fmt"
	"go-huya-danmu/barrage"
)

func main() {

	client := barrage.Client{RoomId: "518512", AppId: "155868690397009450", Key: ""}

	//client.Connect(barrage.MESSAGE_NOTICE, func(call interface{}) {
	//	b := &barrage.MessageNotice{}
	//	barrage.ToStruct(call, b)
	//
	//	fmt.Println(b.SendNick)
	//})
	//
	//client.Connect(barrage.ITEM_NOTICE , func(call interface{}) {
	//
	//	i := &barrage.ItemNotice{}
	//	barrage.ToStruct(call , i)
	//
	//	fmt.Println(i.ItemName)
	//})

	barrage.ConnectMessage(client, func(message *barrage.MessageNotice) {
		//弹幕消息
		fmt.Println(message.SendNick)
	})

	barrage.ConnectItemMessage(client, func(message *barrage.ItemNotice) {
		//礼物消息
		fmt.Println(message.ItemName)
	})

	barrage.ConnectVipMessage(client, func(message *barrage.VipNotice) {
		//贵族进场
		fmt.Println(message.NobleName)
	})
	//lock
	barrage.Sync()

}

package main

import (
	"go-huya-danmu/barrage"
	"fmt"
)

func main() {

	config := barrage.Client{RoomId: "518512", AppId: "155868690397009450", Key: ""}

	config.Connect(barrage.MESSAGE_NOTICE, func(call interface{}) {
		b := &barrage.MessageNotice{}
		barrage.ToStruct(call, b)

		fmt.Println(b.SendNick)
	})

	barrage.Sync()

}

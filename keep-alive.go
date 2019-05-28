package main

import (
	"golang.org/x/net/websocket"
	"time"
	"fmt"
)

func keep(ws *websocket.Conn)  {

	if nil!=ws {

		for ; ; {
			fmt.Println("发送心跳包")
			ws.Write([]byte("ping"))
			time.Sleep(time.Duration(5) * time.Second)
		}
	}

}
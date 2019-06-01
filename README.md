# go-huya-danmu
🐯Get the bullet screen of the huya.com / 虎牙弹幕获取 / Golang version


[Java version](https://github.com/AlexJialene/huya-danmu) | [LICENSE:Apache-2.0](https://github.com/AlexJialene/go-huya-danmu/blob/master/LICENSE)

## preface

To apply this app you must apply for developer qualification on the open platform (http://open.huya.com).

## Start
```
go get github.com/AlexJialene/go-huya-danmu
```

```
barrage.ConnectMessage(client , func(message *barrage.MessageNotice) {
    //弹幕消息
    fmt.Println(message.SendNick)
})

barrage.ConnectItemMessage(client , func(message *barrage.ItemNotice) {
    //礼物消息
    fmt.Println(message.ItemName)
})

barrage.ConnectVipMessage(client , func(message *barrage.VipNotice) {
    //贵族进场
    fmt.Println(message.NobleName)
})

//lock
barrage.Sync()

```
The end of the program must use the Sync（`barrage.Sync()`） method to block the main thread

You can refer to the [example.go]() file.

----

|参数名称	| 类型	|描述|
|:-|:-:|-:|
|roomId|int|房间号|
|sendNick|string|发言人昵称|
|senderAvatarUrl|string|发言人头像|
|senderGender|string|发言人性别， 0：女，1：男|
|showMode|string|显示类型：0.公屏和弹幕 1.公屏 2.弹幕|
|content|string	|发言内容|
|nobleLevel|int|贵族等级|
|fansLevel|int粉丝等级|
|badgeName|string|粉丝徽章名称（普通弹幕，不包括上电视弹幕|


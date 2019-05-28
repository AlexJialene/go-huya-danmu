package main

import (
	"fmt"
	"./dto"
)

//callback
type Callback interface {

	MessageNotice(notice *dto.MessageNotice)

}

type CallbackImpl struct {

}

func (c CallbackImpl) MessageNotice(notice *dto.MessageNotice)  {

	fmt.Println("用户["+notice.SendNick+"]:"+notice.Content)


}

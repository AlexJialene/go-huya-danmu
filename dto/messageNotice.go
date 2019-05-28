package dto

type MessageNotice struct {
	BadgeName       string `json:"badgeName"`
	Content         string `json:"content"`
	FansLevel       int    `json:"fansLevel"`
	NobleLevel      int    `json:"nobleLevel"`
	Roomid          int    `json:"roomid"`
	SendNick        string `json:"sendNick"`
	SenderAvatarUrl string `json:"senderAvatarUrl"`
	SenderGender    int    `json:"senderGender"`
	ShowMode        int    `json:"showMode"`
}

type ResponseData struct {
	StatusCode int `json:"statusCode"`

	StatusMsg string `json:"statusMsg"`

	Data *MessageNotice `json:"data"`
}

type Test struct {
	Name string `json:"nick"`
}

type Test1 struct {
	Type int `json:"type"`
	T *Test `json:"t"`
}

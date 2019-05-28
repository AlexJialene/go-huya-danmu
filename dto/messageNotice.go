package dto

type MessageNotice struct {
	BadgeName       string `json:"badgeName"`
	Content         string
	FansLevel       int
	NobleLevel      int
	Roomid          int
	SendNick        string `json:"sendNick"`
	SenderAvatarUrl string
	SenderGender    int
	ShowMode        int
}

type ResponseData struct {
	StatusCode int `json:"statusCode"`

	StatusMsg string `json:"statusMsg"`

	Data MessageNotice `json:"data"`
}

type Test struct {
	Name string `json:"nick"`
}

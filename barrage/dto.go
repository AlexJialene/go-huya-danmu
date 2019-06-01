package barrage

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

type ItemNotice struct {
	ItemName          string `json:"itemName"`
	PresenterNick     string `json:"presenterNick"`
	RoomId            int    `json:"roomId"`
	SendItemComboHits int    `json:"sendItemComboHits"`
	SendItemCount     int    `json:"sendItemCount"`
	SendNick          string `json:"sendNick"`
	SenderAvatarurl   string `json:"senderAvatarurl"`
}

type VipNotice struct {
	BadgeName     string `json:"badgeName"`
	FansLevel     int    `json:"fansLevel"`
	NobleLevel    int    `json:"nobleLevel"`
	NobleName     string `json:"nobleName"`
	RoomId        int    `json:"roomId"`
	UserAvatarUrl string `json:"userAvatarUrl"`
	UserNick      string `json:"userNick"`
	WeekRank      int    `json:"weekRank"`
	GuardLevel    int    `json:"guardLevel"`
}

type ResponseData struct {
	StatusCode int `json:"statusCode"`

	StatusMsg string `json:"statusMsg"`

	Data interface{} `json:"data"`
}

package funddto

type FundReqeust struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
	Desc  string `json:"desc" form:"desc"`
	Goals int `json:"goals" form:"goals"`
}

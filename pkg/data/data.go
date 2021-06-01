package data

const PROJECT_ID = "auctionee"

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Good struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

type GoodsList struct {
	Login string  `json:"login"`
	Goods []*Good `json:"goods"`
}

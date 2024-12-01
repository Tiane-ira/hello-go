package domain

type Page struct {
	NoPage  bool  `form:"noPage" json:"noPage"`
	CurPage int64 `form:"page" json:"page"`
	Size    int64 `form:"size" json:"size"`
}

type IdReq struct {
	Id uint `json:"id" form:"id" uri:"id"`
}

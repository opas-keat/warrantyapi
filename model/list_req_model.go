package model

type ListReq struct {
	Offset int    `query:"offset" default:"0"`
	Limit  int    `query:"limit" default:"50"`
	Order  string `query:"order"`
}

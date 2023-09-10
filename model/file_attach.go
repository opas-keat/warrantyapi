package model

type FileAttach struct {
	ID        string `json:"id"`
	CreatedBy string `json:"-"`
	FileName  string `json:"file_name" form:"file_name"`
	FileType  string `json:"file_type"`
	FileSize  int64  `json:"file_size"`
	Module    string `json:"module"`
	LinkType  string `json:"link_type" form:"link_type"`
	LinkId    string `json:"link_id" form:"link_id"`
}

type FileAttachResponse struct {
	ID        string `json:"id"`
	CreatedBy string `json:"-"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileSize  int64  `json:"file_size"`
	Module    string `json:"module"`
	LinkType  string `json:"link_type"`
	LinkId    string `json:"link_id"`
}

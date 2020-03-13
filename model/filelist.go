package model

//上传文件信息
type FormFile struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Url      string `json:"url"`
	ThumbUrl string `json:"thumbUrl"`
}

package model

type FileStorageItem struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Component string `json:"component" xorm:"not null VARCHAR(255)"`
	BaseUrl   string `json:"base_url" xorm:"not null VARCHAR(1024)"`
	Path      string `json:"path" xorm:"not null VARCHAR(1024)"`
	Type      string `json:"type" xorm:"VARCHAR(255)"`
	Size      int    `json:"size" xorm:"INT(11)"`
	Name      string `json:"name" xorm:"VARCHAR(255)"`
	UploadIp  string `json:"upload_ip" xorm:"VARCHAR(15)"`
	CreatedAt int    `json:"created_at" xorm:"not null INT(11)"`
}

func (t FileStorageItem) TableName() string {
	return "ty_file_storage_item"
}

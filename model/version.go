package model

type Version struct {
	Vid       int    `json:"vid" xorm:"not null pk autoincr INT(11)"`
	Title     int    `json:"title" xorm:"not null INT(11)"`
	Version   string `json:"version" xorm:"default '' VARCHAR(50)"`
	Desc      string `json:"desc" xorm:"not null TEXT"`
	DownUrl   string `json:"down_url" xorm:"not null VARCHAR(200)"`
	BaseUrl   string `json:"base_url" xorm:"default '' VARCHAR(200)"`
	CreatedAt int    `json:"created_at" xorm:"default 0 INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"default 0 INT(11)"`
}

func (t Version) TableName() string {
	return "ty_version"
}

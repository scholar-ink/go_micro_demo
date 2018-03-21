package model

type LoadingImg struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string `json:"name" xorm:"default '' VARCHAR(250)"`
	Path      string `json:"path" xorm:"not null default '' VARCHAR(250)"`
	BaseUrl   string `json:"base_url" xorm:"VARCHAR(250)"`
	CreatedAt int    `json:"created_at" xorm:"INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"INT(11)"`
}

func (t LoadingImg) TableName() string {
	return "ty_loading_img"
}

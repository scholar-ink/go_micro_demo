package model

type AuthItemChild struct {
	Parent string `json:"parent" xorm:"not null pk VARCHAR(64)"`
	Child  string `json:"child" xorm:"not null pk index VARCHAR(64)"`
}

func (t AuthItemChild) TableName() string {
	return "ty_auth_item_child"
}

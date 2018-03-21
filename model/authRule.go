package model

type AuthRule struct {
	Name      string `json:"name" xorm:"not null pk VARCHAR(64)"`
	Data      string `json:"data" xorm:"TEXT"`
	CreatedAt int    `json:"created_at" xorm:"INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"INT(11)"`
}

func (t AuthRule) TableName() string {
	return "ty_auth_rule"
}

package model

type AuthItem struct {
	Name        string `json:"name" xorm:"not null pk VARCHAR(64)"`
	Type        int    `json:"type" xorm:"not null index INT(11)"`
	Description string `json:"description" xorm:"TEXT"`
	RuleName    string `json:"rule_name" xorm:"index VARCHAR(64)"`
	Data        string `json:"data" xorm:"TEXT"`
	CreatedAt   int    `json:"created_at" xorm:"INT(11)"`
	UpdatedAt   int    `json:"updated_at" xorm:"INT(11)"`
}

func (t AuthItem) TableName() string {
	return "ty_auth_item"
}

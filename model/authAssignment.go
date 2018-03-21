package model

type AuthAssignment struct {
	ItemName  string `json:"item_name" xorm:"not null pk VARCHAR(64)"`
	UserId    string `json:"user_id" xorm:"not null pk VARCHAR(64)"`
	CreatedAt int    `json:"created_at" xorm:"INT(11)"`
}

func (t AuthAssignment) TableName() string {
	return "ty_auth_assignment"
}

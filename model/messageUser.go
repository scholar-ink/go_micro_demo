package model

type MessageUser struct {
	Fid       int `json:"fid" xorm:"not null pk autoincr INT(11)"`
	Mid       int `json:"mid" xorm:"not null INT(11)"`
	Uid       int `json:"uid" xorm:"default 0 INT(11)"`
	CreatedAt int `json:"created_at" xorm:"default 0 INT(11)"`
	UpdatedAt int `json:"updated_at" xorm:"default 0 INT(11)"`
}

func (t MessageUser) TableName() string {
	return "ty_message_user"
}

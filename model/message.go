package model

type Message struct {
	Mid       int    `json:"mid" xorm:"not null pk autoincr INT(11)"`
	Title     string `json:"title" xorm:"default '' VARCHAR(255)"`
	Uid       int    `json:"uid" xorm:"not null default 0 index index index index index INT(11)"`
	Uname     string `json:"uname" xorm:"default '' VARCHAR(255)"`
	Uavatar   string `json:"uavatar" xorm:"default '' VARCHAR(255)"`
	Type      int    `json:"type" xorm:"default 1 INT(11)"`
	DataType  int    `json:"data_type" xorm:"default 0 INT(11)"`
	DataId    int    `json:"data_id" xorm:"default 0 INT(11)"`
	LinkUrl   string `json:"link_url" xorm:"not null default '' VARCHAR(255)"`
	Data      string `json:"data" xorm:"default '' VARCHAR(255)"`
	Content   string `json:"content" xorm:"default '' VARCHAR(255)"`
	CreatedAt int    `json:"created_at" xorm:"default 0 INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"default 0 INT(11)"`
}

func (t Message) TableName() string {
	return "ty_message"
}

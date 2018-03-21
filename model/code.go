package model

type Code struct {
	Cid        int    `json:"cid" xorm:"not null pk autoincr INT(10)"`
	Code       int    `json:"code" xorm:"default 0 INT(10)"`
	Mobile     string `json:"mobile" xorm:"default '' VARCHAR(50)"`
	Type       int    `json:"type" xorm:"default 1 TINYINT(1)"`
	Status     int    `json:"status" xorm:"default 0 TINYINT(3)"`
	CreateTime int    `json:"create_time" xorm:"default 0 INT(11)"`
}

func (t Code) TableName() string {
	return "ty_code"
}

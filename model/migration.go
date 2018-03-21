package model

type Migration struct {
	Version   string `json:"version" xorm:"not null pk VARCHAR(180)"`
	ApplyTime int    `json:"apply_time" xorm:"INT(11)"`
}

func (t Migration) TableName() string {
	return "ty_migration"
}

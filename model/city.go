package model

type City struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name    string `json:"name" xorm:"not null CHAR(15)"`
	Pinyin  string `json:"pinyin" xorm:"not null VARCHAR(150)"`
	IsOpen  int    `json:"is_open" xorm:"TINYINT(3)"`
	Direct  int    `json:"direct" xorm:"TINYINT(3)"`
	Area    int    `json:"area" xorm:"index TINYINT(3)"`
	Hot     int    `json:"hot" xorm:"TINYINT(3)"`
	Pid     int    `json:"pid" xorm:"not null default 0 index INT(10)"`
	Ucfirst string `json:"ucfirst" xorm:"not null index CHAR(1)"`
	Sort    int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

func (t City) TableName() string {
	return "ty_city"
}

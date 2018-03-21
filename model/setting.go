package model

type Setting struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParentId   int    `json:"parent_id" xorm:"not null default 0 index INT(11)"`
	Code       string `json:"code" xorm:"not null index VARCHAR(32)"`
	Type       string `json:"type" xorm:"not null VARCHAR(32)"`
	StoreRange string `json:"store_range" xorm:"VARCHAR(255)"`
	StoreDir   string `json:"store_dir" xorm:"VARCHAR(255)"`
	Value      string `json:"value" xorm:"TEXT"`
	SortOrder  int    `json:"sort_order" xorm:"not null default 50 index INT(11)"`
}

func (t Setting) TableName() string {
	return "ty_setting"
}

package model

type SystemLog struct {
	Id       int64   `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Level    int     `json:"level" xorm:"index INT(11)"`
	Category string  `json:"category" xorm:"index VARCHAR(255)"`
	LogTime  float64 `json:"log_time" xorm:"DOUBLE"`
	Prefix   string  `json:"prefix" xorm:"TEXT"`
	Message  string  `json:"message" xorm:"TEXT"`
}

func (t SystemLog) TableName() string {
	return "ty_system_log"
}

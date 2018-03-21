package model

type Member struct {
	Uid           int    `json:"uid" xorm:"not null pk autoincr INT(10)"`
	Mobile        string `json:"mobile" xorm:"not null default '' VARCHAR(30)"`
	Password      string `json:"password" xorm:"not null VARCHAR(100)"`
	Nickname      string `json:"nickname" xorm:"default '' VARCHAR(100)"`
	Gender        int    `json:"gender" xorm:"default 1 TINYINT(1)"`
	Birthday      int    `json:"birthday" xorm:"default 0 INT(10)"`
	Hobby         string `json:"hobby" xorm:"default '' VARCHAR(255)"`
	Sign          string `json:"sign" xorm:"default '' VARCHAR(255)"`
	AvatarPath    string `json:"avatar_path" xorm:"default '' VARCHAR(255)"`
	AvatarBaseUrl string `json:"avatar_base_url" xorm:"VARCHAR(255)"`
	Address       string `json:"address" xorm:"default '' VARCHAR(255)"`
	ClientId      string `json:"client_id" xorm:"default '' VARCHAR(255)"`
	UpdatedAt     int    `json:"updated_at" xorm:"default 0 INT(10)"`
	CreatedAt     int    `json:"created_at" xorm:"default 0 INT(10)"`
	LastIp        string `json:"last_ip" xorm:"default '' VARCHAR(255)"`
	LastLogin     int    `json:"last_login" xorm:"default 0 INT(10)"`
	Lng           string `json:"lng" xorm:"default '' VARCHAR(255)"`
	Lat           string `json:"lat" xorm:"default '' VARCHAR(255)"`
	Cover         string `json:"cover" xorm:"default '' VARCHAR(255)"`
	Banned        int    `json:"banned" xorm:"default 1 TINYINT(1)"`
	AttentionNum  int    `json:"attention_num" xorm:"default 0 INT(11)"`
	BattentionNum int    `json:"battention_num" xorm:"default 0 INT(11)"`
	SnsNum        int    `json:"sns_num" xorm:"default 0 INT(11)"`
	Province      int    `json:"province" xorm:"default 0 INT(11)"`
	City          int    `json:"city" xorm:"default 0 INT(11)"`
	Town          int    `json:"town" xorm:"default 0 INT(11)"`
	Salesman      int    `json:"salesman" xorm:"default 0 TINYINT(11)"`
}

func (t Member) TableName() string {
	return "ty_member"
}

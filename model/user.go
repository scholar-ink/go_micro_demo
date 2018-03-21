package model

type User struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username           string `json:"username" xorm:"not null unique VARCHAR(255)"`
	AuthKey            string `json:"auth_key" xorm:"not null VARCHAR(32)"`
	PasswordHash       string `json:"password_hash" xorm:"not null VARCHAR(255)"`
	PasswordResetToken string `json:"password_reset_token" xorm:"VARCHAR(255)"`
	Email              string `json:"email" xorm:"not null VARCHAR(255)"`
	Role               string `json:"role" xorm:"not null default 'user' index VARCHAR(64)"`
	Status             int    `json:"status" xorm:"not null default 10 index SMALLINT(6)"`
	CreatedAt          int    `json:"created_at" xorm:"not null index INT(11)"`
	UpdatedAt          int    `json:"updated_at" xorm:"not null INT(11)"`
}

func (t User) TableName() string {
	return "ty_user"
}

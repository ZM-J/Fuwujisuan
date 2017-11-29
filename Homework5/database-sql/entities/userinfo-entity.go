package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `xorm:"pk autoincr"` //语义标签
	UserName   string     `xorm:"notnull unique"`
	DepartName string     `xorm:"notnull"`
	CreateAt   *time.Time `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName should not be null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}

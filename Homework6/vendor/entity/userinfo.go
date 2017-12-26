package entity

import "time"

// UserInfo .
type UserInfo struct {
	UID        int        `table:"userinfo" column:"uid"`
	UserName   string     `table:"userinfo" column:"username"`
	DepartName string     `table:"userinfo" column:"departname"`
	Created    *time.Time `table:"userinfo" column:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName should not be null!")
	}
	if u.Created == nil {
		t := time.Now()
		u.Created = &t
	}
	return &u
}

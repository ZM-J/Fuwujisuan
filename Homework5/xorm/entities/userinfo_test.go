package entities

import (
	"testing"
	"time"
)

func Test_insert(t *testing.T) {
	tn := time.Now()
	u := &UserInfo{UID:233,
		UserName:"带带大师兄",
		DepartName:"带带工作室",
		CreateAt:&tn,
	}
	UserInfoService.Save(u)
	t.Log("Save",u.UserName,"successfully.")
}
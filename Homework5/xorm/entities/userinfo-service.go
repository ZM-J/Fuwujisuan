package entities

import "fmt"

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	var everyone []UserInfo
	err := engine.Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	userQuery := &UserInfo{UID: id}
	_, err := engine.Get(userQuery)
	checkErr(err)
	return userQuery
}

// Count .
func (*UserInfoAtomicService) CountNum() int {
	user := new(UserInfo)
	c, err := engine.Count(user)
	checkErr(err)
	return int(c)
}

package main

import (
	"dao"
	"entity"
	"fmt"
	"reflect"
)

//通过组合模式继承
type IDaoUser interface {
	dao.IDaoBase
}

type daoUser struct {
	dao.DaoBase
}

// 声明一个实例
var daoUserInstance IDaoUser

func DaoUser() IDaoUser {
	if daoUserInstance == nil {
		daoUserInstance = &daoUser{dao.DaoBase{EntityType: reflect.TypeOf(new(entity.UserInfo)).Elem()}}
		daoUserInstance.Init()
	}
	return daoUserInstance
}

func main() {
	daoUserInstance = DaoUser()
	var u = entity.UserInfo{UserName: "SunXiaoChuan"}
	user := entity.NewUserInfo(u)
	user.DepartName = "huya"
	err := daoUserInstance.Save(user)
	if err != nil {
		panic(err)
	}
	pEveryOne, err := daoUserInstance.Find()
	for index := 0; index < pEveryOne.Len(); index++ {
		item := reflect.ValueOf(pEveryOne.Front().Value).Interface()
		fmt.Printf("result: %v\n", reflect.ValueOf(item))
		pEveryOne.Remove(pEveryOne.Front())
	}

}

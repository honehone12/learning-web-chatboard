package user

import (
	"chatboard/common"
	"chatboard/models"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	engine, err := common.OpenDB("chatboard")
	if err != nil {
		t.Error(err)
	}
	OpenService(engine)
	res := CallService(&common.Message{
		Service:  common.ServiceCall,
		FuncType: CreateUser,
		Data: models.User{
			Name:     "TestingTaro",
			Email:    "TestingTaro@yapoo.com",
			Password: "testttest#110",
		},
	})
	if resData, ok := res.Data.(int64); ok {
		if resData != 1 {
			t.Errorf("response was %v", resData)
		}
	} else {
		t.Errorf("response was %s", res.Data.(error).Error())
	}
}

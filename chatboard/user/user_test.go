package user

import (
	"chatboard/common"
	"chatboard/message"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	engine, err := common.OpenDB("chatboard")
	if err != nil {
		t.Error(err)
	}
	OpenService(engine)
	res := CallService(&message.Message{
		Service:  message.ServiceCall,
		FuncType: CreateUser,
		Data: User{
			Name:     "TestingTaro",
			Email:    "TestingTaro@yapoo.com",
			Password: "testttest#110",
		},
	})
	resData, ok := res.Data.(int64)
	if ok {
		if resData != 1 {
			t.Errorf("response was %v", resData)
		}
	} else {
		t.Errorf("response was %s", res.Data.(error).Error())
	}
}

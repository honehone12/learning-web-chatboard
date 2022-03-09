package thread

import (
	"chatboard/common"
	"testing"
)

/////////////////////////////////////////
// need user before testing

func Test_CreateThread(t *testing.T) {
	engine, err := common.OpenDB("chatboard")
	if err != nil {
		t.Error(err)
	}
	OpenService(engine)
	res := CallService(&common.Message{
		Service:  common.ServiceCall,
		FuncType: CreateThread,
		Data: common.Contribution{
			Content:  "Testooo",
			UserID:   0,
			UserName: "TestingTaro",
		},
	})
	if resData, ok := res.Data.(int64); ok {
		if resData != 1 {
			t.Errorf("response was %v", resData)
		}
	} else {
		t.Errorf("response was %s\n", res.Data.(error).Error())
	}
}

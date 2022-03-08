package thread

import (
	"chatboard/common"
	"chatboard/message"
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
	res := CallService(&message.Message{
		Service:  message.ServiceCall,
		FuncType: CreateThread,
		Data: message.Pair{
			ID:   0,
			Data: "Testingooo",
		},
	})
	resData, ok := res.Data.(int64)
	if ok {
		if resData != 1 {
			t.Errorf("response was %v", resData)
		}
	} else {
		t.Errorf("response was %s\n", res.Data.(error).Error())
	}
}

package user

import (
	"chatboard/common"
	"chatboard/message"
	"errors"
	"time"

	"xorm.io/xorm"
)

var engine *xorm.Engine

// engin can be separated from thread db
func OpenService(dbEngine *xorm.Engine) {
	engine = dbEngine
}

const (
	Unknown message.FuncTypeT = iota
	CreateUser
	CreateSession
	GetSessionByUUID
)

func CallService(msg *message.Message) *message.Message {
	if msg.Service == message.ServiceCall {
		return callServiceInternal(msg.FuncType, &msg.Data)
	} else {
		return &message.Message{
			Service:  message.Respose,
			FuncType: Unknown,
			Data:     errors.New("recieved responce message as service call"),
		}
	}
}

func convertType(data *interface{}, target interface{}) (err error) {
	switch target := target.(type) {
	case *User:
		data, ok := (*data).(User)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into User")
		}
	default:
		err = errors.New("unknown convert target")
	}
	return
}

func callServiceInternal(funcType message.FuncTypeT, data *interface{}) *message.Message {
	var err error
	var resFuncType message.FuncTypeT = Unknown
	var resData interface{}
	switch funcType {
	case CreateUser:
		resFuncType = CreateUser
		var user User
		if err = convertType(data, &user); err != nil {
			break
		}
		resData, err = createUser(&user)
	case CreateSession:
		resFuncType = CreateSession
		var user User
		if err = convertType(data, &user); err != nil {
			break
		}
		resData, err = createSession(&user)
	case GetSessionByUUID:
		resFuncType = GetSessionByUUID
		var uuid string
		if err = message.ConvertType(data, &uuid); err != nil {
			break
		}
		resData, err = getSessionByUUID(uuid)
	default:
		err = errors.New("recieved unknown function request")
	}
	if err != nil {
		resData = err
	}
	return &message.Message{
		Service:  message.Respose,
		FuncType: resFuncType,
		Data:     resData,
	}
}

//////////////////////////////////////////////////
// is good way encrypting user data here??
func createUser(user *User) (affected int64, err error) {
	user.UuId = common.NewUUIDString()
	user.Password = common.Encrypt(user.Password)
	user.CreatedAt = time.Now()
	affected, err = engine.Table("users").InsertOne(user)
	return
}

//////////////////////////////////////////////////
// these methods have to work with gin/session
func createSession(user *User) (affected int64, err error) {
	session := Session{
		UuId:      common.NewUUIDString(),
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	affected, err = engine.Table("sessions").InsertOne(&session)
	return
}

func getSessionByUUID(uuid string) (session, err error) {

	return
}

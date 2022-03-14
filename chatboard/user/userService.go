package user

import (
	"chatboard/common"
	"chatboard/models"
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
	Unknown common.FuncTypeT = iota
	CreateUser
	CreateSession
	GetSessionByUUID
	GetUserByUUID
	GetUserByEmail
	DeleteSessionByUUID
)

func CallService(msg *common.Message) *common.Message {
	if msg.Service == common.ServiceCall {
		return callServiceInternal(msg.FuncType, &msg.Data)
	} else {
		return &common.Message{
			Service:  common.Response,
			FuncType: Unknown,
			Data:     errors.New("recieved responce message as service call"),
		}
	}
}

func callServiceInternal(funcType common.FuncTypeT, data *interface{}) *common.Message {
	var err error
	var resFuncType common.FuncTypeT = Unknown
	var resData interface{}
	switch funcType {
	case CreateUser:
		resFuncType = CreateUser
		var user models.User
		if err = common.ConvertType(data, &user); err != nil {
			break
		}
		resData, err = createUser(&user)
	case CreateSession:
		resFuncType = CreateSession
		var user models.User
		if err = common.ConvertType(data, &user); err != nil {
			break
		}
		_, resData, err = createSession(&user)
	case GetSessionByUUID:
		resFuncType = GetSessionByUUID
		var uuid string
		if err = common.ConvertType(data, &uuid); err != nil {
			break
		}
		var ok bool
		if ok, resData, err = getSessionByUUID(uuid); err != nil {
			break
		} else if !ok {
			err = errors.New("no sessions found")
		}
	case GetUserByUUID:
		resFuncType = GetUserByUUID
		var uuid string
		if err = common.ConvertType(data, &uuid); err != nil {
			break
		}
		var ok bool
		if ok, resData, err = getUserByUUID(uuid); err != nil {
			break
		} else if !ok {
			err = errors.New("no user found")
		}
	case GetUserByEmail:
		resFuncType = GetUserByEmail
		var email string
		if err = common.ConvertType(data, &email); err != nil {
			break
		}
		var ok bool
		if ok, resData, err = getUserByEmail(email); err != nil {
			break
		} else if !ok {
			err = errors.New("no user found")
		}
	case DeleteSessionByUUID:
		resFuncType = DeleteSessionByUUID
		var uuid string
		if err = common.ConvertType(data, &uuid); err != nil {
			break
		}
		resData, err = deleteSessionByUUID(uuid)
	default:
		err = errors.New("recieved unknown function request")
	}
	if err != nil {
		resData = err
	}
	return &common.Message{
		Service:  common.Response,
		FuncType: resFuncType,
		Data:     resData,
	}
}

//////////////////////////////////////////////////
// works on tls only
func createUser(user *models.User) (affected int64, err error) {
	user.UuId = common.NewUUIDString()
	user.Password = common.Encrypt(user.Password)
	user.CreatedAt = time.Now()
	affected, err = engine.Table("users").InsertOne(user)
	return
}

//////////////////////////////////////////////////
// these methods have to work with gin/session
func createSession(user *models.User) (affected int64, session *models.Session, err error) {
	session = &models.Session{
		UuId:      common.NewUUIDString(),
		Name:      user.Name,
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	affected, err = engine.Table("sessions").InsertOne(&session)
	return
}

func getSessionByUUID(uuid string) (ok bool, session *models.Session, err error) {
	session = &models.Session{UuId: uuid}
	ok, err = engine.Table("sessions").Get(session)
	return
}

func getUserByUUID(uuid string) (ok bool, user *models.User, err error) {
	user = &models.User{UuId: uuid}
	ok, err = getUser(user)
	return
}

func getUserByEmail(email string) (ok bool, user *models.User, err error) {
	user = &models.User{Email: email}
	ok, err = getUser(user)
	return
}

func getUser(user *models.User) (ok bool, err error) {
	ok, err = engine.Table("users").Get(user)
	return
}

func deleteSessionByUUID(uuid string) (affected int64, err error) {
	del := models.Session{UuId: uuid}
	affected, err = engine.Table("sessions").Delete(&del)
	return
}

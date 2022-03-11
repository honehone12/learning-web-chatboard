package thread

import (
	"chatboard/common"
	"chatboard/models"
	"errors"
	"time"

	"xorm.io/xorm"
)

var engine *xorm.Engine

// engin can be separated from user db
func OpenService(dbEngine *xorm.Engine) {
	engine = dbEngine
}

const (
	Unknown common.FuncTypeT = iota
	GetNumReplies
	GetAllPostsInThread
	GetAllThreads
	GetThreadByUUID
	CreateThread
	CreatePost
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
	case GetNumReplies:
		resFuncType = GetNumReplies
		var id int
		if err = common.ConvertType(data, &id); err != nil {
			break
		}
		resData, err = getNumReplies(id)
	case GetAllPostsInThread:
		resFuncType = GetAllPostsInThread
		var id uint
		if err = common.ConvertType(data, &id); err != nil {
			break
		}
		resData, err = getAllPostsInThread(id)
	case GetAllThreads:
		resFuncType = GetAllThreads
		resData, err = getAllThreads()
	case GetThreadByUUID:
		resFuncType = GetThreadByUUID
		var uuid string
		if err = common.ConvertType(data, &uuid); err != nil {
			break
		}
		var ok bool
		if ok, resData, err = getThreadByUUID(uuid); err != nil {
			break
		} else if !ok {
			err = errors.New("no threads found")
		}
	case CreateThread:
		resFuncType = CreateThread
		var contrib common.Contribution
		if err = common.ConvertType(data, &contrib); err != nil {
			break
		}
		resData, err = createThread(
			contrib.Content,
			contrib.UserID,
			contrib.UserName,
		)
	case CreatePost:
		resFuncType = CreatePost
		var contrib common.Contribution
		if err = common.ConvertType(data, &contrib); err != nil {
			break
		}
		resData, err = createPost(
			contrib.Content,
			contrib.ThreadID,
			contrib.UserID,
			contrib.UserName,
		)
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

func getNumReplies(id int) (int64, error) {
	return engine.Table("posts").Where("thread_id = ?", id).Count()
}

func getAllPostsInThread(id uint) (posts []models.Post, err error) {
	engine.Table("posts").Where("thread_id = ?").Find(&posts)
	return
}

///////////////////////////////////////////////
// here should be changed
// like get 10
func getAllThreads() (threads []models.Thread, err error) {
	err = engine.Table("threads").Desc("created_at").Find(&threads)
	return
}

func getThreadByUUID(uuid string) (ok bool, thread *models.Thread, err error) {
	thread = &models.Thread{UuId: uuid}
	ok, err = engine.Table("threads").Get(thread)
	return
}

func createThread(
	topic string,
	userID uint,
	userName string,
) (affected int64, err error) {
	ins := models.Thread{
		UuId:      common.NewUUIDString(),
		Topic:     topic,
		Owner:     userName,
		UserId:    userID,
		CreatedAt: time.Now(),
	}
	affected, err = engine.Table("threads").InsertOne(&ins)
	return
}

func createPost(
	body string,
	threadID uint,
	userID uint,
	userName string,
) (affected int64, err error) {
	ins := models.Post{
		UuId:        common.NewUUIDString(),
		Body:        body,
		Contributor: userName,
		UserId:      userID,
		ThreadId:    threadID,
		CreatedAt:   time.Now(),
	}
	affected, err = engine.Table("posts").InsertOne(&ins)
	return
}

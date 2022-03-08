package thread

import (
	"chatboard/common"
	"chatboard/message"
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
	Unknown message.FuncTypeT = iota
	GetNumReplies
	GetAllPostsInThread
	GetAllThreads
	GetThreadByUUID
	CreateThread
	CreatePost
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

func callServiceInternal(funcType message.FuncTypeT, data *interface{}) *message.Message {
	var err error
	var resFuncType message.FuncTypeT = Unknown
	var resData interface{}
	switch funcType {
	case GetNumReplies:
		resFuncType = GetNumReplies
		var id int
		if err = message.ConvertType(data, &id); err != nil {
			break
		}
		resData, err = getNumReplies(id)
	case GetAllPostsInThread:
		resFuncType = GetAllPostsInThread
		var id int
		if err = message.ConvertType(data, &id); err != nil {
			break
		}
		resData, err = getAllPostsInThread(id)
	case GetAllThreads:
		resFuncType = GetAllThreads
		resData, err = getAllThreads()
	case GetThreadByUUID:
		resFuncType = GetThreadByUUID
		var uuid string
		if err = message.ConvertType(data, &uuid); err != nil {
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
		var pair message.Pair
		if err = message.ConvertType(data, &pair); err != nil {
			break
		}
		var id int
		var topic string
		if err = message.ConvertType(&pair.ID, &id); err != nil {
			break
		}
		if err = message.ConvertType(&pair.Data, &topic); err != nil {
			break
		}
		resData, err = createThread(topic, id)
	case CreatePost:
		resFuncType = CreatePost
		var pair message.Pair
		if err = message.ConvertType(data, &pair); err != nil {
			break
		}
		var idPair message.IDPair
		var topic string
		if err = message.ConvertType(&pair.ID, &idPair); err != nil {
			break
		}
		if err = message.ConvertType(&pair.Data, &topic); err != nil {
			break
		}
		resData, err = createPost(topic, idPair.TableID, idPair.UserID)
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

func getNumReplies(id int) (int64, error) {
	return engine.Table("posts").Where("thread_id = ?", id).Count()
}

func getAllPostsInThread(id int) (posts []Post, err error) {
	rows, err := engine.Table("posts").Rows(&Post{ThreadId: id})
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		err = rows.Scan(post)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func getAllThreads() (threads []Thread, err error) {
	err = engine.Table("threads").Desc("created_at").Find(&threads)
	return
}

func getThreadByUUID(uuid string) (ok bool, thread *Thread, err error) {
	thread = &Thread{UuId: uuid}
	ok, err = engine.Table("threads").Get(thread)
	return
}

func createThread(topic string, userID int) (affected int64, err error) {
	ins := Thread{
		UuId:      common.NewUUIDString(),
		Topic:     topic,
		UserId:    userID,
		CreatedAt: time.Now(),
	}
	affected, err = engine.Table("threads").InsertOne(&ins)
	return
}

func createPost(body string, id int, userID int) (affected int64, err error) {
	ins := Post{
		UuId:      common.NewUUIDString(),
		Body:      body,
		UserId:    userID,
		ThreadId:  id,
		CreatedAt: time.Now(),
	}
	affected, err = engine.Table("posts").InsertOne(&ins)
	return
}

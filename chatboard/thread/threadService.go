package thread

import (
	"chatboard/message"
	"errors"

	"xorm.io/xorm"
)

var engine *xorm.Engine

// engin can be separated from user db
func OpenService(dbEngine *xorm.Engine) {
	engine = dbEngine
}

const (
	Unknown message.FuncTypeT = iota
	CountNumReplies
	GetAllPostsInThread
	CreateThread
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
	switch funcType {
	case CountNumReplies:
		resFuncType = CountNumReplies
		var id int
		if err = convertType(data, &id); err == nil {
			if res, err := countNumReplies(id); err == nil {
				return &message.Message{
					Service:  message.Respose,
					FuncType: resFuncType,
					Data:     res,
				}
			}
		}
	case GetAllPostsInThread:
		resFuncType = GetAllPostsInThread
		var id int
		if err = convertType(data, &id); err == nil {
			if res, err := getAllPostsInThread(id); err == nil {
				return &message.Message{
					Service:  message.Respose,
					FuncType: resFuncType,
					Data:     res,
				}
			}
		}
	default:
		err = errors.New("recieved unknown function request")
	}
	return &message.Message{
		Service:  message.Respose,
		FuncType: resFuncType,
		Data:     err,
	}
}

func convertType(data *interface{}, target interface{}) (err error) {
	switch target := target.(type) {
	case *int:
		data, ok := (*data).(int)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into int")
		}
	default:
		err = errors.New("unknown convert target")
	}
	return
}

func countNumReplies(id int) (int64, error) {
	return engine.Table("post").Where("thread_id = ?", id).Count()
}

func getAllPostsInThread(id int) (posts []Post, err error) {
	rows, err := engine.Table("post").Rows(&Post{ThreadID: id})
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

func createThread() {

}

package thread

import "time"

type Thread struct {
	Id        int       `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Topic     string    `xorm:"topic"`
	UserId    int       `xorm:"user_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Post struct {
	Id        int       `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Body      string    `xorm:"body"`
	UserId    int       `xorm:"user_id"`
	ThreadId  int       `xorm:"thread_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

func (thread *Thread) When() string {
	return thread.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

func (post *Post) When() string {
	return post.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

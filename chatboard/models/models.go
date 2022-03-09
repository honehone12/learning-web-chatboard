package models

import "time"

type User struct {
	Id        uint      `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"email"`
	Password  string    `xorm:"password"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Session struct {
	Id        uint      `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Email     string    `xorm:"email"`
	UserId    uint      `xorm:"user_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Thread struct {
	Id         uint      `xorm:"id"`
	UuId       string    `xorm:"uu_id"`
	Topic      string    `xorm:"topic"`
	NumReplies uint      `xorm:"num_replies"`
	Owner      string    `xorm:"owner"`
	UserId     uint      `xorm:"user_id"`
	CreatedAt  time.Time `xorm:"created_at"`
}

type Post struct {
	Id          uint      `xorm:"id"`
	UuId        string    `xorm:"uu_id"`
	Body        string    `xorm:"body"`
	Contributor string    `xorm:"contributor"`
	UserId      uint      `xorm:"user_id"`
	ThreadId    uint      `xorm:"thread_id"`
	CreatedAt   time.Time `xorm:"created_at"`
}

func (thread *Thread) When() string {
	return thread.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

func (post *Post) When() string {
	return post.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

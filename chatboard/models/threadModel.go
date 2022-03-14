package models

import "time"

type Thread struct {
	Id         uint      `xorm:"pk autoincr 'id'"`
	UuId       string    `xorm:"not null unique 'uu_id'"`
	Topic      string    `xorm:"TEXT 'topic'"`
	NumReplies uint      `xorm:"num_replies"`
	Owner      string    `xorm:"owner"`
	UserId     uint      `xorm:"user_id"`
	CreatedAt  time.Time `xorm:"not null 'created_at'"`
}

type Post struct {
	Id          uint      `xorm:"ok autoincr 'id'"`
	UuId        string    `xorm:"not null unique 'uu_id'"`
	Body        string    `xorm:"TEXT 'body'"`
	Contributor string    `xorm:"contributor"`
	UserId      uint      `xorm:"user_id"`
	ThreadId    uint      `xorm:"thread_id"`
	CreatedAt   time.Time `xorm:"not null 'created_at'"`
}

func (thread *Thread) When() string {
	return thread.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

func (post *Post) When() string {
	return post.CreatedAt.Format("2006/Jan/2 at 3:04pm")
}

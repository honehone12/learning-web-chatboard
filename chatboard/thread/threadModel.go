package thread

import "time"

type Thread struct {
	ID        int       `xorm:"id"`
	UUID      string    `xorm:"uuid"`
	Topic     string    `xorm:"topic"`
	UserID    int       `xorm:"user_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Post struct {
	ID        int       `xorm:"id"`
	UUID      string    `xorm:"uuid"`
	Body      string    `xorm:"body"`
	UserID    int       `xorm:"user_id"`
	ThreadID  int       `xorm:"thread_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

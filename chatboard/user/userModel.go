package user

import "time"

type User struct {
	ID        int       `xorm:"id"`
	UUID      string    `xorm:"uuid"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"email"`
	Password  string    `xorm:"password"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Session struct {
	ID        int       `xorm:"id"`
	UUID      string    `xorm:"uuid"`
	Email     string    `xorm:"email"`
	UserID    int       `xorm:"user_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

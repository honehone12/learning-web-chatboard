package user

import "time"

type User struct {
	Id        int       `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"email"`
	Password  string    `xorm:"password"`
	CreatedAt time.Time `xorm:"created_at"`
}

type Session struct {
	Id        int       `xorm:"id"`
	UuId      string    `xorm:"uu_id"`
	Email     string    `xorm:"email"`
	UserId    int       `xorm:"user_id"`
	CreatedAt time.Time `xorm:"created_at"`
}

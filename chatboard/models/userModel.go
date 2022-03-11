package models

import "time"

type User struct {
	Id        uint      `xorm:"pk autoincr 'id'"`
	UuId      string    `xorm:"not null unique 'uu_id'"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"not null unique 'email'"`
	Password  string    `xorm:"not null 'password'"`
	CreatedAt time.Time `xorm:"not null 'created_at'"`
}

type Session struct {
	Id        uint      `xorm:"pk autoincr 'id'"`
	UuId      string    `xorm:"not null unique 'uu_id'"`
	Name      string    `xorm:"name"`
	Email     string    `xorm:"email"`
	UserId    uint      `xorm:"user_id"`
	CreatedAt time.Time `xorm:"not null 'created_at'"`
}

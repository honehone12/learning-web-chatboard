package common

import (
	"chatboard/models"
	"errors"
)

// communications between different services
// are done with Message

type ServiceT uint8

const (
	ServiceCall ServiceT = iota
	Respose     ServiceT = iota
)

type FuncTypeT uint8

type Message struct {
	Service  ServiceT
	FuncType FuncTypeT
	// prefer converting to []byte
	Data interface{}
}

type Contribution struct {
	ThreadID uint
	Content  string
	UserID   uint
	UserName string
}

func ConvertType(data *interface{}, target interface{}) (err error) {
	switch target := target.(type) {
	case *uint:
		data, ok := (*data).(uint)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into uint")
		}
	case *string:
		data, ok := (*data).(string)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into string")
		}
	case *models.User:
		data, ok := (*data).(models.User)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into User")
		}
	case *models.Session:
		data, ok := (*data).(models.Session)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into Session")
		}
	case *models.Thread:
		data, ok := (*data).(models.Thread)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into Thread")
		}
	case *models.Post:
		data, ok := (*data).(models.Post)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into Thread")
		}
	case *Contribution:
		data, ok := (*data).(Contribution)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into Contribution")
		}
	default:
		err = errors.New("unknown convert target")
	}
	return
}

package message

import (
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

type Pair struct {
	ID   interface{}
	Data interface{}
}

type IDPair struct {
	UserID  int
	TableID int
}

func ConvertType(data *interface{}, target interface{}) (err error) {
	switch target := target.(type) {
	case *int:
		data, ok := (*data).(int)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into int")
		}
	case *string:
		data, ok := (*data).(string)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into string")
		}
	case *Pair:
		data, ok := (*data).(Pair)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into Pair")
		}
	case *IDPair:
		data, ok := (*data).(IDPair)
		if ok {
			*target = data
		} else {
			err = errors.New("failed to convert data into IDPair")
		}
	default:
		err = errors.New("unknown convert target")
	}
	return
}

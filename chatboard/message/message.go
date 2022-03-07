package message

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

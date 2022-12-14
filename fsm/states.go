package fsm

type State interface {
	ID() string
}

type SimpleState struct {
	Id string
}

func NewSimpleState(id string) *SimpleState {
	return &SimpleState{Id: id}
}

func (n *SimpleState) ID() string { return n.Id }

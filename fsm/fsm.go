package fsm

const MAX_NODES = 100
const MAX_LINKS = 100

type Context = map[string]string

type FSM struct {
	Nodes []Node
}

func NewFsm(nodes []Node) *FSM {
	fsm := FSM{Nodes: nodes}
	return &fsm
}

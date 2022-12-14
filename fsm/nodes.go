package fsm

import (
	"log"
)

type Node interface {
	State
	Transition
	Run(Context) Context
	IsTerminal() bool
}

type PassNode struct {
	State      SimpleState
	Action     NoOp
	Transition *SimpleTransition
	isTerminal bool
}

func NewPassNode(id string, to string, isTerminal bool) *PassNode {
	var t SimpleTransition
	if !isTerminal {
		t = *NewSimpleTransition(to)
	}
	n := &PassNode{
		State:      *NewSimpleState(id),
		Action:     NoOp{},
		Transition: &t,
		isTerminal: isTerminal,
	}
	return n
}

func (n *PassNode) ID() string {
	return n.State.ID()
}

func (n *PassNode) To(ctx Context) string {
	return n.Transition.To(ctx)
}

func (n *PassNode) IsTerminal() bool {
	return n.isTerminal
}

func (n *PassNode) Run(context Context) Context {
	log.Printf("Running node %s\n", n.ID())
	return context
}

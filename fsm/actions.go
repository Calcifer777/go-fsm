package fsm

import (
	"fmt"
)

type Action interface {
	Run(ctx Context) Context
}

type NoOp struct{}

func (e *NoOp) Run(ctx Context) Context { return ctx }

type Echo struct {
	message string
}

func (e *Echo) Run(ctx Context) Context {
	fmt.Println(e.message)
	return ctx
}

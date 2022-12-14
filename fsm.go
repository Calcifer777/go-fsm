package main

import "log"

type StateType int

const (
	Start StateType = iota
	Inner
	Terminal
)

type State struct {
	id        string
	stateType StateType
}

type Context map[string]string
type Action func(Context) Context

type Event struct {
	name   string
	action Action
}

type Rule func(ctx Context) bool

type Transition struct {
	from  *State
	to    *State
	rule  Rule
	event Event
}

type Fsm struct {
	states      []State
	transitions []Transition
}

type Runner struct {
	fsm          Fsm
	context      Context
	currentState *State
}

func (r *Runner) Next() bool {
	if r.currentState.stateType == Terminal {
		return false
	}
	// Find relevant transition
	var transition Transition
	for _, t := range r.fsm.transitions {
		if t.from == r.currentState && t.rule(r.context) {
			transition = t
		}
	}
	// Fire Event
	log.Printf("Firing event %s from state %s", transition.event.name, r.currentState.id)
	r.context = transition.event.action(r.context)
	r.currentState = transition.to
	return true
}

func (r *Runner) Run() {
	var ok bool
	for {
		ok = r.Next()
		if !ok {
			break
		}
	}
}

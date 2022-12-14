package main

import (
	"log"
)

func AlwaysFire(ctx Context) bool { return true }

func NoOp(ctx Context) Context { return ctx }

func main() {

	noOp := Event{name: "no-op", action: NoOp}

	states := []State{
		{id: "start", stateType: Start},
		{id: "step1", stateType: Other},
		{id: "end", stateType: Terminal},
	}

	transitions := []Transition{
		{from: &states[0], to: &states[1], rule: AlwaysFire, event: noOp},
		{from: &states[1], to: &states[2], rule: AlwaysFire, event: noOp},
	}

	fsm := Fsm{states: states, transitions: transitions}

	runner := Runner{
		fsm:          fsm,
		context:      nil,
		currentState: &states[0],
	}

	log.Printf("FSM starting state: %s", runner.currentState.id)
	runner.Run()
	log.Printf("FSM ending state: %s", runner.currentState.id)

}

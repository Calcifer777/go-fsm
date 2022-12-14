package main

import (
	"log"
	"testing"
)

func AlwaysFire(ctx Context) bool                       { return true }
func Choice(key string, value string, ctx Context) bool { return ctx[key] == value }

func NoOp(ctx Context) Context { return ctx }

func TestSimpleFsm(t *testing.T) {
	noOp := Event{name: "no-op", action: NoOp}

	states := []State{
		{id: "start", stateType: Start},
		{id: "step1", stateType: Inner},
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
	if runner.currentState.id != "end" {
		t.Errorf("Expected terminal state in FSM, got %s", runner.currentState.id)
	}
}

func TestFork(t *testing.T) {
	noOp := Event{name: "no-op", action: NoOp}

	states := []State{
		{id: "start", stateType: Start},
		{id: "on", stateType: Inner},
		{id: "off", stateType: Inner},
		{id: "end", stateType: Terminal},
	}

	ifOn := func(ctx Context) bool { return Choice("status", "turn on", ctx) }
	ifOff := func(ctx Context) bool { return Choice("status", "turn off", ctx) }

	transitions := []Transition{
		{from: &states[0], to: &states[1], rule: AlwaysFire, event: noOp},
		{from: &states[0], to: &states[2], rule: ifOn, event: noOp},
		{from: &states[1], to: &states[3], rule: ifOff, event: noOp},
		{from: &states[2], to: &states[3], rule: AlwaysFire, event: noOp},
	}

	fsm := Fsm{states: states, transitions: transitions}

	runner := Runner{
		fsm:          fsm,
		context:      nil,
		currentState: &states[0],
	}

	ok := runner.Next()
	if !ok || runner.currentState.id != "on" {
		t.Errorf("Expected landing state 'on', got '%s'", runner.currentState.id)
	}

}

package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

type FsmState struct {
	value string
	id    int64
}

func (s FsmState) ID() int64 { return s.id }

type FsmEdge struct {
	F, T FsmState
	id   int64
}

func (e FsmEdge) From() graph.Node         { return e.F }
func (e FsmEdge) To() graph.Node           { return e.T }
func (e FsmEdge) ReversedLine() graph.Line { return FsmEdge{F: e.T, T: e.F} }
func (e FsmEdge) ID() int64                { return e.id }

type Fsm struct {
	currentState FsmState
	g            *multi.DirectedGraph
	nodeCounter  int64
	edgeCounter  int64
}

func InitFsm() *Fsm {
	s := &Fsm{}
	s.g = multi.NewDirectedGraph()
	return s
}

func (fsm *Fsm) WithNode(stateName string) FsmState {
	s := FsmState{value: stateName, id: fsm.nodeCounter}
	fsm.currentState = s
	fsm.g.AddNode(s)
	fsm.nodeCounter++
	return fsm.currentState
}

func (fsm *Fsm) WithLink(s1, s2 FsmState) *Fsm {
	fsm.g.SetLine(FsmEdge{F: s1, T: s2, id: fsm.edgeCounter})
	return fsm
}

func main1() {
	state := FsmState{value: "s1", id: 0}
	i := interface{}(state)
	_, ok := i.(graph.Node)
	fmt.Println(ok)
	s := InitFsm()
	n1 := s.WithNode("n1")
	n2 := s.WithNode("n2")
	s.WithLink(n1, n2)
}

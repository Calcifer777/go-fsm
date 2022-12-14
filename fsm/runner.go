package fsm

import "errors"

type FsmRunner struct {
	Fsm           FSM
	CurrentNodeId string
}

func NewFsmRunner(fsm FSM, s string) *FsmRunner {
	return &FsmRunner{
		Fsm:           fsm,
		CurrentNodeId: s,
	}
}

func (runner *FsmRunner) RunStep(ctx Context) (Context, bool, error) {
	var node Node
	for _, n := range runner.Fsm.Nodes {
		if n.ID() == runner.CurrentNodeId {
			node = n
		}
	}
	if node == nil {
		return nil, false, errors.New("current node not found in the FSM")
	} else if node.IsTerminal() {
		return nil, false, nil
	} else {
		out := node.Run(ctx)
		runner.CurrentNodeId = node.To(ctx)
		return out, true, nil
	}
}

func (runner *FsmRunner) Run(ctx Context) (Context, error) {
	for {
		ctx, run, err := runner.RunStep(ctx)
		if err != nil {
			return nil, err
		}
		if !run {
			return ctx, nil
		}
	}
}

package fsm

type Transition interface {
	To(context Context) string
}

type SimpleTransition struct {
	to string
}

func (t *SimpleTransition) To(context Context) string { return t.to }

func NewSimpleTransition(to string) *SimpleTransition {
	return &SimpleTransition{to: to}
}

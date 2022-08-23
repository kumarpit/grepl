package fsm

import "strconv"

type StateGenerator interface {
	Next() State
	NextAccepting() State
}

type NumericStateGenerator struct {
	counter int
}

func (sg *NumericStateGenerator) Next() State {
	state := NewState(strconv.Itoa(sg.counter))
	sg.counter++
	return state
}

func (sg *NumericStateGenerator) NextAccepting() State {
	state := NewAcceptingState(strconv.Itoa(sg.counter))
	sg.counter++
	return state
}
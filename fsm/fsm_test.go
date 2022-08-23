package fsm

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var (
	openState   = NewAcceptingState("open")
	closedState = NewAcceptingState("closed")
	brokenState = NewState("broken")
)

var (
	open    = "open"
	close   = "close"
	kick    = "kick"
	invalid = "foo"
)

func TestFSM(t *testing.T) {
	initial := openState
	machine := New(initial, []Transition {
		{event: open, source: closedState, nextState: openState},
		{event: close, source: openState, nextState: closedState},
		{event: kick, source: closedState, nextState: brokenState},
	})
	
	result := machine.Run([]string{close, open, close, kick})
	assert.False(t, result)
}
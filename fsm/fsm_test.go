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

func Test(t *testing.T) {
	initial := openState
	machine, err := New(inital, []Transition {
		{event: open, source: closedState, nextState: openState},
		{event: close, source: openState, nextState: closedState},
		{event: kick, source: closedState, nextState: brokenState}
	})
	assert.NoError(t, err)
	result := machine.Run([]string{close, open, close})
	assert.True(t, result)
}
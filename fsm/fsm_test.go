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
		{Event: open, Source: closedState, NextState: openState},
		{Event: close, Source: openState, NextState: closedState},
		{Event: kick, Source: closedState, NextState: brokenState},
	})
	
	result := machine.Run([]string{close, open, close, kick})
	assert.False(t, result)
}
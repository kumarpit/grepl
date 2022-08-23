package regex2fsm

import (
	"github.com/kumarpit/grepl/fsm"
	"regexp/syntax"
)

type Parser struct {
	stateGenerator fsm.StateGenerator
}

func (p Parser) getNextState(isAccepting bool) fsm.State {
	if isAccepting {
		return p.stateGenerator.NextAccepting()
	}

	return p.stateGenerator.Next()
}

func (p Parser) Convert(pattern string) (*fsm.StateMachine, error) {
	regexTree, err := syntax.Parse(pattern, syntax.POSIX)
	if err != nil {
		return nil, err
	}

	initialState := p.stateGenerator.Next()
	transitions  := p.parseTree(initialState, regexTree, true)

	return fsm.New(initialState, transitions), nil
}

func (p Parser) parseTree(currentState fsm.State, tree *syntax.Regexp, isAccepting bool) []fsm.Transition {
	return nil
}
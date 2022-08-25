package regex2fsm

import (
	"github.com/kumarpit/grepl/fsm"
	"regexp/syntax"
	"fmt"
)

type Parser struct {
	stateGenerator fsm.StateGenerator
}

func New() *Parser {
	return &Parser{
		stateGenerator: &fsm.stateGenerator{},
	}
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

func (p Parser) ParseTree(currentState fsm.State, tree *syntax.Regexp, isAccepting bool) []fsm.Transition {
	switch tree.Op {
	case syntax.OpAlternate:
		return g.ParseAlternate(currentState, tree, isAccepting)
	
	case syntax.OpLiteral:
		return g.ParseLiteral(currentState, tree, isAccepting)
	
	case syntax.OpStar:
		return g.ParseStar(currentState, tree, isAccepting)
	
	case syntax.OpPlus:
		return g.ParsePlus(currentState, tree, isAccepting)

	case syntax.OpConcat:
		return g.ParseConcat(currentState, tree, isAccepting)

	case syntax.OpCharClass:
		return g.ParseCharClass(currentState, tree, isAccepting)

	default:
		panic(fmt.Sprintf("unsupported operation: %s", tree.op))
	}
}
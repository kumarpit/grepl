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
		// !!!
		stateGenerator: &fsm.NumericStateGenerator{},
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
	transitions  := p.ParseTree(initialState, regexTree, true)

	return fsm.New(initialState, transitions), nil
}

func (p Parser) ParseTree(currentState fsm.State, tree *syntax.Regexp, isAccepting bool) []fsm.Transition {
	switch tree.Op {
	case syntax.OpAlternate:
		return p.ParseAlternate(currentState, tree, isAccepting)
	
	case syntax.OpLiteral:
		return p.ParseLiteral(currentState, tree, isAccepting)
	
	case syntax.OpStar:
		return p.ParseStar(currentState, tree, isAccepting)
	
	case syntax.OpPlus:
		return p.ParsePlus(currentState, tree, isAccepting)

	case syntax.OpConcat:
		return p.ParseConcat(currentState, tree, isAccepting)

	case syntax.OpCharClass:
		return p.ParseCharClass(currentState, tree, isAccepting)

	default:
		panic(fmt.Sprintf("unsupported operation: %s", tree.Op))
	}
}
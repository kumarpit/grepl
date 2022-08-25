package regex2fsm

import (
	"regexp/syntax"
	"github.com/kumarpit/grepl/fsm"
)

func (g Parser) ParseAlternate(currentState fsm.State, alternate *syntax.Regexp, isAccepting bool) []fsm.Transition {
	left  := ParseTree(currentState, alternate.Sub[0], isAccepting)
	right := ParseTree(currentState, alternate.Sub[1], isAccepting)
	return append(left, right...)
}

func (g Parser) ParseLiteral(currentState fsm.State, literal *syntax.Regexp, isAccepting bool) []fsm.Transition {
	transitions := []fsm.Transition{}
	last := currentState

	for i, c := range literal.Rune {
		isLast := i == len(literal.Rune) - 1
		nextState := g.getNextState(isAccepting && isLast)
		transitions = append(transitions, fsm.Transition {
			event: string(c),
			source: fsm.State(last),
			nextState: nextState,
		})
		last = nextState
	}

	return transitions
}

// zero or more
func (g Parser) ParseStar(currentState fsm.State, star *syntax.Regexp, isAccepting bool) []fsm.Transition {
	// make current state accepting
	// add self transition to new state
	if isAccepting {
		currentState = currentState.MakeAccepting()
	}
	nextState := g.getNextState(isAccepting);

	return []fsm.transitions {
		{
			event: string(star.Sub[0].Rune[0]),
			source: currentState,
			nextState: nextState,
		},
		{
			event: string(star.Sub[0].Rune[0]),
			source: nextState,
			nextState: nextState,
		},
	}
}

// one or more
func (g Parser) ParsePlus(currentState fsm.State, plus *syntax.Regexp, isAccepting bool) []fsm.Transition {
	nextState := g.getNextState(isAccepting)
	return []fsm.transition {
		{
			event: string(plus.Sub[0].Rune[0]),
			source: currentState,
			nextState: nextState,
		},
		{
			event: string(plus.Sub[0].Rune[0]),
			source: nextState,
			nextState: nextState,
		},
	}
}

func (g Parser) ParseConcat(currentState fsm.State, concat *syntax.Regexp, isAccepting bool) []fsm.Transition {
	source := currentState
	transitions := []fsm.transition{}
	for i, exp := range concat.Sub {
		isLast := i == len(concat.Sub) - 1
		subTransitions := g.ParseTree(source, exp, isAccepting && isLast)
		source = subTransitions[len(subTransitions) - 1].nextState
		transitions = append(transitions, subTransitions...)
	}

	return transitions
}

func (g Parser) ParseCharClass(currentState fsm.State, charClass *syntax.Regexp, isAccepting bool) []fsm.Transition {
	
}

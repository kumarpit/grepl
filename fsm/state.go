package fsm

type State struct {
	value string
	isAccepting bool
}

func NewAcceptingState(value string) State {
	return State{
		value: value,
		isAccepting: true,
	}
}

func NewState(value string) State {
	return State{
		value: value,
		isAccepting: false,
	}
}

func MakeAccepting(state State) State {
	return NewAcceptingState(state.Value())
}

func (state State) Equal(other State) bool {
	return state.value == other.value
}

func (state State) Value() string {
	return state.value
}

func (state State) Accepting() bool {
	return state.isAccepting
}
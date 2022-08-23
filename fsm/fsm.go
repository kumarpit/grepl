package fsm

type Transition struct {
	event string
	source State
	nextState State
}

type StateMachine struct {
	initialState State
	currentState State
	transitions []Transition
}

func New(initial State, transitions []Transition) *StateMachine {
	return &StateMachine {
		initial,
		initial,
		transitions,
	}
}

func (m *StateMachine) Run(events []string) bool {
	for _, event := range events {
		transition := m.findTransition(event)
		if transition == nil {
			break
		}
		m.currentState = transition.nextState
	}

	return m.currentState.Accepting()
}

func (m *StateMachine) Reset() {
	m.currentState = m.initialState
}

func (m *StateMachine) findTransition(event string) *Transition {
	for _, t := range m.transitions {
		if t.source.Equal(m.currentState) && t.event == event {
			return &t
		}	
	}
	
	return nil
}
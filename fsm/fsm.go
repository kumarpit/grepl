package fsm

type Transition struct {
	Event string
	Source State
	NextState State
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
		m.currentState = transition.NextState
	}

	return m.currentState.Accepting()
}

func (m *StateMachine) Reset() {
	m.currentState = m.initialState
}

func (m *StateMachine) findTransition(event string) *Transition {
	for _, t := range m.transitions {
		if t.Source.Equal(m.currentState) && t.Event == event {
			return &t
		}	
	}
	
	return nil
}
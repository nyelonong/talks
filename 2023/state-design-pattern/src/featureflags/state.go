package main

type state interface {
	updateKey() error
	approveKey() error
	rejectKey() error
}

const (
	roleUser string = "user"
	roleLead string = "lead"
)

type featureflag struct {
	keyActive    state
	keyRequested state
	keyRejected  state
	keyApproved  state

	currentState state

	key   string
	value bool

	role string
}

func checkKey(key string, value bool, role string) *featureflag {
	ff := &featureflag{
		key:   key,
		value: value,
		role:  role,
	}

	keyActiveState := &keyActiveState{
		ff: ff,
	}

	keyRequestedState := &keyRequestedState{
		ff: ff,
	}

	keyApprovedState := &keyApprovedState{
		ff: ff,
	}

	keyRejectedState := &keyRejectedState{
		ff: ff,
	}

	ff.setState(keyActiveState)
	ff.keyActive = keyActiveState
	ff.keyRequested = keyRequestedState
	ff.keyApproved = keyApprovedState
	ff.keyRejected = keyRejectedState

	return ff
}

func (c *featureflag) updateKey() error {
	return c.currentState.updateKey()
}

func (c *featureflag) approveKey() error {
	return c.currentState.approveKey()
}

func (c *featureflag) rejectKey() error {
	return c.currentState.rejectKey()
}

func (c *featureflag) setState(s state) {
	c.currentState = s
}

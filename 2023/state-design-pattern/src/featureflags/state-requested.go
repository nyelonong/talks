package main

import "fmt"

type keyRequestedState struct {
	ff *featureflag
}

func (u *keyRequestedState) updateKey() error {
	return fmt.Errorf("can not update requested key")
}

func (u *keyRequestedState) approveKey() error {
	if u.ff.role != roleLead {
		return fmt.Errorf("only lead can approve key")
	}

	fmt.Println("Key approved.")
	u.ff.setState(u.ff.keyApproved)
	return nil
}

func (u *keyRequestedState) rejectKey() error {
	if u.ff.role != roleLead {
		return fmt.Errorf("only lead can reject key")
	}

	fmt.Println("Key Rejected.")
	u.ff.setState(u.ff.keyRejected)
	return nil
}

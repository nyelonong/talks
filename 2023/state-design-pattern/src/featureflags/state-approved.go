package main

import "fmt"

type keyApprovedState struct {
	ff *featureflag
}

func (u *keyApprovedState) updateKey() error {
	fmt.Println("Proposal key update is placed.")
	u.ff.setState(u.ff.keyRequested)
	return nil
}

func (u *keyApprovedState) approveKey() error {
	return fmt.Errorf("key is already approved")
}

func (u *keyApprovedState) rejectKey() error {
	return fmt.Errorf("no key in requested state")
}

package main

import "fmt"

type keyRejectedState struct {
	ff *featureflag
}

func (u *keyRejectedState) updateKey() error {
	fmt.Println("Proposal key update is placed.")
	u.ff.setState(u.ff.keyRequested)
	return nil
}

func (u *keyRejectedState) approveKey() error {
	return fmt.Errorf("no key in requested state")
}

func (u *keyRejectedState) rejectKey() error {
	return fmt.Errorf("key is already rejected")
}

package main

import "fmt"

type keyActiveState struct {
	ff *featureflag
}

func (u *keyActiveState) updateKey() error {
	fmt.Println("Proposal key update is placed.")
	u.ff.setState(u.ff.keyRequested)
	return nil
}

func (u *keyActiveState) approveKey() error {
	return fmt.Errorf("key is not in requested state")
}

func (u *keyActiveState) rejectKey() error {
	return fmt.Errorf("key is not in requested state")
}

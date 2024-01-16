package main

import (
	"fmt"
)

func main() {
	// Do not store
	p := 101
	q := 23

	// Public Key
	n := p * q

	// User/Prover Master Secret
	s := [3]int{5, 7, 3}

	// Verfier generate 3 random number for each request
	a := [3]int{1, 0, 1}

	// Prover generate random number for each request
	r := 14

	// Prover compuite X
	x := pow(r, 2) % n

	var v [3]int
	stemp := r
	ytemp := x
	for i, si := range s {
		stemp = stemp * pow(si, a[i])

		v[i] = pow(si, 2) % n
		ytemp = ytemp * pow(v[i], a[i])
	}
	y := stemp % n

	verif := ytemp % n
	fmt.Printf("Verifier = %d\n", verif)

	proof := pow(y, 2) % n
	fmt.Printf("Profer = %d\n", proof)

	fmt.Printf("Valid = %t\n", proof == verif)
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

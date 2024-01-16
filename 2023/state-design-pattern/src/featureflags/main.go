package main

import "log"

func main() {
	flag := checkKey("feature-x", true, roleUser)

	// if err := flag.updateKey(); err != nil {
	// 	log.Println(err)
	// }

	if err := flag.approveKey(); err != nil {
		log.Println(err)
	}
}

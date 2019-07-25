package utils

import "log"

// CheckErr revisa si hay un error
func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
		panic(err)
	}
}

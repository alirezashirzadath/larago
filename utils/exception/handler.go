package exception

import "fmt"

func Exception(errorType string, message string) {

	fmt.Println("START ERROR IN TYPE " + errorType)
	fmt.Println(message)
	fmt.Println("END ERROR IN TYPE " + errorType)

	//TODO Implement other logics
}

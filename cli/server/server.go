package main

import (
	"github.com/txcary/investment/restful"
	"fmt"
)

func main() {
	restful.StartServer("8080")	
	err := restful.Wait()
	fmt.Println(err)
}
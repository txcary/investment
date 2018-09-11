package main

import (
	"github.com/txcary/investment/restful"
)

func main() {
	restful.StartServer()	
	restful.Wait()
}
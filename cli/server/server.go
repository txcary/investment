package main

import (
	"github.com/txcary/investment/restful"
	"time"
)

func main() {
	restful.StartServer()	
	time.Sleep(1000*time.Second)
}
package main
import(
	"github.com/txcary/investment/restful"
)

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"time"
)

func clientGet(id string) {
	resp, err := http.Get("http://localhost:8080/stock/"+id)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func startServer() {
	restful.StartServer("8080")
}


func main() {
	if len(os.Args) > 1{
		startServer()
		time.Sleep(1000)	
		clientGet(os.Args[1])
	}else{
		fmt.Println("Usage: stockvalue <id>")
	}
}
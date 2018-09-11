package restful
import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func clientGet() {
	//resp, err := http.Get("http://localhost:8080/stock/{00700}")
	resp, err := http.Get("http://localhost:8080/stock/00700")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func ExampleStartServer() {
	StartServer()
	clientGet()
	//output:
	//Listening 8080
	//MyId=00700
	//TODO
}
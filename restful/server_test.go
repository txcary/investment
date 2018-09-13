package restful
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
)

const (
	serverPort string = "8090"
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

func clientPost() {
	resBytes, err := httpPostJson([]byte(`{"name":"test"}`), "http://localhost:"+serverPort+"/portfolio/getjson");
	if err==nil {
		fmt.Println(string(resBytes))
	}
}

func httpPostJson(postBody []byte, url string) ([]byte, error) {
	var resp *http.Response
	resp, err := http.Post(url, "application/json", bytes.NewReader(postBody))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

func ExampleStartServer() {
	StartServer(serverPort)
	clientPost()
	//clientGet()

	//output:
	//Listening 8080
	//MyId=00700
	//TODO
}
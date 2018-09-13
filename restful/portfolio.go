package restful
import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type TestStruct struct {
	Name string
	Value string
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cmd := vars["cmd"]
	fmt.Println("command:", cmd)

	data, err := ioutil.ReadAll(r.Body)
	if err==nil {
		fmt.Println("json:", string(data))
	}
	
	res := new(TestStruct)
	res.Name = "ResName"
	res.Value = "ResValue"

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
}

func (obj *Server) InitPortfolio() {
	obj.router.HandleFunc("/portfolio/{cmd}", PortfolioHandler)
}
package restful
import (
	"errors"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"fmt"
	"github.com/txcary/investment/portfolio"
)

func getErrorMessage(err error) []byte {
	if err == nil {
		return []byte(`{"code"=0}`)
	} else {
		return []byte(`{"code":1, "msg":"`+err.Error()+`"}`)
	}
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cmd := vars["cmd"]
	fmt.Println("command:", cmd)

	jsonIn, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		w.Write(getErrorMessage(err))
		return
	}
	fmt.Println("json-in:", string(jsonIn))
/*	
	res := new(TestStruct)
	res.Name = "ResName"
	res.Value = "ResValue"
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
*/

	if cmd == "putjson" {
		err = portfolio.Instance().PutJson(jsonIn)
		w.Write(getErrorMessage(err))
		return
	}
	if cmd == "getjson" {
		outputJson, err := portfolio.Instance().GetJson(jsonIn)
		if err!=nil {
			w.Write(getErrorMessage(err))
			return
		}
		fmt.Println("json-out:", string(outputJson))
		w.Write(outputJson)
		return
	}
	w.Write(getErrorMessage(errors.New(cmd+" not Supported!")))
}

func (obj *Server) InitPortfolio() {
	obj.router.HandleFunc("/portfolio/{cmd}", PortfolioHandler)
}
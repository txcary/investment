package restful
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/txcary/investment/stock"
	"encoding/json"
	"fmt"
)

func StockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("MyId="+id)		

	builder := stock.BuilderInstance()
	stock := builder.Build(id)
	err := json.NewEncoder(w).Encode(stock)
	if err != nil {
		panic(err)
	}
}

func (obj *Server) InitStock() {
	obj.router.HandleFunc("/stock/{id}", StockHandler)
}
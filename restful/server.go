package restful
import(
	"net/http"
	"runtime"
	"fmt"
	"sync"
	"errors"
	"github.com/gorilla/mux"
)

const (
	serverPort string = "8080"
)

type Server struct {
	router *mux.Router	
	wg sync.WaitGroup
}

var serverObj *Server

func (obj *Server) Init() {
	obj.router = mux.NewRouter().StrictSlash(true)
	obj.InitStock()
	obj.InitRoot()
}

func StartServer() {
	if serverObj == nil {
		serverObj := new(Server)
		serverObj.Init()
		go func() {
			wg.Add(1)
			fmt.Println("Listening "+serverPort)
			http.ListenAndServe(":"+serverPort, serverObj.router)	
			wg.Down()
		}()
		runtime.Gosched()
	}
}

func Wait() err {
	if serverObj == nil {
		return errors.New("Server not started!")
	}
	serverObj.wg.Wait()
}
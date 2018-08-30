package restful
import(
	"net/http"
	"runtime"
	"fmt"
	"github.com/gorilla/mux"
)

const (
	serverPort string = "8080"
)

type Server struct {
	router *mux.Router	
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
			fmt.Println("Listening "+serverPort)
			http.ListenAndServe(":"+serverPort, serverObj.router)	
		}()
		runtime.Gosched()
	}
}
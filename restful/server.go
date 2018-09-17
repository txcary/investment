package restful
import(
	"net/http"
	"runtime"
	"sync"
	"errors"
	"github.com/gorilla/mux"
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
	obj.InitPortfolio()
}

func StartServer(serverPort string) {
	if serverObj == nil {
		serverObj = new(Server)
		serverObj.Init()
		go func() {
			serverObj.wg.Add(1)
			http.ListenAndServe(":"+serverPort, serverObj.router)	
			serverObj.wg.Done()
		}()
		runtime.Gosched()
	}
}

func Wait() error {
	if serverObj == nil {
		return errors.New("Server not started!")
	}
	serverObj.wg.Wait()
	return nil
}
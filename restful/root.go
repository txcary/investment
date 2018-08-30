package restful
import (
	"net/http"
	"github.com/txcary/investment/config"
)

func (obj *Server) getRootPath() (path string) {
	path = config.Instance().GetString("restful", "root")
	return
}

func (obj *Server) InitRoot() {
	rootPath := obj.getRootPath()	
	obj.router.Handle("/", http.FileServer(http.Dir(rootPath)))
}
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
	//obj.router.Handle("/web/", http.FileServer(http.Dir(rootPath)))
	obj.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(rootPath))))
}
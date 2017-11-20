package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// 新服务器的设置，返回一个服务器
func NewServer() *negroni.Negroni {

	// 调用render.New方法来初始化一个JsonFormatter（Json格式化器）
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	// 初始化一个negroni实例以及一个mux
	n := negroni.Classic()
	mx := mux.NewRouter()

	// 将路径/xuming/{name}路由到mux中
	initRoutes(mx, formatter)

	// 可以被用作http.Handler，因为类型Router实现了http.Handler的接口
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// 建立一个HandleFunc给mux, 支持GET方法，路由路径即为目的路径/xuming/{name}
	mx.HandleFunc("/xuming/{name}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	// 用闭包来将render.Render应用到一个HandleFunc中
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]
		formatter.JSON(w, http.StatusOK, struct{ PlusOneSecond string }{name+" has donated 1s to the elder."})
	}
}
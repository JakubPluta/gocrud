package router

import (
	"net/http"

	"github.com/JakubPluta/gocrud/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello World"))
	})
	router.GET("/api/posts", postController.GetAll)
	router.GET("/api/posts/:postId", postController.GetById)
	router.POST("/api/posts", postController.Create)
	router.PATCH("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)
	return router
}

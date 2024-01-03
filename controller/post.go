package controller

import (
	"net/http"

	"github.com/JakubPluta/gocrud/data/request"
	"github.com/JakubPluta/gocrud/data/response"
	"github.com/JakubPluta/gocrud/helpers"
	"github.com/JakubPluta/gocrud/service"
	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (controller *PostController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postCreateRequest := request.PostCreateRequest{}
	helpers.ReadRequestBody(r, &postCreateRequest)
	controller.PostService.Create(r.Context(), postCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (controller *PostController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postUpdateRequest := request.PostUpdateRequest{}
	helpers.ReadRequestBody(r, &postUpdateRequest)

	postId := params.ByName("postId")
	postUpdateRequest.Id = postId

	controller.PostService.Update(r.Context(), postUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (controller *PostController) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	results := controller.PostService.GetAll(r.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   results,
	}
	helpers.WriteResponseBody(w, webResponse)
}
func (controller *PostController) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	postId := params.ByName("postId")

	result := controller.PostService.GetById(r.Context(), postId)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (controller *PostController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	postId := params.ByName("postId")

	controller.PostService.Delete(r.Context(), postId)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helpers.WriteResponseBody(w, webResponse)
}

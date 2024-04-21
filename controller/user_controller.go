package controller

import (
	"net/http"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
	"simplejobportal/helper"
	"simplejobportal/service"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadRequestBody(requests, &userCreateRequest)

	if !helper.Valid(userCreateRequest.Email) {
		webResp := response.WebResponse{
			Code:   400,
			Status: "Invalid Email Format",
		}

		helper.WriteResponseBody(writer, webResp)
		return
	}

	resultToken := controller.UserService.Create(requests.Context(), userCreateRequest)
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultToken,
	}

	helper.WriteResponseBody(writer, webResp)
}

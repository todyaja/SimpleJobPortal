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

	resultToken, err := controller.UserService.Create(requests.Context(), userCreateRequest)
	if err != nil {
		webResp := response.WebResponse{
			Code:   200,
			Status: err.Error(),
			Data:   resultToken,
		}

		helper.WriteResponseBody(writer, webResp)
		return
	}
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultToken,
	}

	helper.WriteResponseBody(writer, webResp)
}

func (controller *UserController) Login(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	userLoginRequest := request.UserLoginRequest{}
	helper.ReadRequestBody(requests, &userLoginRequest)

	if !helper.Valid(userLoginRequest.Email) {
		webResp := response.WebResponse{
			Code:   400,
			Status: "Invalid Email Format",
		}

		helper.WriteResponseBody(writer, webResp)
		return
	}

	resultToken, err := controller.UserService.Login(requests.Context(), userLoginRequest)
	if err != nil {
		webResp := response.WebResponse{
			Code:   200,
			Status: "wrong password",
		}

		helper.WriteResponseBody(writer, webResp)
		return
	}
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultToken,
	}

	helper.WriteResponseBody(writer, webResp)
}

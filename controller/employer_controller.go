package controller

import (
	"net/http"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
	"simplejobportal/helper"
	"simplejobportal/service"

	"github.com/julienschmidt/httprouter"
)

type EmployerController struct {
	EmployerService service.EmployerService
}

func NewEmployerController(employerService service.EmployerService) *EmployerController {
	return &EmployerController{EmployerService: employerService}
}

func (controller *EmployerController) CreateJob(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	employerCreateRequest := request.JobCreateRequest{}
	helper.ReadRequestBody(requests, &employerCreateRequest)

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	controller.EmployerService.CreateJob(requests.Context(), employerCreateRequest, token)
	webResp := response.WebResponse{
		Code:   201,
		Status: "OK",
	}

	helper.WriteResponseBody(writer, webResp)
}

func (controller *EmployerController) SeeJobApplicant(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	viewApplicantRequest := request.ViewApplicantRequest{}
	helper.ReadRequestBody(requests, &viewApplicantRequest)

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	data := controller.EmployerService.SeeJobApplicant(requests.Context(), viewApplicantRequest, token)

	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	if data.JobId == 0 {
		webResp = response.WebResponse{
			Code:   200,
			Status: "data is empty, either you are not the job creator or there is no applicants yet!",
		}
	}

	helper.WriteResponseBody(writer, webResp)
}
func (controller *EmployerController) ProcessApplicant(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	processApplicantRequest := request.ProcessApplicantRequest{}
	helper.ReadRequestBody(requests, &processApplicantRequest)

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	err = controller.EmployerService.ProcessApplicant(requests.Context(), processApplicantRequest, token)

	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
	}
	if err != nil {
		webResp = response.WebResponse{
			Code:   400,
			Status: err.Error(),
		}
	}
	helper.WriteResponseBody(writer, webResp)
}

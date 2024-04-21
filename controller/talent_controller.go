package controller

import (
	"net/http"
	"simplejobportal/data/request"
	"simplejobportal/data/response"
	"simplejobportal/helper"
	"simplejobportal/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TalentController struct {
	TalentService service.TalentService
}

func NewTalentController(employerService service.TalentService) *TalentController {
	return &TalentController{TalentService: employerService}
}

func (controller *TalentController) SeeJob(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	data := controller.TalentService.SeeJob(requests.Context(), token)
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	helper.WriteResponseBody(writer, webResp)
}

func (controller *TalentController) ApplyJob(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	talentCreateRequest := request.ApplyJobRequest{}
	helper.ReadRequestBody(requests, &talentCreateRequest)

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	controller.TalentService.ApplyJob(requests.Context(), talentCreateRequest, token)
	webResp := response.WebResponse{
		Code:   201,
		Status: "OK",
	}

	helper.WriteResponseBody(writer, webResp)
}

func (controller *TalentController) SeeApplicationDetail(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	applicationId := params.ByName("applicationId")
	applicationIdInt, err := strconv.Atoi(applicationId)
	helper.PanicIfError(err)

	token, err := helper.ReadBearerToken(requests)
	helper.PanicIfError(err)

	data := controller.TalentService.SeeApplicationDetail(requests.Context(), applicationIdInt, token)
	webResp := response.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   data,
	}

	helper.WriteResponseBody(writer, webResp)
}

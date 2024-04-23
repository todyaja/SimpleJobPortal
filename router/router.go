package router

import (
	"fmt"
	"net/http"
	"simplejobportal/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController *controller.UserController,
	employerController *controller.EmployerController,
	talentController *controller.TalentController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Welcome Home")
	})

	//User Controller
	router.POST("/api/register", userController.Create)
	router.POST("/api/login", userController.Login)

	//Employer Controller
	router.POST("/api/employer/job", employerController.CreateJob)
	router.PUT("/api/employer/job/:jobId", employerController.UpdateJob)
	router.DELETE("/api/employer/job/:jobId", employerController.DeleteJob)
	router.POST("/api/employer/view-applicant", employerController.SeeJobApplicant)
	router.PUT("/api/employer/process-applicant", employerController.ProcessApplicant)

	//Talent Controller
	router.GET("/api/talent/job", talentController.SeeJob)
	router.POST("/api/talent/apply-job", talentController.ApplyJob)
	router.GET("/api/talent/my-application/:applicationId", talentController.SeeApplicationDetail)
	router.POST("/api/talent/withdraw-application", talentController.WithdrawApplication)

	return router
}

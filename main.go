package main

import (
	"fmt"
	"net/http"
	"simplejobportal/config"
	"simplejobportal/controller"
	"simplejobportal/helper"
	"simplejobportal/repository"
	"simplejobportal/router"
	"simplejobportal/service"
)

func main() {
	fmt.Println("Start Server")

	db := config.DatabaseConnection()

	// User setup
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserController(userService)

	// Employer setup
	employerRepository := repository.NewEmployerRepository(db)
	employerService := service.NewEmployerServiceImpl(employerRepository)
	employerController := controller.NewEmployerController(employerService)

	// Talent setup
	talentRepository := repository.NewTalentRepository(db)
	talentService := service.NewTalentServiceImpl(talentRepository)
	talentController := controller.NewTalentController(talentService)

	routes := router.NewRouter(userController, employerController, talentController)
	server := http.Server{Addr: "localhost:1234", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

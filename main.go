package main

import (
	"Golang-Concurrency/handler"
	"Golang-Concurrency/routes"
	"Golang-Concurrency/service"
)

var (
	carDetailsService    service.CarDetailService     = service.NewCarDetailService()
	CarDetailsController handler.CarDetailsController = handler.NewCarDetailsController(carDetailsService)
	httpRouter           routes.Router                = routes.NewChRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/carDestails", CarDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}

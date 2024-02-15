package handler

import (
	"Golang-Concurrency/service"
	"encoding/json"
	"net/http"
)

type controller struct {
}

var (
	carDetailsService service.CarDetailService
)

type CarDetailsController interface {
	GetCarDetails(res http.ResponseWriter, req *http.Request)
}

func NewCarDetailsController(service service.CarDetailService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetail()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

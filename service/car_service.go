package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

type CarServiceImpl struct {
}

func NewCarService() CarService {
	return &CarServiceImpl{}
}

var (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

func (*CarServiceImpl) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)

	//call internal api
	resp, _ := client.Get(carServiceUrl)
	//write respon to the channel
	carDataChannel <- resp

}

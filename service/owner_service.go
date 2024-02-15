package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	GetData()
}

type OwnerServiceImpl struct {
}

func NewOwnerService() OwnerService {
	return &OwnerServiceImpl{}
}

var (
	ownerServiceUrl = "https://myfakeapi.com/api/user/1"
)

func (*OwnerServiceImpl) GetData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)

	//call internal api
	resp, _ := client.Get(ownerServiceUrl)
	//write respon to the channel
	carDataChannel <- resp

}

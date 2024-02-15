package service

import (
	"Golang-Concurrency/models"
	"encoding/json"
	"net/http"
)

type CarDetailService interface {
	GetDetail() models.CarDetails
}

type CarDetailServiceImpl struct{}

func NewCarDetailService() CarDetailService {
	return &CarDetailServiceImpl{}
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

func (us *CarDetailServiceImpl) GetDetail() models.CarDetails {
	// goroutine call fake api
	go carService.FetchData()
	// goroutine call fake apis
	go ownerService.GetData()
	car, _ := getCarData()
	owner, _ := getOwnerData()

	return models.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (models.Car, error) {
	r1 := <-carDataChannel
	var car models.Car
	if err := json.NewDecoder(r1.Body).Decode(&car); err != nil {
		return car, err
	}
	return car, nil
}
func getOwnerData() (models.Owner, error) {
	r1 := <-ownerDataChannel
	var owner models.Owner
	if err := json.NewDecoder(r1.Body).Decode(&owner); err != nil {
		return owner, err
	}
	return owner, nil
}

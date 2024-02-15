package models

type CarDetails struct {
	ID        int    `json:"id"`
	Brand     string `json:"car"`
	Model     string `json:"car_model"`
	Year      int    `json:"car_model_year"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

package main

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type RestaurantRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

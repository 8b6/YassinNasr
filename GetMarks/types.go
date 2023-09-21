package main

type mark struct {
	RestaurantName string `json:"restaurantName"`
	Position       string `json:"position"`
	Category       string `json:"category"`
}

type markers struct {
	Status       bool              `json:"status"`
	ResponseCode string            `json:"responseCode"`
	marks        map[string]string `json:"marks"`
}

type markersResponse struct {
	Status          bool    `json:"status"`
	ResponseCode    string  `json:"responseCode"`
	ResponseMessage string  `json:"responseMessage"`
	markers         markers `json:"reservationID"`
}

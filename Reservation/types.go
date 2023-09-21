package main

type Reservation struct {
	ReservationID      string `json:"reservationID"`
	CostumerID         string `json:"costumerID"`
	RestaurantID       string `json:"restaurantID"`
	TableType          string `json:"tableType"`
	ReservetionStartAt string `json:"reservetionStartAt"`
	ReservationEndAt   string `json:"reservationEndAt"`
}

type ResponseBodyForResrvation struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReservationID   string `json:"reservationID"`
}

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type UpdateReservation struct {
	ServiceCode  string `json:"serviceCode"`
	ServicePrice string `json:"servicePrice"`
	ServiceType  string `json:"ServiceType"`
	Description  string `json:"description"`
	HomePage     string `json:"homePage"`
}

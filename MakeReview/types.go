package main

type ResponseBody struct {
	Status          bool   `json:"status"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

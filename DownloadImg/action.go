package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func downloadImgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	folder := vars["folder"]
	pic := vars["pic"]
	IMGsFolderPath := "D:/nippon work/API's/DownloadImg/"
	//IMGsFolderPath:="./"
	URL := IMGsFolderPath + folder + "/" + pic

	//reqBody, mapR := validateUpdateUserInformationApi(r)
	// if len(mapR) > 0 {
	// 	var rBody ResponseBody
	// 	rBody.Status = false
	// 	rBody.ResponseCode = "104"
	// 	rBody.ResponseMessage = "please check your parameter"
	// 	responseMessage(w, rBody)
	// } else {
	downloadImgController(w, URL)
	// }

}

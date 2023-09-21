package main

import (
	"io/ioutil"
	"net/http"
)

func downloadImgController(w http.ResponseWriter, URL string) {

	readedPic, err := ioutil.ReadFile(URL)
	if err == nil {

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(readedPic)
		return

	} else {
		println("no pic")
		w.WriteHeader(http.StatusBadRequest)

	}

}

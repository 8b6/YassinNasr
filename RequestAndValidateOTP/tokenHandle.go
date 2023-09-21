package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var Secret string = "jdnfksdmfksd"

func authToken(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accsessDetails, err := ExtractTokenMetadata(r)
		if err != nil {
			writeErrorLog("Eror when getting token meta data: ", err.Error())
			fmt.Println("Eror when getting token meta data: ", err.Error())
			if err.Error() == "Token is expired" {
				var rBody ResponseBody
				rBody.Status = false
				rBody.ResponseCode = "122"
				rBody.ResponseMessage = "The token expired"
				responseMessage(w, rBody)
				return
			}
			var rBody ResponseBody
			rBody.Status = false
			rBody.ResponseCode = "123"
			rBody.ResponseMessage = "The token invalid"
			responseMessage(w, rBody)
			return
		}
		checktokenMsg := checkToken(accsessDetails)
		if checktokenMsg == "invald token" {

			var rBody ResponseBody
			rBody.Status = false
			rBody.ResponseCode = "102"
			rBody.ResponseMessage = "The token invalid "
			responseMessage(w, rBody)
			return
		}

		if checktokenMsg == "expired token" {

			var rBody ResponseBody
			rBody.Status = false
			rBody.ResponseCode = "102"
			rBody.ResponseMessage = "The token expired "
			responseMessage(w, rBody)
			return
		}

		fmt.Printf("ACCESS_Details: %+v \n", accsessDetails)
		fn(w, r)

	}
}

func checkToken(accsessDetails *AccessDetails) string {
	if accsessDetails.UserId == 1 {
		if accsessDetails.Exp > 10 {
			return "checked correct token"
		} else {
			return "expired token"
		}
	} else {
		return "invald token"

	}

}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("HERE Not OK MEthod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Printf("Claims: %v %10f %v %v \n", claims, claims["exp"], ok, token.Valid)
	if ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			Exp:    int64(exp),
			UserId: userId,
		}, nil
	}
	return nil, err
}

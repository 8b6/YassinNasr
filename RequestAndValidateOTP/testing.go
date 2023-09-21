package main

import "fmt"

func testing() {
	//testCheckOTPInvalid()
	//testcheckOTPused()
	testDetialValidateOTP()
}

func testCheckOTPInvalid() {
	conn, err := SQLConnection()
	if err == nil {

		check := checkOTPInvalid(conn, "249117617113", "1", "1111")
		//CheckOTPInvalid must be true
		fmt.Println("CheckOTPInvalid is ", check)

		check1 := checkOTPInvalid(conn, "249117617113", "12", "1111")
		//CheckOTPInvalid must be false
		fmt.Println("CheckOTPInvalid is ", check1)
	}
}

func testcheckOTPused() {
	conn, err := SQLConnection()
	if err == nil {

		check := checkOTPused(conn, "249117617113", "1", "1111")
		// checkOTPused must be true
		fmt.Println(" checkOTPused is ", check)

		check1 := checkOTPused(conn, "249117617113", "12", "1111")
		// checkOTPused must be false
		fmt.Println(" checkOTPused is ", check1)
	}
}

func testCheckTicketExpired() {
	conn, err := SQLConnection()
	if err == nil {

		check := checkOTPexpird(conn, "249117617113", "1", "1111")
		//checkTicketExpired must be true
		fmt.Println("checkTicketExpired is ", check)

		check1 := checkOTPexpird(conn, "249117617113", "12", "6374")
		//checkTicketExpired must be false
		fmt.Println("checkTicketExpired is ", check1)
	}
}

func testDetialValidateOTP() {
	conn, err := SQLConnection()
	if err == nil {

		bol, resp := detialValidateOTP(conn, "249117617113", "1", "1111")
		//detialValidateOTP must be true
		fmt.Println("detialValidateOTP is ", bol)
		fmt.Println(fmt.Sprintf(" response: %+v", resp), bol)

		// check1 := detialValidateOTP(conn, "249117617113", "12", "6374")
		// //detialValidateOTP must be false
		// fmt.Println("detialValidateOTP is ", check1)
	}
}

package models

import (
	"net/http"
	"fmt"
)

var bal = 0
var clickStatus = 0

type Data struct {
	Balance int
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	if clickStatus == 0 {
		bal++
	}
	if clickStatus == 1 {
		bal += 2
	}
	if clickStatus == 2 {
		bal += 4
	}
	if clickStatus == 3 {
		bal += 8
	}
	fmt.Println("Balance:", bal)
	http.Redirect(w, r, "/", 302)
}
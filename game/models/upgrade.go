package models

import (
	"net/http"
	"fmt"
)

func cUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	if clickStatus == 0 && bal >= 10{
		bal = bal - 10
		clickStatus = 1
		fmt.Println("Click Upgraded! :D")
		fmt.Println("Remaining Balance:", bal)
		fmt.Println("Next upgrade is 50 coins!")
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 0 && bal < 10 {
		fmt.Println("Insufficent funds :(")
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 1 && bal >= 50{
		bal = bal - 50
		clickStatus = 2
		fmt.Println("Click Upgraded! :D")
		fmt.Println("Remaining Balance:", bal)
		fmt.Println("Next upgrade is 100 coins!")
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 1 && bal < 50 {
		fmt.Println("Insufficent funds :(")
		http.Redirect(w, r, "/", 302)
		return
	}

	if clickStatus == 2 && bal >= 100{
		bal = bal - 100
		clickStatus = 3
		fmt.Println("Click Upgraded! :D")
		fmt.Println("Remaining Balance:", bal)
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 2 && bal < 100 {
		fmt.Println("Insufficent funds :(")
		http.Redirect(w, r, "/", 302)
		return
	}
	fmt.Println("No Available Upgrades :(")
	http.Redirect(w, r, "/", 302)
}
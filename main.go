package main

import (
	"final-project/database"
	"final-project/router"
	"net/http"
)

func main() {
	database.StartDB()
	rtr := router.StartApp()

	http.ListenAndServe(":8080", rtr)
}

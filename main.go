package main

import (
	"net/http"
	"user-rest-app/router"
)

func main() {
	// Create a new instance of the logger
	// 
	r:= router.SetupRootRouter()

	port := ":5010"
	http.ListenAndServe(port, r)

}

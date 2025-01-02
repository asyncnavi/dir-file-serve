package main

import (
	"net/http"
)


var ROOT_FOLDER string

func init() {
	ROOT_FOLDER = "./routes/"
}

func main() {

	http.ListenAndServe(":3000", nil)
}

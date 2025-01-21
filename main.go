package main

import (
	"net/http"
)


var (
	projectDir string
)

func init() {
	projectDir = "./routes/"
}

func main() {

	http.ListenAndServe(":3000", nil)
}

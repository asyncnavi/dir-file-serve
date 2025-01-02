package user

import "net/http"


func Hanlder(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Hello this world"))
}

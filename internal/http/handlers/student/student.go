package student

import "net/http"

func Create(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to students Api"))
}
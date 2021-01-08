package transaction

import "net/http"

func Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))

}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))

}
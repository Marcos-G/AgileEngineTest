package account

import "net/http"

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))

}
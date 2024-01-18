package emulator

import (
	"net/http"
)

func (api objectAPIHandlers) HoleWorld(w http.ResponseWriter, r *http.Request) {
	// objectAPI := api.ObjectAPI()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}


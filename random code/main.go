package random

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8000"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

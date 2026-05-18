package http_server

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHTTPServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Обработка запроса /ping")
		w.Write([]byte("Привет docker\n"))
	})

	err := http.ListenAndServe(":5050", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

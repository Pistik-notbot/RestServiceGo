package main

import (
	"fmt"

	"go.mod/http_server"
)

func main() {
	if err := http_server.StartHTTPServer(); err != nil {
		fmt.Printf("Ошибка при запуске HTTP сервера: %v\n", err)
	}
	fmt.Println("HTTP сервер успешно запущен на порту 5050")

}

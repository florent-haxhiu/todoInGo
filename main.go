package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"florent-haxhiu/todoInGo/internal/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Server now up and running")
	http.ListenAndServe(":3000", router.Router())
}

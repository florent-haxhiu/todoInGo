package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"florent-haxhiu/todoInGo/internal/logger"
	"florent-haxhiu/todoInGo/internal/router"
)

func main() {
	logger.InitializeDefault()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server := &http.Server{
		Addr:    ":3000",
		Handler: router.Router(),
	}

	logger.InfoMsg("Server starting", "addr", ":3000")
	err = server.ListenAndServe()

	if err != nil {
		logger.ErrorMsg("Server failed to start", "error", err)
	}
}

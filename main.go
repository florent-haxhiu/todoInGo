package main

import (
	"net/http"

	"florent-haxhiu/todoInGo/internal/router"
)

func main() {
	http.ListenAndServe(":3000", router.Router())
}

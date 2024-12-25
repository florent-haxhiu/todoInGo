package main

import (
	"net/http"

	"florent-haxhiu/todo-go/internal/router"
)

func main() {
	http.ListenAndServe(":3000", router.Router())
}

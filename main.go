package main

import (
	"example.com/test_task/dependencies"
	"example.com/test_task/http"
)

func main() {
	dependencies.Bind()
	http.Start()
}

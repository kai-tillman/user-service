package main

import (
	api "github.com/kai-tillman/user-service/internal/user-service"
)

func main() {

	api.NewAPIService().Start()
}

package main

import (
	"tt/router"
)

func main() {
	router.Echo.Logger.Fatal(router.Echo.Start(":1323"))
}

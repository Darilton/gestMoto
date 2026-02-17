package main

import (
	"gestmoto/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()

	r.Run(":8080")
	fmt.Println("Hello, World!")
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	ports := [3]string{"8080", "8081", "8082"}

	serverFarm.startServerFarm(ports)
}

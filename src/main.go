package main

import (
	"fmt"
	"github.com/rohitamar/load-balancer/serverFarm"
)

func main() {
	ports := [3]string{"8080", "8081", "8082"}
	serverFarm.startServerFarm(ports)
}

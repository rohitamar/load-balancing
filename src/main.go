package main

import (
	"github.com/rohitamar/load-balancing/serverFarm"
)

func main() {
	ports := [3]string{"8080", "8081", "8082"}
	serverFarm.startServerFarm(ports)
}

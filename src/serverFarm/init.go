package serverFarm

import (
	"fmt"
	"net/http"
	"sync"
)

func startServer(port string, waitGroup *sync.WaitGroup, errCh chan<- error) {
	defer waitGroup.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Curretly in ", port)
	})

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		errCh <- err
	}
}

func startServerFarm(ports []string) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(ports))

	errCh := make(chan error, len(ports))
	for _, port := range ports {
		go startServer(port, &waitGroup, errCh)
	}

	waitGroup.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"gchat/client/ws_client"
	"sync"
)

func main() {
	maxUsers := flag.Int("maxUsers", 1, "Number of users to simulate")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(*maxUsers)

	// Create a channel to control the order of execution
	orderChan := make(chan struct{}, 1)

	// Seed the channel to allow the first goroutine to proceed
	orderChan <- struct{}{}

	for i := 1; i <= *maxUsers; i++ {
		go func(userId int64) {
			defer wg.Done()

			// Wait for permission to execute
			<-orderChan

			client := ws_client.WSClient{
				UserId:   userId,
				DeviceId: userId,
				Seq:      0,
			}
			fmt.Printf("UserId: %d, DeviceId: %d, Seq: %d\n", client.UserId, client.DeviceId, client.Seq)
			client.Login()
			client.Start()

			// Allow the next goroutine to proceed
			orderChan <- struct{}{}
			select {}
		}(int64(i))
	}

	wg.Wait()
	fmt.Println("All goroutines have executed in order.")
}

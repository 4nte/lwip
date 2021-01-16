package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"os/user"
	"github.com/go-redis/redis/v8"
	"time"
)
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	// Get private address
	var privateAddress string
	for _, addr := range addrs {
		// Check if address starts with 10.20
		if !strings.HasPrefix(addr.String(), "10.20") {
			continue
		}

		privateAddress = strings.TrimSuffix(addr.String(),"/24")
		break
	}

	// Get user name
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	userName := strings.TrimPrefix(user.HomeDir, `C:\Users\`)



	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			fmt.Println("ping")
			rdb.Set(context.Background(), "user/"+userName, privateAddress, 10 * time.Second)
		}
	}
}
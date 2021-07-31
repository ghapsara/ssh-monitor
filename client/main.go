package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	ssh "sshps"
	"time"
)

func main() {
	s := ssh.New()
	serverAddr := os.Getenv("SERVER_ADDR")
	if len(serverAddr) == 0 {
		serverAddr = "http://host.docker.internal:9999"
	}

	for {
		session := s.GetTotalSession()

		s, err := json.Marshal(session)
		if err != nil {
			fmt.Println("client %w", err)
		}

		fmt.Println(string(s))

		fmt.Println(serverAddr)

		resp, err := http.Post(serverAddr+"/save", "application/json", bytes.NewBuffer(s))
		if err != nil {
			fmt.Println("client", err)
		}
		fmt.Println(resp)

		time.Sleep(1 * time.Second)
	}
}

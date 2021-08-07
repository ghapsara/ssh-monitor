package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	ssh "sshps"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverAddr := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")

	s := ssh.New()

	for {
		session := s.GetTotalSession()

		s, err := json.Marshal(session)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(serverAddr+"/save", "application/json", bytes.NewBuffer(s))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)

		time.Sleep(1 * time.Second)
	}
}

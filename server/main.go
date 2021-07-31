package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	ssh "sshps"

	"sshps/storage"
)

type Server struct {
	Storage storage.Storage
}

func main() {
	storage := storage.NewLocalStorage()

	server := New(&storage)

	httpserver := http.Server{
		Addr: os.Getenv("SERVER_ADDR"),
	}

	http.HandleFunc("/save", server.save)
	http.HandleFunc("/view", server.view)

	log.Fatal(httpserver.ListenAndServe())
}

func New(storage storage.Storage) Server {
	return Server{
		Storage: storage,
	}
}

func (s *Server) save(rw http.ResponseWriter, r *http.Request) {
	// call save to database
	var session ssh.Session
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &session)
	if err != nil {
		fmt.Println(err)
	}

	err = s.Storage.Save(session)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("session saved", err)
}

func (s *Server) view(rw http.ResponseWriter, r *http.Request) {
	// view ssh sessions
	// read aggregated ssh session data
	// output
	session, err := s.Storage.Read()

	if err != nil {
		fmt.Fprintf(rw, "err")
	}

	resp := ""
	for _, s := range session {
		resp += fmt.Sprintf("%s had %d attempts\n", s.Hostname, s.Total)
	}

	fmt.Fprintf(rw, resp)
}

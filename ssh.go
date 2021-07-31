package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Session struct {
	Hostname string `json:"hostname"`
	Total    int    `json:"session"`
}

var sshLogRegex = regexp.MustCompile(`sshd\[\d+\]`)

func New() Session {
	// get hostname
	hostname, _ := os.Hostname()

	return Session{
		Hostname: hostname,
	}
}

// GetTotalSession returns last 2 minutes ssh logs
func (s *Session) GetTotalSession() Session {
	command := "journalctl -t sshd"
	// command := "cat sshlog"
	c, b := exec.Command("/bin/bash", "-c", command), new(strings.Builder)
	c.Stdout = b
	// err := c.Run()
	c.Run()
	// if err != nil {
	// 	return nil
	// }

	sessionID := sshLogRegex.FindAllString(b.String(), -1)
	fmt.Println(b.String())

	s.Total = len(sessionID)

	return *s
}

// func

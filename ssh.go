package ssh

import (
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

func (s *Session) GetTotalSession() (Session, error) {
	command := "journalctl -t sshd"
	c, b := exec.Command("/bin/bash", "-c", command), new(strings.Builder)
	c.Stdout = b
	err := c.Run()

	sessions := sshLogRegex.FindAllString(b.String(), -1)
	s.Total = calculateSessions(sessions)

	return *s, err
}

func calculateSessions(s []string) int {
	sessionDict := make(map[string]bool, len(s))
	for _, v := range s {
		sessionDict[v] = true
	}

	return len(sessionDict)
}

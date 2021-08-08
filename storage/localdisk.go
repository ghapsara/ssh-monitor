package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	ssh "sshps"
)

type LocalStorage struct {
	Filename string
}

func NewLocalStorage() LocalStorage {
	var storageFilename = "data.json"

	f, err := os.Create(storageFilename)
	if err != nil {
		log.Fatal("init storage error", err)
	}

	defer f.Close()

	return LocalStorage{
		Filename: storageFilename,
	}
}

func (l *LocalStorage) Read() ([]ssh.Session, error) {
	var sessions []ssh.Session

	data, err := ioutil.ReadFile(l.Filename)
	if err != nil {
		return sessions, err
	}

	err = json.Unmarshal(data, &sessions)

	return sessions, nil
}

func (l *LocalStorage) Save(s ssh.Session) error {
	sessions, err := l.Read()
	if err != nil {
		return err
	}

	exist := false
	for i, v := range sessions {
		if v.Hostname == s.Hostname {
			sessions[i].Total = s.Total
			exist = true
		}
	}

	if !exist && len(s.Hostname) > 0 {
		sessions = append(sessions, s)
	}

	data, err := json.Marshal(sessions)
	if err != nil {
		return err
	}

	err = os.WriteFile(l.Filename, data, 0644)
	return err
}

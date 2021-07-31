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

func (l *LocalStorage) Save(s ssh.Session) error {
	data, err := ioutil.ReadFile(l.Filename)
	if err != nil {
		return err
	}

	var sessions []ssh.Session

	err = json.Unmarshal(data, &sessions)

	exist := false
	for _, v := range sessions {
		if v.Hostname == s.Hostname {
			v.Total = s.Total
			exist = true
		}
	}

	if !exist && len(s.Hostname) > 0 {
		sessions = append(sessions, s)
	}

	data, err = json.Marshal(sessions)
	if err != nil {
		return err
	}

	err = os.WriteFile(l.Filename, data, 0644)
	return err
}

func (l *LocalStorage) Read() ([]ssh.Session, error) {
	return []ssh.Session{}, nil
}

package storage

import ssh "sshps"

type Storage interface {
	Save(ssh.Session) error
	Read() ([]ssh.Session, error)
}

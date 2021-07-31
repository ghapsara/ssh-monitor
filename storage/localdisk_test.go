package storage

import (
	ssh "sshps"
	"testing"
)

func TestLocalStorage_Save(t *testing.T) {
	type args struct {
		s ssh.Session
	}
	tests := []struct {
		name    string
		l       *LocalStorage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			l: &LocalStorage{},
			args: args{
				s: ssh.Session{
					Hostname: "node-3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalStorage{}
			if err := l.Save(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("LocalStorage.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

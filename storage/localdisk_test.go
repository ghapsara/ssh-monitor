package storage

import (
	"reflect"
	ssh "sshps"
	"testing"
)

func TestLocalStorage_Save(t *testing.T) {
	type args struct {
		s ssh.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				s: ssh.Session{
					Hostname: "node-3",
					Total:    3,
				},
			},
		},
		{
			args: args{
				s: ssh.Session{
					Hostname: "node-3",
					Total:    4,
				},
			},
		},
	}
	l := NewLocalStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := l.Save(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("LocalStorage.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalStorage_Read(t *testing.T) {
	type fields struct {
		Filename string
	}
	tests := []struct {
		name     string
		fields   fields
		sessions []ssh.Session
		want     []ssh.Session
		wantErr  bool
	}{
		{
			sessions: []ssh.Session{
				{
					Hostname: "host-1",
					Total:    2,
				},
				{
					Hostname: "host-1",
					Total:    5,
				},
			},
			want: []ssh.Session{
				{
					Hostname: "host-1",
					Total:    5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLocalStorage()
			for _, v := range tt.sessions {
				l.Save(v)
			}
			got, err := l.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalStorage.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalStorage.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

package ssh

import "testing"

func Test_calculateSessions(t *testing.T) {
	tests := []struct {
		name     string
		sessions []string
		want     int
	}{
		{
			sessions: []string{"sshd[52]", "sshd[52]", "sshd[640]", "sshd[640]", "sshd[640]", "sshd[640]", "sshd[673]"},
			want:     3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSessions(tt.sessions); got != tt.want {
				t.Errorf("calculateSessions() = %v, want %v", got, tt.want)
			}
		})
	}
}

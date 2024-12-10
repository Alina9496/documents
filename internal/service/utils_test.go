package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkLogin(t *testing.T) {
	tests := []struct {
		name  string
		login string
		want  bool
	}{
		{
			name:  "lenght < 8",
			login: "login",
			want:  false,
		},
		{
			name:  "there are no number",
			login: "loginlogin",
			want:  false,
		},
		{
			name:  "there are no letters",
			login: "123456789",
			want:  false,
		},
		{
			name:  "login corrected",
			login: "login345",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkLogin(tt.login)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_checkPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "lenght < 8",
			password: "pass",
			want:     false,
		},
		{
			name:     "there are no number",
			password: "password",
			want:     false,
		},
		{
			name:     "there are no lower letter",
			password: "PASS1234",
			want:     false,
		},
		{
			name:     "there are no upper letter",
			password: "pass1234",
			want:     false,
		},
		{
			name:     "there are no symbol",
			password: "Passw345",
			want:     false,
		},
		{
			name:     "password corrected",
			password: "Passw_345",
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkPassword(tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}

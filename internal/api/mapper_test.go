package api

import (
	"testing"

	"github.com/Alina9496/documents/internal/domain"
	v1 "github.com/Alina9496/documents/pkg/api/v1"
	"github.com/stretchr/testify/assert"
)

func Test_toDomainUser(t *testing.T) {
	tests := []struct {
		name string
		req  v1.User
		want *domain.User
	}{
		{
			name: "convert to domain.User",
			req: v1.User{
				Login:    "Login",
				Password: "Password",
			},
			want: &domain.User{
				Login:    "Login",
				Password: "Password",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toDomainUser(tt.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toLoginResp(t *testing.T) {
	tests := []struct {
		name  string
		login string
		want  v1.RespLogin
	}{
		{
			name:  "convert to v1.RespLogin",
			login: "login",
			want: v1.RespLogin{
				Login: "login",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toLoginResp(tt.login)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toTokenResp(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  v1.RespToken
	}{
		{
			name:  "convert to v1.RespToken",
			token: "CxBiwVruDAD8kp8jgeOY",
			want: v1.RespToken{
				Token: "CxBiwVruDAD8kp8jgeOY",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toTokenResp(tt.token)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toLogOutTokenResp(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  map[string]bool
	}{
		{
			name:  "convert to map",
			token: "CxBiwVruDAD8kp8jgeOY",
			want:  map[string]bool{"CxBiwVruDAD8kp8jgeOY": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toLogOutTokenResp(tt.token)
			assert.Equal(t, tt.want, got)
		})
	}
}

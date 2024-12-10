package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/repo"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/Alina9496/tool/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	repo    *MockRepository
	cache   *MockCache
	service *Service
}

func (s *ServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.repo = NewMockRepository(ctrl)
	s.cache = NewMockCache(ctrl)
	s.service = New(s.repo, s.cache, logger.New(""))
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) Test_isUserExists() {
	ctx := context.Background()
	user := &domain.User{
		Login:    "login345",
		Password: "Passw_345",
	}
	id := uuid.New()
	tests := []struct {
		name  string
		ctx   context.Context
		user  *domain.User
		want  bool
		calls func()
	}{
		{
			name: "user not found",
			ctx:  ctx,
			user: user,
			want: false,
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(uuid.Nil, errors.ErrUnsupported)
			},
		},
		{
			name: "user found",
			ctx:  ctx,
			user: user,
			want: true,
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(id, nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got := s.service.isUserExists(tt.ctx, tt.user)
			s.Equal(tt.want, got)
		})
	}
}

func (s *ServiceSuite) Test_Registration() {
	ctx := context.Background()
	user := &domain.User{
		Login:    "login345",
		Password: "Passw_345",
	}
	id := uuid.New()
	tests := []struct {
		name  string
		ctx   context.Context
		user  *domain.User
		want  string
		err   error
		calls func()
	}{
		{
			name:  "user equal nil",
			ctx:   ctx,
			user:  nil,
			want:  "",
			err:   ErrUserIsNil,
			calls: func() {},
		},
		{
			name: "login incorrect",
			ctx:  ctx,
			user: &domain.User{
				Login:    "login",
				Password: "Passw_345",
			},
			want:  "",
			err:   ErrUserLoginIncorected,
			calls: func() {},
		},
		{
			name: "password incorrect",
			ctx:  ctx,
			user: &domain.User{
				Login:    "login345",
				Password: "passw345",
			},
			want:  "",
			err:   ErrUserPasswordIncorected,
			calls: func() {},
		},
		{
			name: "error user alredy exist",
			ctx:  ctx,
			user: user,
			want: "",
			err:  fmt.Errorf("error when check user: %w", ErrUserExists),
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(id, nil)
			},
		},
		{
			name: "error registration",
			ctx:  ctx,
			user: user,
			want: "",
			err:  fmt.Errorf("error when registration user: %w", ErrRegistrationUser),
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(uuid.Nil, nil)
				s.repo.EXPECT().Registration(ctx, user).Return(errors.ErrUnsupported)
			},
		},
		{
			name: "user was successfully registration",
			ctx:  ctx,
			user: user,
			want: user.Login,
			err:  nil,
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(uuid.Nil, nil)
				s.repo.EXPECT().Registration(ctx, user).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got, err := s.service.Registration(tt.ctx, tt.user)
			s.Equal(tt.want, got)
			s.Equal(tt.err, err)
		})
	}
}

func (s *ServiceSuite) Test_Authentication() {
	ctx := context.Background()
	user := &domain.User{
		Login:    "login345",
		Password: "Passw_345",
	}
	id := uuid.New()
	tests := []struct {
		name  string
		ctx   context.Context
		user  *domain.User
		want  string
		err   error
		calls func()
	}{
		{
			name:  "user equal nil",
			ctx:   ctx,
			user:  nil,
			err:   ErrUserIsNil,
			calls: func() {},
		},
		{
			name: "login incorrect",
			ctx:  ctx,
			user: &domain.User{
				Login:    "login",
				Password: "Passw_345",
			},
			err:   ErrUserLoginIncorected,
			calls: func() {},
		},
		{
			name: "password incorrect",
			ctx:  ctx,
			user: &domain.User{
				Login:    "login345",
				Password: "passw345",
			},
			err:   ErrUserPasswordIncorected,
			calls: func() {},
		},
		{
			name: "error user not found",
			ctx:  ctx,
			user: user,
			err:  fmt.Errorf("error when check user: %w", ErrUserNotFound),
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(uuid.Nil, nil)
			},
		},
		{
			name: "error authentication",
			ctx:  ctx,
			user: user,
			err:  fmt.Errorf("error when authentication user: %w", ErrAuthenticationUser),
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(id, nil)
				s.repo.EXPECT().Authentication(ctx, user).Return(errors.ErrUnsupported)
			},
		},
		{
			name: "user was successfully authentication",
			ctx:  ctx,
			user: user,
			err:  nil,
			calls: func() {
				s.repo.EXPECT().CheckUser(ctx, user).Return(id, nil)
				s.repo.EXPECT().Authentication(ctx, user).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			_, err := s.service.Authentication(tt.ctx, tt.user)
			s.Equal(tt.err, err)
		})
	}
}

func (s *ServiceSuite) Test_LogOut() {
	ctx := context.Background()
	token := "CxBiwVruDAD8kp8jgeOY"
	tests := []struct {
		name    string
		ctx     context.Context
		token   string
		wantErr error
		calls   func()
	}{
		{
			name:    "token not found",
			ctx:     ctx,
			token:   token,
			wantErr: ErrTokenNotFound,
			calls: func() {
				s.repo.EXPECT().LogOut(ctx, token).Return(repo.ErrTokenNotFound)
			},
		},
		{
			name:    "error logout user",
			ctx:     ctx,
			token:   token,
			wantErr: fmt.Errorf("error when logout user: %w", ErrLogOutUser),
			calls: func() {
				s.repo.EXPECT().LogOut(ctx, token).Return(errors.ErrUnsupported)
			},
		},
		{
			name:    "logout user",
			ctx:     ctx,
			token:   token,
			wantErr: nil,
			calls: func() {
				s.repo.EXPECT().LogOut(ctx, token).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			err := s.service.LogOut(tt.ctx, tt.token)
			s.Equal(tt.wantErr, err)
		})
	}
}

func (s *ServiceSuite) Test_Upload() {
	userID := uuid.New()
	documentID := uuid.New()
	ctx := context.Background()
	tests := []struct {
		name     string
		ctx      context.Context
		document *dto.Document
		want     string
		err      error
		calls    func()
	}{
		{
			name: "success",
			ctx:  ctx,
			document: &dto.Document{
				Name:    "name",
				Token:   "token",
				Mime:    "image/jpeg",
				Content: []byte(""),
				Grant:   []string{"login"},
				Public:  true,
			},
			want: "name",
			err:  nil,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(userID, nil)
				s.cache.EXPECT().Set(gomock.Any(), userID, gomock.Any())
				s.repo.EXPECT().ExecTx(ctx, gomock.Any()).DoAndReturn(
					func(ctx context.Context, fn func(ctx context.Context) error) error {
						return fn(ctx)
					},
				)
				s.repo.EXPECT().Save(ctx, gomock.Any()).Return(documentID, nil)
				s.repo.EXPECT().AddGrant(ctx, gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got, err := s.service.Upload(tt.ctx, tt.document)
			s.Equal(tt.want, got)
			s.Equal(tt.err, err)
		})
	}
}

func (s *ServiceSuite) Test_GetDocument() {
	userID := uuid.New()
	ownerUserID := uuid.New()
	documentID := uuid.New()
	ctx := context.Background()
	now := time.Now()
	tests := []struct {
		name       string
		ctx        context.Context
		documentID uuid.UUID
		token      string
		want       *domain.Document
		err        error
		calls      func()
	}{
		{
			name:       "success",
			ctx:        ctx,
			documentID: documentID,
			token:      "token",
			want: &domain.Document{
				ID:        documentID,
				UserID:    ownerUserID,
				Name:      "name",
				Mime:      "image/jpeg",
				Content:   "",
				Grant:     []string{"login"},
				CreatedAt: now,
				Public:    false,
			},
			err: nil,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetDocument(ctx, documentID).Return(&domain.Document{
					ID:        documentID,
					UserID:    ownerUserID,
					Name:      "name",
					Mime:      "image/jpeg",
					Content:   "",
					Grant:     []string{"login"},
					CreatedAt: now,
					Public:    false,
				}, nil)
				s.cache.EXPECT().Set(gomock.Any(), &domain.Document{
					ID:        documentID,
					UserID:    ownerUserID,
					Name:      "name",
					Mime:      "image/jpeg",
					Content:   "",
					Grant:     []string{"login"},
					CreatedAt: now,
					Public:    false,
				}, gomock.Any())
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(userID, nil)
				s.cache.EXPECT().Set(gomock.Any(), userID, gomock.Any())
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUser(ctx, userID).Return(&domain.User{Login: "login"}, nil)
				s.cache.EXPECT().Set(gomock.Any(), &domain.User{Login: "login"}, gomock.Any())
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().CheckGrant(ctx, documentID, "login").Return(true, nil)
				s.cache.EXPECT().Set(gomock.Any(), true, gomock.Any())
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got, err := s.service.GetDocument(tt.ctx, tt.documentID, tt.token)
			s.Equal(tt.want, got)
			s.Equal(tt.err, err)
		})
	}
}

func (s *ServiceSuite) Test_GetDocuments() {
	ownerUserID := uuid.New()
	documentID1, documentID2, documentID3 := uuid.New(), uuid.New(), uuid.New()
	ctx := context.Background()
	now := time.Now()

	tests := []struct {
		name   string
		ctx    context.Context
		filter *dto.GetDocumentsRequest
		want   []domain.Document
		err    error
		calls  func()
	}{
		{
			name: "success",
			ctx:  ctx,
			filter: &dto.GetDocumentsRequest{
				Token: "token",
				Login: "login",
				Key:   "mime",
				Value: "image/jpeg",
				Limit: 3,
			},
			want: []domain.Document{
				{
					ID:        documentID3,
					UserID:    ownerUserID,
					Name:      "1",
					Mime:      "image/jpeg",
					Content:   "",
					Grant:     []string{"login"},
					CreatedAt: now.AddDate(0, 0, -1),
					Public:    false,
				},
				{
					ID:        documentID1,
					UserID:    ownerUserID,
					Name:      "1",
					Mime:      "image/jpeg",
					Content:   "",
					Grant:     []string{"login"},
					CreatedAt: now,
					Public:    false,
				},
				{
					ID:        documentID2,
					UserID:    ownerUserID,
					Name:      "2",
					Mime:      "image/jpeg",
					Content:   "",
					Grant:     []string{"login"},
					CreatedAt: now,
					Public:    false,
				},
			},
			err: nil,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(ownerUserID, nil)
				s.cache.EXPECT().Set(gomock.Any(), ownerUserID, gomock.Any())
				s.repo.EXPECT().GetDocuments(ctx, gomock.Any()).Return([]domain.Document{
					{
						ID:        documentID1,
						UserID:    ownerUserID,
						Name:      "1",
						Mime:      "image/jpeg",
						Content:   "",
						Grant:     []string{"login"},
						CreatedAt: now,
						Public:    false,
					},
					{
						ID:        documentID2,
						UserID:    ownerUserID,
						Name:      "2",
						Mime:      "image/jpeg",
						Content:   "",
						Grant:     []string{"login"},
						CreatedAt: now,
						Public:    false,
					},
					{
						ID:        documentID3,
						UserID:    ownerUserID,
						Name:      "1",
						Mime:      "image/jpeg",
						Content:   "",
						Grant:     []string{"login"},
						CreatedAt: now.AddDate(0, 0, -1),
						Public:    false,
					},
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got, err := s.service.GetDocuments(tt.ctx, tt.filter)
			s.Equal(tt.want, got)
			s.Equal(tt.err, err)
		})
	}
}

func (s *ServiceSuite) Test_DeleteDocument() {
	ctx := context.Background()
	id := uuid.New()
	userID := uuid.New()
	tests := []struct {
		name  string
		ctx   context.Context
		id    uuid.UUID
		token string
		want  uuid.UUID
		err   error
		calls func()
	}{
		{
			name:  "success",
			ctx:   ctx,
			id:    id,
			token: "token",
			want:  id,
			err:   nil,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(userID, nil)
				s.cache.EXPECT().Set(gomock.Any(), userID, gomock.Any())
				s.repo.EXPECT().DeleteDocument(ctx, id, userID).Return(id, nil)
			},
		},
		{
			name:  "user not found",
			ctx:   ctx,
			id:    id,
			token: "token",
			want:  uuid.Nil,
			err:   ErrUserNotFound,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(uuid.Nil, ErrUserNotFound)
			},
		},
		{
			name:  "document not found",
			ctx:   ctx,
			id:    id,
			token: "token",
			want:  uuid.Nil,
			err:   ErrDocumentNotFound,
			calls: func() {
				s.cache.EXPECT().Get(gomock.Any()).Return(nil, false)
				s.repo.EXPECT().GetUserID(ctx, "token").Return(userID, nil)
				s.cache.EXPECT().Set(gomock.Any(), userID, gomock.Any())
				s.repo.EXPECT().DeleteDocument(ctx, id, userID).Return(uuid.Nil, ErrDocumentNotFound)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.calls()
			got, err := s.service.DeleteDocument(tt.ctx, tt.id, tt.token)
			s.Equal(tt.want, got)
			s.Equal(tt.err, err)
		})
	}
}

//go:generate mockgen -source=interfaces.go -destination=./mock_interfaces.go -package=service
package service

import (
	"context"
	"time"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/google/uuid"
)

type Repository interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
	Registration(ctx context.Context, user *domain.User) error
	CheckUser(ctx context.Context, user *domain.User) (uuid.UUID, error)
	Authentication(ctx context.Context, user *domain.User) error
	GetUserID(ctx context.Context, token string) (uuid.UUID, error)
	LogOut(ctx context.Context, token string) error
	Save(ctx context.Context, document *domain.Document) (uuid.UUID, error)
	AddGrant(ctx context.Context, grant *domain.Grant) error
	GetDocument(ctx context.Context, id uuid.UUID) (*domain.Document, error)
	CheckGrant(ctx context.Context, documentID uuid.UUID, login string) (bool, error)
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetDocuments(ctx context.Context, filter *dto.GetDocuments) ([]domain.Document, error)
	DeleteDocument(ctx context.Context, id, userID uuid.UUID) (uuid.UUID, error)
}

type Cache interface {
	Set(k string, x any, d time.Duration)
	Get(k string) (any, bool)
}

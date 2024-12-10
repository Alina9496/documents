package api

import (
	"context"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/google/uuid"
)

type Service interface {
	Registration(ctx context.Context, user *domain.User) (string, error)
	Authentication(ctx context.Context, user *domain.User) (string, error)
	LogOut(ctx context.Context, token string) error
	Upload(ctx context.Context, document *dto.Document) (name string, err error)
	GetDocument(ctx context.Context, id uuid.UUID, token string) (*domain.Document, error)
	GetDocuments(ctx context.Context, filter *dto.GetDocumentsRequest) ([]domain.Document, error)
	DeleteDocument(ctx context.Context, id uuid.UUID, token string) (uuid.UUID, error)
}

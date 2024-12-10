package service

import (
	"encoding/base64"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/google/uuid"
)

func toDocument(userID uuid.UUID, document *dto.Document) *domain.Document {
	if document == nil {
		return nil
	}

	return &domain.Document{
		Name:    document.Name,
		UserID:  userID,
		Mime:    document.Mime,
		Content: base64.StdEncoding.EncodeToString(document.Content),
		Grant:   document.Grant,
		Public:  document.Public,
	}
}

func toGrant(login string, userID, documentID uuid.UUID) *domain.Grant {
	return &domain.Grant{
		UserID:         userID,
		DocumentID:     documentID,
		GrantUserLogin: login,
	}
}

func toGetDocuments(userID uuid.UUID, filter *dto.GetDocumentsRequest) *dto.GetDocuments {
	return &dto.GetDocuments{
		UserID: userID,
		Login:  filter.Login,
		Key:    filter.Key,
		Value:  filter.Value,
		Limit:  filter.Limit,
	}
}

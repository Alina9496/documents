package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Alina9496/documents/internal/domain"
	service "github.com/Alina9496/documents/internal/service"
	"github.com/Alina9496/documents/internal/service/dto"
	v1 "github.com/Alina9496/documents/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

func toDomainUser(req v1.User) *domain.User {
	return &domain.User{
		Login:    req.Login,
		Password: req.Password,
	}
}

func toLoginResp(login string) v1.RespLogin {
	return v1.RespLogin{
		Login: login,
	}
}

func toTokenResp(token string) v1.RespToken {
	return v1.RespToken{
		Token: token,
	}
}

func toFile(c *gin.Context) (*dto.Document, error) {
	metaData := c.Request.FormValue("meta")
	var req v1.Meta
	err := json.Unmarshal([]byte(metaData), &req)
	if err != nil || !req.IsValid() {
		return nil, errInvalidMetaData
	}

	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &dto.Document{
		Name:    req.Name,
		Token:   req.Token,
		Mime:    req.Mime,
		Content: body,
		Grant:   req.Grant,
		Public:  req.Public,
	}, nil
}

func toUploadResponse(name string) v1.UploadResponse {
	return v1.UploadResponse{
		Data: v1.Data{
			File: name,
		},
	}
}
func toGetDocumentsRequest(c *gin.Context) (*dto.GetDocumentsRequest, error) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return nil, errInvalidLimit
	}
	req := &dto.GetDocumentsRequest{
		Token: getUserTokenFromContext(c),
		Login: c.Query("login"),
		Key:   c.Query("key"),
		Value: c.Query("value"),
		Limit: limit,
	}

	return req, req.IsValid()
}

func toGetDocumentsResp(documents []domain.Document) v1.GetDocumentsResp {
	var resp v1.GetDocumentsResp
	resp.DataDocuments.Docs = make([]v1.Document, 0, len(documents))
	for _, doc := range documents {
		resp.DataDocuments.Docs = append(resp.DataDocuments.Docs, v1.Document{
			ID:      doc.ID.String(),
			Name:    doc.Name,
			Mime:    doc.Mime,
			File:    true,
			Public:  doc.Public,
			Created: doc.CreatedAt.Format(time.DateTime),
			Grant:   doc.Grant,
		})
	}
	return resp
}

func toLogOutTokenResp(token string) map[string]bool {
	return map[string]bool{token: true}
}

func errToHttpStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	case errors.Is(err, errAdminUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, service.ErrUserLoginIncorected),
		errors.Is(err, service.ErrUserPasswordIncorected),
		errors.Is(err, service.ErrUserIsNil):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrUserNotFound),
		errors.Is(err, service.ErrDocumentNotFound),
		errors.Is(err, service.ErrDocumentsNotFound),
		errors.Is(err, service.ErrTokenNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrNoAccess):
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

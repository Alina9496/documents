package dto

import "github.com/google/uuid"

type Document struct {
	Name    string
	Token   string
	Mime    string
	Content []byte
	Grant   []string
	Public  bool
}

type GetDocumentsRequest struct {
	Token string
	Login string
	Key   string
	Value string
	Limit int
}

type GetDocuments struct {
	UserID uuid.UUID
	Login  string
	Key    string
	Value  string
	Limit  int
}

func (g *GetDocumentsRequest) IsValid() error {
	if g.Limit < 1 {
		return ErrInvalidLimit
	}

	if g.Key != "name" && g.Key != "mime" {
		return ErrInvalidKey
	}

	if g.Value == "" {
		return ErrEmptyValue
	}

	return nil
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Login    string
	Password string
	Token    string
}

type Document struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Mime      string
	Content   string
	Grant     []string
	CreatedAt time.Time
	Public    bool
}

type Grant struct {
	UserID         uuid.UUID
	DocumentID     uuid.UUID
	GrantUserLogin string
	CreatedAt      time.Time
}

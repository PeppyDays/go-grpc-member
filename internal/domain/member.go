package domain

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID             uuid.UUID
	Email          Email
	Password       Password
	CreatedAt      time.Time
	LastSignedInAt time.Time
}

func NewMember(email Email, password Password) *Member {
	return &Member{
		Email:    email,
		Password: password,
	}
}

func (m *Member) Register() {
	m.ID, _ = uuid.NewUUID()
	m.CreatedAt = time.Now().UTC()
}

func (m *Member) SignIn() {
	m.LastSignedInAt = time.Now().UTC()
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID             string
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
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().UTC()
}

func (m *Member) SignIn() {
	m.LastSignedInAt = time.Now().UTC()
}

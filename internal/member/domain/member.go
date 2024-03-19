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

func NewMember(id string, email Email, password Password, createdAt time.Time, LastSignedInAt time.Time) *Member {
	return &Member{}
}

func (m *Member) Register(email Email, password Password) {
	m.ID = uuid.NewString()
	m.Email = email
	m.Password = password
	m.CreatedAt = time.Now().UTC()
}

func (m *Member) SignIn() {
	m.LastSignedInAt = time.Now().UTC()
}

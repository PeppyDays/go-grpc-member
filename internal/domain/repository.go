package domain

import (
	"context"
	"github.com/google/uuid"
)

type Reader interface {
	FindMemberByID(ctx context.Context, id uuid.UUID) (*Member, error)
}

type Writer interface {
	Save(ctx context.Context, m *Member) error
}

type ReadWriter interface {
	Reader
	Writer
}

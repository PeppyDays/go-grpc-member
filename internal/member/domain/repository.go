package domain

import "context"

type Reader interface {
	FindMemberByID(ctx context.Context, id string) (*Member, error)
}

type Writer interface {
	Save(ctx context.Context, m *Member) error
}

type ReadWriter interface {
	Reader
	Writer
}

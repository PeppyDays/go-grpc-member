package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/peppydays/go-grpc-member/internal/member/domain"
)

type MemoryRepository struct {
	items map[uuid.UUID]domain.Member
	sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		items: make(map[uuid.UUID]domain.Member),
	}
}

func (r *MemoryRepository) FindMemberByID(_ context.Context, id uuid.UUID) (*domain.Member, error) {
	if member, ok := r.items[id]; ok {
		return &member, nil
	}
	return nil, fmt.Errorf("member id %v not found in repository", id)
}

func (r *MemoryRepository) Save(_ context.Context, m *domain.Member) error {
	r.Lock()
	defer r.Unlock()
	r.items[m.ID] = *m
	return nil
}

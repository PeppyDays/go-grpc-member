package service

import (
	"context"
	"fmt"

	"github.com/peppydays/go-grpc-member/internal/member/domain"
)

type MemberService struct {
	repository domain.ReadWriter
}

func NewMemberService(repository domain.ReadWriter) *MemberService {
	return &MemberService{
		repository: repository,
	}
}

func (s *MemberService) RegisterMember(ctx context.Context, email string, password string) (*domain.Member, error) {
	validatedEmail, err := domain.NewEmail(email)
	if err != nil {
		return nil, err
	}

	validatedPassword, err := domain.NewPassword(password)
	if err != nil {
		return nil, err
	}

	member := domain.NewMember(validatedEmail, validatedPassword)
	member.Register()

	if err = s.repository.Save(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to register a member: %w", err)
	}

	return member, nil
}

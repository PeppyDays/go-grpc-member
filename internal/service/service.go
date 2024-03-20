package service

import (
	"context"
	"fmt"
	domain2 "github.com/peppydays/go-grpc-member/internal/domain"
)

type MemberService struct {
	repository domain2.ReadWriter
}

func NewMemberService(repository domain2.ReadWriter) *MemberService {
	return &MemberService{
		repository: repository,
	}
}

func (s *MemberService) SignUp(ctx context.Context, email string, password string) (*domain2.Member, error) {
	validatedEmail, err := domain2.NewEmail(email)
	if err != nil {
		return nil, err
	}

	validatedPassword, err := domain2.NewPassword(password)
	if err != nil {
		return nil, err
	}

	member := domain2.NewMember(validatedEmail, validatedPassword)
	member.Register()

	if err = s.repository.Save(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to register a member: %w", err)
	}

	return member, nil
}

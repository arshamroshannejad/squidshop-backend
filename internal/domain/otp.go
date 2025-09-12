package domain

import "context"

type OTPService interface {
	Generate(ctx context.Context, phone string) (string, error)
	Verify(ctx context.Context, phone, code string) (bool, error)
}

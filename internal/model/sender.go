package model

import "context"

type IVerifySender interface {
	SendCode(ctx context.Context, to string) (string, error)
	SendBindingMailCode(ctx context.Context, to string) (string, error)
	SendBindingCompletionMail(ctx context.Context, to string)
	SendUpMailCode(ctx context.Context, to string) (string, error)
	SendUpCompletionMail(ctx context.Context, to string) error
}

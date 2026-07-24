package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternal interface {
	CreateWallet(ctx context.Context, userId int) (*external.Wallet, error)
	SendEmail(ctx context.Context, recipient, templateName string, placehilder map[string]string) error
}

package interfaces

import (
	"context"
	"ewallet-ums/external"
)

type IExternalWallet interface {
	CreateWallet(ctx context.Context, userId int) (*external.Wallet, error)
}

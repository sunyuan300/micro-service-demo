package transport

import (
	"context"
)

type Server interface {
	Start(context.Context) error
	Stop(ctx context.Context) error
}
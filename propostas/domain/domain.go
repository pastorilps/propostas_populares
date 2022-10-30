package domain

import "context"

type DbConnect interface {
	CreateAttatchments(ctx context.Context) (err error)
}

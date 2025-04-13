package repository

import (
	"context"
)

type UserRepository[Model any] interface {
	FindAll(ctx context.Context) ([]Model, error)
	FindByName(ctx context.Context, name string) (Model, error)
	FindByID(ctx context.Context, id string) (Model, error)
	Insert(ctx context.Context, user Model) (Model, error)
	DeleteByID(ctx context.Context, id string) error
	Delete(ctx context.Context, user Model) error
	Update(ctx context.Context, user Model) error
}

package repository

import "context"

type Repository interface {
	GetDBType() string
}

type FindAll[Model any] interface {
	FindAll(ctx context.Context) ([]Model, error)
}

type FindByID[Model any] interface {
	FindByID(ctx context.Context, id string) (Model, error)
}

type FindByName[Model any] interface {
	FindByName(ctx context.Context, name string) (Model, error)
}

type Insert[Model any] interface {
	Insert(ctx context.Context, user Model) (Model, error)
}

type DeleteByID[Model any] interface {
	DeleteByID(ctx context.Context, id string) error
}

type Delete[Model any] interface {
	Delete(ctx context.Context, user Model) error
}

type Update[Model any] interface {
	Update(ctx context.Context, user Model) error
}

const cacheDuration = 10800 // 3 hours => 10800 seconds

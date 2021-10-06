package news

import "context"

type Domain struct {
	Article interface{}
}

type Repository interface {
	GetByCategory(ctx context.Context, name string) (Domain, error)
}

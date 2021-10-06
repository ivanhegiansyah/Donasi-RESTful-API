package news

import "context"

type Domain struct {
	Name        string
	Author      string
	Title       string
	Description string
	URL         string
}

type Repository interface {
	GetByName(ctx context.Context, name string) (Domain, error)
}

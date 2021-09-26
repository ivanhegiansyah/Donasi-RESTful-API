package donationtypes

import (
	"context"
	"time"
)

type Domain struct {
	Id               int
	Name string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	AddDonationType(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	AddDonationType(ctx context.Context, domain Domain, id int) (Domain, error)
}
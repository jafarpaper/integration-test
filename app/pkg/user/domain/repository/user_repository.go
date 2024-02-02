package repository

import (
	"context"
	"github.com/google/uuid"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"integration-test/app/pkg/user/domain/model/request"
)

type UserRepository interface {
	Get(ctx context.Context, req request.GetRequest) ([]*aggregate.User, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*aggregate.User, error)
	Create(ctx context.Context, req *aggregate.User) error
	Update(ctx context.Context, req *aggregate.User) error
	Delete(ctx context.Context, req *aggregate.User) error
}

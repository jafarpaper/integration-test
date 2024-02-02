package usecase

import (
	"context"
	"github.com/google/uuid"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"integration-test/app/pkg/user/domain/model/request"
)

type UserUseCase interface {
	Get(ctx context.Context, req request.GetRequest) ([]*aggregate.User, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*aggregate.User, error)
	Create(ctx context.Context, req request.UserRequest) error
	Update(ctx context.Context, req request.UserRequest, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

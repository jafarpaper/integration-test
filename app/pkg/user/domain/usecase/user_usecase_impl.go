package usecase

import (
	"context"
	"github.com/google/uuid"
	"integration-test/app/pkg/user/domain/exception"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"integration-test/app/pkg/user/domain/model/request"
	"integration-test/app/pkg/user/domain/repository"
)

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(userRepository repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{userRepository: userRepository}
}

func (u userUseCaseImpl) Get(ctx context.Context, req request.GetRequest) ([]*aggregate.User, int, error) {
	data, i, err := u.userRepository.Get(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return data, i, nil
}

func (u userUseCaseImpl) FindById(ctx context.Context, id uuid.UUID) (*aggregate.User, error) {
	data, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, exception.NotFoundError
	}

	return data, nil
}

func (u userUseCaseImpl) Create(ctx context.Context, req request.UserRequest) error {
	data := aggregate.ConvertToUser(req)
	err := u.userRepository.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCaseImpl) Update(ctx context.Context, req request.UserRequest, id uuid.UUID) error {
	data := aggregate.ConvertToUserUpdate(req, id)
	err := u.userRepository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCaseImpl) Delete(ctx context.Context, id uuid.UUID) error {
	data, err := u.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

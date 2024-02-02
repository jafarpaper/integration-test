package user

import (
	"github.com/go-playground/validator/v10"
	"integration-test/app/pkg/user/domain/repository"
	"integration-test/app/pkg/user/domain/usecase"
	"integration-test/app/pkg/user/infrastructure/persistence"
	"integration-test/app/pkg/user/interface/controller"
	"integration-test/utils/arangodb"
	"sync"
)

var once sync.Once

var (
	requestValidator *validator.Validate
	userUseCase      usecase.UserUseCase
	userRepository   repository.UserRepository
)

func InitializeUser(db arangodb.ArangoDB) {
	once.Do(func() {
		requestValidator = validator.New()
		userRepository = persistence.NewUserArangoRepository(db)
		userUseCase = usecase.NewUserUseCaseImpl(userRepository)
	})
}

func NewHttpUserController() *controller.UserController {
	return controller.NewHttpUserController(userUseCase, requestValidator)
}

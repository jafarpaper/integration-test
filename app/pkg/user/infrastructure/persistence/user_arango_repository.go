package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/noldwidjaja/slate/arango"
	"integration-test/app/pkg/user/domain/model/aggregate"
	"integration-test/app/pkg/user/domain/model/request"
	"integration-test/app/pkg/user/domain/repository"
	"integration-test/utils/constant"
)

type UserArangoRepository struct {
	arangoRepositoryInterface arango.ArangoBaseRepositoryInterface
}

func NewUserArangoRepository(db arango.ArangoDB) repository.UserRepository {
	return &UserArangoRepository{arangoRepositoryInterface: arango.NewArangoBaseRepository(db, constant.CollectionUser)}
}

func (u *UserArangoRepository) Get(ctx context.Context, req request.GetRequest) ([]*aggregate.User, int, error) {
	var users []*aggregate.User

	query := fmt.Sprintf(`For data in %s `, constant.CollectionUser)

	countQuery := fmt.Sprintf(`RETURN COUNT(FOR data in %s `, constant.CollectionUser)
	filter := ""

	if req.Filter != nil && len(req.Filter) > 0 {
		filter += " FILTER "
		for key, value := range req.Filter {
			filter += fmt.Sprintf("data.%s LIKE '%%%s%%' && ", key, value)
		}
		filter = filter[:len(filter)-4]
	}

	if len(req.Filter) == 0 && req.FilterAll != "" {
		filter += fmt.Sprintf(" FILTER data LIKE '%%%s%%' ", req.FilterAll)
	}

	query += filter

	countQuery += filter
	countQuery += " RETURN 1)"

	if req.Sort != nil && len(req.Sort) > 0 {
		query += " SORT "
		for field, direction := range req.Sort {
			query += fmt.Sprintf("data.%s %s, ", field, direction)
		}
		query = query[:len(query)-2]
	}

	req.Page = (req.Page * req.Size) - req.Size
	query += fmt.Sprintf(" LIMIT %d, %d ", req.Page, req.Size)
	query += "Return data"

	cursor, err := u.arangoRepositoryInterface.DB().Query(ctx, query, map[string]interface{}{})
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close()

	for cursor.HasMore() {
		var data aggregate.User
		_, err = cursor.ReadDocument(ctx, &data)
		if err != nil {
			return nil, 0, fmt.Errorf("error get all user arango repository : %w", err)
		}

		users = append(users, &data)
	}

	countCursor, err := u.arangoRepositoryInterface.DB().Query(ctx, countQuery, map[string]interface{}{})
	defer countCursor.Close()

	var totalCount int
	if countCursor.HasMore() {
		_, err = countCursor.ReadDocument(ctx, &totalCount)
		if err != nil {
			return nil, 0, fmt.Errorf("error reading count result: %w", err)
		}
	}

	return users, totalCount, nil
}

func (u *UserArangoRepository) FindById(ctx context.Context, id uuid.UUID) (*aggregate.User, error) {
	var collection []aggregate.User

	err := u.arangoRepositoryInterface.Where("id", id).Get(&collection)
	if err != nil {
		return nil, err
	}

	if &collection[0] == nil {
		return nil, nil
	}
	return &collection[0], nil
}

func (u *UserArangoRepository) Create(ctx context.Context, req *aggregate.User) error {
	err := u.arangoRepositoryInterface.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserArangoRepository) Update(ctx context.Context, req *aggregate.User) error {
	err := u.arangoRepositoryInterface.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserArangoRepository) Delete(ctx context.Context, req *aggregate.User) error {
	err := u.arangoRepositoryInterface.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

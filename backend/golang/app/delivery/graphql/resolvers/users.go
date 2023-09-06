package resolvers

import (
	"context"

	"github.com/taimoor99/three-tier-golang/app/delivery/graphql/generated"
	"github.com/taimoor99/three-tier-golang/app/entities"
)

type usersResolver struct{ *Resolver }

// Users returns generated.UsersResolver implementation.
func (r *Resolver) Users() generated.UsersResolver { return &usersResolver{r} }

func (r *usersResolver) ID(ctx context.Context, obj *entities.Users) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user entities.UsersCreateInput) (string, error) {
	userId, err := r.UserModel.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*entities.Users, error) {
	user, err := r.UserModel.FindUserByID(ctx, id)
	if err != nil {
		return &entities.Users{}, err
	}
	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context, limit, offset int) ([]entities.Users, error) {
	user, err := r.UserModel.GetAllUsers(ctx, int64(limit), int64(offset))
	if err != nil {
		return []entities.Users{}, err
	}
	return user, nil
}

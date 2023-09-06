package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/taimoor99/three-tier-golang/app/entities"
)

func GetSession(ctx context.Context, db string) (*mongo.Database, error) {
	// Connect to our mongo
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI("mongodb://mongodb:27017/"))

	if err != nil {
		return nil, err
	}

	return client.Database(db), err
}

type user struct {
	mgo *mongo.Database
}

type UserRepository interface {
	FindUserById(ctx context.Context, id string) (entities.Users, error)
	FindUserByEmail(ctx context.Context, email string) (entities.Users, error)
	CreateUser(ctx context.Context, users entities.Users) (entities.Users, error)
	GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error)
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &user{
		mgo: db,
	}
}

func (u *user) FindUserById(ctx context.Context, id string) (entities.Users, error) {
	var user entities.Users
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Users{}, err
	}
	if err := u.mgo.Collection(user.UsersCollection()).
		FindOne(ctx, bson.M{"_id": objID}).Decode(&user); err != nil {
		return entities.Users{}, err
	}
	return user, nil
}

func (u *user) FindUserByEmail(ctx context.Context, email string) (entities.Users, error) {
	var user entities.Users
	if err := u.mgo.Collection(user.UsersCollection()).
		FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return entities.Users{}, err
	}
	return user, nil
}

func (u *user) CreateUser(ctx context.Context, user entities.Users) (entities.Users, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.mgo.Collection(user.UsersCollection()).InsertOne(ctx, user)
	if err != nil {
		return entities.Users{}, err
	}
	return user, nil
}

func (u *user) GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error) {
	var users []entities.Users
	option := options.FindOptions{
		Limit: &limit,
		Skip: &offset,
	}
	rows, err := u.mgo.Collection(entities.Users{}.UsersCollection()).Find(ctx, nil, &option)
	if err != nil {
		return []entities.Users{}, err
	}

	if err := rows.Decode(&users); err != nil {
		return []entities.Users{}, err
	}

	return users, nil
}
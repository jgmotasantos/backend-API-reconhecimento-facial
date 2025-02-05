package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) FindOneByEmail(ctx context.Context, email string) (*User, bool) {

	filter := bson.M{"email": email}
	existingUser := &User{}

	// caso nao tenha erro, significa que a busca funcionou e existe
	// um user cadastrado com esse email
	err := r.collection.FindOne(ctx, filter).Decode(existingUser)

	if err != nil {

		return nil, false

	}

	return existingUser, true

}

func (r *UserRepository) FindOneByID(ctx context.Context, userID string) (*User, bool) {

	filter := bson.M{"_id": userID}
	existingUser := &User{}

	// caso nao tenha erro, significa que a busca funcionou e existe
	// um user cadastrado com esse email
	err := r.collection.FindOne(ctx, filter).Decode(existingUser)

	if err != nil {

		return nil, false

	}

	return existingUser, true

}

/////////////////////////
// POST /auth/register //
/////////////////////////

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {

	_, found := r.FindOneByEmail(ctx, user.Email)

	if found {
		return ErrEmailAlreadyExists
	}

	_, err := r.collection.InsertOne(ctx, user)

	return err

}

// DELETE /auth/delete

func (r *UserRepository) DeleteUser(ctx context.Context, userId string) error {

	_, found := r.FindOneByID(ctx, userId)

	if !found {

		return ErrUserNotFound

	}

    filter := bson.M{"_id": userId}

    _, err := r.collection.DeleteOne(ctx, filter)

    return err
}
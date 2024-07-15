package user

import (
	"context"
	"fmt"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/storage/mongostorage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	SaveUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, id string, u *User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
	SaveOrUpdateUser(ctx context.Context, u *User) error
	LogUserAction(ctx context.Context, log *UserActionLog) error
}

type mongoStorage struct {
	client            *mongo.Client
	database          string
	userCollection    string
	actionsCollection string
}

func NewMongoStorage(config mongostorage.Config) (Storage, error) {
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to ensure connection is established
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return &mongoStorage{
		client:            client,
		database:          config.Database,
		userCollection:    config.CollectionPrefix + "_users",
		actionsCollection: config.CollectionPrefix + "_user_actions",
	}, nil
}

func (s *mongoStorage) SaveUser(ctx context.Context, u *User) error {
	collection := s.client.Database(s.database).Collection(s.userCollection)
	u.LastUpdated = time.Now()
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (s *mongoStorage) UpdateUser(ctx context.Context, id string, u *User) error {
	collection := s.client.Database(s.database).Collection(s.userCollection)
	u.LastUpdated = time.Now()
	filter := bson.M{"id": id}
	update := bson.M{"$set": u}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (s *mongoStorage) GetUserByID(ctx context.Context, id string) (*User, error) {
	collection := s.client.Database(s.database).Collection(s.userCollection)
	filter := bson.M{"id": id}
	var u User
	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &u, nil
}

func (s *mongoStorage) SaveOrUpdateUser(ctx context.Context, u *User) error {
	collection := s.client.Database(s.database).Collection(s.userCollection)
	u.LastUpdated = time.Now()

	filter := bson.M{"id": u.ID}
	update := bson.M{"$set": u}
	opts := options.Update().SetUpsert(true) // This option will insert the document if it does not exist

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to save or update user: %w", err)
	}

	return nil
}

func (s *mongoStorage) LogUserAction(ctx context.Context, log *UserActionLog) error {
	collection := s.client.Database(s.database).Collection(s.actionsCollection)
	log.Timestamp = time.Now()
	_, err := collection.InsertOne(ctx, log)
	if err != nil {
		return fmt.Errorf("failed to log user action: %w", err)
	}
	return nil
}

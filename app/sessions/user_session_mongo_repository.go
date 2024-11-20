package sessions

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userSessionMongoDBRepository struct {
	collection *mongo.Collection
}

func NewUserSessionRepository(collection *mongo.Collection) *userSessionMongoDBRepository {
	return &userSessionMongoDBRepository{collection: collection}
}

func (r *userSessionMongoDBRepository) Create(session UserSession) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, session)

	return err
}

func (r *userSessionMongoDBRepository) Update(session *UserSession) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"external_id": session.ExternalID}, bson.M{"$set": session})

	return err
}

func (r *userSessionMongoDBRepository) GetByExternalID(externalID uint64) (*UserSession, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session UserSession

	err := r.collection.FindOne(ctx, bson.M{"external_id": strconv.FormatUint(externalID, 10)}).Decode(&session)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = SessionDoesNotExist
		}
		return &UserSession{}, err
	}

	return &session, nil
}

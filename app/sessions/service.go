package sessions

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserSessionService interface {
	StartSession(userSession UserSession) error

	FinishSession(externalSessionID string) error
}

type userSessionService struct {
	collection *mongo.Collection
}

func NewUserSessionService(collection *mongo.Collection) UserSessionService {
	return &userSessionService{collection: collection}
}

func (s *userSessionService) StartSession(userSession UserSession) error {
	log.Printf("%+v", userSession)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, userSession)

	return err
}

func (s *userSessionService) FinishSession(externalSessionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.UpdateOne(
		ctx,
		bson.M{"external_id": externalSessionID},
		bson.D{
			{
				"$set", bson.D{
					{"finished_at", time.Now()},
				},
			},
		},
	)

	return err
}

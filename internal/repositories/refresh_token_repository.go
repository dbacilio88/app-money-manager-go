package repositories

import (
	"context"
	"control-financiero/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefreshTokenRepository struct {
	collection *mongo.Collection
}

func NewRefreshTokenRepository(db *mongo.Database) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		collection: db.Collection("refresh_tokens"),
	}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, refreshToken *models.RefreshToken) error {
	refreshToken.ID = primitive.NewObjectID()
	refreshToken.CreatedAt = time.Now()
	
	_, err := r.collection.InsertOne(ctx, refreshToken)
	return err
}

func (r *RefreshTokenRepository) FindByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.collection.FindOne(ctx, bson.M{"token": token, "revoked": false}).Decode(&refreshToken)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepository) Revoke(ctx context.Context, token string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"token": token},
		bson.M{"$set": bson.M{"revoked": true}},
	)
	return err
}

func (r *RefreshTokenRepository) RevokeAllByUsuario(ctx context.Context, usuarioID primitive.ObjectID) error {
	_, err := r.collection.UpdateMany(
		ctx,
		bson.M{"usuarioId": usuarioID},
		bson.M{"$set": bson.M{"revoked": true}},
	)
	return err
}

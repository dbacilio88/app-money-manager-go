package repositories

import (
	"context"
	"control-financiero/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransaccionRepository struct {
	collection *mongo.Collection
}

func NewTransaccionRepository(db *mongo.Database) *TransaccionRepository {
	return &TransaccionRepository{
		collection: db.Collection("transacciones"),
	}
}

func (r *TransaccionRepository) Create(ctx context.Context, transaccion *models.Transaccion) error {
	transaccion.ID = primitive.NewObjectID()
	transaccion.CreatedAt = time.Now()
	transaccion.UpdatedAt = time.Now()
	
	_, err := r.collection.InsertOne(ctx, transaccion)
	return err
}

func (r *TransaccionRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Transaccion, error) {
	var transaccion models.Transaccion
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaccion)
	if err != nil {
		return nil, err
	}
	return &transaccion, nil
}

func (r *TransaccionRepository) FindByUsuario(ctx context.Context, usuarioID primitive.ObjectID) ([]*models.Transaccion, error) {
	opts := options.Find().SetSort(bson.D{{Key: "fecha", Value: -1}})
	cursor, err := r.collection.Find(ctx, bson.M{"usuarioId": usuarioID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transacciones []*models.Transaccion
	if err := cursor.All(ctx, &transacciones); err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (r *TransaccionRepository) FindByUsuarioAndRange(ctx context.Context, usuarioID primitive.ObjectID, start, end time.Time) ([]*models.Transaccion, error) {
	filter := bson.M{
		"usuarioId": usuarioID,
		"fecha": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}
	
	opts := options.Find().SetSort(bson.D{{Key: "fecha", Value: -1}})
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transacciones []*models.Transaccion
	if err := cursor.All(ctx, &transacciones); err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (r *TransaccionRepository) Update(ctx context.Context, transaccion *models.Transaccion) error {
	transaccion.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": transaccion.ID},
		bson.M{"$set": transaccion},
	)
	return err
}

func (r *TransaccionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

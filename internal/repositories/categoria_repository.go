package repositories

import (
	"context"
	"control-financiero/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoriaRepository struct {
	collection *mongo.Collection
}

func NewCategoriaRepository(db *mongo.Database) *CategoriaRepository {
	return &CategoriaRepository{
		collection: db.Collection("categorias"),
	}
}

func (r *CategoriaRepository) Create(ctx context.Context, categoria *models.Categoria) error {
	categoria.ID = primitive.NewObjectID()
	categoria.CreatedAt = time.Now()
	categoria.UpdatedAt = time.Now()
	
	_, err := r.collection.InsertOne(ctx, categoria)
	return err
}

func (r *CategoriaRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Categoria, error) {
	var categoria models.Categoria
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&categoria)
	if err != nil {
		return nil, err
	}
	return &categoria, nil
}

func (r *CategoriaRepository) FindAll(ctx context.Context, usuarioID *primitive.ObjectID) ([]*models.Categoria, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"usuarioId": nil},
			{"usuarioId": usuarioID},
		},
	}
	
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categorias []*models.Categoria
	if err := cursor.All(ctx, &categorias); err != nil {
		return nil, err
	}
	return categorias, nil
}

func (r *CategoriaRepository) Update(ctx context.Context, categoria *models.Categoria) error {
	categoria.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": categoria.ID},
		bson.M{"$set": categoria},
	)
	return err
}

func (r *CategoriaRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

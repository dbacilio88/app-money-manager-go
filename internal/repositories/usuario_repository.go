package repositories

import (
	"context"
	"control-financiero/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
	collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) *UsuarioRepository {
	return &UsuarioRepository{
		collection: db.Collection("usuarios"),
	}
}

func (r *UsuarioRepository) Create(ctx context.Context, usuario *models.Usuario) error {
	usuario.ID = primitive.NewObjectID()
	usuario.CreatedAt = time.Now()
	usuario.UpdatedAt = time.Now()
	
	_, err := r.collection.InsertOne(ctx, usuario)
	return err
}

func (r *UsuarioRepository) FindByEmail(ctx context.Context, email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) FindByGoogleID(ctx context.Context, googleID string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.collection.FindOne(ctx, bson.M{"googleId": googleID}).Decode(&usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) FindAll(ctx context.Context) ([]*models.Usuario, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var usuarios []*models.Usuario
	if err := cursor.All(ctx, &usuarios); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UsuarioRepository) Update(ctx context.Context, usuario *models.Usuario) error {
	usuario.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": usuario.ID},
		bson.M{"$set": usuario},
	)
	return err
}

func (r *UsuarioRepository) UpdateEstado(ctx context.Context, id primitive.ObjectID, estado string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"estado": estado, "updatedAt": time.Now()}},
	)
	return err
}

func (r *UsuarioRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

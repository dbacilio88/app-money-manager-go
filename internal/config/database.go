package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verificar conexión
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func InitializeCollections(client *mongo.Client, cfg *Config) error {
	db := client.Database(cfg.MongoDB)

	// Crear índices para usuarios
	usersCollection := db.Collection("usuarios")
	_, err := usersCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "googleId", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
	})
	if err != nil {
		return err
	}

	// Crear índices para transacciones
	transaccionesCollection := db.Collection("transacciones")
	_, err = transaccionesCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "usuarioId", Value: 1}, {Key: "fecha", Value: -1}},
		},
	})
	if err != nil {
		return err
	}

	// Crear índices para refresh tokens
	refreshTokensCollection := db.Collection("refresh_tokens")
	_, err = refreshTokensCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "token", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "expiresAt", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	})
	if err != nil {
		return err
	}

	// Inicializar roles por defecto
	rolesCollection := db.Collection("roles")
	count, err := rolesCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		roles := []interface{}{
			bson.M{
				"nombre":   "admin",
				"permisos": []string{"all"},
			},
			bson.M{
				"nombre":   "user",
				"permisos": []string{"read", "write"},
			},
		}
		_, err = rolesCollection.InsertMany(context.Background(), roles)
		if err != nil {
			return err
		}
		log.Println("✅ Roles por defecto creados")
	}

	// Crear categorías por defecto
	categoriasCollection := db.Collection("categorias")
	count, err = categoriasCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		categorias := []interface{}{
			bson.M{"nombre": "Salario", "tipo": "ingreso", "color": "#10b981", "usuarioId": nil},
			bson.M{"nombre": "Freelance", "tipo": "ingreso", "color": "#3b82f6", "usuarioId": nil},
			bson.M{"nombre": "Alimentación", "tipo": "egreso", "color": "#ef4444", "usuarioId": nil},
			bson.M{"nombre": "Transporte", "tipo": "egreso", "color": "#f59e0b", "usuarioId": nil},
			bson.M{"nombre": "Entretenimiento", "tipo": "egreso", "color": "#8b5cf6", "usuarioId": nil},
			bson.M{"nombre": "Servicios", "tipo": "egreso", "color": "#ec4899", "usuarioId": nil},
			bson.M{"nombre": "Alquiler", "tipo": "egreso", "color": "#6366f1", "usuarioId": nil},
			bson.M{"nombre": "Préstamo", "tipo": "ambos", "color": "#14b8a6", "usuarioId": nil},
		}
		_, err = categoriasCollection.InsertMany(context.Background(), categorias)
		if err != nil {
			return err
		}
		log.Println("✅ Categorías por defecto creadas")
	}

	log.Println("✅ Colecciones e índices inicializados")
	return nil
}

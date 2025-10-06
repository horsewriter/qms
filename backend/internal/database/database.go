package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quality-system/internal/models"
)

// DB is a struct that holds the MongoDB client.
type DB struct {
	Client *mongo.Client
}

// NewDB creates a new DB instance and connects to MongoDB.
func NewDB() *DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	return &DB{Client: client}
}

func (db *DB) getCollection(collectionName string) *mongo.Collection {
	return db.Client.Database("quality_system").Collection(collectionName)
}

// Employee database functions

func (db *DB) GetEmployees(ctx context.Context) ([]models.Employee, error) {
	collection := db.getCollection("employees")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []models.Employee
	if err = cursor.All(ctx, &employees); err != nil {
		return nil, err
	}

	return employees, nil
}

func (db *DB) CreateEmployee(ctx context.Context, employee models.Employee) (*mongo.InsertOneResult, error) {
	collection := db.getCollection("employees")
	return collection.InsertOne(ctx, employee)
}

func (db *DB) DeleteEmployee(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := db.getCollection("employees")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

// Area database functions

func (db *DB) GetAreas(ctx context.Context) ([]models.Area, error) {
	collection := db.getCollection("areas")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var areas []models.Area
	if err = cursor.All(ctx, &areas); err != nil {
		return nil, err
	}

	return areas, nil
}

func (db *DB) CreateArea(ctx context.Context, area models.Area) (*mongo.InsertOneResult, error) {
	collection := db.getCollection("areas")
	return collection.InsertOne(ctx, area)
}

func (db *DB) DeleteArea(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := db.getCollection("areas")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

// Level database functions

func (db *DB) GetLevels(ctx context.Context) ([]models.Level, error) {
	collection := db.getCollection("levels")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var levels []models.Level
	if err = cursor.All(ctx, &levels); err != nil {
		return nil, err
	}

	return levels, nil
}

func (db *DB) CreateLevel(ctx context.Context, level models.Level) (*mongo.InsertOneResult, error) {
	collection := db.getCollection("levels")
	return collection.InsertOne(ctx, level)
}

func (db *DB) DeleteLevel(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := db.getCollection("levels")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

// PartNumber database functions

func (db *DB) GetPartNumbers(ctx context.Context) ([]models.PartNumber, error) {
	collection := db.getCollection("part_numbers")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var partNumbers []models.PartNumber
	if err = cursor.All(ctx, &partNumbers); err != nil {
		return nil, err
	}

	return partNumbers, nil
}

func (db *DB) CreatePartNumber(ctx context.Context, partNumber models.PartNumber) (*mongo.InsertOneResult, error) {
	collection := db.getCollection("part_numbers")
	return collection.InsertOne(ctx, partNumber)
}

func (db *DB) DeletePartNumber(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := db.getCollection("part_numbers")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

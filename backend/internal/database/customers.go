package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"quality-system/internal/models"
)

func (db *DB) GetCustomers(ctx context.Context) ([]models.Customer, error) {
	collection := db.getCollection("customers")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customers []models.Customer
	if err = cursor.All(ctx, &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

func (db *DB) CreateCustomer(ctx context.Context, customer models.Customer) (*mongo.InsertOneResult, error) {
	collection := db.getCollection("customers")
	return collection.InsertOne(ctx, customer)
}

func (db *DB) UpdateCustomer(ctx context.Context, id primitive.ObjectID, customer models.Customer) (*mongo.UpdateResult, error) {
	collection := db.getCollection("customers")
	update := bson.M{"$set": bson.M{"customerID": customer.CustomerID, "name": customer.Name}}
	return collection.UpdateOne(ctx, bson.M{"_id": id}, update)
}

func (db *DB) DeleteCustomer(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := db.getCollection("customers")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Employee represents an employee in the system.
// Note: "Level" is intentionally omitted as it is now a separate, managed entity.
// The connection between an Employee and a Level will be handled through a reference
// if needed in the future, but for now, they are independent.
type Employee struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Number string             `bson:"number" json:"number"`
}

// Area represents a designated work area within the facility.
// This struct is used for managing areas where defective materials might be found.
type Area struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
}

// Level represents a classification level, which can be used for different purposes,
// such as employee seniority, material priority, or defect severity.
type Level struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
}

// PartNumber represents a specific part, along with its customer association.
// This struct is crucial for tracking parts throughout the manufacturing and quality process.
type PartNumber struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Number     string             `bson:"number" json:"number"`
	Customer   string             `bson:"customer" json:"customer"`
	CustomerID string             `bson:"customerID" json:"customerID"` // New field for Customer ID
}

type Customer struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CustomerID string             `bson:"customerID" json:"customerID"`
	Name       string             `bson:"name" json:"name"`
}

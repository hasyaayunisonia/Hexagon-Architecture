package db

import "go.mongodb.org/mongo-driver/mongo"

// DB is contains functions for products db.
type DB struct {
	db      *mongo.Database
	products string
}

// New to create new products db.
func New(db *mongo.Database, products string) *DB {
	return &DB{
		db:      db,
		products: products,
	}
}

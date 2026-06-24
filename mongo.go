// Package dbmongo is togo's MongoDB driver plugin.
//
// MongoDB is a document store, not a database/sql backend — togo's SQL ORM (sqlc +
// Atlas + `togo make:resource`) still targets Postgres/MySQL/SQLite. This plugin
// connects a *mongo.Client from DATABASE_URL during boot and exposes it via
// Client(), for document-store workloads alongside the SQL kernel. Install with
// `togo new --db mongodb` or `togo install togo-framework/db-mongodb`.
package dbmongo

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/togo-framework/togo"
)

var (
	mu     sync.RWMutex
	client *mongo.Client
)

// Client returns the connected MongoDB client, or nil before the plugin has booted
// (or when DATABASE_URL is unset).
func Client() *mongo.Client {
	mu.RLock()
	defer mu.RUnlock()
	return client
}

func init() {
	togo.RegisterProviderFunc("db-mongodb", togo.PriorityService, func(*togo.Kernel) error {
		uri := os.Getenv("DATABASE_URL")
		if uri == "" {
			return nil // no Mongo configured — leave Client() nil
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}
		mu.Lock()
		client = c
		mu.Unlock()
		return nil
	})
}

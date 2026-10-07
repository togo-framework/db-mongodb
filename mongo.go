// Package dbmongo is togo's MongoDB driver plugin.
//
// MongoDB is a document store, not a database/sql backend — togo's SQL ORM (sqlc +
// Atlas + `togo make:resource`) still targets Postgres/MySQL/SQLite. This plugin
// connects a *mongo.Client from MONGODB_URL (or a mongodb:// DATABASE_URL) during boot and exposes it via
// Client(), for document-store workloads alongside the SQL kernel. Install with
// `togo new --db mongodb` or `togo install togo-framework/db-mongodb`.
package dbmongo

import (
	"context"
	"os"
	"strings"
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
// (or when no MongoDB URL is configured).
func Client() *mongo.Client {
	mu.RLock()
	defer mu.RUnlock()
	return client
}

// mongoURI is MONGODB_URL, else DATABASE_URL, and only when it is a MongoDB URL: next to
// the SQL kernel DATABASE_URL is usually the Postgres/MySQL one, which is not ours to dial.
func mongoURI() string {
	for _, k := range []string{"MONGODB_URL", "DATABASE_URL"} {
		if u := os.Getenv(k); strings.HasPrefix(u, "mongodb://") || strings.HasPrefix(u, "mongodb+srv://") {
			return u
		}
	}
	return ""
}

func init() {
	togo.RegisterProviderFunc("db-mongodb", togo.PriorityService, func(*togo.Kernel) error {
		uri := mongoURI()
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

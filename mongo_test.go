package dbmongo

import "testing"

func TestMongoURI(t *testing.T) {
	cases := []struct {
		name, mongo, database, want string
	}{
		{"nothing set", "", "", ""},
		{"sql DATABASE_URL is not ours", "", "postgres://u:p@db:5432/app", ""},
		{"mongo DATABASE_URL", "", "mongodb://db:27017/app", "mongodb://db:27017/app"},
		{"srv DATABASE_URL", "", "mongodb+srv://cluster.example/app", "mongodb+srv://cluster.example/app"},
		{"MONGODB_URL next to a sql DATABASE_URL", "mongodb://m:27017/app", "postgres://u:p@db:5432/app", "mongodb://m:27017/app"},
		{"MONGODB_URL wins", "mongodb://m:27017/a", "mongodb://d:27017/b", "mongodb://m:27017/a"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv("MONGODB_URL", c.mongo)
			t.Setenv("DATABASE_URL", c.database)
			if got := mongoURI(); got != c.want {
				t.Fatalf("mongoURI() = %q, want %q", got, c.want)
			}
		})
	}
}

# db-mongodb

togo's **MongoDB** driver plugin. Connects a `*mongo.Client` from `DATABASE_URL`
during boot and exposes it via `dbmongo.Client()`.

> MongoDB is a document store, not a SQL backend. togo's ORM (sqlc + Atlas +
> `togo make:resource`) still targets Postgres/MySQL/SQLite — this plugin adds a
> Mongo client alongside it for document-store workloads.

```bash
togo new my-app --db mongodb         # wires it for you
# or, in an existing app:
togo install togo-framework/db-mongodb
```

```
DB_DRIVER=mongodb
DATABASE_URL=mongodb://user:pass@localhost:27017/my-app
```

```go
import dbmongo "github.com/togo-framework/db-mongodb"
coll := dbmongo.Client().Database("my-app").Collection("events")
```

MIT

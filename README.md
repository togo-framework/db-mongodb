<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo"></p>

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

---

## 💎 Premium sponsors

togo is proudly sponsored by **ID8 Media** and **One Studio**.

<p align="center">
  <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
</p>

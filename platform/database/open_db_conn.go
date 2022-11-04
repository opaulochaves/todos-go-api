package database

import "github.com/opchav/todos-go-api/app/queries"

type Queries struct {
  *queries.TodoQueries
}

func OpenDBConn() (*Queries, error) {
  db, err := PostgreSQLConnection()
  if err != nil {
    return nil, err
  }

  return &Queries{
    TodoQueries: &queries.TodoQueries{DB: db},
  }, nil
}

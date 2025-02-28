package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"skillsrock-test/api"
	"skillsrock-test/config"
	"skillsrock-test/repository/tasks"
	tasks2 "skillsrock-test/service/tasks"
)

func main() {
	cfg := config.Read()
	db := newDatabase(cfg)
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		panic(err)
	}

	tasksRepository := tasks.NewRepository(db)
	tasksService := tasks2.NewService(tasksRepository)
	server := api.NewServer(cfg, tasksService)
	server.ListenAndServe()
}

func newDatabase(cfg config.Settings) *sqlx.DB {
	pool, err := pgxpool.New(context.Background(), cfg.Database.PostgresConnection)
	if err != nil {
		panic(err)
	}

	return sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
}

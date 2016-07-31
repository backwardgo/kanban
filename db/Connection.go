package db

import (
	"log"

	"github.com/backwardgo/kanban/env"
	"github.com/jmoiron/sqlx"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

type Connection interface {
	Begin() (Transaction, error)
	Run(UnitOfWork) error
}

func NewConnection() (Connection, error) {
	return NewConnectionWithURL(env.DATABASE_URL)
}

func NewConnectionWithDB(sqlxDB *sqlx.DB) Connection {
	runner.MustPing(sqlxDB.DB)
	sqlxDB.MustExec(`set timezone='UTC'`)

	return &connection{
		runnerDB: runner.NewDBFromSqlx(sqlxDB),
	}
}

func NewConnectionWithURL(databaseURL string) (Connection, error) {
	sqlxDB, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	return NewConnectionWithDB(sqlxDB), nil
}

type connection struct {
	runnerDB *runner.DB
}

func (c *connection) Begin() (Transaction, error) {
	runnerTx, err := c.runnerDB.Begin()
	if err != nil {
		return nil, err
	}

	return newTransaction(runnerTx), nil
}

func (c *connection) Run(unitOfWork UnitOfWork) error {
	t, err := c.Begin()
	if err != nil {
		return err
	}

	if err = unitOfWork(t); err != nil {
		if re := t.Rollback(); re != nil {
			log.Printf("Rollback failed %v", re)
		}

		return err
	}

	return t.Commit()
}

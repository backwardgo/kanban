package db

import (
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

type Transaction interface {
	Commit() error
	Rollback() error

	AfterCommit(func())
	AfterRollback(func())

	insertInto(table string) *dat.InsertBuilder
	selekt(columns ...string) *dat.SelectBuilder
	update(table string) *dat.UpdateBuilder
}

func newTransaction(runnerTx *runner.Tx) *transaction {
	return &transaction{
		afterCommit:   []func(){},
		afterRollback: []func(){},
		runnerTx:      runnerTx,
	}
}

type transaction struct {
	afterCommit   []func()
	afterRollback []func()
	runnerTx      *runner.Tx
}

func (t *transaction) AfterCommit(cb func()) {
	t.afterCommit = append(t.afterCommit, cb)
}

func (t *transaction) AfterRollback(cb func()) {
	t.afterRollback = append(t.afterRollback, cb)
}

func (t *transaction) Rollback() error {
	if err := t.runnerTx.Rollback(); err != nil {
		return err
	}

	for _, afterRollback := range t.afterRollback {
		afterRollback()
	}

	return nil
}

func (t *transaction) Commit() error {
	if err := t.runnerTx.Commit(); err != nil {
		return err
	}

	for _, afterCommit := range t.afterCommit {
		afterCommit()
	}

	return nil
}

func (t *transaction) insertInto(table string) *dat.InsertBuilder {
	return t.runnerTx.InsertInto(table)
}

func (t *transaction) selekt(columns ...string) *dat.SelectBuilder {
	return t.runnerTx.Select(columns...)
}

func (t *transaction) update(table string) *dat.UpdateBuilder {
	return t.runnerTx.Update(table)
}

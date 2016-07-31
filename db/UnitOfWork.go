package db

type UnitOfWork func(Transaction) error

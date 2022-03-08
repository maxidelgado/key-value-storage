package kvs

import (
	"errors"
)

var (
	errKeyNotSet     = errors.New("key not set")
	errNoTransaction = errors.New("no transaction")
)

type TransactionalKVS interface {
	kvs
	Begin()
	Commit() error
	Rollback() error
}

type kvs interface {
	Set(string, string)
	Get(string) (string, error)
	Delete(string) error
	Count(string) int
}

func New() TransactionalKVS {
	return transactionalKVS{
		transactions: &transactionStack{},
		storage:      newMapKVS(),
	}
}

type transactionalKVS struct {
	transactions *transactionStack
	storage      MapKVS
}

func (k transactionalKVS) Set(key, value string) {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		k.storage.Set(key, value)
		return
	}
	tx.Set(key, value)
}

func (k transactionalKVS) Get(key string) (string, error) {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Get(key)
	}
	return tx.Get(key)
}

func (k transactionalKVS) Delete(key string) error {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Delete(key)
	}
	return tx.Delete(key)
}

func (k transactionalKVS) Count(value string) int {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Count(value)
	}
	return tx.Count(value)
}

func (k transactionalKVS) Begin() {
	k.transactions.push(newMapKVS())
}

func (k transactionalKVS) Commit() error {
	currentTx, err := k.transactions.pop()
	if err != nil {
		return err
	}

	parentTx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		k.storage.ApplyChanges(currentTx.GetMap())
		return nil
	}

	parentTx.ApplyChanges(currentTx.GetMap())
	return nil
}

func (k transactionalKVS) Rollback() error {
	_, err := k.transactions.pop()
	return err
}

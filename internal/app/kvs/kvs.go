package kvs

import (
	"errors"
)

var (
	errKeyNotSet     = errors.New("key not set")
	errNoTransaction = errors.New("no transaction")
)

type KVS interface {
	Set(string, string)
	Get(string) (string, error)
	Delete(string) error
	Count(string) int
	Begin()
	Commit() error
	Rollback() error
}

func New() KVS {
	return kvs{
		transactions: &transactionStack{},
		storage:      newStorage(),
	}
}

type kvs struct {
	transactions *transactionStack
	storage      storage
}

func (k kvs) Set(key, value string) {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		k.storage.Set(key, value)
		return
	}
	tx.Set(key, value)
}

func (k kvs) Get(key string) (string, error) {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Get(key)
	}
	return tx.Get(key)
}

func (k kvs) Delete(key string) error {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Delete(key)
	}
	return tx.Delete(key)
}

func (k kvs) Count(value string) int {
	tx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		return k.storage.Count(value)
	}
	return tx.Count(value)
}

func (k kvs) Begin() {
	k.transactions.push(newStorage())
}

func (k kvs) Commit() error {
	currentTx, err := k.transactions.pop()
	if err != nil {
		return err
	}

	parentTx, err := k.transactions.peek()
	if errors.Is(err, errNoTransaction) {
		k.storage.ApplyChanges(currentTx.GetChanges())
		return nil
	}

	parentTx.ApplyChanges(currentTx.GetChanges())
	return nil
}

func (k kvs) Rollback() error {
	_, err := k.transactions.pop()
	return err
}

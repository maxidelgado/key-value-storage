package kvs

import (
	"errors"
)

var (
	errKeyNotSet     = errors.New("key not set")
	errNoTransaction = errors.New("no transaction")
)

// KVS exposes the public methods for the key-value-storage implementation.
type KVS interface {
	// Set a key/value pair.
	Set(string, string)
	// Get the value for a given key.
	Get(string) (string, error)
	// Delete the key/value pair for a given key.
	Delete(string) error
	// Count the number of occurrences for a given value.
	Count(string) int
	// Begin a new transaction. Nested transactions are allowed.
	Begin()
	// Commit the current transaction.
	Commit() error
	// Rollback the current transaction.
	Rollback() error
}

func New() KVS {
	return kvs{
		stack:  &stack{},
		mainTx: newTransaction(),
	}
}

type kvs struct {
	// the stack where all the living transactions are stored.
	stack *stack
	// the mainTransaction where all the data is stored at the end. Its livecycle is the same as the app.
	mainTx transaction
}

func (k kvs) Set(key, value string) {
	// Check if there's a living transaction in the stack.
	tx, err := k.stack.peek()
	if errors.Is(err, errNoTransaction) {
		// If not, just proceed to set to the main tx.
		k.mainTx.set(key, value)
		return
	}
	tx.set(key, value)
}

func (k kvs) Get(key string) (string, error) {
	tx, err := k.stack.peek()
	if errors.Is(err, errNoTransaction) {
		return k.mainTx.get(key)
	}
	return tx.get(key)
}

func (k kvs) Delete(key string) error {
	tx, err := k.stack.peek()
	if errors.Is(err, errNoTransaction) {
		return k.mainTx.delete(key)
	}
	return tx.delete(key)
}

func (k kvs) Count(value string) int {
	tx, err := k.stack.peek()
	if errors.Is(err, errNoTransaction) {
		return k.mainTx.count(value)
	}
	return tx.count(value)
}

func (k kvs) Begin() {
	// Push a new tx to the stack.
	k.stack.push(newTransaction())
}

func (k kvs) Commit() error {
	// Take the latest tx in the stack and delete it.
	currentTx, err := k.stack.pop()
	if err != nil {
		return err
	}

	// Check if there's a previous tx in the stack (would be the parent tx).
	parentTx, err := k.stack.peek()
	if errors.Is(err, errNoTransaction) {
		// Apply the changes to the parent tx.
		k.mainTx.applyChanges(currentTx.storage)
		return nil
	}

	// If not, apply the changes to the main tx.
	parentTx.applyChanges(currentTx.storage)
	return nil
}

func (k kvs) Rollback() error {
	// Delete the latest tx in the stack.
	_, err := k.stack.pop()
	return err
}

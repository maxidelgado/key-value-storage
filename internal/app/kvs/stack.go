package kvs

// stack is an ordered, linked collection of transactions.
type stack []transaction

// isEmpty return true if the stack is empty.
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

// push a tx to the end of the stack.
func (s *stack) push(storage transaction) {
	*s = append(*s, storage)
}

// peek returns the latest tx in the stack. If any tx wasn't found, return an error.
func (s *stack) peek() (transaction, error) {
	if s.isEmpty() {
		return transaction{}, errNoTransaction
	}
	return (*s)[len(*s)-1], nil
}

// pop returns the latest tx in the stack and deletes it. If any tx wasn't found, return an error.
func (s *stack) pop() (transaction, error) {
	if s.isEmpty() {
		return transaction{}, errNoTransaction
	}
	latest := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return latest, nil
}

package kvs

type transactionStack []storage

func (s *transactionStack) isEmpty() bool {
	return len(*s) == 0
}

func (s *transactionStack) push(storage storage) {
	*s = append(*s, storage)
}

func (s *transactionStack) peek() (storage, error) {
	if s.isEmpty() {
		return storage{}, errNoTransaction
	}
	return (*s)[len(*s)-1], nil
}

func (s *transactionStack) pop() (storage, error) {
	if s.isEmpty() {
		return storage{}, errNoTransaction
	}
	latest := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return latest, nil
}

package kvs

type transactionStack []MapKVS

func (s *transactionStack) isEmpty() bool {
	return len(*s) == 0
}

func (s *transactionStack) push(storage MapKVS) {
	*s = append(*s, storage)
}

func (s *transactionStack) peek() (MapKVS, error) {
	if s.isEmpty() {
		return nil, errNoTransaction
	}
	return (*s)[len(*s)-1], nil
}

func (s *transactionStack) pop() (MapKVS, error) {
	if s.isEmpty() {
		return nil, errNoTransaction
	}
	latest := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return latest, nil
}

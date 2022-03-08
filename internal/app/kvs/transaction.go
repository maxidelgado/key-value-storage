package kvs

func newTransaction() transaction {
	return transaction{
		storage:  make(map[string]string),
		countMap: make(map[string]int),
	}
}

// transaction stores key/value pairs while it's alive.
type transaction struct {
	storage  map[string]string
	countMap map[string]int
}

func (s transaction) set(key, value string) {
	s.storage[key] = value
	s.countMap[value] += 1
}

func (s transaction) get(key string) (string, error) {
	value, found := s.storage[key]
	if !found {
		return "", errKeyNotSet
	}
	return value, nil
}

func (s transaction) delete(key string) error {
	if value, found := s.storage[key]; found {
		delete(s.storage, key)
		s.countMap[value] -= 1
		return nil
	}
	return errKeyNotSet
}

func (s transaction) count(value string) int {
	return s.countMap[value]
}

// applyChanges is to merge a given changesMap (provided by a child's tx) to the internal map of the current tx.
func (s transaction) applyChanges(changesMap map[string]string) {
	for key, value := range changesMap {
		_ = s.delete(key)
		s.set(key, value)
	}
}

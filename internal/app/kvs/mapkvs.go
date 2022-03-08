package kvs

type MapStorage map[string]string

type MapKVS interface {
	kvs
	GetMap() MapStorage
	ApplyChanges(MapStorage)
}

func newMapKVS() MapKVS {
	return mapKVS{
		storage: make(MapStorage),
		count:   make(map[string]int),
	}
}

type mapKVS struct {
	storage MapStorage
	count   map[string]int
}

func (s mapKVS) Set(key, value string) {
	s.storage[key] = value
	s.count[value] += 1
}

func (s mapKVS) Get(key string) (string, error) {
	value, found := s.storage[key]
	if !found {
		return "", errKeyNotSet
	}
	return value, nil
}

func (s mapKVS) Delete(key string) error {
	if value, found := s.storage[key]; found {
		delete(s.storage, key)
		s.count[value] -= 1
		return nil
	}
	return errKeyNotSet
}

func (s mapKVS) Count(value string) int {
	return s.count[value]
}

func (s mapKVS) ApplyChanges(m MapStorage) {
	for key, value := range m {
		_ = s.Delete(key)
		s.Set(key, value)
	}
}

func (s mapKVS) GetMap() MapStorage {
	return s.storage
}

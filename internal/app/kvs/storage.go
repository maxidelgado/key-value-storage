package kvs

func newStorage() storage {
	return storage{
		memoryMap: make(map[string]string),
		count:     make(map[string]int),
	}
}

type storage struct {
	memoryMap map[string]string
	count     map[string]int
}

func (s storage) Set(key, value string) {
	s.memoryMap[key] = value
	s.count[value] += 1
}

func (s storage) Get(key string) (string, error) {
	value, found := s.memoryMap[key]
	if !found {
		return "", errKeyNotSet
	}
	return value, nil
}

func (s storage) Delete(key string) error {
	if value, found := s.memoryMap[key]; found {
		delete(s.memoryMap, key)
		s.count[value] -= 1
		return nil
	}
	return errKeyNotSet
}

func (s storage) Count(value string) int {
	return s.count[value]
}

func (s storage) ApplyChanges(m map[string]string) {
	for key, value := range m {
		_ = s.Delete(key)
		s.Set(key, value)
	}
}

func (s storage) GetChanges() map[string]string {
	return s.memoryMap
}

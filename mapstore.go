package toaster

import (
	"fmt"
)

type MapStore map[string]*Toast

func (ms *MapStore) Add(toast *Toast) (string, error) {
	id := generateUniqueID()
	(*ms)[id] = toast
	return id, nil
}

func (ms MapStore) Get(key string) (*Toast, error) {
	t, ok := ms[key]
	if !ok {
		return t, fmt.Errorf("toaster: '%s' key not found", key)
	}
	return t, nil
}

func (ms *MapStore) Delete(key string) error {
	_, ok := (*ms)[key]
	if !ok {
		return fmt.Errorf("toaster: '%s' key not found", key)
	}

	delete(*ms, key)
	return nil
}

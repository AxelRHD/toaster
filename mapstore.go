package toaster

import (
	"fmt"
)

type MapStore struct {
	Messages            map[string]*Toast
	ToastTemplate       string
	HyperscriptTemplate string
}

func CreateMapStore() MapStore {
	return MapStore{
		ToastTemplate:       toastTemplate,
		HyperscriptTemplate: jsTemplate,
	}
}

func (ms *MapStore) Add(toast *Toast) (string, error) {
	id := generateUniqueID()

	(*ms).Messages[id] = toast
	return id, nil
}

func (ms MapStore) Get(key string) (*Toast, error) {
	t, ok := ms.Messages[key]
	if !ok {
		return t, fmt.Errorf("toaster: '%s' key not found", key)
	}
	return t, nil
}

func (ms *MapStore) Delete(key string) error {
	_, ok := (*ms).Messages[key]
	if !ok {
		return fmt.Errorf("toaster: '%s' key not found", key)
	}

	delete(ms.Messages, key)
	return nil
}

func (ms MapStore) GetToastTempl() string {
	return ms.ToastTemplate
}

func (ms *MapStore) SetToastTempl(tmpl string) {
	ms.ToastTemplate = tmpl
}

func (ms MapStore) GetJsTempl() string {
	return ms.HyperscriptTemplate
}

func (ms *MapStore) SetJsTempl(tmpl string) {
	ms.HyperscriptTemplate = tmpl
}

package toaster

import (
	"fmt"
)

type MapStore struct {
	Messages            map[string]*Toast
	ToastTemplate       string
	HyperscriptTemplate string
}

func CreateMapStore() *MapStore {
	return &MapStore{
		Messages:            make(map[string]*Toast),
		ToastTemplate:       toastTemplate,
		HyperscriptTemplate: jsTemplate,
	}
}

func (ms MapStore) New(msg string) *Toast {
	return &Toast{
		Message:             msg,
		Location:            LOC_TOP_RIGHT,
		Icon:                true,
		ToastTemplate:       ms.ToastTemplate,
		HyperscriptTemplate: ms.HyperscriptTemplate,
	}
}

func (ms *MapStore) Save(toast *Toast) (string, error) {
	id := generateUniqueID()

	toast.ID = id

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

func (ms MapStore) GetHyperTempl() string {
	return ms.HyperscriptTemplate
}

func (ms *MapStore) SetHyperTempl(tmpl string) {
	ms.HyperscriptTemplate = tmpl
}

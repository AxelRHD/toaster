package toaster

import (
	"fmt"
	"strings"
	"time"

	"github.com/devjefster/GoShortUniqueID/idgen"
)

type Toast struct {
	ID            string        `json:"id" db:"id"`
	Title         string        `json:"title,omitempty" db:"title"`
	Message       string        `json:"message,omitempty" db:"message"`
	Location      ToastLocation `json:"location,omitempty" db:"location"`
	Icon          bool          `json:"icon" db:"icon"`
	Dismissable   bool          `json:"dismissable" db:"dismissable"`
	Type          ToastType     `json:"type,omitempty" db:"msg_type"`
	CreatedAt     time.Time     `json:"created_at" db:"-"`
	CreatedAtStr  string        `json:"-" db:"created_at"`
	ToastTemplate string        `json:"-" db:"-"`
	JsTemplate    string        `json:"-" db:"-"`
}

type Store interface {
	Add(*Toast) (string, error)
	Get(string) (*Toast, error)
	Delete(string) error
	GetToastTempl() string
	SetToastTempl(string)
	GetJsTempl() string
	SetJsTempl(string)
}

type ToastLocation string
type ToastType string

const (
	LOC_TOP_LEFT      ToastLocation = "top-left"
	LOC_TOP_CENTER    ToastLocation = "top-center"
	LOC_TOP_RIGHT     ToastLocation = "top-right" // default
	LOC_BOTTOM_LEFT   ToastLocation = "bottom-left"
	LOC_BOTTOM_CENTER ToastLocation = "bottom-center"
	LOC_BOTTOM_RIGHT  ToastLocation = "bottom-right"

	TYPE_SUCCESS ToastType = "success"
	TYPE_WARNING ToastType = "warning"
	TYPE_ERROR   ToastType = "error"
	TYPE_INFO    ToastType = "info"
)

const (
	toastTemplate = "butterup.toast({ %s })"
	jsTemplate    = "on load %s remove me"
)

func New(msg string) *Toast {
	return &Toast{
		Message:  msg,
		Location: LOC_TOP_RIGHT,
		Icon:     true,
	}
}

func (t *Toast) SetTitle(title string) *Toast {
	t.Title = title
	return t
}

func (t *Toast) SetMessage(msg string) *Toast {
	t.Message = msg
	return t
}

func (t *Toast) SetErrMessage(msg error) *Toast {
	t.Message = msg.Error()
	t.Type = TYPE_ERROR
	return t
}

func (t *Toast) SetLocation(loc ToastLocation) *Toast {
	t.Location = loc
	return t
}

func (t *Toast) ShowIcon(icn bool) *Toast {
	t.Icon = icn
	return t
}

func (t *Toast) SetDismissable(dis bool) *Toast {
	t.Dismissable = dis
	return t
}

func (t *Toast) SetType(tp ToastType) *Toast {
	t.Type = tp
	return t
}

func (t Toast) Render() string {
	var parts []string

	if t.Title != "" {
		parts = append(parts, fmt.Sprintf("title: '%v'", t.Title))
	}

	if t.Message != "" {
		parts = append(parts, fmt.Sprintf("message: '%v'", t.Message))
	}

	if t.Location != "" {
		parts = append(parts, fmt.Sprintf("location: '%v'", t.Location))
	}

	parts = append(parts, fmt.Sprintf("icon: %v", t.Icon))

	parts = append(parts, fmt.Sprintf("dismissable: %v", t.Dismissable))

	if t.Type != "" {
		parts = append(parts, fmt.Sprintf("type: '%v'", t.Type))
	}

	fmt.Println("templ:", t.ToastTemplate)
	return fmt.Sprintf(t.ToastTemplate, strings.Join(parts, ", "))
}

func (t Toast) RenderHyperscript() string {
	msg := t.Render()

	return fmt.Sprintf(t.JsTemplate, msg)
}

func generateUniqueID() string {
	return idgen.New(8, "", "").Generate()
}

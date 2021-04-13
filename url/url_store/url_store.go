package url_store

import (
	"lld/url_shortner/url"
)

type URLStore interface {
	Init() error
	Get(short string) (url.URLObject, error)
	Store(url.URLObject) error
	Delete(url.URLObject) error

	GetUnique() (string, error)
}

package reader

import (
	"os"
)

type IReader interface {
	List() ([]string, error)
	ReadDir(dirname string) ([]os.DirEntry, error)
}

type Reader struct{}

func NewReader() IReader {
	return &Reader{}
}

func (r Reader) List() ([]string, error) {
	templatesName := make([]string, 0)

	entries, err := r.ReadDir("template")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			templatesName = append(templatesName, entry.Name())
		}
	}

	return templatesName, nil
}

func (r Reader) ReadDir(dirname string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

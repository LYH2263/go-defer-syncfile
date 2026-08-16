package save

import (
	"os"
)

type file interface {
	Write([]byte) (int, error)
	Sync() error
	Close() error
}

func writeTo(f file, data []byte) (err error) {
	defer f.Close()
	if _, err = f.Write(data); err != nil {
		return err
	}
	defer f.Sync() // BUG: Sync error discarded
	return nil
}

// WriteFile writes data using os file; tests inject via writeTo.
func WriteFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return writeTo(f, data)
}

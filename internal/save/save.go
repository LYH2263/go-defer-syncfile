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
	// Defers run LIFO: Close (deferred first) runs last, after Sync below —
	// the correct order for durability. Each is wrapped in a closure so its
	// error propagates to the named result `err`, without overwriting an
	// earlier error (e.g. a Write or Sync failure must not be masked by Close).
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	if _, err = f.Write(data); err != nil {
		return err
	}
	defer func() {
		if serr := f.Sync(); serr != nil && err == nil {
			err = serr
		}
	}()
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

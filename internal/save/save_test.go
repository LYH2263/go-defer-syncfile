package save

import (
	"errors"
	"testing"
)

type fake struct {
	syncErr error
}

func (f *fake) Write(p []byte) (int, error) { return len(p), nil }
func (f *fake) Sync() error                   { return f.syncErr }
func (f *fake) Close() error                  { return nil }

func TestSyncErrorSurfaces(t *testing.T) {
	boom := errors.New("sync failed")
	err := writeTo(&fake{syncErr: boom}, []byte("x"))
	if !errors.Is(err, boom) {
		t.Fatalf("got %v", err)
	}
}

package poppler

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	t.Logf("version = '%s'\n", version)
	if version == "" {
		t.Fail()
	}
}

func TestGetBackend(t *testing.T) {
	backend := GetBackend()
	t.Logf("backend = #%d - '%s'\n", backend, backend)
	if backend == BackendUnknown {
		t.Fail()
	}
}

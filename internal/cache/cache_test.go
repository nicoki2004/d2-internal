package cache

import (
	"os"
	"path/filepath"
	"testing"
)

type sample struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestFileCache_SaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cache.json")

	fc := NewFileCache()
	in := sample{Foo: "ok", Bar: 42}
	if err := fc.Save(path, in); err != nil {
		t.Fatalf("save: %v", err)
	}

	var out sample
	if err := fc.Load(path, &out); err != nil {
		t.Fatalf("load: %v", err)
	}

	if out != in {
		t.Fatalf("unexpected data: %+v", out)
	}
}

func TestGetCacheFileName_Default(t *testing.T) {
	old := os.Getenv("CACHE_FILE_NAME")
	if err := os.Unsetenv("CACHE_FILE_NAME"); err != nil {
		t.Fatalf("unset env: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("CACHE_FILE_NAME", old)
	})

	if got := GetChacheFileName(); got != "profile_cache.json" {
		t.Fatalf("expected default cache name, got %s", got)
	}
}

func TestGetCacheFileName_EnvOverride(t *testing.T) {
	old := os.Getenv("CACHE_FILE_NAME")
	if err := os.Setenv("CACHE_FILE_NAME", "custom.json"); err != nil {
		t.Fatalf("set env: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("CACHE_FILE_NAME", old)
	})

	if got := GetChacheFileName(); got != "custom.json" {
		t.Fatalf("expected env cache name, got %s", got)
	}
}

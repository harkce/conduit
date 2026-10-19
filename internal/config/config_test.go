package config

import "testing"

func TestLoadDefaultAddr(t *testing.T) {
	t.Setenv("ADDR", "")
	if got := env("UNSET_TEST_KEY_XYZ", ":3000"); got != ":3000" {
		t.Fatalf("expected %q, got %q", ":3000", got)
	}
}

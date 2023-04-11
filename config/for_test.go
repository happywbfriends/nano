package config

import "testing"

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

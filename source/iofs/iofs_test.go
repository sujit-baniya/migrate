//go:build go1.16
// +build go1.16

package iofs_test

import (
	"testing"

	"github.com/sujit-baniya/migrate/source/iofs"
	st "github.com/sujit-baniya/migrate/source/testing"
)

func Test(t *testing.T) {
	// reuse the embed.FS set in example_test.go
	d, err := iofs.New(fs, "testdata/migrations")
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}

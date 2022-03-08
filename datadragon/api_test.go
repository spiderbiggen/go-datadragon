//go:build integration
// +build integration

package datadragon

import (
	"context"
	"testing"
)

func TestDataDragon_Languages(t *testing.T) {
	t.Run("Languages", func(t *testing.T) {
		d := New()
		got, err := d.Languages(context.Background())
		if err != nil {
			t.Errorf("Languages() error = %v", err)
			return
		}
		if len(got) == 0 {
			t.Errorf("Languages() expected at least one language")
			return
		}
	})
}

func TestDataDragon_Locales(t *testing.T) {
	t.Run("Locales", func(t *testing.T) {
		d := New()
		got, err := d.Locales(context.Background())
		if err != nil {
			t.Errorf("Locales() error = %v", err)
			return
		}
		if len(got) == 0 {
			t.Errorf("Locales() expected at least one locale")
			return
		}
	})
}

func TestDataDragon_Versions(t *testing.T) {
	t.Run("Versions", func(t *testing.T) {
		d := New()
		got, err := d.Versions(context.Background())
		if err != nil {
			t.Errorf("Versions() error = %v", err)
			return
		}
		if len(got) == 0 {
			t.Errorf("Versions() expected at least one version")
			return
		}
	})
}

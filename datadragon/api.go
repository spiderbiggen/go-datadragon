package datadragon

import (
	"context"
)

func (d DataDragon) Versions(ctx context.Context) (versions []string, err error) {
	err = d.apiJson(ctx, "versions.json", &versions)
	return
}

func (d DataDragon) Languages(ctx context.Context) ([]string, error) {
	return d.Locales(ctx)
}

func (d DataDragon) Locales(ctx context.Context) (locales []string, err error) {
	err = d.cdnJson(ctx, "languages.json", &locales)
	return
}

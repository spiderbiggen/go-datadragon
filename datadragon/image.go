package datadragon

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

var (
	ErrNoChampion = errors.New("no champion specified")
	ErrNoSpell    = errors.New("no spell specified")
)

func (d DataDragon) ImageSplash(ctx context.Context, champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.ImageSplashRaw(ctx, champion, skin))
}

func (d DataDragon) ImageSplashRaw(ctx context.Context, champion string, skin uint8) (io.ReadCloser, error) {
	if champion == "" {
		return nil, ErrNoChampion
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("img/champion/splash/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) ImageLoading(ctx context.Context, champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.ImageLoadingRaw(ctx, champion, skin))
}

func (d DataDragon) ImageLoadingRaw(ctx context.Context, champion string, skin uint8) (io.ReadCloser, error) {
	if champion == "" {
		return nil, ErrNoChampion
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("img/champion/loading/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) ImageChampionSquare(ctx context.Context, champion string, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.ImageChampionSquareRaw(ctx, champion, requestConfig...))
}

func (d DataDragon) ImageChampionSquareRaw(ctx context.Context, champion string, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	if champion == "" {
		return nil, ErrNoChampion
	}
	c, err := d.mergeConfig(requestConfig, RequireVersion())
	if err != nil {
		return nil, err
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("%s/img/champion/%s.png", c.Version, champion))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) ImagePassive(ctx context.Context, champion ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.ImagePassiveRaw(ctx, champion, requestConfig...))
}

func (d DataDragon) ImagePassiveRaw(ctx context.Context, passive ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	if passive.Full == "" {
		return nil, ErrNoChampion
	}
	c, err := d.mergeConfig(requestConfig, RequireVersion())
	if err != nil {
		return nil, err
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("%s/img/passive/%s", c.Version, passive.Full))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) ImageSpell(ctx context.Context, spell ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.ImageSpellRaw(ctx, spell, requestConfig...))
}

func (d DataDragon) ImageSpellRaw(ctx context.Context, spell ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	if spell.Full == "" {
		return nil, ErrNoSpell
	}
	c, err := d.mergeConfig(requestConfig, RequireVersion())
	if err != nil {
		return nil, err
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("%s/img/spell/%s", c.Version, spell.Full))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) ImageItem(ctx context.Context, key uint16, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.ImageItemRaw(ctx, key, requestConfig...))
}

func (d DataDragon) ImageItemRaw(ctx context.Context, key uint16, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	c, err := d.mergeConfig(requestConfig, RequireVersion())
	if err != nil {
		return nil, err
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("%s/img/item/%d.png", c.Version, key))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) readJpeg(body io.ReadCloser, err error) (image.Image, error) {
	return d.readImage(body, err, jpeg.Decode)
}

func (d DataDragon) readPng(body io.ReadCloser, err error) (image.Image, error) {
	return d.readImage(body, err, png.Decode)
}

func (d DataDragon) readImage(body io.ReadCloser, err error, decoder func(io.Reader) (image.Image, error)) (image.Image, error) {
	if err != nil {
		return nil, err
	}
	defer d.closeBody(body)
	img, err := decoder(body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

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
)

func (d DataDragon) SplashImage(ctx context.Context, champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.SplashImageRaw(ctx, champion, skin))
}

func (d DataDragon) SplashImageRaw(ctx context.Context, champion string, skin uint8) (io.ReadCloser, error) {
	if champion == "" {
		return nil, ErrNoChampion
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("img/champion/splash/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) LoadingImage(ctx context.Context, champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.LoadingImageRaw(ctx, champion, skin))
}

func (d DataDragon) LoadingImageRaw(ctx context.Context, champion string, skin uint8) (io.ReadCloser, error) {
	if champion == "" {
		return nil, ErrNoChampion
	}
	resp, err := d.cdnRequest(ctx, fmt.Sprintf("img/champion/loading/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d DataDragon) SquareChampionImage(ctx context.Context, champion string, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.SquareChampionImageRaw(ctx, champion, requestConfig...))
}

func (d DataDragon) SquareChampionImageRaw(ctx context.Context, champion string, requestConfig ...RequestConfig) (io.ReadCloser, error) {
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

func (d DataDragon) PassiveImage(ctx context.Context, champion ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.PassiveImageRaw(ctx, champion, requestConfig...))
}

func (d DataDragon) PassiveImageRaw(ctx context.Context, passive ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
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

func (d DataDragon) SpellImage(ctx context.Context, spell ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.SpellImageRaw(ctx, spell, requestConfig...))
}

func (d DataDragon) SpellImageRaw(ctx context.Context, spell ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	if spell.Full == "" {
		return nil, ErrNoChampion
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

package datadragon

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func (d *DataDragon) SplashImage(champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.SplashImageRaw(champion, skin))
}

func (d *DataDragon) SplashImageRaw(champion string, skin uint8) (io.ReadCloser, error) {
	resp, err := d.cdnRequest(fmt.Sprintf("img/champion/splash/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *DataDragon) LoadingImage(champion string, skin uint8) (image.Image, error) {
	return d.readJpeg(d.LoadingImageRaw(champion, skin))
}

func (d *DataDragon) LoadingImageRaw(champion string, skin uint8) (io.ReadCloser, error) {
	resp, err := d.cdnRequest(fmt.Sprintf("img/champion/loading/%s_%d.jpg", champion, skin))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *DataDragon) SquareChampionImage(champion string, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.SquareChampionImageRaw(champion, requestConfig...))
}

func (d *DataDragon) SquareChampionImageRaw(champion string, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	c := d.mergeConfig(requestConfig)
	resp, err := d.cdnRequest(fmt.Sprintf("%s/img/champion/%s.png", c.Version, champion))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *DataDragon) PassiveImage(champion ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.PassiveImageRaw(champion, requestConfig...))
}

func (d *DataDragon) PassiveImageRaw(champion ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	c := d.mergeConfig(requestConfig)
	resp, err := d.cdnRequest(fmt.Sprintf("%s/img/passive/%s", c.Version, champion.Full))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *DataDragon) SpellImage(spell ChampionImage, requestConfig ...RequestConfig) (image.Image, error) {
	return d.readPng(d.SpellImageRaw(spell, requestConfig...))
}

func (d *DataDragon) SpellImageRaw(spell ChampionImage, requestConfig ...RequestConfig) (io.ReadCloser, error) {
	c := d.mergeConfig(requestConfig)
	resp, err := d.cdnRequest(fmt.Sprintf("%s/img/spell/%s", c.Version, spell.Full))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (d *DataDragon) readJpeg(body io.ReadCloser, err error) (image.Image, error) {
	return d.readImage(body, err, jpeg.Decode)
}

func (d *DataDragon) readPng(body io.ReadCloser, err error) (image.Image, error) {
	return d.readImage(body, err, png.Decode)
}

func (d *DataDragon) readImage(body io.ReadCloser, err error, decoder func(io.Reader) (image.Image, error)) (image.Image, error) {
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

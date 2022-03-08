package datadragon

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("page not found")
)

func (d DataDragon) apiJson(ctx context.Context, path string, v interface{}) error {
	resp, err := d.apiRequest(ctx, path)
	if err != nil {
		return err
	}
	defer d.closeBody(resp.Body)
	return json.NewDecoder(resp.Body).Decode(v)
}

func (d DataDragon) realmsJson(ctx context.Context, path string, v interface{}) error {
	resp, err := d.realmsRequest(ctx, path)
	if err != nil {
		return err
	}
	defer d.closeBody(resp.Body)
	return json.NewDecoder(resp.Body).Decode(v)
}

func (d DataDragon) cdnJson(ctx context.Context, path string, v interface{}) error {
	resp, err := d.cdnRequest(ctx, path)
	if err != nil {
		return err
	}
	defer d.closeBody(resp.Body)
	return json.NewDecoder(resp.Body).Decode(v)
}

func (d DataDragon) apiRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		d.closeBody(resp.Body)
		return nil, ErrNotFound
	}
	return resp, err
}

func (d DataDragon) realmsRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, realmsUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		d.closeBody(resp.Body)
		return nil, ErrNotFound
	}
	return resp, err
}

func (d DataDragon) cdnRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cdnUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if resp.StatusCode >= 403 && resp.StatusCode <= 404 {
		d.closeBody(resp.Body)
		return nil, ErrNotFound
	}
	return resp, err
}

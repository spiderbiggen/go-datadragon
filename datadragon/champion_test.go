//go:build integration

package datadragon

import (
	"context"
	"testing"
)

func TestDataDragon_Champion(t *testing.T) {
	type fields struct {
		opts []opt
	}
	type args struct {
		ctx    context.Context
		id     string
		params []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Aatrox",
			fields: fields{opts: []opt{WithVersion("12.4.1")}},
			args: args{
				ctx: context.Background(),
				id:  "Aatrox",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.fields.opts...)
			_, err := d.Champion(tt.args.ctx, tt.args.id, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Champion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataDragon_Champions(t *testing.T) {
	type fields struct {
		opts []opt
	}
	type args struct {
		ctx    context.Context
		params []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "All Champions",
			fields: fields{opts: []opt{WithVersion("12.4.1")}},
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.fields.opts...)
			_, err := d.Champions(tt.args.ctx, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Champions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDataDragon_ChampionsFull(t *testing.T) {
	type fields struct {
		opts []opt
	}
	type args struct {
		ctx    context.Context
		params []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "All Champions",
			fields: fields{opts: []opt{WithVersion("12.4.1")}},
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.fields.opts...)
			_, err := d.ChampionsFull(tt.args.ctx, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChampionsFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

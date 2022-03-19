//go:build integration

package datadragon

import (
	"context"
	"testing"
)

func TestDataDragon_ImageLoading(t *testing.T) {
	type args struct {
		ctx      context.Context
		champion string
		skin     uint8
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Success Aatrox",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     0,
			},
		},
		{
			name: "Success Aatrox other skin",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     1,
			},
		},
		{
			name: "Success cho'gath",
			args: args{
				ctx:      context.Background(),
				champion: "Chogath",
				skin:     0,
			},
		},
		{
			name: "Invalid no champion",
			args: args{
				ctx:      context.Background(),
				champion: "",
				skin:     0,
			},
			wantErr: ErrNoChampion,
		},
		{
			name: "Invalid other skin",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     255,
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			got, err := d.ImageLoading(tt.args.ctx, tt.args.champion, tt.args.skin)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImageLoading() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImageLoading() got = %v", got)
				return
			}
		})
	}
}

func TestDataDragon_ImagePassive(t *testing.T) {
	type fields struct {
		version *string
	}
	type args struct {
		ctx           context.Context
		passive       ChampionImage
		requestConfig []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Success Aatrox",
			args: args{
				ctx:     context.Background(),
				passive: ChampionImage{Full: "Aatrox_Passive.png"},
			},
		},
		{
			name: "Success cho'gath",
			args: args{
				ctx:     context.Background(),
				passive: ChampionImage{Full: "GreenTerror_TailSpike.png"},
			},
		},
		{
			name: "Success Aatrox other version",
			args: args{
				ctx:     context.Background(),
				passive: ChampionImage{Full: "Aatrox_Passive.png"},
				requestConfig: []RequestConfig{
					{Version: testVersion2},
				},
			},
		},
		{
			name: "Success Aatrox with locale",
			args: args{
				ctx:           context.Background(),
				passive:       ChampionImage{Full: "Aatrox_Passive.png"},
				requestConfig: []RequestConfig{{Locale: "testLocale2"}},
			},
		},
		{
			name: "Success custom version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:           context.Background(),
				passive:       ChampionImage{Full: "Aatrox_Passive.png"},
				requestConfig: []RequestConfig{{Version: testVersion2}},
			},
		},
		{
			name: "Invalid image",
			args: args{
				ctx:     context.Background(),
				passive: ChampionImage{Full: "aatrox.png"},
			},
			wantErr: ErrNotFound,
		},
		{
			name: "Invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:     context.Background(),
				passive: ChampionImage{Full: "Aatrox_Passive.png"},
			},
			wantErr: ErrNoVersion,
		},
		{
			name: "Invalid no champion",
			args: args{
				passive: ChampionImage{},
			},
			wantErr: ErrNoChampion,
		},
		{
			name: "Invalid no champion with invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				passive:       ChampionImage{},
				requestConfig: []RequestConfig{{Version: ""}},
			},
			wantErr: ErrNoChampion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			if tt.fields.version != nil {
				d.Version = *tt.fields.version
			}
			got, err := d.ImagePassive(tt.args.ctx, tt.args.passive, tt.args.requestConfig...)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImagePassive() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImagePassive() got = %v", got)
				return
			}
		})
	}
}

func TestDataDragon_ImageSpell(t *testing.T) {
	type fields struct {
		version *string
	}
	type args struct {
		ctx           context.Context
		spell         ChampionImage
		requestConfig []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Success Aatrox Q",
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "AatroxQ.png"},
			},
		},
		{
			name: "Success Aatrox W",
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "AatroxW.png"},
			},
		},
		{
			name: "Success Aatrox E",
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "AatroxE.png"},
			},
		},
		{
			name: "Success Aatrox R",
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "AatroxR.png"},
			},
		},
		{
			name: "Success custom version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:           context.Background(),
				spell:         ChampionImage{Full: "AatroxQ.png"},
				requestConfig: []RequestConfig{{Version: testVersion2}},
			},
		},
		{
			name: "Success Aatrox Q with locale",
			args: args{
				ctx:           context.Background(),
				spell:         ChampionImage{Full: "AatroxQ.png"},
				requestConfig: []RequestConfig{{Locale: "testLocale2"}},
			},
		},
		{
			name: "Invalid image",
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "aatrox.png"},
			},
			wantErr: ErrNotFound,
		},
		{
			name: "Invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:   context.Background(),
				spell: ChampionImage{Full: "AatroxQ.png"},
			},
			wantErr: ErrNoVersion,
		},
		{
			name: "Invalid no champion",
			args: args{
				spell: ChampionImage{},
			},
			wantErr: ErrNoSpell,
		},
		{
			name: "Invalid no champion with invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				spell:         ChampionImage{},
				requestConfig: []RequestConfig{{Version: ""}},
			},
			wantErr: ErrNoSpell,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			if tt.fields.version != nil {
				d.Version = *tt.fields.version
			}
			got, err := d.ImageSpell(tt.args.ctx, tt.args.spell, tt.args.requestConfig...)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImageSpell() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImageSpell() got = %v", got)
				return
			}
		})
	}
}

func TestDataDragon_ImageSplash(t *testing.T) {
	type fields struct {
		version *string
	}
	type args struct {
		ctx      context.Context
		champion string
		skin     uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Success Aatrox",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     0,
			},
		},
		{
			name: "Success Aatrox other skin",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     1,
			},
		},
		{
			name: "Success cho'gath",
			args: args{
				ctx:      context.Background(),
				champion: "Chogath",
				skin:     0,
			},
		},
		{
			name: "Invalid no champion",
			args: args{
				ctx:      context.Background(),
				champion: "",
				skin:     0,
			},
			wantErr: ErrNoChampion,
		},
		{
			name: "Invalid other skin",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				skin:     255,
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			if tt.fields.version != nil {
				d.Version = *tt.fields.version
			}
			got, err := d.ImageSplash(tt.args.ctx, tt.args.champion, tt.args.skin)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImageSplash() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImageSplash() got = %v", got)
				return
			}
		})
	}
}

func TestDataDragon_ImageChampionSquare(t *testing.T) {
	type fields struct {
		version *string
	}
	type args struct {
		ctx           context.Context
		champion      string
		requestConfig []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Success Aatrox",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
			},
		},
		{
			name: "Success cho'gath",
			args: args{
				ctx:      context.Background(),
				champion: "Chogath",
			},
		},
		{
			name: "Success Aatrox other version",
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
				requestConfig: []RequestConfig{
					{Version: testVersion2},
				},
			},
		},
		{
			name: "Success Aatrox with locale",
			args: args{
				ctx:           context.Background(),
				champion:      "Aatrox",
				requestConfig: []RequestConfig{{Locale: "testLocale2"}},
			},
		},
		{
			name: "Success custom version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:           context.Background(),
				champion:      "Aatrox",
				requestConfig: []RequestConfig{{Version: testVersion2}},
			},
		},
		{
			name: "Invalid image",
			args: args{
				ctx:      context.Background(),
				champion: "aatrox",
			},
			wantErr: ErrNotFound,
		},
		{
			name: "Invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:      context.Background(),
				champion: "Aatrox",
			},
			wantErr: ErrNoVersion,
		},
		{
			name: "Invalid no champion",
			args: args{
				champion: "",
			},
			wantErr: ErrNoChampion,
		},
		{
			name: "Invalid no champion with invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				champion:      "",
				requestConfig: []RequestConfig{{Version: ""}},
			},
			wantErr: ErrNoChampion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			if tt.fields.version != nil {
				d.Version = *tt.fields.version
			}
			got, err := d.ImageChampionSquare(tt.args.ctx, tt.args.champion, tt.args.requestConfig...)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImageChampionSquare() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImageChampionSquare() got = %v", got)
				return
			}
		})
	}
}

func TestDataDragon_ImageItem(t *testing.T) {
	type fields struct {
		version *string
	}
	type args struct {
		ctx           context.Context
		item          uint16
		requestConfig []RequestConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Success Boots of speed",
			args: args{
				ctx:  context.Background(),
				item: 1001,
			},
		},
		{
			name: "Success faerie charm",
			args: args{
				ctx:  context.Background(),
				item: 1004,
			},
		},
		{
			name: "Success Boots of speed other version",
			args: args{
				ctx:  context.Background(),
				item: 1001,
				requestConfig: []RequestConfig{
					{Version: testVersion2},
				},
			},
		},
		{
			name: "Success custom version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:           context.Background(),
				item:          1001,
				requestConfig: []RequestConfig{{Version: testVersion2}},
			},
		},
		{
			name: "Invalid image",
			args: args{
				ctx:  context.Background(),
				item: 0,
			},
			wantErr: ErrNotFound,
		},
		{
			name: "Invalid version",
			fields: fields{
				version: &emptyString,
			},
			args: args{
				ctx:  context.Background(),
				item: 1001,
			},
			wantErr: ErrNoVersion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithVersion(testVersion))
			if tt.fields.version != nil {
				d.Version = *tt.fields.version
			}
			got, err := d.ImageItem(tt.args.ctx, tt.args.item, tt.args.requestConfig...)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("ImageItem() error = %v, wantErr = %v", err, tt.wantErr)
					return
				}
				return
			}
			if got == nil {
				t.Errorf("ImageItem() got = %v", got)
				return
			}
		})
	}
}

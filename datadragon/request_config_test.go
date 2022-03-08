package datadragon

import (
	"reflect"
	"testing"
)

func TestDataDragon_mergeConfig(t *testing.T) {
	type fields struct {
		Locale  string
		Version string
	}
	type args struct {
		params     []RequestConfig
		validators []configValidator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RequestConfig
		wantErr error
	}{
		{
			name:    "Empty",
			want:    &RequestConfig{},
			wantErr: nil,
		},
		{
			name: "No RequestConfig",
			fields: fields{
				Locale:  "en_US",
				Version: testVersion,
			},
			want:    &RequestConfig{Version: testVersion, Locale: "en_US"},
			wantErr: nil,
		},
		{
			name: "No base",
			args: args{
				params: []RequestConfig{
					{Locale: "en_US", Version: testVersion},
				},
			},
			want:    &RequestConfig{Version: testVersion, Locale: "en_US"},
			wantErr: nil,
		},
		{
			name: "Same",
			fields: fields{
				Locale:  "en_US",
				Version: testVersion,
			},
			args: args{
				params: []RequestConfig{
					{Locale: "en_US", Version: testVersion},
				},
			},
			want:    &RequestConfig{Version: testVersion, Locale: "en_US"},
			wantErr: nil,
		},
		{
			name: "Different",
			fields: fields{
				Locale:  "en_US",
				Version: testVersion,
			},
			args: args{
				params: []RequestConfig{
					{Locale: "testLocale2", Version: testVersion2},
				},
			},
			want:    &RequestConfig{Version: testVersion2, Locale: "testLocale2"},
			wantErr: nil,
		},
		{
			name: "Ignore everything but the first",
			fields: fields{
				Locale:  "en_US",
				Version: testVersion,
			},
			args: args{
				params: []RequestConfig{
					{Locale: "testLocale2", Version: testVersion2},
					{Locale: "en_US", Version: testVersion},
					{Locale: "en_US", Version: testVersion},
				},
			},
			want:    &RequestConfig{Version: testVersion2, Locale: "testLocale2"},
			wantErr: nil,
		},
		{
			name:    "RequireVersion empty config",
			args:    args{validators: []configValidator{RequireVersion()}},
			want:    nil,
			wantErr: ErrNoVersion,
		},
		{
			name:    "RequireLocale empty config",
			args:    args{validators: []configValidator{RequireLocale()}},
			want:    nil,
			wantErr: ErrNoLocale,
		},
		{
			name:    "RequireVersion before RequireLocale empty config",
			args:    args{validators: []configValidator{RequireVersion(), RequireLocale()}},
			want:    nil,
			wantErr: ErrNoVersion,
		},
		{
			name: "RequireLocale before RequireVersion empty version",
			fields: fields{
				Locale:  "en_US",
				Version: "",
			},
			args:    args{validators: []configValidator{RequireLocale(), RequireVersion()}},
			want:    nil,
			wantErr: ErrNoVersion,
		},
		{
			name: "RequireVersion before RequireLocale empty locale",
			fields: fields{
				Locale:  "",
				Version: "en_US",
			},
			args:    args{validators: []configValidator{RequireVersion(), RequireLocale()}},
			want:    nil,
			wantErr: ErrNoLocale,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataDragon{
				Locale:  tt.fields.Locale,
				Version: tt.fields.Version,
			}
			got, err := d.mergeConfig(tt.args.params, tt.args.validators...)
			if err != tt.wantErr {
				t.Errorf("mergeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

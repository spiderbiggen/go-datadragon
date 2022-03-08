package datadragon

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDataDragon_InRealm(t *testing.T) {
	type fields struct {
		Client  *http.Client
		Region  Region
		Locale  string
		Version string
	}
	type args struct {
		realm *Realm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   DataDragon
	}{
		{
			name: "Apply version from datadragon",
			fields: fields{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale,
				Version: testVersion,
			},
			args: args{
				realm: &Realm{
					DataDragon: testVersion2,
				},
			},
			want: DataDragon{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale,
				Version: testVersion2,
			},
		},
		{
			name: "Apply locale from datadragon",
			fields: fields{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale,
				Version: testVersion,
			},
			args: args{
				realm: &Realm{
					Locale: testLocale2,
				},
			},
			want: DataDragon{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale2,
				Version: testVersion,
			},
		},
		{
			name: "Apply locale and version from datadragon",
			fields: fields{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale,
				Version: testVersion,
			},
			args: args{
				realm: &Realm{
					Locale:     testLocale2,
					DataDragon: testVersion2,
				},
			},
			want: DataDragon{
				Client:  &http.Client{Timeout: 10 * time.Second},
				Region:  NA1,
				Locale:  testLocale2,
				Version: testVersion2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DataDragon{
				Client:  tt.fields.Client,
				Region:  tt.fields.Region,
				Locale:  tt.fields.Locale,
				Version: tt.fields.Version,
			}
			d.InRealm(tt.args.realm)
			if !reflect.DeepEqual(d, tt.want) {
				t.Errorf("InRealm() d = %v", d)
			}
		})
	}
}

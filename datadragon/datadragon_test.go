package datadragon

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

const testVersion = "12.4.1"
const testVersion2 = "12.5.1"

const testLocale = "en_US"
const testLocale2 = "de_DE"

var emptyString = ""

func TestNew(t *testing.T) {
	type args struct {
		opts []opt
	}
	tests := []struct {
		name          string
		args          args
		want          DataDragon
		invalidRegion bool
	}{
		{
			name: "No opts",
			want: DataDragon{Client: &http.Client{}, Region: NA1, Locale: "en_US", Version: ""},
		},
		{
			name: "With version",
			args: args{
				opts: []opt{
					WithVersion(testVersion),
				},
			},
			want: DataDragon{Client: &http.Client{}, Region: NA1, Locale: testLocale, Version: testVersion},
		},
		{
			name: "With region",
			args: args{
				opts: []opt{
					WithRegion(EUW1),
				},
			},
			want: DataDragon{Client: &http.Client{}, Region: EUW1, Locale: testLocale, Version: ""},
		},
		{
			name: "With locale",
			args: args{
				opts: []opt{
					WithLocale(testLocale2),
				},
			},
			want: DataDragon{Client: &http.Client{}, Region: NA1, Locale: testLocale2, Version: ""},
		},
		{
			name: "With timeout",
			args: args{
				opts: []opt{
					WithTimeout(10 * time.Second),
				},
			},
			want: DataDragon{Client: &http.Client{Timeout: 10 * time.Second}, Region: NA1, Locale: testLocale, Version: ""},
		},
		{
			name: "With client",
			args: args{
				opts: []opt{
					WithClient(&http.Client{Timeout: 10 * time.Second}),
				},
			},
			want: DataDragon{Client: &http.Client{Timeout: 10 * time.Second}, Region: NA1, Locale: testLocale, Version: ""},
		},
		{
			name: "With invalid region",
			args: args{
				opts: []opt{
					WithRegion(Region(255)),
				},
			},
			want:          DataDragon{Client: &http.Client{}, Region: Region(255), Locale: testLocale, Version: ""},
			invalidRegion: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.opts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if got.Region.isValid() == tt.invalidRegion {
				t.Errorf("Region() = %v, invalidRegion %v", got.Region, tt.invalidRegion)
			}
		})
	}
}

package datadragon

import "testing"

func TestRegion_Realm(t *testing.T) {
	tests := []struct {
		name string
		r    Region
		want string
	}{
		{name: "BR1 is 'br'", r: BR1, want: "br"},
		{name: "EUN1 is 'eune'", r: EUN1, want: "eune"},
		{name: "EUW1 is 'euw'", r: EUW1, want: "euw"},
		{name: "JP1 is 'jp'", r: JP1, want: "jp"},
		{name: "KR is 'kr'", r: KR, want: "kr"},
		{name: "LA1 is 'lan'", r: LA1, want: "lan"},
		{name: "LA2 is 'las'", r: LA2, want: "las"},
		{name: "NA1 is 'na'", r: NA1, want: "na"},
		{name: "OC1 is 'oce'", r: OC1, want: "oce"},
		{name: "TR1 is 'tr'", r: TR1, want: "tr"},
		{name: "RU is 'ru'", r: RU, want: "ru"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Realm(); got != tt.want {
				t.Errorf("Realm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegion_isValid(t *testing.T) {
	tests := []struct {
		name string
		r    Region
		want bool
	}{
		{name: "BR1", r: BR1, want: true},
		{name: "EUN1", r: EUN1, want: true},
		{name: "EUW1", r: EUW1, want: true},
		{name: "JP1", r: JP1, want: true},
		{name: "KR", r: KR, want: true},
		{name: "LA1", r: LA1, want: true},
		{name: "LA2", r: LA2, want: true},
		{name: "NA1", r: NA1, want: true},
		{name: "OC1", r: OC1, want: true},
		{name: "TR1", r: TR1, want: true},
		{name: "RU", r: RU, want: true},
		{name: "Nothing", r: Region(255), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.isValid(); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

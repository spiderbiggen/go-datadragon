package datadragon

type Region uint8

const (
	// NA1 is the North American server
	NA1 Region = iota
	// EUW1 is the EU West server
	EUW1
	// EUN1 is the EU Nordic & East server
	EUN1
	// BR1 is the Brazilian server
	BR1
	// JP1 is the Japanese server
	JP1
	// KR is the Korean server
	KR
	// LA1 is the Northern Latin American server
	LA1
	// LA2 is the Southern Latin American server
	LA2
	// OC1 is the Oceania server
	OC1
	// TR1 is the Turkish server
	TR1
	// RU is the Russian server
	RU
)

var realms = map[Region]string{
	BR1:  "br",
	EUN1: "eune",
	EUW1: "euw",
	JP1:  "jp",
	KR:   "kr",
	LA1:  "lan",
	LA2:  "las",
	NA1:  "na",
	OC1:  "oce",
	TR1:  "tr",
	RU:   "ru",
}

func (r Region) isValid() bool {
	_, found := realms[r]
	return found
}

func (r Region) Realm() string {
	val, _ := realms[r]
	return val
}

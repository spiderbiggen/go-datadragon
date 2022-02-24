package datadragon

type Realm struct {
	N struct {
		Item        string `json:"item"`
		Rune        string `json:"rune"`
		Mastery     string `json:"mastery"`
		Summoner    string `json:"summoner"`
		Champion    string `json:"champion"`
		ProfileIcon string `json:"profileicon"`
		Map         string `json:"map"`
		Language    string `json:"language"`
		Sticker     string `json:"sticker"`
	} `json:"n"`
	Version        string `json:"v"`
	Locale         string `json:"l"`
	Cdn            string `json:"cdn"`
	DataDragon     string `json:"dd"`
	Lg             string `json:"lg"`
	Css            string `json:"css"`
	ProfileIconMax int    `json:"profileiconmax"`
	Store          string `json:"store,omitempty"`
}

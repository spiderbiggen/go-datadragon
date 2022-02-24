package datadragon

import (
	"encoding/json"
	"errors"
	"fmt"
)

type championsResponse struct {
	Type    string              `json:"type"`
	Format  string              `json:"format"`
	Version string              `json:"version"`
	Data    map[string]Champion `json:"database"`
}

type Champion struct {
	Version string        `json:"version"`
	Id      string        `json:"id"`
	Key     uint16        `json:"key,string"`
	Name    string        `json:"name"`
	Title   string        `json:"title"`
	Blurb   string        `json:"blurb"`
	Info    ChampionInfo  `json:"info"`
	Image   ChampionImage `json:"image"`
	Lore    *string       `json:"lore"`
	Tags    []string      `json:"tags"`
	Partype string        `json:"partype"`
	Stats   ChampionStats `json:"stats"`
	// These are only available for the championFull request
	Skins     []ChampionSkin   `json:"skins"`
	AllyTips  []string         `json:"allytips"`
	EnemyTips []string         `json:"enemytips"`
	Spells    []ChampionSpell  `json:"spells"`
	Passive   *ChampionPassive `json:"passive"`
}

type ChampionInfo struct {
	Attack     uint8 `json:"attack"`
	Defense    uint8 `json:"defense"`
	Magic      uint8 `json:"magic"`
	Difficulty uint8 `json:"difficulty"`
}

type ChampionImage struct {
	Full    string `json:"full"`
	Sprite  string `json:"sprite"`
	Group   string `json:"group"`
	OffsetX uint16 `json:"x"`
	OffsetY uint16 `json:"y"`
	Width   uint8  `json:"w"`
	Height  uint8  `json:"h"`
}

type ChampionSkin struct {
	Id      uint16 `json:"id,string"`
	Num     uint8  `json:"num"`
	Name    string `json:"name"`
	Chromas bool   `json:"chromas"`
}

type ChampionStats struct {
	Hp                   float32 `json:"hp"`
	HpPerLevel           float32 `json:"hpperlevel"`
	Mp                   float32 `json:"mp"`
	MpPerLevel           float32 `json:"mpperlevel"`
	MoveSpeed            float32 `json:"movespeed"`
	Armor                float32 `json:"armor"`
	ArmorPerLevel        float32 `json:"armorperlevel"`
	SpellBlock           float32 `json:"spellblock"`
	SpellBlockPerLevel   float32 `json:"spellblockperlevel"`
	AttackRange          float32 `json:"attackrange"`
	HpRegen              float32 `json:"hpregen"`
	HpRegenPerLevel      float32 `json:"hpregenperlevel"`
	MpRegen              float32 `json:"mpregen"`
	MpRegenPerLevel      float32 `json:"mpregenperlevel"`
	Crit                 float32 `json:"crit"`
	CritPerLevel         float32 `json:"critperlevel"`
	AttackDamage         float32 `json:"attackdamage"`
	AttackDamagePerLevel float32 `json:"attackdamageperlevel"`
	AttackSpeedPerLevel  float32 `json:"attackspeedperlevel"`
	AttackSpeed          float32 `json:"attackspeed"`
}

type ChampionSpell struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tooltip     string `json:"tooltip"`
	Leveltip    struct {
		Label  []string `json:"label"`
		Effect []string `json:"effect"`
	} `json:"leveltip"`
	Maxrank      int           `json:"maxrank"`
	Cooldown     []float32     `json:"cooldown"`
	CooldownBurn string        `json:"cooldownBurn"`
	Cost         []int         `json:"cost"`
	CostBurn     string        `json:"costBurn"`
	Effect       [][]float32   `json:"effect"`
	EffectBurn   []string      `json:"effectBurn"`
	CostType     string        `json:"costType"`
	Maxammo      string        `json:"maxammo"`
	Range        []int         `json:"range"`
	RangeBurn    string        `json:"rangeBurn"`
	Image        ChampionImage `json:"image"`
	Resource     string        `json:"resource"`
}

type ChampionPassive struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       ChampionImage `json:"image"`
}

func (d *DataDragon) Champions(params ...RequestConfig) (champions map[string]Champion, err error) {
	config := d.mergeConfig(params)

	if config.Version == "" {
		err = errors.New("no api version specified") // TODO define errors
		return
	}

	resp, err := d.cdnRequest(fmt.Sprintf("%s/database/%s/champion.json", config.Version, config.Locale))
	if err != nil {
		return
	}
	defer d.closeBody(resp.Body)
	var championResponse championsResponse
	err = json.NewDecoder(resp.Body).Decode(&championResponse)
	return championResponse.Data, err
}

func (d *DataDragon) ChampionsFull(params ...RequestConfig) (champions map[string]Champion, err error) {
	config := d.mergeConfig(params)

	if config.Version == "" {
		err = errors.New("no api version specified") // TODO define errors
		return
	}

	resp, err := d.cdnRequest(fmt.Sprintf("%s/database/%s/championFull.json", config.Version, config.Locale))
	if err != nil {
		return
	}
	defer d.closeBody(resp.Body)
	var championResponse championsResponse
	err = json.NewDecoder(resp.Body).Decode(&championResponse)
	return championResponse.Data, err
}

func (d *DataDragon) Champion(id string, params ...RequestConfig) (champion Champion, err error) {
	if id == "" {
		err = errors.New("no id specified") // TODO define errors
		return
	}

	config := d.mergeConfig(params)

	if config.Version == "" {
		err = errors.New("no api version specified") // TODO define errors
		return
	}

	resp, err := d.cdnRequest(fmt.Sprintf("%s/database/%s/champion/%s.json", config.Version, config.Locale, id))
	if err != nil {
		return
	}
	defer d.closeBody(resp.Body)
	var championsResponse championsResponse
	if err = json.NewDecoder(resp.Body).Decode(&championsResponse); err != nil {
		return
	}
	champion, found := championsResponse.Data[id]
	if !found {
		err = errors.New("champion not found")
		return
	}
	return
}

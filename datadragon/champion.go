package datadragon

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNoId             = errors.New("no id specified")
	ErrChampionNotFound = errors.New("champion not found")
)

type championsResponse struct {
	Type    string              `json:"type"`
	Format  string              `json:"format"`
	Version string              `json:"Version"`
	Data    map[string]Champion `json:"data"`
}
type championsFullResponse struct {
	Type    string                  `json:"type"`
	Format  string                  `json:"format"`
	Version string                  `json:"Version"`
	Data    map[string]ChampionFull `json:"data"`
}

type Champion struct {
	Version string        `json:"Version"`
	Id      string        `json:"id"`
	Key     uint16        `json:"key,string"`
	Name    string        `json:"name"`
	Title   string        `json:"title"`
	Blurb   string        `json:"blurb"`
	Image   ChampionImage `json:"image"`
	Tags    []string      `json:"tags"`
	Partype string        `json:"partype"`
	Stats   ChampionStats `json:"stats"`
}

type ChampionFull struct {
	Champion
	Info      ChampionInfo    `json:"info"`
	Lore      string          `json:"lore"`
	Skins     []ChampionSkin  `json:"skins"`
	AllyTips  []string        `json:"allytips"`
	EnemyTips []string        `json:"enemytips"`
	Spells    []ChampionSpell `json:"spells"`
	Passive   ChampionPassive `json:"passive"`
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
	Id      uint32 `json:"id,string"`
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
	CritChance           float32 `json:"crit"`
	CritChancePerLevel   float32 `json:"critperlevel"`
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
	LevelTip    struct {
		Label  []string `json:"label"`
		Effect []string `json:"effect"`
	} `json:"leveltip"`
	MaxRank      int           `json:"maxrank"`
	CoolDown     []float32     `json:"cooldown"`
	CoolDownBurn string        `json:"cooldownBurn"`
	Cost         []int         `json:"cost"`
	CostBurn     string        `json:"costBurn"`
	Effect       [][]float32   `json:"effect"`
	EffectBurn   []string      `json:"effectBurn"`
	CostType     string        `json:"costType"`
	MaxAmmo      string        `json:"maxammo"`
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

func (d DataDragon) Champions(ctx context.Context, params ...RequestConfig) (map[string]Champion, error) {
	config, err := d.mergeConfig(params, RequireVersion(), RequireLocale())
	if err != nil {
		return nil, err
	}

	var championResponse championsResponse
	err = d.cdnJson(ctx, fmt.Sprintf("%s/data/%s/champion.json", config.Version, config.Locale), &championResponse)
	if err != nil {
		return nil, err
	}
	return championResponse.Data, nil
}

func (d DataDragon) ChampionsFull(ctx context.Context, params ...RequestConfig) (map[string]ChampionFull, error) {
	config, err := d.mergeConfig(params, RequireVersion(), RequireLocale())
	if err != nil {
		return nil, err
	}

	var championResponse championsFullResponse
	err = d.cdnJson(ctx, fmt.Sprintf("%s/data/%s/championFull.json", config.Version, config.Locale), &championResponse)
	if err != nil {
		return nil, err
	}
	return championResponse.Data, nil
}

func (d *DataDragon) Champion(ctx context.Context, id string, params ...RequestConfig) (*ChampionFull, error) {
	if id == "" {
		return nil, ErrNoId
	}

	config, err := d.mergeConfig(params, RequireVersion())
	if err != nil {
		return nil, err
	}

	var championsResponse championsFullResponse
	err = d.cdnJson(ctx, fmt.Sprintf("%s/data/%s/champion/%s.json", config.Version, config.Locale, id), &championsResponse)
	if err != nil {
		if err == ErrNotFound {
			return nil, ErrChampionNotFound
		}
		return nil, err
	}
	champion, found := championsResponse.Data[id]
	if !found {
		return nil, ErrChampionNotFound
	}
	return &champion, nil
}

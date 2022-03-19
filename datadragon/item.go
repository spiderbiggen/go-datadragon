package datadragon

import (
	"context"
	"fmt"
	"strconv"
)

type itemResponse struct {
	Type    string          `json:"type"`
	Version string          `json:"version"`
	Data    map[string]Item `json:"data"`
	Groups  []ItemGroup     `json:"groups"`
	Tree    []ItemTree      `json:"tree"`
}

type ItemResponse struct {
	Items  map[uint16]Item
	Groups []ItemGroup
	Tree   []ItemTree
}

func (r ItemResponse) ItemSlice() []Item {
	result := make([]Item, 0, len(r.Items))
	for _, item := range r.Items {
		result = append(result, item)
	}
	return result
}

type Item struct {
	Name string `json:"name"`
	// ID is the item map key
	ID string `json:"-"`
	// Key is the item map key as a number
	Key         uint16 `json:"-"`
	Group       string `json:"group"`
	Description string `json:"description"`
	Colloq      string `json:"colloq"`
	Plaintext   string `json:"plaintext"`
	Gold        struct {
		Base        uint16 `json:"base"`
		Total       uint16 `json:"total"`
		Sell        uint16 `json:"sell"`
		Purchasable bool   `json:"purchasable"`
	} `json:"gold"`
	Consumed         bool     `json:"consumed"`
	Stacks           int      `json:"stacks"`
	Depth            uint8    `json:"depth"`
	ConsumeOnFull    bool     `json:"consumeOnFull"`
	From             []string `json:"from"`
	Into             []string `json:"into"`
	SpecialRecipe    int      `json:"specialRecipe"`
	InStore          bool     `json:"inStore"`
	HideFromAll      bool     `json:"hideFromAll"`
	RequiredChampion string   `json:"requiredChampion"`
	RequiredAlly     string   `json:"requiredAlly"`
	Stats            struct {
		FlatHPPoolMod         *uint16 `json:"FlatHPPoolMod,omitempty"`
		FlatMPPoolMod         *uint16 `json:"FlatMPPoolMod,omitempty"`
		FlatMPRegenMod        *uint16 `json:"FlatMPRegenMod,omitempty"`
		FlatArmorMod          *uint16 `json:"FlatArmorMod,omitempty"`
		FlatPhysicalDamageMod *uint16 `json:"FlatPhysicalDamageMod,omitempty"`
		FlatMagicDamageMod    *uint16 `json:"FlatMagicDamageMod,omitempty"`
		FlatMovementSpeedMod  *uint16 `json:"FlatMovementSpeedMod,omitempty"`
		FlatAttackSpeedMod    *uint16 `json:"FlatAttackSpeedMod,omitempty"`
		FlatCritDamageMod     *uint16 `json:"FlatCritDamageMod,omitempty"`
		FlatBlockMod          *uint16 `json:"FlatBlockMod,omitempty"`
		FlatSpellBlockMod     *uint16 `json:"FlatSpellBlockMod,omitempty"`
		FlatEXPBonus          *uint16 `json:"FlatEXPBonus,omitempty"`
		FlatEnergyRegenMod    *uint16 `json:"FlatEnergyRegenMod,omitempty"`
		FlatEnergyPoolMod     *uint16 `json:"FlatEnergyPoolMod,omitempty"`
		// Special Flat
		FlatHPRegenMod    *float32 `json:"FlatHPRegenMod,omitempty"`
		FlatCritChanceMod *float32 `json:"FlatCritChanceMod,omitempty"`
		// Percentages
		PercentHPPoolMod         *uint16  `json:"PercentHPPoolMod,omitempty"`
		PercentMPPoolMod         *uint16  `json:"PercentMPPoolMod,omitempty"`
		PercentHPRegenMod        *float32 `json:"PercentHPRegenMod,omitempty"`
		PercentMPRegenMod        *uint16  `json:"PercentMPRegenMod,omitempty"`
		PercentArmorMod          *uint16  `json:"PercentArmorMod,omitempty"`
		PercentPhysicalDamageMod *uint16  `json:"PercentPhysicalDamageMod,omitempty"`
		PercentMagicDamageMod    *uint16  `json:"PercentMagicDamageMod,omitempty"`
		PercentMovementSpeedMod  *float32 `json:"PercentMovementSpeedMod,omitempty"`
		PercentAttackSpeedMod    *float32 `json:"PercentAttackSpeedMod,omitempty"`
		PercentDodgeMod          *uint16  `json:"PercentDodgeMod,omitempty"`
		PercentCritChanceMod     *uint16  `json:"PercentCritChanceMod,omitempty"`
		PercentCritDamageMod     *uint16  `json:"PercentCritDamageMod,omitempty"`
		PercentBlockMod          *uint16  `json:"PercentBlockMod,omitempty"`
		PercentSpellBlockMod     *uint16  `json:"PercentSpellBlockMod,omitempty"`
		PercentEXPBonus          *uint16  `json:"PercentEXPBonus,omitempty"`
		PercentLifeStealMod      *float32 `json:"PercentLifeStealMod,omitempty"`
		PercentSpellVampMod      *uint16  `json:"PercentSpellVampMod,omitempty"`
	} `json:"stats"`
	Effect map[string]string `json:"effect"`
	Tags   []string          `json:"tags"`
	Maps   map[string]bool   `json:"maps"`
}

type ItemGroup struct {
	ID              string `json:"id"`
	MaxGroupOwnable int8   `json:"MaxGroupOwnable,string"`
}

type ItemTree struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}

func (d DataDragon) Items(ctx context.Context, params ...RequestConfig) (*ItemResponse, error) {
	config, err := d.mergeConfig(params, RequireVersion(), RequireLocale())
	if err != nil {
		return nil, err
	}

	var response itemResponse
	err = d.cdnJson(ctx, fmt.Sprintf("%s/data/%s/item.json", config.Version, config.Locale), &response)
	if err != nil {
		return nil, err
	}

	result := ItemResponse{
		Items:  make(map[uint16]Item, len(response.Data)),
		Groups: response.Groups,
		Tree:   response.Tree,
	}
	for s, item := range response.Data {
		i, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return nil, err
		}

		item.ID = s
		item.Key = uint16(i)
		result.Items[item.Key] = item
	}
	return &result, nil
}

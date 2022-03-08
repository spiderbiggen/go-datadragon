package datadragon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidRegion = errors.New("invalid region")
)

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
}

// InRealm applies values from the given Realm to the current DataDragon instance.
func (d *DataDragon) InRealm(realm *Realm) {
	if realm == nil {
		return
	}
	if realm.Locale != "" {
		d.Locale = realm.Locale
	}
	if realm.DataDragon != "" {
		d.Version = realm.DataDragon
	}
}

func (d *DataDragon) Realm(ctx context.Context, optRegion ...Region) (*Realm, error) {
	region := d.Region
	if len(optRegion) > 0 {
		if !optRegion[0].isValid() {
			return nil, ErrInvalidRegion
		}
		region = optRegion[0]
	}
	if !region.isValid() {
		return nil, ErrInvalidRegion
	}
	resp, err := d.realmsRequest(ctx, fmt.Sprintf("%s.json", region.Realm()))
	if err != nil {
		return nil, err
	}
	defer d.closeBody(resp.Body)

	var realm Realm
	err = json.NewDecoder(resp.Body).Decode(&realm)
	if err != nil {
		return nil, err
	}
	return &realm, nil
}

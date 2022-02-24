package datadragon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (d *DataDragon) Versions() (versions []string, err error) {
	resp, err := d.apiRequest("versions.json")
	if err != nil {
		return nil, err
	}
	defer d.closeBody(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&versions)
	return
}

func (d *DataDragon) Languages() (languages []string, err error) {
	resp, err := d.cdnRequest("languages.json")
	if err != nil {
		return
	}
	defer d.closeBody(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&languages)
	return
}

func (d *DataDragon) Realm(optRegion ...Region) (realm Realm, err error) {
	region := d.region
	if len(optRegion) > 0 && optRegion[0].isValid() {
		region = optRegion[0]
	}
	resp, err := d.realmsRequest(fmt.Sprintf("%s.json", region.Realm()))
	if err != nil {
		return
	}
	defer d.closeBody(resp.Body)

	if err = json.NewDecoder(resp.Body).Decode(&realm); err != nil {
		buf, err2 := ioutil.ReadAll(resp.Body)
		log.Println(buf, err2)
		return
	}
	return
}

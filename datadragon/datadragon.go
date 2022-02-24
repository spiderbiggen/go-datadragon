package datadragon

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DataDragon struct {
	apiUrl    string
	realmsUrl string
	cdnUrl    string
	client    *http.Client
	version   string
	locale    string
	region    Region
}

type opt func(d *DataDragon)

// WithTimeout sets the timeout for the http client.
func WithTimeout(timeout time.Duration) func(d *DataDragon) {
	return func(d *DataDragon) { d.client.Timeout = timeout }
}

// WithRegion sets the default region for the DataDragon API
func WithRegion(region Region) func(d *DataDragon) {
	return func(d *DataDragon) {
		if region.isValid() {
			d.region = region
		}
	}
}

// WithLanguage applies the given string as locale.
// If validate is true it will check against the locales on Data Dragon
func WithLanguage(language string, validate ...bool) func(d *DataDragon) {
	return func(d *DataDragon) {
		if len(validate) > 0 && validate[0] {
			locales, err := d.Languages()
			if err != nil {
				return
			}
			for _, s := range locales {
				if s == language {
					d.locale = language
				}
			}
			return
		}
		d.locale = language
	}
}

// New creates a DataDragon instance with sensible defaults. Opts can be used to override the default values.
// Provided opts:
//   WithLanguage
//   WithRegion
//   WithTimeout
func New(opts ...opt) DataDragon {
	api := DataDragon{
		apiUrl:    "https://ddragon.leagueoflegends.com/api",
		realmsUrl: "https://ddragon.leagueoflegends.com/realms",
		cdnUrl:    "https://ddragon.leagueoflegends.com/cdn",
		client: &http.Client{
			Timeout: time.Second * 15,
		},
		region: NA1,
		locale: "en_US",
	}

	for _, apply := range opts {
		apply(&api)
	}

	return api
}

// Version will only be set if ApplyRealm was previously called with either of UseDataDragonVersion or UseClientVersion
func (d *DataDragon) Version() string {
	return d.version
}

type realmOpt func(d *DataDragon, r *Realm)

// UseLanguage applies the language from the Realm to the DataDragon instance
func UseLanguage() func(d *DataDragon, r *Realm) {
	return func(d *DataDragon, r *Realm) {
		if r.Locale != "" {
			d.locale = r.Locale
		}
	}
}

// UseDataDragonVersion sets the default api version to the DataDragon version of the given Realm
func UseDataDragonVersion() func(d *DataDragon, r *Realm) {
	return func(d *DataDragon, r *Realm) {
		if r.DataDragon != "" {
			d.version = r.DataDragon
		}
	}
}

// UseClientVersion sets the default api version to the client version of the given Realm
func UseClientVersion() func(d *DataDragon, r *Realm) {
	return func(d *DataDragon, r *Realm) {
		if r.Version != "" {
			d.version = r.Version
		}
	}
}

// ApplyRealm applies values from the given Realm to the current DataDragon instance.
// opts can be used to override more than just the cdnUrl.
// Provided opts:
//   UseLanguage
//   UseDataDragonVersion
//   UseClientVersion
func (d *DataDragon) ApplyRealm(realm *Realm, opts ...realmOpt) {
	if realm == nil {
		return
	}
	if realm.Cdn != "" {
		_, err := url.Parse(realm.Cdn)
		if err == nil {
			d.cdnUrl = realm.Cdn
		}
	}
	for _, apply := range opts {
		apply(d, realm)
	}
}

type RequestConfig struct {
	Version string
	Locale  string
}

func (d *DataDragon) mergeConfig(params []RequestConfig) RequestConfig {
	config := RequestConfig{
		Version: d.version,
		Locale:  d.locale,
	}
	if len(params) > 0 {
		param := params[0]
		if param.Version != "" {
			config.Version = param.Version
		}
		if param.Locale != "" {
			config.Locale = param.Locale
		}
	}
	return config
}

func (d *DataDragon) apiRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, d.apiUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(req)
}

func (d *DataDragon) realmsRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, d.realmsUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(req)
}

func (d *DataDragon) cdnRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, d.cdnUrl+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(req)
}

func (d *DataDragon) closeBody(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		// TODO better logging
		log.Println(err)
	}
}

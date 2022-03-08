package datadragon

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseUrl   = "https://ddragon.leagueoflegends.com"
	apiUrl    = baseUrl + "/api"
	realmsUrl = baseUrl + "/realms"
	cdnUrl    = baseUrl + "/cdn"
)

type DataDragon struct {
	Client  *http.Client
	Region  Region
	Locale  string
	Version string
}

type opt func(d *DataDragon)

// WithTimeout sets the timeout for the http client.
func WithTimeout(timeout time.Duration) func(d *DataDragon) {
	return func(d *DataDragon) { d.Client.Timeout = timeout }
}

// WithClient sets the http client.
func WithClient(client *http.Client) func(d *DataDragon) {
	return func(d *DataDragon) { d.Client = client }
}

// WithRegion sets the default Region for the DataDragon API
func WithRegion(region Region) func(d *DataDragon) {
	return func(d *DataDragon) { d.Region = region }
}

// WithLocale applies the given string as Locale.
func WithLocale(locale string) func(d *DataDragon) {
	return func(d *DataDragon) { d.Locale = locale }
}

// WithVersion sets the default version for the http client.
func WithVersion(version string) func(d *DataDragon) {
	return func(d *DataDragon) { d.Version = version }
}

// New creates a DataDragon instance with some defaults. Opts can be used to override the default values.
// Provided opts:
//   WithLocale
//   WithRegion
//   WithTimeout
func New(opts ...opt) DataDragon {
	api := DataDragon{
		Client: &http.Client{},
		Region: NA1,
		Locale: "en_US",
	}

	for _, apply := range opts {
		apply(&api)
	}

	return api
}

func (d DataDragon) closeBody(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		// TODO better logging
		log.Println(err)
	}
}

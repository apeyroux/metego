/*
API: https://openmeteoforecast.org/wiki/API/0.1
API: http://wiki.openstreetmap.org/wiki/Nominatim

ex: "http://nominatim.openstreetmap.org/search?q=bessans&format=json&limit=1"
ex: "http://api.ometfn.net/0.1/forecast/eu12/44.5,6.32/now.json"
*/

package metego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL_BASE_NOMINATIM = "http://nominatim.openstreetmap.org"
	URL_BASE_OMET      = "http://api.ometfn.net"
)

type Nominatim struct {
	Place_id     string
	Licence      string
	Osm_type     string
	Osm_id       string
	Boundingbox  []string
	Lat          string
	Lon          string
	Display_name string
	Class        string
	Type         string
	Importance   string
	Icon         string
}

type Grid struct {
	x       int
	y       int
	x_error int
	y_error int
}

type OMet struct {
	Doc                   string
	License               string
	Domain                string
	Run                   string
	Grid                  Grid
	Ntimes                int
	Times                 []int
	Temp                  []int
	Rh                    []int
	Low_clouds            []int
	Medium_clouds         []int
	High_clouds           []int
	Precipitations        []int
	Pblh                  []int
	Pressure              []int
	Wind_10m_ground_speed []int
	Wind_10m_ground_dir   []int
	Wind_1000m_msl_speed  []int
	Wind_1000m_msl_dir    []int
	Wind_2000m_msl_speed  []int
	Wind_2000m_msl_dir    []int
	Wind_3000m_msl_speed  []int
	Wind_3000m_msl_dir    []int
	Wind_4000m_msl_speed  []int
	Wind_4000m_msl_dir    []int
	Status                string
	Msg                   string
	Srv                   string
}

func NewNominatim(q string) (*Nominatim, error) {
	var nr []Nominatim

	urlws := fmt.Sprintf("%s/search?q=%s&format=json&limit=1", URL_BASE_NOMINATIM, q)

	resp, err := http.Get(urlws)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &nr)

	return &nr[0], _
}

func NewOMet(nominatim *Nominatim) (*OMet, error) {
	var or OMet

	urlws := fmt.Sprintf("%s/0.1/forecast/eu12/%s,%s/now.json", URL_BASE_OMET, nominatim.Lat, nominatim.Lon)

	resp, err := http.Get(urlws)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &or)

	return &or, nil
}

package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Village        string `json:"village"`
	Neighbourhood  string `json:"neighbourhood"`
	Suburb         string `json:"suburb"`
	Borough        string `json:"borough"`
	CityDistrict   string `json:"city_district"`
	Administrative string `json:"administrative"`
}

type ReverseGeocodeResponse struct {
	Address Address `json:"address"`
}

// ReverseGeocode melakukan reverse geocoding untuk mendapatkan Kecamatan dan Kelurahan
func ReverseGeocode(lat float64, lon float64) (string, string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&format=json&zoom=18", lat, lon)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("User-Agent", "YourAppName")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch reverse geocoding data, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var result ReverseGeocodeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", err
	}

	address := result.Address
	kelurahan := address.Village
	if kelurahan == "" {
		kelurahan = address.Neighbourhood
	}
	if kelurahan == "" {
		kelurahan = address.Suburb
	}

	kecamatan := address.Borough
	if kecamatan == "" {
		kecamatan = address.CityDistrict
	}
	if kecamatan == "" {
		kecamatan = address.Administrative
	}
	if kelurahan == address.Suburb && kelurahan != "" {
		kecamatan = address.Suburb
	}
	if len(kelurahan) > 0 && (kelurahan[0] == 'r' || kelurahan[0] == 'R') {
		kelurahan = address.Suburb
	}

	return kecamatan, kelurahan, nil
}

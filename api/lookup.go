package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchManufacturers(web *http.Client, host string, year int) ([]string, error) {
	url := fmt.Sprintf("%s/lookup/manufacturers/%v", host, year)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	var result []string
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchModels(web *http.Client, host string, year int, manufacturer string) ([]string, error) {
	url := fmt.Sprintf("%s/lookup/models/%v/%s", host, year, manufacturer)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	var result []string
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchTrims(web *http.Client, host string, year int, manufacturer, model string) ([]string, error) {
	url := fmt.Sprintf("%s/lookup/trim/%v/%s/%s", host, year, manufacturer, model)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	var result []string
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

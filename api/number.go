package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/vin/core"
	"net/http"
)

func ValidateVIN(web *http.Client, host, vin string) (bool, error) {
	url := fmt.Sprintf("%s/validate/%s", host, vin)
	resp, err := web.Get(url)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	result := false
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func LookupVIN(web *http.Client, host, vin string) (core.VIN, error) {
	url := fmt.Sprintf("%s/lookup/%s", host, vin)
	resp, err := web.Get(url)

	if err != nil {
		return core.VIN{}, err
	}

	defer resp.Body.Close()

	result := core.VIN{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

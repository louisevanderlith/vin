package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/vin/core"
	"io/ioutil"
	"net/http"
)

func ValidateVIN(web *http.Client, host, vin string) (bool, error) {
	url := fmt.Sprintf("%s/validate/%s", host, vin)
	resp, err := web.Get(url)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return false, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := false
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func LookupVIN(web *http.Client, host, vin string) (hsk.Record, error) {
	url := fmt.Sprintf("%s/lookup/%s", host, vin)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewRecord(core.VIN{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

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

func FetchVIN(web *http.Client, host string, k hsk.Key) (core.VIN, error) {
	url := fmt.Sprintf("%s/admin/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.VIN{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return core.VIN{}, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := core.VIN{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchLatestVINs(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/admin/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewResultPage(core.VIN{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

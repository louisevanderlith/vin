package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/vin/core"
	"net/http"
)

func FetchRegion(web *http.Client, host string, k hsk.Key) (core.Region, error) {
	url := fmt.Sprintf("%s/regions/%s", host, k.String())
	resp, err := web.Get(url)

	if err != nil {
		return core.Region{}, err
	}

	defer resp.Body.Close()

	result := core.Region{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

func FetchRegions(web *http.Client, host, pagesize string) (records.Page, error) {
	url := fmt.Sprintf("%s/regions/%s", host, pagesize)
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := records.NewResultPage(core.Region{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}

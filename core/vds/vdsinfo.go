package vds

import (
	"fmt"
	"log"
)

type VDSAnalyzer func(vds string, obj *VDSInfo) (interface{}, error)

type VDSInfo struct {
	Code string //6 Characters of the VDS
}

var analyzers map[string]VDSAnalyzer

func init() {
	analyzers = make(map[string]VDSAnalyzer)
	analyzers["BMW"] = AnalyseBMW
	analyzers["Toyota"] = AnalyseToyota
}

func FindVDSInfo(make string, unique string, years []int) (interface{}, error) {
	vdsStr := unique[3:8]

	result := &VDSInfo{Code: vdsStr}
	analyzer, ok := analyzers[make]

	if !ok {
		return nil, fmt.Errorf("no analyzer found for %s", make)
	}

	tmp, err := analyzer(vdsStr, result)

	if err != nil {
		return nil, err
	}

	log.Printf("%v\n", tmp)
	return nil, nil
}

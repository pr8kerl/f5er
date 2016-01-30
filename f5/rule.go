package f5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBRawValues struct {
	VerificationStatus string `json:"verificationStatus"`
}

type LBRule struct {
	Name         string      `json:"name"`
	Partition    string      `json:"partition"`
	Fullpath     string      `json:"fullPath"`
	Generation   int         `json:"generation"`
	ApiAnonymous string      `json:"apiAnonymous"`
	ApiRawValues LBRawValues `json:"apiRawValues"`
}

type LBRules struct {
	Items []LBRule `json:"items"`
}

func (f *Device) ShowRules() (error, *LBRules) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/rule"
	res := LBRules{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowRule(rname string) (error, *LBRule) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/rule/" + rule
	res := LBRule{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddRule() (error, *LBRule) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/rule"
	res := LBRule{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// post the request
	err, resp := f.sendRequest(u, POST, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdateRule(rname string) (error, *LBRule) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/rule/" + rule
	res := LBRule{}
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// put the request
	err, resp := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeleteRule(rname string) (error, *Response) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/rule/" + rule
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

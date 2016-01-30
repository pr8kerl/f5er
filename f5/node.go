package f5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBNodeFQDN struct {
	AddressFamily string `json:"addressFamily"`
	AutoPopulate  string `json:"autopopulate"`
	DownInterval  int    `json:"downInterval"`
	Interval      int    `json:"interval"`
}

type LBNode struct {
	Name            string     `json:"name"`
	Partition       string     `json:"partition"`
	FullPath        string     `json:"fullPath"`
	Generation      int        `json:"generation"`
	Address         string     `json:"address,omitEmpty"`
	ConnectionLimit int        `json:"connectionLimit"`
	Fqdn            LBNodeFQDN `json:"fqdn"`
	Logging         string     `json:"logging"`
	Monitor         string     `json:"monitor"`
	RateLimit       string     `json:"rateLimit"`
	Session         string     `json:"session,omitEmpty"`
	State           string     `json:"state,omitEmpty"`
}

type LBNodeRef struct {
	Link  string   `json:"selfLink"`
	Items []LBNode `json":items"`
}

type LBNodes struct {
	Items []LBNode `json:"items"`
}

func (f *Device) ShowNodes() (error, *LBNodes) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node"
	res := LBNodes{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowNode(nname string) (error, *LBNode) {

	//u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddNode() (error, *LBNode) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node"
	res := LBNode{}
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

func (f *Device) UpdateNode(nname string) (error, *LBNode) {

	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}
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

func (f *Device) DeleteNode(nname string) (error, *Response) {

	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

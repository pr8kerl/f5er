package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	//	"github.com/kr/pretty"
)


type LBVirtualPolicy struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
}

type LBVirtualPoliciesRef struct {
	Items []LBVirtualPolicy `json":items"`
}

type LBVirtualProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
	Context   string `json:"context"`
}

type LBVirtualProfileRef struct {
	Items []LBVirtualProfile `json":items"`
}

type LBVirtual struct {
	Name             string               `json:"name"`
	FullPath         string               `json:"fullPath"`
	Partition        string               `json:"partition"`
	Destination      string               `json:"destination"`
	Pool             string               `json:"pool"`
	AddressStatus    string               `json:"addressStatus"`
	AutoLastHop      string               `json:"autoLasthop"`
	CmpEnabled       string               `json:"cmpEnabled"`
	ConnectionLimit  int                  `json:"connectionLimit"`
	Enabled          bool                 `json:"enabled"`
	IpProtocol       string               `json:"ipProtocol"`
	Source           string               `json:"source"`
	SourcePort       string               `json:"sourcePort"`
	SynCookieStatus  string               `json:"synCookieStatus"`
	TranslateAddress string               `json:"translateAddress"`
	TranslatePort    string               `json:"translatePort"`
	Profiles         LBVirtualProfileRef  `json:"profilesReference"`
	Policies         LBVirtualPoliciesRef `json:"policiesReference"`
}

type LBVirtuals struct {
	Items []LBVirtual
}

func showVirtuals() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual"
	res := LBVirtuals{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		fmt.Printf("%s\n", v.FullPath)
	}

}

func showVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addVirtual() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual"
	res := LBVirtual{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a virtual struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// post the request
	err, resp := SendRequest(u, POST, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func updateVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + vname
	res := LBVirtual{}
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a virtual struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// put the request
	err, resp := SendRequest(u, PUT, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func deleteVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + vname
	res := json.RawMessage{}

	err, resp := SendRequest(u, DELETE, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, vname)
	}

}

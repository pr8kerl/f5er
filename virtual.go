package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	//	"github.com/kr/pretty"
)

/*
	{
		"kind": "tm:ltm:virtual:virtualstate",
		"name": "sit_mydotapi_443_vs",
		"partition": "DMZ",
		"fullPath": "/DMZ/sit_mydotapi_443_vs",
		"generation": 98,
		"selfLink": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~sit_mydotapi_443_vs?ver=11.6.0",
		"addressStatus": "yes",
		"autoLasthop": "default",
		"cmpEnabled": "yes",
		"connectionLimit": 0,
		"destination": "/DMZ/10.60.60.146%6:443",
		"enabled": true,
		"gtmScore": 0,
		"ipProtocol": "tcp",
		"mask": "255.255.255.255",
		"mirror": "disabled",
		"mobileAppTunnel": "disabled",
		"nat64": "disabled",
		"pool": "/DMZ/audmzagw-sit_mydotapi_443_pool",
		"rateLimit": "disabled",
		"rateLimitDstMask": 0,
		"rateLimitMode": "object",
		"rateLimitSrcMask": 0,
		"source": "0.0.0.0%6/0",
		"sourceAddressTranslation": {
			"type": "none"
		},
		"sourcePort": "preserve",
		"synCookieStatus": "not-activated",
		"translateAddress": "enabled",
		"translatePort": "enabled",
		"vlansDisabled": true,
		"vsIndex": 145,

	"policiesReference": {
        "link":"https://localhost/mgmt/tm/ltm/virtual/~DMZ-Legacy~audmzisa-sit.store.myob.com.au_443_vs/policies?ver=11.6.0",
		"isSubcollection": true
        "items":[
          {
            "kind":"tm:ltm:virtual:policies:policiesstate",
            "name":"SNAT_VLAN8_Nodes",
            "partition":"DMZ-Legacy",
            "fullPath":"/DMZ-Legacy/SNAT_VLAN8_Nodes",
            "generation":1,
            "selfLink":"https://localhost/mgmt/tm/ltm/virtual/~DMZ-Legacy~audmzisa-sit.store.myob.com.au_443_vs/policies/~DMZ- Legacy~SNAT_VLAN8_Nodes?ver=11.6.0"
          }
        ]
    },
    "profilesReference": {
        "link": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~audmzagw-sit-bettacart-443-vs/profiles?ver=11.6.0",
        "isSubcollection": true,
        "items": [
            {
                "kind": "tm:ltm:virtual:profiles:profilesstate",
                "name": "sit.store.myob.com.au",
                "partition": "Common",
                "fullPath": "/Common/sit.store.myob.com.au",
                "generation": 42,
                "selfLink": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~audmzagw-sit-bettacart-443-vs/profiles/~Common~sit.store.myob.com.au?ver=11.6.0",
                "context": "serverside"
            },
            {
                "kind": "tm:ltm:virtual:profiles:profilesstate",
                "name": "tcp",
                "partition": "Common",
                "fullPath": "/Common/tcp",
                "generation": 42,
                "selfLink": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~audmzagw-sit-bettacart-443-vs/profiles/~Common~tcp?ver=11.6.0",
                "context": "all"
            }
        ]
    }



	}
*/

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

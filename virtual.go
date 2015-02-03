package main

import (
	"fmt"
	"log"
	"strings"
	//	"github.com/kr/pretty"
)

type LBVirtual struct {
	Name        string `json:"name"`
	Path        string `json:"fullPath"`
	Destination string `json:"destination"`
	Pool        string `json:"pool"`
}

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
			"link": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~sit_mydotapi_443_vs/policies?ver=11.6.0",
			"isSubcollection": true
		},
		"profilesReference": {
			"link": "https://localhost/mgmt/tm/ltm/virtual/~DMZ~sit_mydotapi_443_vs/profiles?ver=11.6.0",
			"isSubcollection": true
		}
	}


*/

type LBVirtuals struct {
	Items []LBVirtual
}

func showVirtuals() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual"
	res := LBVirtuals{}

	err := GetRequest(u, &res)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Items {
		fmt.Printf("%s\n", v.Path)
	}

}

func showVirtual(vname string) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err := GetRequest(u, &res)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(&res)

}

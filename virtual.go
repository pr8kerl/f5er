package main

import (
	"fmt"
	"log"
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
		fmt.Printf("virtual:\t%s\n", v.Path)
	}

}

func showVirtual(vname string) {

	u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/~" + partition + "~" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err := GetRequest(u, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("virtual name:\t%s\n", res.Name)
	fmt.Printf("path:\t%s\n", res.Path)
	fmt.Printf("destination:\t%s\n", res.Destination)
	fmt.Printf("pool:\t%s\n", res.Pool)

	/*
		for i, member := range res.MemberRef.Items {
			fmt.Printf("\tmember %d name:\t\t%s\n", i, member.Name)
			fmt.Printf("\tmember %d address:\t%s\n", i, member.Address)
			fmt.Printf("\tmember %d state:\t\t%s\n", i, member.State)
		}
	*/

}

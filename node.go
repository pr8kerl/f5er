package main

import (
	"fmt"
	"strings"
	//	"github.com/kr/pretty"
	"log"
)

/*
{
				  "kind":"tm:ltm:node:nodecollectionstate",
					"selfLink":"https://localhost/mgmt/tm/ltm/node?ver=11.6.0",
					"items":[
					  {
                    "kind":"tm:ltm:node:nodestate",
										"name":"audmzagw03-audmzapi",
										"partition":"DMZ-Legacy",
										"fullPath":"/DMZ-Legacy/audmzagw03-audmzapi",
										"generation":1,
										"selfLink":"https://localhost/mgmt/tm/ltm/node/~DMZ-Legacy~audmzagw03-audmzapi?ver=11.6.0",
										"address":"10.60.8.147%5",
										"connectionLimit":0,
										"dynamicRatio":1,
										"ephemeral":"false",
										"fqdn":{"addressFamily":"ipv4","autopopulate":"disabled","downInterval":5,"interval":3600},
										"logging":"disabled",
										"monitor":"default",
										"rateLimit":"disabled",
										"ratio":1,
										"session":"user-enabled",
										"state":"unchecked"
						},
*/

type LBNodeFQDN struct {
	AddressFamily string         `json:"addressFamily"`
	AutoPopulate  string         `json:"autopopulate"`
	DownInterval  int            `json:"downInterval"`
	Interval      int            `json:"interval"`
	Items         []LBPoolMember `json":items"`
}

type LBNode struct {
	Name            string     `json:"name"`
	Partition       string     `json:"partition"`
	Fullpath        string     `json:"fullPath"`
	Generation      int        `json:"generation"`
	Address         string     `json:"address"`
	ConnectionLimit int        `json:"connectionLimit"`
	Fqdn            LBNodeFQDN `json:"fqdn"`
	Logging         string     `json:"logging"`
	Monitor         string     `json:"monitor"`
	RateLimit       string     `json:"rateLimit"`
	Session         string     `json:"session"`
	State           string     `json:"state"`
}

type LBNodeRef struct {
	Link  string   `json:"selfLink"`
	Items []LBNode `json":items"`
}

type LBNodes struct {
	Items []LBNode `json:"items"`
}

func showNodes() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/node"
	res := LBNodes{}

	err := GetRequest(url, &res)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Items {
		//fmt.Printf("pool:\t%s\n", v.Fullpath)
		fmt.Printf("%s\n", v.Fullpath)
	}
}

func showNode(nname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}

	err := GetRequest(u, &res)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(&res)

}

func createNode(nname string) {
	fmt.Printf("%s\n", nname)
}

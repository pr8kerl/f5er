package main

import (
	"fmt"
	"strings"
	//	"github.com/kr/pretty"
	"encoding/json"
	"io/ioutil"
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
	AddressFamily string `json:"addressFamily"`
	AutoPopulate  string `json:"autopopulate"`
	DownInterval  int    `json:"downInterval"`
	Interval      int    `json:"interval"`
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

// a node struct but with only the postable fields
// used to create a node
type LBNodePost struct {
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
}

func showNodes() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/node"
	res := LBNodes{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
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

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}
	printResponse(&res)

}

func addNode() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/node"
	res := LBNode{}
	body := LBNodePost{}

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
	err = PostRequest(u, &body, &res)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(&res)

}

func deleteNode(nname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
	result := json.RawMessage{}

	err, resp := DeleteRequest(u, &result)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	} else {
		log.Printf("%d: %s deleted successfully\n", resp.Status(), nname)
	}

}

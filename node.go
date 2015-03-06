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

func showNodes() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/node"
	res := LBNodes{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		//fmt.Printf("pool:\t%s\n", v.FullPath)
		fmt.Printf("%s\n", v.FullPath)
	}
}

func showNode(nname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addNode() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/node"
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
	err, resp := PostRequest(u, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func updateNode(nname string) {

	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
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
	err, resp := PutRequest(u, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
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
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, nname)
	}

}

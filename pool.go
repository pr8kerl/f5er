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
		"kind": "tm:ltm:pool:members:membersstate",
		"name": "audmzbilltweb04-pdv:443",
		"partition": "DMZ",
		"fullPath": "/DMZ/audmzbilltweb04-pdv:443",
		"generation": 1,
		"selfLink": "https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-pdv_443_pool/members/~DMZ~audmzbilltweb04-pdv:443?ver=11.6.0",
		"address": "10.60.61.214%6",
		"connectionLimit": 0,
		"dynamicRatio": 1,
		"ephemeral": "false",
		"fqdn": {
			"autopopulate": "disabled"
		},
		"inheritProfile": "enabled",
		"logging": "disabled",
		"monitor": "default",
		"priorityGroup": 0,
		"rateLimit": "disabled",
		"ratio": 1,
		"session": "monitor-enabled",
		"state": "up"
	}

*/
// a pool member
type LBPoolMember struct {
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Address         string `json:"address"`
	ConnectionLimit int    `json:"connectionLimit"`
	DynamicRatio    int    `json:"dynamicRatio"`
	Ephemeral       string `json:"ephemeral"`
	InheritProfile  string `json:"inheritProfile"`
	Logging         string `json:"logging"`
	Monitor         string `json:"monitor"`
	PriorityGroup   int    `json:"priorityGroup"`
	RateLimit       string `json:"rateLimit"`
	Ratio           int    `json:"ratio"`
	Session         string `json:"session"`
	State           string `json:"state"`
}

// a pool member reference - just a link and an array of pool members
type LBPoolMemberRef struct {
	Link  string         `json:"link"`
	Items []LBPoolMember `json":items"`
}

type LBPoolMembers struct {
	Link  string         `json:"selfLink"`
	Items []LBPoolMember `json":items"`
}

// used by online/offline
type MemberState struct {
	State   string `json:"state"`
	Session string `json:"session"`
}

/*
{
	"kind": "tm:ltm:pool:poolstate",
	"name": "audmzbilltweb-pdv_443_pool",
	"partition": "DMZ",
	"fullPath": "/DMZ/audmzbilltweb-pdv_443_pool",
	"generation": 1,
	"selfLink": "https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-pdv_443_pool?expandSubcollections=true\u0026ver=11.6.0",
	"allowNat": "yes",
	"allowSnat": "yes",
	"ignorePersistedWeight": "disabled",
	"ipTosToClient": "pass-through",
	"ipTosToServer": "pass-through",
	"linkQosToClient": "pass-through",
	"linkQosToServer": "pass-through",
	"loadBalancingMode": "round-robin",
	"minActiveMembers": 0,
	"minUpMembers": 0,
	"minUpMembersAction": "failover",
	"minUpMembersChecking": "disabled",
	"monitor": "/Common/tcp ",
	"queueDepthLimit": 0,
	"queueOnConnectionLimit": "disabled",
	"queueTimeLimit": 0,
	"reselectTries": 0,
	"serviceDownAction": "none",
	"slowRampTime": 10,
	"membersReference": {

*/

type LBPool struct {
	Name                   string          `json:"name"`
	FullPath               string          `json:"fullPath"`
	Generation             int             `json:"generation"`
	AllowNat               string          `json:"allowNat"`
	AllowSnat              string          `json:"allowSnat"`
	IgnorePersistedWeight  string          `json:"ignorePersistedWeight"`
	IpTosToClient          string          `json:"ipTosToClient"`
	IpTosToServer          string          `json:"ipTosToServer"`
	LinkQosToClient        string          `json:"linkQosToClient"`
	LinkQosToServer        string          `json:"linkQosToServer"`
	LoadBalancingMode      string          `json:"loadBalancingMode"`
	MinActiveMembers       int             `json:"minActiveMembers"`
	MinUpMembers           int             `json:"minUpMembers"`
	MinUpMembersAction     string          `json:"minUpMembersAction"`
	MinUpMembersChecking   string          `json:"minUpMembersChecking"`
	Monitor                string          `json:"monitor"`
	QueueDepthLimit        int             `json:"queueDepthLimit"`
	QueueOnConnectionLimit string          `json:"queueOnConnectionLimit"`
	QueueTimeLimit         int             `json:"queueTimeLimit"`
	ReselectTries          int             `json:"reselectTries"`
	ServiceDownAction      string          `json:"serviceDownAction"`
	SlowRampTime           int             `json:"slowRampTime"`
	MemberRef              LBPoolMemberRef `json:"membersReference"`
}

type LBPools struct {
	Items []LBPool `json:"items"`
}

func showPools() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/pool"
	res := LBPools{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		//fmt.Printf("pool:\t%s\n", v.FullPath)
		fmt.Printf("%s\n", v.FullPath)
	}
}

func showPool(pname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"
	res := LBPool{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addPool() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/pool"
	res := LBPool{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a pool struct
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

func updatePool(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool
	res := LBPool{}
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

func deletePool(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool
	result := json.RawMessage{}

	err, resp := DeleteRequest(u, &result)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, pname)
	}
}

func showPoolMembers(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	//	member := strings.Replace(pmember, "/", "~", -1)
	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members?expandSubcollections=true"
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res.Items)

}

func addPoolMembers(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}
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

func updatePoolMembers(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}
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

func deletePoolMembers(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	result := json.RawMessage{}

	err, resp := DeleteRequest(u, &result)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, pname)
	}
}

func onlinePoolMember(mname string) {

	pmember := strings.Replace(mname, "/", "~", -1)
	pool := strings.Replace(f5Pool, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members/" + pmember
	res := json.RawMessage{}

	/*
	   {"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)
	   {"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)
	   {"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)
	*/
	body := MemberState{"user-up", "user-enabled"}

	// put the request
	err, resp := PutRequest(u, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s online\n", resp.HttpResponse().Status, mname)
	}
	printResponse(&res)

}

func offlinePoolMember(mname string) {

	pmember := strings.Replace(mname, "/", "~", -1)
	pool := strings.Replace(f5Pool, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members/" + pmember
	res := json.RawMessage{}

	/*
	   {"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)
	   {"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)
	   {"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)
	*/
	var body MemberState
	if now {
		body = MemberState{"user-down", "user-disabled"}
	} else {
		body = MemberState{"user-up", "user-disabled"}
	}

	// put the request
	err, resp := PutRequest(u, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s offline\n", resp.HttpResponse().Status, mname)
	}
	printResponse(&res)

}

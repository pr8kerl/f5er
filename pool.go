package main

import (
	"fmt"
	"strings"
	//	"github.com/kr/pretty"
	"log"
)

// a pool member
type LBPoolMember struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	Fullpath  string `json:"fullPath"`
	Address   string `json:"address"`
	State     string `json:"state"`
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

type LBPool struct {
	Name              string          `json:"name"`
	Fullpath          string          `json:"fullPath"`
	Generation        int             `json:"generation"`
	AllowNat          string          `json:"allowNat"`
	AllowSnat         string          `json:"allowSnat"`
	LoadBalancingMode string          `json:"loadBalancingMode"`
	Monitor           string          `json:"monitor"`
	MemberRef         LBPoolMemberRef `json:"membersReference"`
}

type LBPools struct {
	Items []LBPool `json:"items"`
}

func showPools() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/pool"
	res := LBPools{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}

	for _, v := range res.Items {
		//fmt.Printf("pool:\t%s\n", v.Fullpath)
		fmt.Printf("%s\n", v.Fullpath)
	}
}

func showPool(pname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"
	res := LBPool{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}
	printResponse(&res)

	/*
		fmt.Printf("pool name:\t%s\n", res.Name)
		fmt.Printf("fullpath:\t%s\n", res.Fullpath)
		fmt.Printf("lb mode:\t%s\n", res.LoadBalancingMode)
		fmt.Printf("monitor:\t%s\n", res.Monitor)

		for i, member := range res.MemberRef.Items {
			fmt.Printf("\tmember %d name:\t\t%s\n", i, member.Name)
			fmt.Printf("\tmember %d address:\t%s\n", i, member.Address)
			fmt.Printf("\tmember %d state:\t\t%s\n", i, member.State)
		}
	*/

}

func addPool(pname string) {
	fmt.Printf("%s\n", pname)
}

func showPoolMembers(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	//	member := strings.Replace(pmember, "/", "~", -1)
	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members?expandSubcollections=true"
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}
	printResponse(&res.Items)

}

func addPoolMember(pname string) {

	pool := strings.Replace(pname, "/", "~", -1)
	//	member := strings.Replace(pmember, "/", "~", -1)
	//u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members?expandSubcollections=true"
	u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}

	// check input is set
	// read in input

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}
	printResponse(&res.Items)

}

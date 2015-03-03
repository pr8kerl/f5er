package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBStack struct {
	Nodes []map[string]json.RawMessage `json:"nodes"`
	Pool  json.RawMessage              `json:"pool"`
}

func showStack() {

	//	xid := strings.Replace(xname, "/", "~", -1)
	//	u := "https://" + f5Host + "/mgmt/tm/transaction/" + xid
	stack := LBStack{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	nres := LBNode{}

	for count, n := range stack.Nodes {
		//fmt.Printf("pool:\t%s\n", v.Fullpath)
		nde := string(n["fullPath"])
		// because its read ffrom a map - strip the quotes
		nde = strings.Replace(nde, "\"", "", -1)
		fmt.Printf("\nnode[%d]: %s\n", count, nde)

		node := strings.Replace(nde, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		err, resp := GetRequest(u, &nres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&nres)

	}

	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\npool: %s\n", jpool.Fullpath)
		pool := strings.Replace(jpool.Fullpath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"

		err, resp := GetRequest(u, &pres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&pres)

	}

}

func addStack() {

	//	u := "https://" + f5Host + "/mgmt/tm/transaction"
	stack := LBStack{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	// post the request
	//	err, resp := PostRequest(u, &body, &res)
	//	if err != nil {
	//		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	//	}
	//	printResponse(&res)
}

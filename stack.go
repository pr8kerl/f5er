package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBStack struct {
	Nodes   []map[string]json.RawMessage `json:"nodes"`
	Pool    json.RawMessage              `json:"pool"`
	Virtual json.RawMessage              `json:"virtual"`
}

type LBTransaction struct {
	TransId int    `json:"transId"`
	Timeout int    `json:"timeoutSeconds"`
	State   string `json:"state"`
}

/*
{
"transId":1389812351,
"state":"STARTED",
"timeoutSeconds":30,
"kind":"tm:transactionstate",
"selfLink":"https://localhost/mgmt/tm/transaction/1389812351?ver=11.5.0"
}

 func SendRequest(u string, method int, sess *napping.Session, pload interface{}, res interface{}) (error, *napping.Response) {
*/

func showStack() {

	InitSession()

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

	// show nodes
	for count, n := range stack.Nodes {
		//fmt.Printf("pool:\t%s\n", v.FullPath)
		nde := string(n["fullPath"])
		// because its read ffrom a map - strip the quotes
		nde = strings.Replace(nde, "\"", "", -1)
		fmt.Printf("\nnode[%d]: %s\n", count, nde)

		node := strings.Replace(nde, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		//err, resp := GetRequest(u, &nres)
		err, resp := SendRequest(u, RESTGET, &sessn, nil, &nres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&nres)

	}

	// show pool
	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\npool: %s\n", jpool.FullPath)
		pool := strings.Replace(jpool.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"

		err, resp := GetRequest(u, &pres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&pres)

	}
	// show virtual
	if len(stack.Virtual) > 0 {

		vres := LBVirtual{}
		virt := LBVirtual{}
		if err := json.Unmarshal(stack.Virtual, &virt); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nvirtual: %s\n", virt.FullPath)
		virtual := strings.Replace(virt.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + virtual + "?expandSubcollections=true"

		err, resp := GetRequest(u, &vres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&vres)

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

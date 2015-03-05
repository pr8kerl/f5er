package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBStack struct {
	Nodes   []json.RawMessage `json:"nodes"`
	Pool    json.RawMessage   `json:"pool"`
	Virtual json.RawMessage   `json:"virtual"`
}

type LBTransaction struct {
	TransId int    `json:"transId"`
	Timeout int    `json:"timeoutSeconds"`
	State   string `json:"state"`
}

type LBEmptyBody struct{}

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

	//	xid := strings.Replace(xname, "/", "~", -1)
	//	u := "https://" + f5Host + "/mgmt/tm/transaction/" + xid
	stack := LBStack{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	// show nodes
	for count, n := range stack.Nodes {

		nres := LBNode{}
		nde := LBNode{}
		if err := json.Unmarshal(n, &nde); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		node := strings.Replace(nde.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		err, resp := SendRequest(u, GET, &sessn, nil, &nres)
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

		log.Printf("\npool: %s\n", jpool.FullPath)
		pool := strings.Replace(jpool.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"

		err, resp := SendRequest(u, GET, &sessn, nil, &pres)
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

		log.Printf("\nvirtual: %s\n", virt.FullPath)
		virtual := strings.Replace(virt.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + virtual + "?expandSubcollections=true"

		sessn.Header.Set("Haribo", "macht kinder froh")
		err, resp := SendRequest(u, GET, &sessn, nil, &vres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&vres)

	}

}

func addStack() {

	stack := LBStack{}
	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatal(err)
	}

	u := "https://" + f5Host + "/mgmt/tm/transaction"
	body := LBEmptyBody{}
	tres := LBTransaction{}
	err, resp := SendRequest(u, POST, &sessn, &body, &tres)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&tres)

	tid := fmt.Sprintf("%d", tres.TransId)
	log.Printf("created transaction id: %s\n", tid)
	sessn.Header.Set("X-F5-REST-Coordination-Id", tid)

	// add nodes
	nres := LBNode{}
	nde := LBNode{}
	for count, n := range stack.Nodes {

		if err := json.Unmarshal(n, &nde); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		node := strings.Replace(nde.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		err, resp := SendRequest(u, POST, &sessn, &nde, &nres)
		if err != nil {
			log.Fatalf("%s : error adding %s : %s\n", resp.HttpResponse().Status, node, err)
		} else {
			log.Printf("%s : %s added\n", resp.HttpResponse().Status, node)
		}
		printResponse(&nres)
	}

	// add pool
	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		log.Printf("\npool: %s\n", jpool.FullPath)
		pool := strings.Replace(jpool.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"

		err, resp := SendRequest(u, GET, &sessn, &jpool, &pres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		}
		printResponse(&pres)

	}

}

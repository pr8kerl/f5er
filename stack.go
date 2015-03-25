package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type LBStack struct {
	Nodes    []json.RawMessage `json:"nodes"`
	Pool     json.RawMessage   `json:"pool"`
	Virtuals []json.RawMessage `json:"virtuals"`
}

type LBTransaction struct {
	TransId int    `json:"transId"`
	Timeout int    `json:"timeoutSeconds"`
	State   string `json:"state"`
}

type LBEmptyBody struct{}

type LBTransactionState struct {
	State string `json:"state"`
}

type LBNodeFQDNUpdate struct {
	DownInterval int `json:"downInterval"`
	Interval     int `json:"interval"`
}

type LBNodeUpdate struct {
	Name            string           `json:"name"`
	Partition       string           `json:"partition"`
	FullPath        string           `json:"fullPath"`
	Generation      int              `json:"generation"`
	ConnectionLimit int              `json:"connectionLimit"`
	Fqdn            LBNodeFQDNUpdate `json:"fqdn"`
	Logging         string           `json:"logging"`
	Monitor         string           `json:"monitor"`
	RateLimit       string           `json:"rateLimit"`
}

/*
{
"transId":1389812351,
"state":"STARTED",
"timeoutSeconds":30,
"kind":"tm:transactionstate",
"selfLink":"https://localhost/mgmt/tm/transaction/1389812351?ver=11.5.0"
}

*/

func showStack() {

	//	xid := strings.Replace(xname, "/", "~", -1)
	//	u := "https://" + f5Host + "/mgmt/tm/transaction/" + xid
	stack := LBStack{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatalf("error reading input json file: %s\n", err)
	}

	// convert json to a stack struct
	err = json.Unmarshal(dat, &stack)
	if err != nil {
		log.Fatalf("error unmarshaling input json file into a stack: %s\n", err)
	}

	// show nodes
	for count, n := range stack.Nodes {

		nres := LBNode{}
		nde := LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &nde)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		node := strings.Replace(nde.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		err, resp := SendRequest(u, GET, &sessn, nil, &nres)
		if err != nil {
			log.Printf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			printResponse(&nres)
		}
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
			log.Printf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			printResponse(&pres)
		}

	}
	// show virtual
	for count, v := range stack.Virtuals {

		vres := LBVirtual{}
		virt := LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}

		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)
		virtual := strings.Replace(virt.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + virtual + "?expandSubcollections=true"

		sessn.Header.Set("Haribo", "macht kinder froh")
		err, resp := SendRequest(u, GET, &sessn, nil, &vres)
		if err != nil {
			log.Printf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			printResponse(&vres)
		}

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
	empty := LBEmptyBody{}
	tres := LBTransaction{}
	err, resp := SendRequest(u, POST, &sessn, &empty, &tres)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : transaction %d created\n", resp.HttpResponse().Status, tres.TransId)
	}

	tid := fmt.Sprintf("%d", tres.TransId)
	// set the transaction header
	sessn.Header.Set("X-F5-REST-Coordination-Id", tid)

	// add nodes
	for count, n := range stack.Nodes {

		nres := LBNode{}
		nde := LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &nde)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		u := "https://" + f5Host + "/mgmt/tm/ltm/node"
		// send the raw json - not the struct - some fields may be omiitted from the input intentionally and the struct will insert empty fields
		err, resp := SendRequest(u, POST, &sessn, &n, &nres)
		if err != nil {
			log.Fatalf("%s : error adding %s : %s\n", resp.HttpResponse().Status, nde.FullPath, err)
		} else {
			log.Printf("%s : node[%d] %s added\n", resp.HttpResponse().Status, count, nde.FullPath)
		}
	}

	// add pool
	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		log.Printf("\npool: %s\n", jpool.FullPath)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool"

		err, resp := SendRequest(u, POST, &sessn, &stack.Pool, &pres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			log.Printf("%s : pool %s added\n", resp.HttpResponse().Status, jpool.FullPath)
		}

	}

	// add virtual
	for count, v := range stack.Virtuals {

		vres := LBVirtual{}
		virt := LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}

		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual"

		err, resp := SendRequest(u, POST, &sessn, &v, &vres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			log.Printf("%s : virtual %s added\n", resp.HttpResponse().Status, virt.FullPath)
		}

	}

	// if we made it here - commit the transaction - remove the transaction header first
	sessn.Header.Del("X-F5-REST-Coordination-Id")

	u = "https://" + f5Host + "/mgmt/tm/transaction/" + tid
	body := LBTransaction{State: "VALIDATING"}
	tres = LBTransaction{}
	err, resp = SendRequest(u, PATCH, &sessn, &body, &tres)
	if err != nil {
		log.Fatalf("\n%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("\n%s : transaction %s committed\n", resp.HttpResponse().Status, tid)
	}

}

func updateStack() {

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
	empty := LBEmptyBody{}
	tres := LBTransaction{}
	err, resp := SendRequest(u, POST, &sessn, &empty, &tres)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	tid := fmt.Sprintf("%d", tres.TransId)
	log.Printf("created transaction id: %s\n", tid)
	// set the transaction header
	sessn.Header.Set("X-F5-REST-Coordination-Id", tid)

	// nodes
	for count, n := range stack.Nodes {

		nres := LBNode{}
		nde := LBNodeUpdate{}
		// convert json to a node struct
		err = json.Unmarshal(n, &nde)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		npath := strings.Replace(nde.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + npath
		err, resp := SendRequest(u, PUT, &sessn, &nde, &nres)
		if err != nil {
			log.Fatalf("%s : %s : %s\n", resp.HttpResponse().Status, nde.FullPath, err)
		} else {
			log.Printf("%s : node[%d] %s updated\n", resp.HttpResponse().Status, count, nde.FullPath)
		}
	}

	// pool
	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		log.Printf("\npool: %s\n", jpool.FullPath)
		pool := strings.Replace(jpool.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool

		err, resp := SendRequest(u, PUT, &sessn, &stack.Pool, &pres)
		if err != nil {
			log.Fatalf("%s : %s : %s\n", resp.HttpResponse().Status, jpool.FullPath, err)
		} else {
			log.Printf("%s : pool %s updated\n", resp.HttpResponse().Status, jpool.FullPath)
		}

	}

	// add virtual
	for count, v := range stack.Virtuals {

		vres := LBVirtual{}
		virt := LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}

		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)
		virtual := strings.Replace(virt.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + virtual

		err, resp := SendRequest(u, PUT, &sessn, &v, &vres)
		if err != nil {
			log.Fatalf("%s : %s : %s\n", resp.HttpResponse().Status, virt.FullPath, err)
		} else {
			log.Printf("%s : virtual %s updated\n", resp.HttpResponse().Status, virt.FullPath)
		}

	}

	// if we made it here - commit the transaction
	sessn.Header.Del("X-F5-REST-Coordination-Id")

	u = "https://" + f5Host + "/mgmt/tm/transaction/" + tid
	body := LBTransaction{State: "VALIDATING"}
	tres = LBTransaction{}
	err, resp = SendRequest(u, PATCH, &sessn, &body, &tres)
	if err != nil {
		log.Fatalf("\n%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("\n%s : transaction %s committed\n", resp.HttpResponse().Status, tid)
	}

}

func deleteStack() {

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
	empty := LBEmptyBody{}
	tres := LBTransaction{}
	err, resp := SendRequest(u, POST, &sessn, &empty, &tres)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	tid := fmt.Sprintf("%d", tres.TransId)
	log.Printf("created transaction id: %s\n", tid)
	// set the transaction header
	sessn.Header.Set("X-F5-REST-Coordination-Id", tid)

	// virtual
	for count, v := range stack.Virtuals {

		vres := LBVirtual{}
		virt := LBVirtual{}
		if err := json.Unmarshal(v, &virt); err != nil {
			log.Fatal(err)
		}

		log.Printf("\nvirtual[%d]: %s\n", count, virt.FullPath)
		virtual := strings.Replace(virt.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/virtual/" + virtual

		err, resp := SendRequest(u, DELETE, &sessn, &v, &vres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			log.Printf("%s : virtual %s deleted\n", resp.HttpResponse().Status, virt.FullPath)
		}

	}

	// pool
	if len(stack.Pool) > 0 {

		pres := LBPool{}
		jpool := LBPool{}
		if err := json.Unmarshal(stack.Pool, &jpool); err != nil {
			log.Fatal(err)
		}

		log.Printf("\npool: %s\n", jpool.FullPath)
		pool := strings.Replace(jpool.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool

		err, resp := SendRequest(u, DELETE, &sessn, &stack.Pool, &pres)
		if err != nil {
			log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
		} else {
			log.Printf("%s : pool %s deleted\n", resp.HttpResponse().Status, jpool.FullPath)
		}

	}

	// if we made it here - commit the transaction
	sessn.Header.Del("X-F5-REST-Coordination-Id")

	u = "https://" + f5Host + "/mgmt/tm/transaction/" + tid
	body := LBTransaction{State: "VALIDATING"}
	tres = LBTransaction{}
	err, resp = SendRequest(u, PATCH, &sessn, &body, &tres)
	if err != nil {
		log.Fatalf("\n%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("\n%s : transaction %s committed\n", resp.HttpResponse().Status, tid)
	}

	// delete nodes outside of transaction - pools depend on them and won't delete otherwise
	for count, n := range stack.Nodes {

		nres := LBNode{}
		nde := LBNode{}
		// convert json to a node struct
		err = json.Unmarshal(n, &nde)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nnode[%d]: %s\n", count, nde.FullPath)

		node := strings.Replace(nde.FullPath, "/", "~", -1)
		u := "https://" + f5Host + "/mgmt/tm/ltm/node/" + node
		err, resp := SendRequest(u, DELETE, &sessn, nil, &nres)
		if err != nil {
			log.Fatalf("%s : error deleting %s : %s\n", resp.HttpResponse().Status, nde.FullPath, err)
		} else {
			log.Printf("%s : node[%d] %s deleted\n", resp.HttpResponse().Status, count, nde.FullPath)
		}
	}

}

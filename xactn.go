package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
{
"transId":1389812351,
"state":"STARTED",
"timeoutSeconds":30,
"kind":"tm:transactionstate",
"selfLink":"https://localhost/mgmt/tm/transaction/1389812351?ver=11.5.0"
}

{
"items":[],
"kind":"tm:transactioncollectionstate",
"selfLink":"https://localhost/mgmt/tm/transaction?ver=11.6.0"
}

*/

type LBXactn struct {
	TransId int    `json:"transId"`
	State   string `json:"state"`
	Timeout int    `json:"timeoutSeconds"`
	Link    string `json:"selfLink"`
}

type LBXactns struct {
	Items []LBXactn `json:"items"`
}

func showXactns() {

	url := "https://" + f5Host + "/mgmt/tm/transaction"
	res := LBXactns{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		//fmt.Printf("pool:\t%s\n", v.Fullpath)
		fmt.Printf("id: %d, state: %s\n", v.TransId, v.State)
	}
}

func showXactn(xname string) {

	xid := strings.Replace(xname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/transaction/" + xid
	res := LBXactn{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addXactn() {

	u := "https://" + f5Host + "/mgmt/tm/transaction"
	res := LBXactn{}
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

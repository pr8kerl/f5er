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
  "kind":"tm:ltm:rule:rulecollectionstate",
  "selfLink":"https://localhost/mgmt/tm/ltm/rule?ver=11.6.0",
  "items":[
       {
          "kind":"tm:ltm:rule:rulestate",
          "name":"HTTP_to_HTTPS_redirect_301",
          "partition":"Common",
          "fullPath":"/Common/HTTP_to_HTTPS_redirect_301",
          "generation":1,
          "selfLink":"https://localhost/mgmt/tm/ltm/rule/~Common~HTTP_to_HTTPS_redirect_301?ver=11.6.0",
          "apiAnonymous":"when HTTP_REQUEST {\n  HTTP::respond 301 Location \"https://[getfield [HTTP::host] : 1][HTTP::uri]\"\n}"
       },
       {
          "kind":"tm:ltm:rule:rulestate",
          "name":"_sys_APM_ExchangeSupport_OA_BasicAuth",
          "partition":"Common",
          "fullPath":"/Common/_sys_APM_ExchangeSupport_OA_BasicAuth",
          "generation":1,
          "selfLink":"https://localhost/mgmt/tm/ltm/rule/~Common~_sys_APM_ExchangeSupport_OA_BasicAuth?ver=11.6.0",
          "apiAnonymous":"nodelete nowrite \n  bla - lots removed  ",
          "apiRawValues":{"verificationStatus":"signature-verified"}
       },
       {
          "kind":"tm:ltm:rule:rulestate",
          "name":"HTTP_to_HTTPS_redirect_301",
          "partition":"DMZ-Legacy",
          "fullPath":"/DMZ-Legacy/HTTP_to_HTTPS_redirect_301",
          "generation":1,
          "selfLink":"https://localhost/mgmt/tm/ltm/rule/~DMZ-Legacy~HTTP_to_HTTPS_redirect_301?ver=11.6.0",
          "apiAnonymous":"when HTTP_REQUEST {\n  HTTP::respond 301 Location \"https://[getfield [HTTP::host] : 1][HTTP::uri]\"\n}"}
  ]
}
*/

type LBRawValues struct {
	VerificationStatus string `json:"verificationStatus"`
}

type LBRule struct {
	Name         string      `json:"name"`
	Partition    string      `json:"partition"`
	Fullpath     string      `json:"fullPath"`
	Generation   int         `json:"generation"`
	ApiAnonymous string      `json:"apiAnonymous"`
	ApiRawValues LBRawValues `json:"apiRawValues"`
}

type LBRules struct {
	Items []LBRule `json:"items"`
}

func showRules() {

	url := "https://" + f5Host + "/mgmt/tm/ltm/rule"
	res := LBRules{}

	err, resp := GetRequest(url, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		fmt.Printf("%s\n", v.Fullpath)
	}
}

func showRule(rname string) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/rule/" + rule
	res := LBRule{}

	err, resp := GetRequest(u, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addRule() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/rule"
	res := LBRule{}
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

func updateRule(rname string) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/rule/" + rule
	res := LBRule{}
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

func deleteRule(rname string) {

	rule := strings.Replace(rname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/rule/" + rule
	result := json.RawMessage{}

	err, resp := DeleteRequest(u, &result)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, rname)
	}

}

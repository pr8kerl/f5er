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
				"kind":"tm:ltm:policy:policycollectionstate",
				"selfLink":"https://localhost/mgmt/tm/ltm/policy?expandSubcollections=true&ver=11.6.0",
				"items":[
				{
								"kind":"tm:ltm:policy:policystate",
								"name":"sit.store.myob-websites",
								"partition":"DMZ-Legacy",
								"fullPath":"/DMZ-Legacy/sit.store.myob-websites",
								"generation":1,
								"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites?ver=11.6.0",
								"controls":["forwarding"],
								"requires":["http","tcp"],
								"strategy":"/Common/first-match",
								"rulesReference":{
												"link":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules?ver=11.6.0",
												"isSubcollection":true,
												"items":[
												{
																"kind":"tm:ltm:policy:rules:rulesstate",
																"name":"\"match /buy\"",
																"fullPath":"\"match /buy\"",
																"generation":1,
																"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/%22match%20~buy%22?ver=11.6.0",
																"ordinal":1,
																"actionsReference":{
																				"link":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/%22match%20~buy%22/actions?ver=11.6.0",
																				"isSubcollection":true,
																				"items":[
																				{
																								"kind":"tm:ltm:policy:rules:actions:actionsstate",
																								"name":"0",
																								"fullPath":"0",
																								"generation":1,
																								"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/%22match%20~buy%22/actions/0?ver=11.6.0",
																								"code":0,
																								"forward":true,
																								"pool":"/DMZ-Legacy/audmzagw-mydotapisit_80_pool",
																								"port":0,
																								"request":true,
																								"select":true,
																								"status":0,
																								"vlanId":0
																				}
																				]
																},
																"conditionsReference":{
																				"link":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/%22match%20~buy%22/conditions?ver=11.6.0",
																				"isSubcollection":true,
																				"items":[
																				{
																								"kind":"tm:ltm:policy:rules:conditions:conditionsstate",
																								"name":"0",
																								"fullPath":"0",
																								"generation":1,
																								"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/%22match%20~buy%22/conditions/0?ver=11.6.0",
																								"all":true,
																								"caseInsensitive":true,
																								"external":true,
																								"httpUri":true,
																								"index":0,
																								"present":true,
																								"remote":true,
																								"request":true,
																								"startsWith":true,
																								"values":["/buy","/Buy"]
																				}
																				]
																}
												},
												{
																"kind":"tm:ltm:policy:rules:rulesstate",
																"name":"v8_nodes",
																"fullPath":"v8_nodes",
																"generation":1,
																"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/v8_nodes?ver=11.6.0",
																"ordinal":2,
																"actionsReference":{
																				"link":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/v8_nodes/actions?ver=11.6.0",
																				"isSubcollection":true,
																				"items":[
																				{
																								"kind":"tm:ltm:policy:rules:actions:actionsstate",
																								"name":"0",
																								"fullPath":"0",
																								"generation":1,
																								"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/v8_nodes/actions/0?ver=11.6.0",
																								"code":0,
																								"forward":true,
																								"port":0,
																								"request":true,
																								"select":true,
																								"snatpool":"/DMZ-Legacy/SNAT_Source_VLAN8",
																								"status":0,
																								"vlanId":0
																				}
																				]
																},
																"conditionsReference":{
																				"link":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/v8_nodes/conditions?ver=11.6.0",
																				"isSubcollection":true,
																				"items":[
																				{
																								"kind":"tm:ltm:policy:rules:conditions:conditionsstate",
																								"name":"0",
																								"fullPath":"0",
																								"generation":1,
																								"selfLink":"https://localhost/mgmt/tm/ltm/policy/~DMZ-Legacy~sit.store.myob-websites/rules/v8_nodes/conditions/0?ver=11.6.0",
																								"caseInsensitive":true,
																								"external":true,
																								"index":0,
																								"present":true,
																								"remote":true,
																								"request":true,
																								"startsWith":true,
																								"tcp":true,
																								"values":["10.60.8."],
																								"vlan":true
																				}
																				]
																}
												}
												]
								}
				}
}


*/

type LBPolicyConditions struct {
	Name            string   `json:"name"`
	FullPath        string   `json:"fullPath"`
	Generation      int      `json:"generation"`
	All             bool     `json:"all"`
	CaseInsensitive bool     `json:"caseInsensitive"`
	External        bool     `json:"external"`
	HttpUri         bool     `json:"httpUri"`
	Index           int      `json:"index"`
	Present         bool     `json:"present"`
	Remote          bool     `json:"remote"`
	Request         bool     `json:"request"`
	StartsWith      bool     `json:"startsWith"`
	Values          []string `json:"values"`
}

type LBPolicyActions struct {
	Name       string `json:"name"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Code       int    `json:"code"`
	Forward    bool   `json:"forward"`
	Pool       string `json:"pool"`
	Port       int    `json:"port"`
	Request    bool   `json:"request"`
	Select     bool   `json:"select"`
	Status     int    `json:"status"`
	VlanId     int    `json:"vlanId"`
}

type LBPolicyConditionsRef struct {
	Items []LBPolicyConditions `json:"items"`
}

type LBPolicyActionsRef struct {
	Items []LBPolicyActions `json:"items"`
}

type LBPolicyRules struct {
	Name          string                `json:"name"`
	FullPath      string                `json:"fullPath"`
	Generation    int                   `json:"generation"`
	Ordinal       int                   `json:"ordinal"`
	ActionsRef    LBPolicyActionsRef    `json:"actionsReference"`
	ConditionsRef LBPolicyConditionsRef `json:"conditionsReference"`
}

type LBPolicyRulesRef struct {
	Items []LBPolicyRules `json:"items"`
}

type LBPolicy struct {
	Name       string           `json:"name"`
	Partition  string           `json:"partition"`
	FullPath   string           `json:"fullPath"`
	Generation int              `json:"generation"`
	Controls   []string         `json:"controls"`
	Requires   []string         `json:"requires"`
	Strategy   string           `json:"strategy"`
	RulesRef   LBPolicyRulesRef `json:"rulesReference"`
}

type LBPolicies struct {
	Items []LBPolicy `json:"items"`
}

func showPolicies() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/policy"
	res := LBPolicies{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		//fmt.Printf("policy:\t%s\n", v.FullPath)
		fmt.Printf("%s\n", v.FullPath)
	}
}

func showPolicy(pname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/policy/~" + partition + "~" + pname + "?expandSubcollections=true"

	policy := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/policy/" + policy + "?expandSubcollections=true"
	res := LBPolicy{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addPolicy() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/policy"
	res := LBPolicy{}
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
	err, resp := SendRequest(u, POST, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func updatePolicy(pname string) {

	policy := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/policy/" + policy
	res := LBPolicy{}
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
	err, resp := SendRequest(u, PUT, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func deletePolicy(pname string) {

	//u := "https://" + f5Host + "/mgmt/tm/ltm/policy/~" + partition + "~" + pname + "?expandSubcollections=true"
	policy := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/policy/" + policy
	res := json.RawMessage{}

	err, resp := SendRequest(u, DELETE, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, pname)
	}

}

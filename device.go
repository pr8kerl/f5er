package main

import (
	"log"
)

type LBDeviceRef struct {
	Link  string          `json:"selfLink"`
	Items []LBDeviceState `json":items"`
}

type LBDeviceState struct {
	Name          string `json:"name"`
	Path          string `json:"fullPath"`
	FailoverState string `json:"failoverState"`
	ManagementIP  string `json:"managementIP"`
}


func showDevice() {

	u := "https://" + f5Host + "/mgmt/tm/cm/device"
	res := LBDeviceRef{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res.Items)

}

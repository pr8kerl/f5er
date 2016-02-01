package f5

import (
	"encoding/json"
	"strings"
)

type LBNodeFQDN struct {
	AddressFamily string `json:"addressFamily"`
	AutoPopulate  string `json:"autopopulate"`
	DownInterval  int    `json:"downInterval"`
	Interval      int    `json:"interval"`
}

type LBNode struct {
	Name            string     `json:"name"`
	Partition       string     `json:"partition"`
	FullPath        string     `json:"fullPath"`
	Generation      int        `json:"generation"`
	Address         string     `json:"address,omitEmpty"`
	ConnectionLimit int        `json:"connectionLimit"`
	Fqdn            LBNodeFQDN `json:"fqdn"`
	Logging         string     `json:"logging"`
	Monitor         string     `json:"monitor"`
	RateLimit       string     `json:"rateLimit"`
	Session         string     `json:"session,omitEmpty"`
	State           string     `json:"state,omitEmpty"`
}

type LBNodeRef struct {
	Link  string   `json:"selfLink"`
	Items []LBNode `json":items"`
}

type LBNodes struct {
	Items []LBNode `json:"items"`
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

func (f *Device) ShowNodes() (error, *LBNodes) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node"
	res := LBNodes{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowNode(nname string) (error, *LBNode) {

	//u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddNode(body *json.RawMessage) (error, *LBNode) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node"
	res := LBNode{}

	// post the request
	err, _ := f.sendRequest(u, POST, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdateNode(nname string, body *json.RawMessage) (error, *LBNode) {

	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := LBNode{}

	// put the request
	err, _ := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeleteNode(nname string) (error, *Response) {

	node := strings.Replace(nname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/node/" + node
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, resp
	}

}

package f5

import (
	"encoding/json"
	"log"
	"strings"
)

// a pool member
type LBPoolMember struct {
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	FullPath        string `json:"fullPath"`
	Address         string `json:"address"`
	ConnectionLimit int    `json:"connectionLimit"`
	DynamicRatio    int    `json:"dynamicRatio"`
	Ephemeral       string `json:"ephemeral"`
	InheritProfile  string `json:"inheritProfile"`
	Logging         string `json:"logging"`
	Monitor         string `json:"monitor"`
	PriorityGroup   int    `json:"priorityGroup"`
	RateLimit       string `json:"rateLimit"`
	Ratio           int    `json:"ratio"`
	Session         string `json:"session"`
	State           string `json:"state"`
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

// used by online/offline
type LBPoolMemberState struct {
	State   string `json:"state"`
	Session string `json:"session"`
}

type LBPool struct {
	Name                   string          `json:"name"`
	FullPath               string          `json:"fullPath"`
	Generation             int             `json:"generation"`
	AllowNat               string          `json:"allowNat"`
	AllowSnat              string          `json:"allowSnat"`
	IgnorePersistedWeight  string          `json:"ignorePersistedWeight"`
	IpTosToClient          string          `json:"ipTosToClient"`
	IpTosToServer          string          `json:"ipTosToServer"`
	LinkQosToClient        string          `json:"linkQosToClient"`
	LinkQosToServer        string          `json:"linkQosToServer"`
	LoadBalancingMode      string          `json:"loadBalancingMode"`
	MinActiveMembers       int             `json:"minActiveMembers"`
	MinUpMembers           int             `json:"minUpMembers"`
	MinUpMembersAction     string          `json:"minUpMembersAction"`
	MinUpMembersChecking   string          `json:"minUpMembersChecking"`
	Monitor                string          `json:"monitor"`
	QueueDepthLimit        int             `json:"queueDepthLimit"`
	QueueOnConnectionLimit string          `json:"queueOnConnectionLimit"`
	QueueTimeLimit         int             `json:"queueTimeLimit"`
	ReselectTries          int             `json:"reselectTries"`
	ServiceDownAction      string          `json:"serviceDownAction"`
	SlowRampTime           int             `json:"slowRampTime"`
	MemberRef              LBPoolMemberRef `json:"membersReference"`
}

type LBPools struct {
	Items []LBPool `json:"items"`
}

func (f *Device) ShowPools() (error, *LBPools) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool"
	res := LBPools{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowPool(pname string) (error, *LBPool) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"
	res := LBPool{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddPool() (error, *LBPool) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool"
	res := LBPool{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a pool struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// post the request
	err, resp := f.sendRequest(u, POST, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdatePool(pname string) (error, *LBPool) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool
	res := LBPool{}
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
	err, resp := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeletePool(pname string) (error, *Response) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &resp
	}

}

func (f *Device) ShowPoolMembers(pname string) (error, *LBPoolMembers) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddPoolMembers(pname string) (error, *LBPoolMembers) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}
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
	err, resp := f.sendRequest(u, POST, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdatePoolMembers(pname string) (error, *LBPoolMembers) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := LBPoolMembers{}
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
	err, resp := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeletePoolMembers(pname string) (error, *Response) {

	pool := strings.Replace(pname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members"
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) OnlinePoolMember(mname string) (error, *Response) {

	pmember := strings.Replace(mname, "/", "~", -1)
	pool := strings.Replace(f5Pool, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members/" + pmember
	res := json.RawMessage{}

	/*
	   {"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)
	   {"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)
	   {"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)
	*/
	body := MemberState{"user-up", "user-enabled"}

	// put the request
	err, resp := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &resp
	}

}

func (f *Device) OfflinePoolMember(mname string) (error, *Response) {

	pmember := strings.Replace(mname, "/", "~", -1)
	pool := strings.Replace(f5Pool, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/pool/" + pool + "/members/" + pmember
	res := json.RawMessage{}

	/*
	   {"state": "user-down", "session": "user-disabled"} (Member Forced Offline in GUI)
	   {"state": "user-up", "session": "user-disabled"} (Member Disabled in GUI)
	   {"state": "user-up", "session": "user-enabled"}  (Member Enabled in GUI)
	*/
	var body MemberState
	if now {
		body = MemberState{"user-down", "user-disabled"}
	} else {
		body = MemberState{"user-up", "user-disabled"}
	}

	// put the request
	err, resp := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

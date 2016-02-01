package f5

import (
	"encoding/json"
	"strings"
)

type LBVirtualPolicy struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
}

type LBVirtualPoliciesRef struct {
	Items []LBVirtualPolicy `json":items"`
}

type LBVirtualProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
	Context   string `json:"context"`
}

type LBVirtualPersistProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	TmDefault string `json:"tmDefault"`
}

type LBVirtualProfileRef struct {
	Items []LBVirtualProfile `json":items"`
}

type LBVirtual struct {
	Name             string                    `json:"name"`
	FullPath         string                    `json:"fullPath"`
	Partition        string                    `json:"partition"`
	Destination      string                    `json:"destination"`
	Pool             string                    `json:"pool"`
	AddressStatus    string                    `json:"addressStatus"`
	AutoLastHop      string                    `json:"autoLasthop"`
	CmpEnabled       string                    `json:"cmpEnabled"`
	ConnectionLimit  int                       `json:"connectionLimit"`
	Enabled          bool                      `json:"enabled"`
	IpProtocol       string                    `json:"ipProtocol"`
	Source           string                    `json:"source"`
	SourcePort       string                    `json:"sourcePort"`
	SynCookieStatus  string                    `json:"synCookieStatus"`
	TranslateAddress string                    `json:"translateAddress"`
	TranslatePort    string                    `json:"translatePort"`
	Profiles         LBVirtualProfileRef       `json:"profilesReference"`
	Policies         LBVirtualPoliciesRef      `json:"policiesReference"`
	Rules            []string                  `json:"rules"`
	Persist          []LBVirtualPersistProfile `json:"persist"`
}

type LBVirtuals struct {
	Items []LBVirtual
}

func (f *Device) ShowVirtuals() (error, *LBVirtuals) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual"
	res := LBVirtuals{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowVirtual(vname string) (error, *LBVirtual) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddVirtual(virt *json.RawMessage) (error, *LBVirtual) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual"
	res := LBVirtual{}

	// post the request
	err, _ := f.sendRequest(u, POST, virt, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdateVirtual(vname string, body *json.RawMessage) (error, *LBVirtual) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname
	res := LBVirtual{}

	// put the request
	err, _ := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeleteVirtual(vname string) (error, *Response) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, resp
	}

}

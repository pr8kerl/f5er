package f5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
{
  "kind":"tm:ltm:monitor:http:httpcollectionstate",
  "selfLink":"https://localhost/mgmt/tm/ltm/monitor/http?ver=11.6.0",
  "items":[
    {"kind":"tm:ltm:monitor:http:httpstate",
     "name":"http",
     "partition":"Common",
     "fullPath":"/Common/http",
     "generation":0,
     "selfLink":"https://localhost/mgmt/tm/ltm/monitor/http/~Common~http?ver=11.6.0",
     "adaptive":"disabled",
     "adaptiveDivergenceType":"relative",
     "adaptiveDivergenceValue":25,
     "adaptiveLimit":200,
     "adaptiveSamplingTimespan":300,
     "destination":"*:*",
     "interval":5,
     "ipDscp":0,
     "manualResume":"disabled",
     "reverse":"disabled",
     "send":"GET /\\r\\n",
     "timeUntilUp":0,
     "timeout":16,
     "transparent":"disabled",
     "upInterval":0},
     {
     "kind":"tm:ltm:monitor:http:httpstate",
     "name":"http_head_f5",
     "partition":"Common",
     "fullPath":"/Common/http_head_f5",
     "generation":0,
     "selfLink":"https://localhost/mgmt/tm/ltm/monitor/http/~Common~http_head_f5?ver=11.6.0",
     "adaptive":"disabled",
     "adaptiveDivergenceType":"relative",
     "adaptiveDivergenceValue":25,
     "adaptiveLimit":200,
     "adaptiveSamplingTimespan":300,
     "defaultsFrom":"/Common/http",
     "destination":"*:*",
     "interval":5,
     "ipDscp":0,
     "manualResume":"disabled",
     "recv":"Server\\:",
     "reverse":"disabled",
     "send":"HEAD / HTTP/1.0\\r\\n\\r\\n",
     "timeUntilUp":0,
     "timeout":16,
     "transparent":"disabled",
     "upInterval":0},
     }
   ]
}

*/

type LBMonitorHttp struct {
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	FullPath                 string `json:"fullPath"`
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	DefaultsFrom             string `json:"defaultsFrom"`
	Destination              string `json:"destination"`
	Interval                 int    `json:"interval"`
	IpDscp                   int    `json:"ipDscp"`
	ManualResume             string `json:"manualResume"`
	Recv                     string `json:"recv"`
	Reverse                  string `json:"reverse"`
	Send                     string `json:"send"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

type LBMonitorHttpRef struct {
	Items []LBMonitorHttp `json":items"`
}

func (f *Device) ShowMonitorsHttp() (error, *LBMonitorHttpRef) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/monitor/http"
	res := LBMonitorHttpRef{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowMonitorHttp(vname string) (error, *LBMonitorHttp) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/monitor/http/" + vname + "?expandSubcollections=true"
	res := LBMonitorHttp{}

	err, resp := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) AddMonitorHttp() (error, *LBMonitorHttp) {

	u := "https://" + f.Hostname + "/mgmt/tm/ltm/monitor/http"
	res := LBMonitorHttp{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a virtual struct
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

func (f *Device) UpdateMonitorHttp(vname string) (error, *LBMonitorHttp) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/monitor/http/" + vname
	res := LBMonitorHttp{}
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a virtual struct
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

func (f *Device) DeleteMonitorHttp(vname string) (error, *Response) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := "https://" + f.Hostname + "/mgmt/tm/ltm/monitor/http/" + vname
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &resp
	}

}

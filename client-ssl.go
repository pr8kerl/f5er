package main

import (
	"fmt"
	"strings"
	//	"github.com/kr/pretty"
	"encoding/json"
	"io/ioutil"
	"log"
)

type LBClientSsl struct {
	Name                            string           `json:"name"`
	Partition                       string           `json:"partition"`
	Fullpath                        string           `json:"fullPath"`
	Generation                      int              `json:"generation"`
	AlertTimeout                    string           `json:"alertTimeout"`
	AllowNonSsl                     string           `json:"allowNonSsl"`
	Authenticate                    string           `json:"authenticate"`
	AuthenticateDepth               int              `json:"authenticateDepth"`
	CacheSize                       int              `json:"cacheSize"`
	CacheTimeout                    int              `json:"cacheTimeout"`
	Cert                            string           `json:"cert"`
	CertExtensionIncludes           []string         `json:"certExtensionIncludes"`
	CertLifespan                    int              `json:"certLifespan"`
	CertLookupByIpaddrPort          string           `json:"certLookupByIpaddrPort"`
	Chain                           string           `json:"chain"`
	Ciphers                         string           `json:"ciphers"`
	DefaultsFrom                    string           `json:"defaultsFrom"`
	ForwardProxyBypassDefaultAction string           `json:"forwardProxyBypassDefaultAction"`
	GenericAlert                    string           `json:"genericAlert"`
	HandshakeTimeout                string           `json:"handshakeTimeout"`
	InheritCertkeychain             string           `json:"inheritCertkeychain"`
	Key                             string           `json:"key"`
	MaxRenegotiationsPerMinute      int              `json:"maxRenegotiationsPerMinute"`
	ModSslMethods                   string           `json:"modSslMethods"`
	Mode                            string           `json:"mode"`
	TmOptions                       []string         `json:"tmOptions"`
	PeerCertMode                    string           `json:"peerCertMode"`
	PeerNoRenegotiateTimeout        string           `json:"peerNoRenegotiateTimeout"`
	ProxySsl                        string           `json:"proxySsl"`
	ProxySslPassthrough             string           `json:"proxySslPassthrough"`
	RenegotiateMaxRecordDelay       string           `json:"renegotiateMaxRecordDelay"`
	RenegotiatePeriod               string           `json:"renegotiatePeriod"`
	RenegotiateSize                 string           `json:"renegotiateSize"`
	Renegotiation                   string           `json:"renegotiation"`
	RetainCertificate               string           `json:"retainCertificate"`
	SecureRenegotiation             string           `json:"secureRenegotiation"`
	SessionMirroring                string           `json:"sessionMirroring"`
	SessionTicket                   string           `json:"sessionTicket"`
	SniDefault                      string           `json:"sniDefault"`
	SniRequire                      string           `json:"sniRequire"`
	SslForwardProxy                 string           `json:"sslForwardProxy"`
	SslForwardProxyBypass           string           `json:"sslForwardProxyBypass"`
	SslSignHash                     string           `json:"sslSignHash"`
	StrictResume                    string           `json:"strictResume"`
	UncleanShutdown                 string           `json:"uncleanShutdown"`
	CertKeyChain                    []LBCertKeyChain `json:"certKeyChain"`
}

type LBCertKeyChain struct {
	Name  string `json:"name"`
	Cert  string `json:"cert"`
	Chain string `json:"chain"`
	Key   string `json:"key"`
}

type LBClientSsls struct {
	Items []LBClientSsl `json:"items"`
}

func showClientSsls() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/client-ssl"
	res := LBClientSsls{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		fmt.Printf("%s\n", v.Fullpath)
	}
}

func showClientSsl(cname string) {

	client := strings.Replace(cname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/client-ssl/" + client
	res := LBClientSsl{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addClientSsl() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/client-ssl"
	res := LBClientSsl{}
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

func updateClientSsl(cname string) {

	client := strings.Replace(cname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/client-ssl/" + client
	res := LBClientSsl{}
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

func deleteClientSsl(cname string) {

	client := strings.Replace(cname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/client-ssl/" + client
	res := json.RawMessage{}

	err, resp := SendRequest(u, DELETE, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, cname)
	}

}

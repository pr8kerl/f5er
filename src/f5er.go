// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released
// under the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for
// details.  Resist intellectual serfdom - the ownership of ideas is akin to
// slavery.

// Example demonstrating use of package napping, with HTTP Basic
// authentictation over HTTPS, to retrieve a Github auth token.
package main

/*

NOTE: This example may only work on *nix systems due to gopass requirements.

*/

import (
	"code.google.com/p/gopass"
	"fmt"
	"github.com/jmcvetta/napping"
	//	"github.com/kr/pretty"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	//	"time"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	//
	// Prompt user for Github username/password
	//
	var username string
	fmt.Printf("f5 username: ")
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		log.Fatal(err)
	}
	passwd, err := gopass.GetPass("f5 password: ")
	if err != nil {
		log.Fatal(err)
	}
	//
	// Compose request
	//
	// http://developer.github.com/v3/oauth/#create-a-new-authorization
	//
	//payload := struct {
	//		Scopes []string `json:"scopes"`
	//		Note   string   `json:"note"`
	//	}{
	//		Scopes: []string{"public_repo"},
	//		Note:   "testing Go napping" + time.Now().String(),
	//	}

	type LBPool struct {
		Name      string `json:"name"`
		Partition string `json:"partition"`
		Fullpath  string `json:"fullPath"`
	}

	/*
		{
						"kind":"tm:ltm:pool:poolstate",
						"name":"audmzbilltweb-sit_443_pool",
						"partition":"DMZ",
						"fullPath":"/DMZ/audmzbilltweb-sit_443_pool",
						"generation":156,
						"selfLink":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool?ver=11.6.0",
						"allowNat":"yes",
						"allowSnat":"yes",
						"ignorePersistedWeight":"disabled",
						"ipTosToClient":"pass-through",
						"ipTosToServer":"pass-through",
						"linkQosToClient":"pass-through",
						"linkQosToServer":"pass-through",
						"loadBalancingMode":"round-robin",
						"minActiveMembers":0,
						"minUpMembers":0,
						"minUpMembersAction":"failover",
						"minUpMembersChecking":"disabled",
						"monitor":"/Common/https ",
						"queueDepthLimit":0,
						"queueOnConnectionLimit":"disabled",
						"queueTimeLimit":0,
						"reselectTries":0,
						"serviceDownAction":"none",
						"slowRampTime":10,
						"membersReference":{
										"link":"https://localhost/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0",
										"isSubcollection":true
						}
		}
	*/

	//
	// Struct to hold response data
	//
	type ResponseUserAgent struct {
		Useragent string `json:"user-agent"`
	}
	//
	// Struct to hold response data
	//
	//	res := struct {
	//		Id        int
	//		Url       string
	//		Scopes    []string
	//		Token     string
	//		App       map[string]string
	//		Note      string
	//		NoteUrl   string `json:"note_url"`
	//		UpdatedAt string `json:"updated_at"`
	//		CreatedAt string `json:"created_at"`
	//	}{}
	//
	// Struct to hold error response
	//
	e := struct {
		Message string
		Errors  []struct {
			Resource string
			Field    string
			Code     string
		}
	}{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//
	// Setup HTTP Basic auth for this session (ONLY use this with SSL).  Auth
	// can also be configured on a per-request basis when using Send().
	//
	s := napping.Session{
		Client:   client,
		Userinfo: url.UserPassword(username, passwd),
	}
	url := "https://10.60.99.241/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool"
	//
	// Send request to server
	//
	res := LBPool{}
	resp, err := s.Get(url, nil, &res, &e)
	if err != nil {
		log.Fatal(err)
	}
	//
	// Process response
	//
	//	println("")
	//	if resp.Status() == 201 {
	//		fmt.Printf("Github auth token: %s\n\n", res.Token)
	//	} else {
	//		fmt.Println("Bad response status from Github server")
	//		fmt.Printf("\t Status:  %v\n", resp.Status())
	//		fmt.Printf("\t Message: %v\n", e.Message)
	//		fmt.Printf("\t Errors: %v\n", e.Message)
	//		pretty.Println(e.Errors)
	//	}
	//	println("")

	//
	// Process response
	//
	println("")
	fmt.Printf("pool fullpath: %s\n\n", res.Fullpath)
	println("")
	fmt.Println("response Status:", resp.Status())
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Header")
	fmt.Println(resp.HttpResponse().Header)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("RawText")
	fmt.Println(resp.RawText())
	println("")
}

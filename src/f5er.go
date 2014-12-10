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
	//	"github.com/kr/pretty"
	"code.google.com/p/gopass"
	"crypto/tls"
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	//"strings"
	//	"time"
)

var (
	f5Host    string
	username  string
	passwd    string
	partition string
	poolname  string
	cfgFile   string = "f5.json"
)

var f5Cmd = &cobra.Command{
	Use:   "f5er",
	Short: "tickle an F5 load balancer using REST",
	Long:  "A utility to create and manage F5 configuration objects",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show F5 objects",
	Long:  "show current state of F5 objects. Show requires an object, eg. f5er show pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			bail("show needs an argument - the object type to show perhaps??")
		}
	},
}

var showPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "show a pool",
	Long:  "show the current state of a pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("f5er show pool")
			showPools()
		} else {
			pname := args[0]
			fmt.Println("f5er show pool " + pname)
			showPool(pname)
		}
	},
}

type httperr struct {
	Message string
	Errors  []struct {
		Resource string
		Field    string
		Code     string
	}
}

type LBPool struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	Fullpath  string `json:"fullPath"`
}

type LBPools struct {
	Items []LBPool `json:"items"`
}

func InitialiseConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	if f5Host != "" {
		viper.Set("f5", f5Host)
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can't find your config file: %s", cfgFile)
	}
	checkRequiredFlag("f5")
	f5Host = viper.GetString("f5")
	username = viper.GetString("username")
	passwd = viper.GetString("passwd")
	partition = viper.GetString("partition")
	poolname = viper.GetString("poolname")
}

func checkRequiredFlag(flg string) {

	if !viper.IsSet(flg) {
		log.SetFlags(0)
		log.Fatalf("\nerror: missing required option --%s\n\n", flg)
	}
	if !viper.IsSet("username") {
		promptForCreds()
	}

}

func promptForCreds() {
	//
	// Prompt user for f5 username/password
	//
	fmt.Printf("No login credentials defined in config - prompting...\n")
	fmt.Printf("f5 username: ")
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		log.Fatal(err)
	}
	passwd, err = gopass.GetPass("f5 password: ")
	if err != nil {
		log.Fatal(err)
	}
}

func bail(msg string) {
	log.SetFlags(0)
	log.Fatalf("\n%s\n\n", msg)
}

func showPools() {

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

	url := "https://" + f5Host + "/mgmt/tm/ltm/pool"
	res := LBPools{}

	//url := "https://10.60.99.241/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool"
	//url := "https://10.60.99.241/mgmt/tm/ltm/pool"
	//url := "https://10.60.99.242/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0"
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := s.Get(url, nil, &res, &e)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res.Items {
		fmt.Printf("pool\npartition:\t%s\n", v.Partition)
		fmt.Printf("name:\t\t%s\n", v.Name)
		fmt.Printf("fullpath:\t\t%s\n\n", v.Fullpath)
	}
	fmt.Println("response Status:", resp.Status())
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Header")
	fmt.Println(resp.HttpResponse().Header)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("RawText")
	fmt.Println(resp.RawText())
	println("")
}

func showPool(pname string) {

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

	//	url := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + poolname + "?\\$expand=*"
	url := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + poolname
	res := LBPool{}

	//url := "https://10.60.99.241/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool"
	//url := "https://10.60.99.241/mgmt/tm/ltm/pool"
	//url := "https://10.60.99.242/mgmt/tm/ltm/pool/~DMZ~audmzbilltweb-sit_443_pool/members?ver=11.6.0"
	//
	// Send request to server
	//
	e := httperr{}
	resp, err := s.Get(url, nil, &res, &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("url:\t%s\n", url)
	fmt.Printf("pool\npartition:\t%s\n", res.Partition)
	fmt.Printf("name:\t\t%s\n", res.Name)
	fmt.Printf("fullpath:\t\t%s\n\n", res.Fullpath)
	fmt.Println("response Status:", resp.Status())
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Header")
	fmt.Println(resp.HttpResponse().Header)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("RawText")
	fmt.Println(resp.RawText())
	println("")
}

func init() {
	f5Cmd.Flags().StringVarP(&f5Host, "f5", "f", "", "IP or hostname of F5 to poke")
	viper.BindPFlag("f5", f5Cmd.Flags().Lookup("f5"))
	f5Cmd.AddCommand(showCmd)
	showCmd.AddCommand(showPoolCmd)
	//	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetFlags(0)
	InitialiseConfig()
}

func main() {
	viper.AutomaticEnv()
	f5Cmd.Execute()
}

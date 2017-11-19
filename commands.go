package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pr8kerl/f5er/f5"
	"github.com/spf13/cobra"
)

var f5Cmd = &cobra.Command{
	Use:   "f5er",
	Short: "tickle an F5 load balancer using REST",
	Long:  "A utility to manage F5 configuration objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkFlags(cmd)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show current version",
	Long:  "show compiled version of f5er binary",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("f5er %s commit %s\n", f5.GetVersion(), commit)
	},
}
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show F5 objects",
	Long:  "show current state of F5 objects. Show requires an object, eg. f5er show pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("all available ltm modules")
			show()
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add F5 objects",
	Long:  "add or create a new F5 object. Add requires an object, eg. f5er add pool",
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update F5 objects",
	Long:  "update an existing F5 object. Update requires an object, eg. f5er update pool",
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete F5 objects",
	Long:  "delete an F5 object. Delete requires an object, eg. f5er delete pool",
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

var offlineCmd = &cobra.Command{
	Use:   "offline",
	Short: "offline a pool member",
	Long:  "offline an F5 pool member. eg. f5er offline poolmember /partition/poolmember",
	Run: func(cmd *cobra.Command, args []string) {
		offline()
	},
}

var onlineCmd = &cobra.Command{
	Use:   "online",
	Short: "online a pool member",
	Long:  "online an F5 pool member. eg. f5er online poolmember /partition/poolmember",
	Run: func(cmd *cobra.Command, args []string) {
		online()
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "get F5 statistics",
	Long:  "get statistics for F5 objects.",
	Run: func(cmd *cobra.Command, args []string) {
		appliance.SetStatsPathPrefix(statsPathPrefix)
		if len(args) < 1 {
			stats()
		}
	},
}

var showDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "show an f5 device",
	Long:  "show the current state of an f5 device",
	Run: func(cmd *cobra.Command, args []string) {
		err, res := appliance.ShowDevice()
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var showPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "show a pool",
	Long:  "show the current state of a pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowPools()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowPool(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "add a pool",
	Long:  "add a new pool",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddPool(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)

	},
}

var updatePoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "update a pool",
	Long:  "update an existing pool",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update pool requires a pool name as an argument (ie /partition/poolname )")
		} else {
			pname := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdatePool(pname, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)

		}
	},
}

var deletePoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "delete a pool",
	Long:  "delete an existing pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete pool requires a pool name as an argument (ie /partition/poolname )")
		} else {
			name := args[0]
			err, res := appliance.DeletePool(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var statsPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "show pool statistics",
	Long:  "show the current statistics of a pool",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err, res := appliance.StatsPools()
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		} else {
			name := args[0]
			err, res := appliance.StatsPool(name)
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		}
	},
}

var showPoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "show pool members",
	Long:  "show the pool members of a given pool",
	Run: func(cmd *cobra.Command, args []string) {
		//		if !viper.IsSet("pool") {
		//			log.Fatal("show poolmember needs a pool to work with - please use --pool flag.")
		//		}
		if len(args) < 1 {
			log.Fatal("show poolmember requires a pool as an argument - in the form of /partition/poolname")
		} else {
			name := args[0]
			err, res := appliance.ShowPoolMembers(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res.Items)
		}
	},
}

var addPoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "add poolmember",
	Long:  "add poolmember",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("add poolmember requires a pool name as an argument (ie /partition/poolname )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.AddPoolMembers(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var updatePoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "update poolmembers",
	Long:  "update existing poolmembers",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update poolmember requires a pool name as an argument (ie /partition/poolname )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdatePoolMembers(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deletePoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "delete poolmembers",
	Long:  "delete existing poolmembers",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("delete poolmember requires a pool name as an argument (ie /partition/poolname )")
		} else {
			name := args[0]
			err, res := appliance.DeletePoolMembers(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var offlinePoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "offline a poolmember",
	Long:  "offline an existing poolmember",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("pool")
		if len(args) < 1 {
			log.Fatal("offline poolmember requires a poolmember name as an argument (ie /partition/poolmember )")
		} else {
			name := args[0]
			err, res := appliance.OfflinePoolMember(f5Pool, name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var onlinePoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "online a poolmember",
	Long:  "online an existing poolmember",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("pool")
		if len(args) < 1 {
			log.Fatal("online poolmember requires a poolmember name as an argument (ie /partition/poolmember )")
		} else {
			name := args[0]
			err, res := appliance.OnlinePoolMember(f5Pool, name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var statsPoolMembersCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "show poolmember statistics",
	Long:  "show the current statistics of all members in a pool",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err, res := appliance.StatsCommonPoolMembers()
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		} else {
			name := args[0]
			err, res := appliance.StatsPoolMembers(name)
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		}
	},
}

var showVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "show a virtual server",
	Long:  "show the current state of a virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowVirtuals()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowVirtual(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "add a virtual server",
	Long:  "add a new virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddVirtual(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "update a virtual server",
	Long:  "update an existing virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update virtual requires a virtual server name as an argument (ie /partition/virtualservername )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdateVirtual(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "delete a virtual server",
	Long:  "delete a virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete virtual requires a virtual server name as an argument (ie /partition/virtualservername )")
		} else {
			name := args[0]
			err, res := appliance.DeleteVirtual(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var statsVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "show virtual server statistics",
	Long:  "show the current statistics of a virtual server",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err, res := appliance.StatsVirtuals()
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		} else {
			name := args[0]
			err, res := appliance.StatsVirtual(name)
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		}
	},
}

var showPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "show a policy",
	Long:  "show the current state of a policy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowPolicies()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowPolicy(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "add a policy",
	Long:  "add a new policy",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddPolicy(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updatePolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "update a policy",
	Long:  "update an existing F5 policy",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update policy requires a policy name as an argument (ie /partition/policy )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdatePolicy(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deletePolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "delete a policy",
	Long:  "delete a policy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete policy requires a policy name as an argument (ie /partition/policy )")
		} else {
			name := args[0]
			err, res := appliance.DeletePolicy(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var showNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "show a node",
	Long:  "show the current state of a node",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowNodes()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowNode(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "add a node",
	Long:  "add a new node",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddNode(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "update a node",
	Long:  "update an existing F5 node",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update node requires a node name as an argument (ie /partition/nodename )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdateNode(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "delete a node",
	Long:  "delete a node",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete node requires a node name as an argument (ie /partition/nodename )")
		} else {
			name := args[0]
			err, res := appliance.DeleteNode(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var statsNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "show node statistics",
	Long:  "show the current statistics of a node",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err, res := appliance.StatsNodes()
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		} else {
			name := args[0]
			err, res := appliance.StatsNode(name)
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		}
	},
}

var showRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "show a rule",
	Long:  "show the details of a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowRules()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowRule(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "add a rule",
	Long:  "add a new rule",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddRule(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "update a rule",
	Long:  "update an existing F5 rule",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update rule requires a rule name as an argument (ie /partition/rulename )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdateRule(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "delete a rule",
	Long:  "delete a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete rule requires a rule name as an argument (ie /partition/rulename )")
		} else {
			name := args[0]
			err, res := appliance.DeleteRule(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var statsRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "show rule statistics",
	Long:  "show the current statistics of a rule",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err, res := appliance.StatsRules()
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		} else {
			name := args[0]
			err, res := appliance.StatsRule(name)
			if err != nil {
				log.Fatal(err)
			}
			for _, datapoint := range res {
				fmt.Printf("%s\n", datapoint.String())
			}
		}
	},
}

var showProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "show profiles",
	Long:  "show profiles .\nProvide a profile type or a profile name with the full path like so: server-ssl/~partition~custom_server_ssl_name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowProfiles()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.Reference.Link)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowProfile(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var showServerSslCmd = &cobra.Command{
	Use:   "server-ssl",
	Short: "show a server-ssl profile",
	Long:  "show the details of a server-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowServerSsls()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowServerSsl(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addServerSslCmd = &cobra.Command{
	Use:   "server-ssl",
	Short: "add a server-ssl profile",
	Long:  "add a new server-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddServerSsl(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateServerSslCmd = &cobra.Command{
	Use:   "server-ssl",
	Short: "update a server-ssl profile",
	Long:  "update an existing F5 server-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update server-ssl requires a server-ssl profile name as an argument (ie /partition/profilename )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdateServerSsl(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteServerSslCmd = &cobra.Command{
	Use:   "server-ssl",
	Short: "delete a server-ssl profile",
	Long:  "delete a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete server-ssl requires a server-ssl profile name as an argument (ie /partition/profilename )")
		} else {
			name := args[0]
			err, res := appliance.DeleteServerSsl(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var showClientSslCmd = &cobra.Command{
	Use:   "client-ssl",
	Short: "show a client-ssl profile",
	Long:  "show the details of a client-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowClientSsls()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowClientSsl(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addClientSslCmd = &cobra.Command{
	Use:   "client-ssl",
	Short: "add a client-ssl profile",
	Long:  "add a new client-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddClientSsl(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateClientSslCmd = &cobra.Command{
	Use:   "client-ssl",
	Short: "update a client-ssl profile",
	Long:  "update an existing F5 client-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update client-ssl requires a client-ssl profile name as an argument (ie /partition/profilename )")
		} else {
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			name := args[0]
			err, res := appliance.UpdateClientSsl(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteClientSslCmd = &cobra.Command{
	Use:   "client-ssl",
	Short: "delete a client-ssl profile",
	Long:  "delete a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete client-ssl requires a client-ssl profile name as an argument (ie /partition/profilename )")
		} else {
			name := args[0]
			err, res := appliance.DeleteClientSsl(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var showMonitorHttpCmd = &cobra.Command{
	Use:   "monitor-http",
	Short: "show a monitor-http profile",
	Long:  "show the details of a monitor-http profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			err, res := appliance.ShowMonitorsHttp()
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Items {
				fmt.Printf("%s\n", v.FullPath)
			}
		} else {
			name := args[0]
			err, res := appliance.ShowMonitorHttp(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var addMonitorHttpCmd = &cobra.Command{
	Use:   "monitor-http",
	Short: "add a monitor-http profile",
	Long:  "add a new monitor-http profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		body := json.RawMessage{}
		// read in input file
		dat, err := ioutil.ReadFile(f5Input)
		if err != nil {
			log.Fatal(err)
		}
		// convert bytes to a json message
		err = json.Unmarshal(dat, &body)
		if err != nil {
			log.Fatal(err)
		}
		err, res := appliance.AddMonitorHttp(&body)
		if err != nil {
			log.Fatal(err)
		}
		appliance.PrintObject(res)
	},
}

var updateMonitorHttpCmd = &cobra.Command{
	Use:   "monitor-http",
	Short: "update a monitor-http profile",
	Long:  "update an existing F5 monitor-http profile",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		if len(args) < 1 {
			log.Fatal("update monitor-http requires a monitor-http profile name as an argument (ie /partition/monitorname )")
		} else {
			name := args[0]
			body := json.RawMessage{}
			// read in input file
			dat, err := ioutil.ReadFile(f5Input)
			if err != nil {
				log.Fatal(err)
			}
			// convert bytes to a json message
			err = json.Unmarshal(dat, &body)
			if err != nil {
				log.Fatal(err)
			}
			err, res := appliance.UpdateMonitorHttp(name, &body)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var deleteMonitorHttpCmd = &cobra.Command{
	Use:   "monitor-http",
	Short: "delete a monitor-http profile",
	Long:  "delete a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("delete monitor-http requires a monitor-http profile name as an argument (ie /partition/profilename )")
		} else {
			name := args[0]
			err, res := appliance.DeleteMonitorHttp(name)
			if err != nil {
				log.Fatal(err)
			}
			appliance.PrintObject(res)
		}
	},
}

var showStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "show a stack transaction",
	Long:  "show a stack transaction",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		showStack()
	},
}

var addStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "add a stack",
	Long:  "add a new stack",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		addStack()
	},
}

var updateStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "update a stack",
	Long:  "update a stack",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		updateStack()
	},
}

var deleteStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "delete a stack",
	Long:  "delete a stack",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		deleteStack()
	},
}

var uploadFileCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload a file",
	Long:  "upload a file to the f5 server",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		dat, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Uploading file", filepath.Base(filename))
		err = appliance.UploadFile(filepath.Base(filename), dat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done")
	},
}

func PrintCerts(cert *f5.SSLCertificate) {
	fmt.Printf("Name: %v Partition: %s\n", cert.Name, cert.Partition)
	fmt.Printf("Issuer: %s\n", cert.Issuer)
	fmt.Printf("Subject: %s\n", cert.Subject)
	fmt.Printf("Strength: %d Curve: %s Type: %s\n", cert.KeySize, cert.CurveName, cert.KeyType)
	fmt.Printf("Checksum: %s\n", cert.Checksum)
	fmt.Printf("Uploaded: %s Expires %s\n", cert.CreateTime, cert.ExpireTime)
}

var showCertCmd = &cobra.Command{
	Use:   "cert",
	Short: "show a certificate",
	Long:  "show a certificate",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("partition then cert are required.")
		}
		partition := args[0]
		cert_name := args[1]
		err, cert := appliance.GetCertificate(partition, cert_name)
		if err != nil {
			log.Fatal(err)
		}
		PrintCerts(cert)
	},
}

var showCertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "show all certificates",
	Long:  "show all certificates",
	Run: func(cmd *cobra.Command, args []string) {
		err, certs := appliance.GetCertificates()
		if err != nil {
			log.Fatal(err)
		}
		for _, cert := range certs.Items {
			fmt.Println("")
			PrintCerts(&cert)
		}
	},
}

var addCertCmd = &cobra.Command{
	Use:   "cert",
	Short: "add a certificate [name, partition, uploaded_file_name]. Note the name should be the same for the key/cert and should not have a suffix (.crt/.key)",
	Long:  "add certificate to f5. note this does not create ssl profiles. uploaded_file_name should be a filename uploaded via f5 upload\nExample: f5 add cert mysite.com Common mysite.com.crt",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("[cert_name, partition, local_file]")
		}
		err, cert := appliance.CreateCertificateFromLocalFile(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		PrintCerts(cert)
	},
}

var addKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "add a certificate key [name, partition, uploaded_file_name]. Note the name should be the same for the key/cert and should not have a suffix (.crt/.key)",
	Long:  "add certificate key to f5. note this does not create ssl profiles. uploaded_file_name should be a filename uploaded via f5 upload\nExample: f5 add key mysite.com Common mysite.com.key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("[key_name, partition, local_file]")
		}
		err, cert := appliance.CreateKeyFromLocalFile(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
		PrintCerts(cert)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs a bash command on the f5",
	Long:  "Runs a bash command on the f5\nExample: f5 run \"ls -al\"",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Too little or too many arguements, Example: f5 run \"ls -al\"")
		}
		err, res := appliance.Run(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.CommandResult)
	},
}

func show() {

	err, mods := appliance.ShowModules()
	if err != nil {
		log.Fatalf("cannot show modules: %s\n", err)
	}
	for _, v := range mods.Items {
		fmt.Printf("module:\t%s\n", v.Reference.Link)
	}

}

func add() {
	fmt.Println("what sort of F5 object would you like to add?")
}

func update() {
	fmt.Println("what sort of F5 object would you like to update?")
}

func delete() {
	fmt.Println("what sort of F5 object would you like to delete?")
}

func offline() {
	fmt.Println("which pool member would you like to offline?")
}

func online() {
	fmt.Println("which pool member would you like to online?")
}

func stats() {

	fmt.Println("what sort of F5 object would you like stats for? (virtual, pool, node or rule)")

	/*
		err, res := appliance.Stats()
		if err != nil {
			log.Fatal(err)
		}
		for _, datapoint := range res {
			fmt.Printf("%s\n", datapoint.String())
		}
		if err != nil {
			log.Fatalf("cannot get statistics: %s\n", err)
		}
	*/

}

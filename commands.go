package main

import (
	"encoding/json"
	"fmt"
	"github.com/pr8kerl/f5er/f5"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var f5Cmd = &cobra.Command{
	Use:   "f5er",
	Short: "tickle an F5 load balancer using REST",
	Long:  "A utility to manage F5 configuration objects",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		checkRequiredFlag("f5")
	//	},
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

var showDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "show an f5 device",
	Long:  "show the current state of an f5 device",
	Run: func(cmd *cobra.Command, args []string) {
		appliance.ShowDevice()
	},
}

var showPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "show a pool",
	Long:  "show the current state of a pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowPools()
		} else {
			name := args[0]
			appliance.ShowPool(name)
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
		appliance.AddPool(&body)
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
			appliance.UpdatePool(pname, &body)
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
			appliance.DeletePool(name)
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
			appliance.ShowPoolMembers(name)
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
			appliance.AddPoolMembers(name, &body)
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
			appliance.UpdatePoolMembers(name, &body)
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
			appliance.DeletePoolMembers(name)
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
			appliance.OfflinePoolMember(f5Pool, name)
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
			appliance.OnlinePoolMember(f5Pool, name)
		}
	},
}

var showVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "show a virtual server",
	Long:  "show the current state of a virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowVirtuals()
		} else {
			name := args[0]
			appliance.ShowVirtual(name)
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
		appliance.AddVirtual(&body)
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
			appliance.UpdateVirtual(name, &body)
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
			appliance.DeleteVirtual(name)
		}
	},
}

var showPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "show a policy",
	Long:  "show the current state of a policy",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowPolicies()
		} else {
			name := args[0]
			appliance.ShowPolicy(name)
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
		appliance.AddPolicy(&body)
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
			appliance.UpdatePolicy(name, &body)
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
			appliance.DeletePolicy(name)
		}
	},
}

var showNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "show a node",
	Long:  "show the current state of a node",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowNodes()
		} else {
			name := args[0]
			appliance.ShowNode(name)
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
		appliance.AddNode(&body)
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
			appliance.UpdateNode(name, &body)
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
			appliance.DeleteNode(name)
		}
	},
}

var showRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "show a rule",
	Long:  "show the details of a rule",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowRules()
		} else {
			name := args[0]
			appliance.ShowRule(name)
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
		appliance.AddRule(&body)
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
			appliance.UpdateRule(name, &body)
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
			appliance.DeleteRule(name)
		}
	},
}

var showServerSslCmd = &cobra.Command{
	Use:   "server-ssl",
	Short: "show a server-ssl profile",
	Long:  "show the details of a server-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowServerSsls()
		} else {
			name := args[0]
			appliance.ShowServerSsl(name)
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
		appliance.AddServerSsl(&body)
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
			appliance.UpdateServerSsl(name, &body)
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
			appliance.DeleteServerSsl(name)
		}
	},
}

var showClientSslCmd = &cobra.Command{
	Use:   "client-ssl",
	Short: "show a client-ssl profile",
	Long:  "show the details of a client-ssl profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowClientSsls()
		} else {
			name := args[0]
			appliance.ShowClientSsl(name)
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
		appliance.AddClientSsl(&body)
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
			appliance.UpdateClientSsl(name, &body)
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
			appliance.DeleteClientSsl(name)
		}
	},
}

var showMonitorHttpCmd = &cobra.Command{
	Use:   "monitor-http",
	Short: "show a monitor-http profile",
	Long:  "show the details of a monitor-http profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			appliance.ShowMonitorsHttp()
		} else {
			name := args[0]
			appliance.ShowMonitorHttp(name)
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
		appliance.AddMonitorHttp(&body)
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
			appliance.UpdateMonitorHttp(name, &body)
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
			appliance.DeleteMonitorHttp(name)
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

func show() {

	u := "https://" + f5Host + "/mgmt/tm/ltm"
	res := f5.LBModules{}
	appliance.ShowModules()

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%d: %s\n", resp.Status(), err)
	}

	for _, v := range res.Items {
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

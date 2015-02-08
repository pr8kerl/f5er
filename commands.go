package main

import (
	"fmt"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	"log"
)

var f5Cmd = &cobra.Command{
	Use:   "f5er",
	Short: "tickle an F5 load balancer using REST",
	Long:  "A utility to manage F5 configuration objects",
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
var showPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "show a pool",
	Long:  "show the current state of a pool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			showPools()
		} else {
			name := args[0]
			showPool(name)
		}
	},
}

var addPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "add a pool",
	Long:  "add a new pool",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		addPool()
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
			name := args[0]
			updatePool(name)
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
			deletePool(name)
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
			showPoolMembers(name)
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
			addPoolMembers(name)
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
			updatePoolMembers(name)
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
			deletePoolMembers(name)
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
			offlinePoolMember(name)
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
			onlinePoolMember(name)
		}
	},
}

var showVirtualCmd = &cobra.Command{
	Use:   "virtual",
	Short: "show a virtual server",
	Long:  "show the current state of a virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			showVirtuals()
		} else {
			name := args[0]
			showVirtual(name)
		}
	},
}

var showNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "show a node",
	Long:  "show the current state of a node",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			showNodes()
		} else {
			name := args[0]
			showNode(name)
		}
	},
}

var addNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "add a node",
	Long:  "add a new node",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlag("input")
		addNode()
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
			updateNode(name)
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
			deleteNode(name)
		}
	},
}

// F5 Module data struct
// to show all available modules when using show without args
type LBModule struct {
	Link string `json:"link"`
}

type LBModuleRef struct {
	Reference LBModule `json:"reference"`
}

type LBModules struct {
	Items []LBModuleRef `json:"items"`
}

func show() {

	u := "https://" + f5Host + "/mgmt/tm/ltm"
	res := LBModules{}

	err, resp := GetRequest(u, &res)
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

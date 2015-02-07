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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete F5 objects",
	Long:  "delete an F5 object. Delete requires an object, eg. f5er delete pool",
	Run: func(cmd *cobra.Command, args []string) {
		delete()
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

var showPoolMembersCmd = &cobra.Command{
	Use:   "poolmembers",
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

var addPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "add a pool",
	Long:  "add a new pool",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		addPool(name)
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

	err := GetRequest(u, &res)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res.Items {
		fmt.Printf("module:\t%s\n", v.Reference.Link)
	}

}

func add() {
	fmt.Println("what sort of F5 object would you like to add?")
}

func delete() {
	fmt.Println("what sort of F5 object would you like to delete?")
}

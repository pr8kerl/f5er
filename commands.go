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
	Long:  "A utility to create and manage F5 configuration objects",
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

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create F5 objects",
	Long:  "create a new F5 object. Create requires an object, eg. f5er create pool",
	Run: func(cmd *cobra.Command, args []string) {
		create()
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

var showPoolMemberCmd = &cobra.Command{
	Use:   "poolmember",
	Short: "show a pool member",
	Long:  "show the details of a pool member",
	Run: func(cmd *cobra.Command, args []string) {
		//		if !viper.IsSet("pool") {
		//			log.Fatal("show poolmember needs a pool to work with - please use --pool flag.")
		//		}
		if len(args) < 1 {
			showPoolMembers()
		} else {
			name := args[0]
			showPoolMember(name)
		}
	},
}

var createPoolCmd = &cobra.Command{
	Use:   "pool",
	Short: "create a pool",
	Long:  "create a new pool",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createPool(name)
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

func create() {

	fmt.Println("what sort of F5 object would you like to create?")

}

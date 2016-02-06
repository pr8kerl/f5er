package main

import (
	"github.com/jmcvetta/napping"
	"github.com/pr8kerl/f5er/f5"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	appliance *f5.Device
	f5Host    string
	username  string
	passwd    string
	cfgFile   string = "f5.json"
	f5Input   string
	f5Pool    string
	session   napping.Session
	transport *http.Transport
	client    *http.Client
	debug     bool
	now       bool
)

func InitialiseConfig() {

	viper.SetConfigFile(cfgFile)
	viper.AddConfigPath(".")
	viper.SetDefault("username", "admin")
	viper.SetDefault("debug", false)
	viper.SetDefault("force", false)

	viper.ReadInConfig()

	viper.SetEnvPrefix("f5")
	viper.BindEnv("device")
	viper.BindEnv("username")
	viper.BindEnv("passwd")

	viper.BindPFlag("device", f5Cmd.PersistentFlags().Lookup("f5"))
	viper.BindPFlag("debug", f5Cmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("pool", onlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("pool", offlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("input", f5Cmd.PersistentFlags().Lookup("input"))

	if f5Cmd.PersistentFlags().Lookup("f5").Changed {
		// use cmdline f5 flag if supplied
		viper.Set("device", f5Host)
	}
	if f5Cmd.PersistentFlags().Lookup("debug").Changed {
		viper.Set("debug", true)
	}
	if f5Cmd.PersistentFlags().Lookup("input").Changed {
		viper.Set("input", f5Input)
	}
	if offlinePoolMemberCmd.Flags().Lookup("pool").Changed {
		viper.Set("pool", f5Pool)
	}
	if offlinePoolMemberCmd.Flags().Lookup("now").Changed {
		viper.Set("now", true)
	}
	if onlinePoolMemberCmd.Flags().Lookup("pool").Changed {
		viper.Set("pool", f5Pool)
	}

	debug = viper.GetBool("debug")
	now = viper.GetBool("now")
	username = viper.GetString("username")
	passwd = viper.GetString("passwd")
	f5Host = viper.GetString("device")

	if username == "" {
		log.Fatalf("\nerror: missing username; use config file or F5_USERNAME environment variable\n\n")
	}
	if passwd == "" {
		log.Fatalf("\nerror: missing password; use config file or F5_PASSWD environment variable\n\n")
	}
	// finally check that f5 is not an empty string (default)
	//	if f5Host != "" {
	//		viper.Set("f5", f5Host)
	//	} else {
	//		f5Host = viper.GetString("f5")
	//	}
	if f5Host == "" {
		log.Fatalf("\nerror: missing f5 device hostname; use config file or F5_DEVICE environment variable\n\n")
	}

}

func checkRequiredFlag(flg string) {
	if !viper.IsSet(flg) {
		log.SetFlags(0)
		log.Fatalf("\nerror: missing required option --%s\n\n", flg)
	}
}

func init() {

	f5Cmd.PersistentFlags().StringVarP(&f5Host, "f5", "f", "", "IP or hostname of F5 to poke")
	f5Cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug output")
	f5Cmd.PersistentFlags().StringVarP(&f5Input, "input", "i", "", "input json f5 configuration")
	offlinePoolMemberCmd.Flags().StringVarP(&f5Pool, "pool", "p", "", "F5 pool name")
	offlinePoolMemberCmd.Flags().BoolVarP(&now, "now", "n", false, "force member offline immediately")
	onlinePoolMemberCmd.Flags().StringVarP(&f5Pool, "pool", "p", "", "F5 pool name")

	// show
	f5Cmd.AddCommand(showCmd)
	showCmd.AddCommand(showPoolCmd)
	showCmd.AddCommand(showPoolMemberCmd)
	showCmd.AddCommand(showVirtualCmd)
	showCmd.AddCommand(showNodeCmd)
	showCmd.AddCommand(showPolicyCmd)
	showCmd.AddCommand(showDeviceCmd)
	showCmd.AddCommand(showRuleCmd)
	showCmd.AddCommand(showProfileCmd)
	showCmd.AddCommand(showClientSslCmd)
	showCmd.AddCommand(showServerSslCmd)
	showCmd.AddCommand(showMonitorHttpCmd)
	showCmd.AddCommand(showStackCmd)

	// add
	f5Cmd.AddCommand(addCmd)
	addCmd.AddCommand(addPoolCmd)
	addCmd.AddCommand(addPoolMemberCmd)
	addCmd.AddCommand(addNodeCmd)
	addCmd.AddCommand(addPolicyCmd)
	addCmd.AddCommand(addVirtualCmd)
	addCmd.AddCommand(addRuleCmd)
	addCmd.AddCommand(addClientSslCmd)
	addCmd.AddCommand(addServerSslCmd)
	addCmd.AddCommand(addMonitorHttpCmd)
	addCmd.AddCommand(addStackCmd)

	// update
	f5Cmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updatePoolCmd)
	updateCmd.AddCommand(updatePoolMemberCmd)
	updateCmd.AddCommand(updateNodeCmd)
	updateCmd.AddCommand(updatePolicyCmd)
	updateCmd.AddCommand(updateVirtualCmd)
	updateCmd.AddCommand(updateRuleCmd)
	updateCmd.AddCommand(updateClientSslCmd)
	updateCmd.AddCommand(updateServerSslCmd)
	updateCmd.AddCommand(updateMonitorHttpCmd)
	updateCmd.AddCommand(updateStackCmd)

	// delete
	f5Cmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deletePoolCmd)
	deleteCmd.AddCommand(deletePoolMemberCmd)
	deleteCmd.AddCommand(deleteNodeCmd)
	deleteCmd.AddCommand(deletePolicyCmd)
	deleteCmd.AddCommand(deleteVirtualCmd)
	deleteCmd.AddCommand(deleteRuleCmd)
	deleteCmd.AddCommand(deleteClientSslCmd)
	deleteCmd.AddCommand(deleteServerSslCmd)
	deleteCmd.AddCommand(deleteMonitorHttpCmd)
	deleteCmd.AddCommand(deleteStackCmd)

	// offline
	f5Cmd.AddCommand(offlineCmd)
	offlineCmd.AddCommand(offlinePoolMemberCmd)

	// online
	f5Cmd.AddCommand(onlineCmd)
	onlineCmd.AddCommand(onlinePoolMemberCmd)

	//	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetFlags(0)

	InitialiseConfig()
	appliance = f5.New(f5Host, username, passwd)
}

func main() {
	//f5Cmd.DebugFlags()
	f5Cmd.Execute()
}

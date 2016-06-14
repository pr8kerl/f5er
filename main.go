package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"github.com/pr8kerl/f5er/f5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	appliance              *f5.Device
	f5Host                 string
	username               string
	passwd                 string
	cfgName                string = "f5"
	f5Input                string
	f5Pool                 string
	session                napping.Session
	transport              *http.Transport
	client                 *http.Client
	debug                  bool
	now                    bool
	stats_path_prefix      string
	stats_show_zero_values bool
)

func initialiseConfig() {

	viper.SetConfigName(cfgName)
	viper.AddConfigPath("$HOME/.f5")
	viper.AddConfigPath(".")
	viper.SetDefault("username", "admin")
	viper.SetDefault("debug", false)
	viper.SetDefault("force", false)
	viper.SetDefault("stats_path_prefix", "f5")
	viper.SetDefault("stats_show_zero_values", false)

	viper.SetEnvPrefix("f5")
	viper.BindEnv("device")
	viper.BindEnv("username")
	viper.BindEnv("passwd")
	viper.BindEnv("debug")

	viper.BindPFlag("f5", f5Cmd.PersistentFlags().Lookup("f5"))
	viper.BindPFlag("debug", f5Cmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("input", f5Cmd.PersistentFlags().Lookup("input"))
	viper.BindPFlag("pool", onlinePoolMemberCmd.Flags().Lookup("pool"))
	viper.BindPFlag("pool", offlinePoolMemberCmd.Flags().Lookup("pool"))

	// ignore errors - may be using environment vars or cmdline args
	viper.ReadInConfig()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: no config file found: %s\n", err)
		os.Exit(1)
	}

}

func checkFlags(cmd *cobra.Command) {

	debug = viper.GetBool("debug")
	now = viper.GetBool("now")
	username = viper.GetString("username")
	passwd = viper.GetString("passwd")
	f5Host = viper.GetString("device")
	stats_path_prefix = viper.GetString("stats_path_prefix")
	stats_show_zero_values = viper.GetBool("stats_show_zero_values")

	if username == "" {
		fmt.Fprint(os.Stderr, "\nerror: missing username; use config file or F5_USERNAME environment variable\n\n")
		os.Exit(1)
	}
	if passwd == "" {
		fmt.Fprint(os.Stderr, "\nerror: missing password; use config file or F5_PASSWD environment variable\n\n")
		os.Exit(1)
	}
	if f5Host == "" {
		fmt.Fprint(os.Stderr, "\nerror: missing f5 device hostname; use config file or F5_DEVICE environment variable\n\n")
		os.Exit(1)
	}

	// this has to be done here inside cobraCommand.Execute() inc case cmd line args are passed.
	// args are only parsed after cobraCommand.Run() - urgh
	appliance = f5.New(f5Host, username, passwd, f5.TOKEN)
	appliance.SetDebug(debug)
	appliance.SetStatsPathPrefix(stats_path_prefix)
	appliance.SetStatsShowZeroes(stats_show_zero_values)

}

func checkRequiredFlag(flg string) {
	if !viper.IsSet(flg) {
		fmt.Fprintf(os.Stdout, "\nerror: missing required option --%s\n\n", flg)
		os.Exit(1)
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

	f5Cmd.AddCommand(statsCmd)
	statsCmd.AddCommand(statsPoolCmd)
	statsCmd.AddCommand(statsVirtualCmd)
	statsCmd.AddCommand(statsNodeCmd)
	statsCmd.AddCommand(statsRuleCmd)

	// read config
	initialiseConfig()

}

func main() {
	//	f5Cmd.DebugFlags()
	f5Cmd.Execute()
}

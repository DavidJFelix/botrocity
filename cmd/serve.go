package cmd

import (
	"github.com/spf13/cobra"
	"github.com/hostables/botrocity/conf"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a mattermost connected bot server",
	Run:   serve,
}

var server string
var team string

func init() {
	// Register subcommand with the root command
	RootCmd.AddCommand(ServeCmd)

	// Register subcommand flags
	ServeCmd.Flags().StringVarP(&server, "server", "s", "", "The Mattermost server to connect to")
	ServeCmd.Flags().StringVarP(&team, "team", "t", "", "The Mattermost team to listen to")

	// Bind the command line flags to config
	conf.Config.BindPFlag("server", ServeCmd.Flags().Lookup("server"))
	conf.Config.BindPFlag("team", ServeCmd.Flags().Lookup("team"))
}

func serve(cmd *cobra.Command, args []string) {
	// TODO: do something here
}

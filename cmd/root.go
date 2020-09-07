package cmd

import "github.com/spf13/cobra"

var (
	logFile string
	log     bool
	rootCmd = &cobra.Command{
		Use:   "virtcurse",
		Short: "manage your hypervisor from a curses interface",
		Long: `manage your hypervisor from a curses interface using Libvirt APIs
and supported drivers: https://libvirt.org/drivers.html`,
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
	)
	rootCmd.PersistentFlags().StringVar(&logFile, "log-file", "./virtcurse.log", "path to store the logs")
	rootCmd.PersistentFlags().BoolVar(&log, "log", false, "enable logging in the `log-file`")
}

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tormath1/virtcurse/components"
)

var (
	logFile, libvirtURI string
	log                 bool
	rootCmd             = &cobra.Command{
		Use:   "virtcurse",
		Short: "manage your hypervisor from a curses interface",
		Long: `manage your hypervisor from a curses interface using Libvirt APIs
and supported drivers: https://libvirt.org/drivers.html`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return components.Execute(libvirtURI)
		},
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
	rootCmd.PersistentFlags().StringVar(&libvirtURI, "libvirt-uri", "qemu:///system", "virsh connect URI")
}

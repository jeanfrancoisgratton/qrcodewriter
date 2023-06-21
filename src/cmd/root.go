// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"qrcodewriter/qrwriter"
)

var version = "0.500-1 (2023.06.17)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "qrcw",
	Short:   "QR Code writer",
	Version: version,
	Long:    `Creates a QR code in a browser from a user-supplied URL.`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		qrwriter.Changelog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clCmd)
	//rootCmd.PersistentFlags().StringVarP(&certs.CertConfigFile, "qrwriter", "e", "defaultCertConfig.json", "environment configuration file.")
}

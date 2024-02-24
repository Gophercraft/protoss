package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root_cmd = &cobra.Command{
	Use:   "protoss file.exe",
	Short: "A BGS / BNet RPC / Protobuf decompiler",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("protoss: executable argument required")
		}
		run_dumper(cmd, args[0])
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := root_cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//	type DumpOptions struct {
//		The directory to output decompiled definitions to
//		Output string
//		The Go package path to prepend to Protocol Buffers packages
//		when generating "go_package" attribute
//		GoPrefix string
//		If true, output additional .json files to assist in ensuring accurate .proto decompilation
//		JSON bool
//		If true, don't output descriptors that are in the google/ path
//		IgnoreGoogle bool
//	}
func init() {
	root_cmd.Flags().StringP("output", "o", "proto", "The directory to output decompiled definitions to")
	root_cmd.Flags().StringP("go_prefix", "g", "github.com/Gophercraft/protoss/extensions", "The Go package path to prepend to Protocol Buffers packages when generating \"go_package\" attribute")
	root_cmd.Flags().BoolP("json_debug", "j", false, "Output additional json debug files to help when improving the accuracy of Protoss")
	root_cmd.Flags().BoolP("ignore_google", "i", true, "ignore descriptor files that are in the google/ path")
}

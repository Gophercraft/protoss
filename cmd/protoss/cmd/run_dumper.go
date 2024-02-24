package cmd

import (
	"fmt"
	"os"

	"github.com/Gophercraft/protoss"
	"github.com/spf13/cobra"
)

func run_dumper(root_cmd *cobra.Command, program string) {
	var opts protoss.DumpOptions

	output, err := root_cmd.Flags().GetString("output")
	if err != nil {
		panic(err)
	}
	go_prefix, err := root_cmd.Flags().GetString("go_prefix")
	if err != nil {
		panic(err)
	}
	debug_json, err := root_cmd.Flags().GetBool("json_debug")
	if err != nil {
		panic(err)
	}
	ignore_google, err := root_cmd.Flags().GetBool("ignore_google")
	if err != nil {
		panic(err)
	}

	opts.JSON = debug_json
	opts.Output = output
	opts.GoPrefix = go_prefix
	opts.IgnoreGoogle = ignore_google

	if err := protoss.DumpBinary(program, &opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

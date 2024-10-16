package main

import (
	"os"

	"github.com/openshift/crd-schema-checker/pkg/cmd/checkmanifests"
	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"
)

func main() {
	command := newCommand()
	code := cli.Run(command)
	os.Exit(code)
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "crd-schema-checker",
		Short: "Check CRD schemas for compatibility",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	//if v := version.Get().String(); len(v) == 0 {
	//	cmd.Version = "<unknown>"
	//} else {
	//	cmd.Version = v
	//}

	cmd.AddCommand(checkmanifests.NewCheckManifestsCommand())

	return cmd
}

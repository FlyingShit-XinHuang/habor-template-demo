package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/client"
)

var (
	getCmd = cobra.Command{
		Use:          "get [flags] <template id>",
		Short:        "Get specified template",
		Long:         "Get specified template",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preRunE(); nil != err {
				return err
			}
			if len(args) < 1 {
				return fmt.Errorf("Missing args")
			}
			getOpts.id = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return getTmpl()
		},
	}

	getOpts struct {
		id string
	}
)

func init() {
	RootCmd.AddCommand(&getCmd)
}

func getTmpl() error {
	if "" == getOpts.id {
		return fmt.Errorf("id should be specified")
	}
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	tmpl, err := c.Get(getOpts.id)
	if nil != err {
		return err
	}
	fmt.Println(tmpl)
	return nil
}

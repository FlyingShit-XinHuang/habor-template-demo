package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/client"
)

var (
	delCmd = cobra.Command{
		Use:          "delete [flags] <template id>",
		Short:        "Delete specified template",
		Long:         "Delete specified template",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preRunE(); nil != err {
				return err
			}
			if len(args) < 1 {
				return fmt.Errorf("Missing args")
			}
			delOpts.id = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return delTmpl()
		},
	}

	delOpts struct {
		id string
	}
)

func init() {
	RootCmd.AddCommand(&delCmd)
}

func delTmpl() error {
	if "" == delOpts.id {
		return fmt.Errorf("id should be specified")
	}
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	err := c.Delete(delOpts.id)
	if nil != err {
		return err
	}
	fmt.Println("Delete successfully")
	return nil
}

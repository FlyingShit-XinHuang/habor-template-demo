package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/client"
)

var (
	updateCmd = cobra.Command{
		Use:          "update [flags] <template id>",
		Short:        "update template from specified file",
		Long:         "update template from specified file",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preRunE(); nil != err {
				return err
			}
			if len(args) < 1 {
				return fmt.Errorf("Missing args")
			}
			updateOpts.id = args[0]
			if "" == updateOpts.file {
				return fmt.Errorf("file should be specified")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return updateTmpl()
		},
	}

	updateOpts struct {
		id   string
		file string
	}
)

func init() {
	RootCmd.AddCommand(&updateCmd)

	updateCmd.Flags().StringVar(&updateOpts.file, "file", "", "Path of template resource file")
}

func updateTmpl() error {
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	err := c.UpdateFromFile(updateOpts.id, updateOpts.file)
	if nil != err {
		return err
	}
	fmt.Println("Update successfully")
	return nil
}

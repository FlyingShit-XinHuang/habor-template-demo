package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/client"
)

var (
	createCmd = cobra.Command{
		Use:          "create",
		Short:        "Create template from specified file",
		Long:         "Create template from specified file",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preRunE(); nil != err {
				return err
			}
			if "" == createOpts.file {
				return fmt.Errorf("file should be specified")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return createTmpl()
		},
	}

	createOpts struct {
		file string
	}
)

func init() {
	RootCmd.AddCommand(&createCmd)

	createCmd.Flags().StringVar(&createOpts.file, "file", "", "Path of template resource file")
}

func createTmpl() error {
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	err := c.CreateFromFile(createOpts.file)
	if nil != err {
		return err
	}
	fmt.Println("Create successfully")
	return nil
}

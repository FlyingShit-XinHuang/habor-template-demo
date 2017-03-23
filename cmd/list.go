package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/api/v1"
	"habor-template-demo/client"
)

var (
	listCmd = cobra.Command{
		Use:          "list",
		Short:        "List all templates",
		Long:         "List all templates",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRunE()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if "workspaces" == listOpts.resource {
				return listWorkspaces()
			}
			return listTmpls()
		},
	}

	listOpts struct {
		offset   int
		limit    int
		resource string
	}
)

func init() {
	RootCmd.AddCommand(&listCmd)

	listCmd.Flags().IntVar(&listOpts.offset, "offset", 0, "Offset of pagination of list operation")
	listCmd.Flags().IntVar(&listOpts.limit, "limit", 20, "Limit of pagination of list operation")
	listCmd.Flags().StringVar(&listOpts.resource, "resource", "templates",
		"Type of resources, only templates and workspaces supported")
}

func listTmpls() error {
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	list, err := c.List(&v1.PaginationOptions{Offset: listOpts.offset, Limit: listOpts.limit})
	if nil != err {
		return err
	}
	fmt.Println(list)
	return nil
}

func listWorkspaces() error {
	c := client.NewWorkspaceClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password)
	list, err := c.List(nil)
	if nil != err {
		return err
	}
	fmt.Println(list)
	return nil
}

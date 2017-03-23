package cmd

import (
	"github.com/spf13/cobra"

	"fmt"

	"habor-template-demo/client"
)

var (
	sendCmd = cobra.Command{
		Use:          "send [flags] <to> <template id>",
		Short:        "Send message(s) with specified template",
		Long:         "Send message(s) with specified template",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := preRunE(); nil != err {
				return err
			}
			if len(args) < 2 {
				return fmt.Errorf("Missing args")
			}
			sendOpts.to = args[0]
			sendOpts.id = args[1]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return sendWithTmpl()
		},
	}

	sendOpts struct {
		id string
		to string
	}
)

func init() {
	RootCmd.AddCommand(&sendCmd)
}

func sendWithTmpl() error {
	if "" == sendOpts.id {
		return fmt.Errorf("id should be specified")
	}
	c := client.NewTemplateClient(RootOpts.ApiKey, RootOpts.User, RootOpts.Password, RootOpts.Workspace)
	location, err := c.SendMsg(sendOpts.id, sendOpts.to)
	if nil != err {
		return err
	}
	fmt.Printf("Message info: %s\n", location)
	return nil
}

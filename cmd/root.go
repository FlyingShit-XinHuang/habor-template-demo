package cmd

import (
	//"flag"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command{
		Use: "demo",
	}
	RootOpts struct {
		ApiKey    string
		User      string
		Password  string
		Workspace string
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&RootOpts.ApiKey,
		"apikey", "k", "", "The API key to query Whispir API")
	RootCmd.PersistentFlags().StringVarP(&RootOpts.User,
		"user", "u", "", "The user name of Whispir")
	RootCmd.PersistentFlags().StringVarP(&RootOpts.Password,
		"password", "p", "", "The password of Whispir")
	RootCmd.PersistentFlags().StringVarP(&RootOpts.Workspace,
		"workspace", "w", "", "The workspace")
}

func preRunE() error {
	if err := checkAuthParams(); nil != err {
		return err
	}
	return nil
}

func checkAuthParams() error {
	switch {
	case "" == RootOpts.ApiKey:
		return fmt.Errorf("apikey should be specified")
	case "" == RootOpts.User:
		return fmt.Errorf("user should be specified")
	case "" == RootOpts.Password:
		return fmt.Errorf("password should be specified")
	}
	//flag.Set("v", "8")
	//flag.Set("logtostderr", "true")
	return nil
}

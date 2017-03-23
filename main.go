package main

import (
	"habor-template-demo/cmd"
	"os"
)

func main() {
	//flag.Parse()
	//c := client.NewTemplateClient("bvtaqrzveu86gpz8khd66gm7", "haborhuang", "07301091Hx", workspace)
	//list, err := c.List(&v1.PaginationOptions{Offset:offset, Limit:limit})
	//if nil != err {
	//	glog.Fatalf("Failed to list templates: %v\n", err)
	//}
	//
	//fmt.Printf("Response: %v\n", list)

	if err := cmd.RootCmd.Execute(); nil != err {
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func buildflags() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	companyptr := addCmd.String("companyname", "google", "name of the company")
	joblinkptr := addCmd.String("joblink", "link", "link of the job role")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'list' commands ")
		os.Exit(1)
	}
	switch os.Args[1] {

	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("add command")
		fmt.Println("name:", *companyptr)
		fmt.Println("joblink:", *joblinkptr)
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("list command")

	default:
		fmt.Println("expected add or list comands")
		os.Exit(1)

	}
}

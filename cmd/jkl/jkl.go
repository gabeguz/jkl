package main

import (
	"flag"
	"fmt"
	"log"

	"otremblay.com/jkl"
)

func main() {
	jkl.FindRCFile()
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Print(usage)
		return
	}
	if err := runcmd(flag.Args()); err != nil {
		log.Println(err)
	}
}

func runcmd(args []string) error {
	switch args[0] {
	case "list":
		lcmd, err := NewListCmd(flag.Args()[1:])
		if err != nil {
			return err
		}
		return lcmd.List()
	case "create":
		ccmd, err := NewCreateCmd(flag.Args()[1:])
		if err != nil {
			return err
		}
		return ccmd.Create()
	case "task":
		tcmd := &TaskCmd{}
		return tcmd.Handle(flag.Args()[1:])
	case "edit":
		ecmd := &EditCmd{}
		return ecmd.Edit(flag.Arg(1))
	case "comment":
		ccmd, err := NewCommentCmd(flag.Args()[1:])
		if err != nil {
			return err
		}
		return ccmd.Comment()
	}
	fmt.Println(usage)
	return nil
}

const usage = `Usage:
jkl [options] <command> [args]

Available commands:

list
create
edit
`

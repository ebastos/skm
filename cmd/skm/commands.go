package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

func initCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize SSH keys store for the first time use",
			Action:  initialize,
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a new SSH key.",
			Action:  create,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "b", Usage: "bits"},
				cli.StringFlag{Name: "C", Usage: "comment"},
			},
		},
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "List all the available SSH keys",
			Action:  list,
		},
		{
			Name:    "use",
			Aliases: []string{"u"},
			Usage:   "Set specific SSH key as default by its alias name",
			Action:  use,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete specific SSH key by alias name",
			Action:  delete,
		},
		{
			Name:    "rename",
			Aliases: []string{"rn"},
			Usage:   "Rename SSH key alias name to a new one",
			Action:  rename,
		},
		{
			Name:    "copy",
			Aliases: []string{"cp"},
			Usage:   "Copy current SSH public key to a remote host",
			Action:  copy,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "p", Usage: "SSH port"},
			},
		},
		{
			Name:    "display",
			Aliases: []string{"dp"},
			Usage:   "Display the current SSH public key",
			Action:  display,
		},
		{
			Name:    "backup",
			Aliases: []string{"b"},
			Usage:   "Backup all SSH keys to an archive file",
			Action:  backup,
		},
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "Restore SSH keys from an existing archive file",
			Action:  restore,
		},
	}
}

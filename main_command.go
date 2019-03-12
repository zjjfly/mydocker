package main

import (
	"com.siemens/zjj/mydocker/container"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `create a container with namespace and cgroup limit
		   mydocker run -it [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			log.Error("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: `Init container process run user's process in container.Do not call it outside`,
	Action: func(context *cli.Context) error {
		log.Infof("Init come on")
		cmd := context.Args().Get(0)
		log.Infof("Init Command %s", cmd)
		return container.RunContainerInitProcess(cmd, nil)
	},
}

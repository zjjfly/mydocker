package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zjjfly/mydocker/cgroups/subsystems"
	"github.com/zjjfly/mydocker/container"
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
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpu limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			log.Error("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		tty := context.Bool("it")
		memoryLimit := context.String("m")
		cpuLimit := context.String("cpushare")
		cpuSetLimit := context.String("cpuset")
		Run(tty, cmdArray, &subsystems.ResourceConfig{
			MemoryLimit: memoryLimit,
			CpuSet:      cpuSetLimit,
			CpuShare:    cpuLimit,
		})
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: `Init container process run user's process in container.Do not call it outside`,
	Action: func(context *cli.Context) error {
		log.Infof("Init come on")
		return container.RunContainerInitProcess()
	},
}

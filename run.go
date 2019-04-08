package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/zjjfly/mydocker/cgroups"
	"github.com/zjjfly/mydocker/cgroups/subsystems"
	"github.com/zjjfly/mydocker/container"
	"os"
	"strings"
)

func Run(tty bool, cmdArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup", res)
	defer cgroupManager.Destroy()
	cgroupManager.Set()
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(cmdArray, writePipe)
	parent.Wait()
}

func sendInitCommand(cmdArray []string, writePipe *os.File) {
	command := strings.Join(cmdArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/zjjfly/mydocker/cgroups"
	"github.com/zjjfly/mydocker/cgroups/subsystems"
	"github.com/zjjfly/mydocker/container"
)

func Run(tty bool, command string, res *subsystems.ResourceConfig) {
	parent := container.NewParentProcess(tty, command)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	parent.Wait()
}

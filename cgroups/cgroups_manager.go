package cgroups

import (
	log "github.com/sirupsen/logrus"
	"github.com/zjjfly/mydocker/cgroups/subsystems"
)

type CgroupManager struct {
	Path           string
	ResourceConfig *subsystems.ResourceConfig
}

func NewCgroupManager(path string, res *subsystems.ResourceConfig) *CgroupManager {
	return &CgroupManager{
		Path:           path,
		ResourceConfig: res,
	}
}

func (c *CgroupManager) Set() error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Set(c.Path, c.ResourceConfig); err != nil {
			log.Error(err.Error())
		}
	}
	return nil
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Apply(c.Path, pid); err != nil {
			log.Error(err.Error())
		}
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}

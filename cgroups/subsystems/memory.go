package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (s *MemorySubsystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set %s cgroup proc fail %v", s.Name(), err)
		}
		return nil
	} else {
		return fmt.Errorf("get %s cgroup %s error: %v", s.Name(), cgroupPath, err)
	}
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		return os.Remove(subsysCgroupPath)
	} else {
		return err
	}
}

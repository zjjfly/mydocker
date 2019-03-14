package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
)

type CpusetSubSystem struct {
}

func (s *CpusetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.CpuSet != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(res.CpuSet), 0644); err != nil {
				return fmt.Errorf("set cgroup cpuset.cpus fail %v", err)
			}
		} else {
			numCPU := runtime.NumCPU()
			cpuset := fmt.Sprintf("0-%d", numCPU-1)
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(cpuset), 0644); err != nil {
				return fmt.Errorf("set cgroup cpuset.cpus fail %v", err)
			}
		}
		//cpuset cgroup在把一个进程加入tasks之前必须在cpuset.cpus和cpuset.mems中都有值，否则会报no space left
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpuset.mems"), []byte(strconv.Itoa(0)), 0644); err != nil {
			return fmt.Errorf("set cgroup cpuset.mems fail %v", err)
		}
		return nil
	} else {
		return err
	}
}

func (s *CpusetSubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}

func (s *CpusetSubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set %s cgroup proc fail %v", s.Name(), err)
		}
		return nil
	} else {
		return fmt.Errorf("get %s cgroup %s error: %v", s.Name(), cgroupPath, err)
	}
}

func (s *CpusetSubSystem) Name() string {
	return "cpuset"
}

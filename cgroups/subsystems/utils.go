package subsystems

import (
	"bufio"
	"fmt"
	"github.com/zjjfly/mydocker/util"
	"os"
	"path"
	"strings"
)

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint(subsystem)
	p := path.Join(cgroupRoot, cgroupPath)
	if !util.IsExist(p) {
		if autoCreate {
			if err := os.Mkdir(p, 0755); err != nil {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		} else {
			return "", fmt.Errorf("cgroup path not exist")
		}
	}
	return p, nil
}

package container

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"strings"
	"syscall"
)

func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command %s", command)
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	cmds := strings.Split(command, " ")
	err := syscall.Exec(cmds[0], cmds[0:], os.Environ())
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

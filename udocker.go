package main

import (
	"os"
	"os/exec"
	"strconv"
)

func main() {
	uid := strconv.Itoa(os.Getuid())
	gid := strconv.Itoa(os.Getgid())
	uid_gid := uid + ":" + gid

	cmdargs := append([]string{"run", "--cap-drop", "ALL", "-u", uid_gid, "--rm", "-i", "--"}, os.Args[1:]...)
	cmd := exec.Command("/usr/bin/docker", cmdargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

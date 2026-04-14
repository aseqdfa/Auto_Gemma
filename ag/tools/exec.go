package tools

import "os/exec"

func RunCmd(cmd string) string {
	out, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out)
}

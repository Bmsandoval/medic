package kube_svc

import (
	"errors"
	"github.com/bmsandoval/medic/bash"
	"github.com/bmsandoval/medic/utils"
	"strings"
)

func GetChoiceShell(pod string) (string, error) {
	err, out, errout := utils.ExecGetOutput(bash.KubeListShellsOfPod(pod))
	if err != nil {
		return "", err
	}
	if errout != "" && errout != "Unable to use a TTY - input is not a terminal or the right kind of file\n" {
		return "", errors.New("error getting pods, maybe kubernetes hasn't been started yet")
	}
	trimmedOutput := strings.TrimSpace(out)
	splitOutput := strings.Split(trimmedOutput,"\n")
	if len(splitOutput) < 1 {
		return "", errors.New("no pods found")
	}

	for _, v := range splitOutput {
		if v == "/bin/bash" {
			return "bash", nil
		}
	}
	return "sh", nil
}

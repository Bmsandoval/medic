package kube_svc

import (
	"medic/bash"
	"medic/utils"
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"strings"
)

func SelectPod() (*string, error) {
	err, out, errout := utils.ExecGetOutput(bash.KubeListPodsByName)
	if err != nil {
		return nil, err
	}
	if errout != "" {
		return nil, errors.New("error getting pods, maybe kubernetes hasn't been started yet")
	}
	trimmedOutput := strings.TrimSpace(out)
	splitOutput := strings.Split(trimmedOutput,"\n")
	if len(splitOutput) < 1 {
		return nil, errors.New("no pods found")
	}
	selected, err := fuzzyfinder.Find(splitOutput,
		func(i int) string {
			return splitOutput[i]
		})
	if err != nil {
		return nil, errors.New("no pod selected, aborting")
	}
	return &splitOutput[selected], nil
}

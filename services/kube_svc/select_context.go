package kube_svc

import (
	"medic/bash"
	"medic/utils"
	"errors"
	"github.com/ktr0731/go-fuzzyfinder"
	"strings"
)

func SelectContext() (*string, error) {
	err, out, errout := utils.ExecGetOutput(bash.KubeListCtxtsByName)
	if err != nil {
		return nil, err
	}
	if errout != "" {
		return nil, errors.New("error getting contexts, maybe kubernetes hasn't been started yet")
	}
	trimmedOutput := strings.TrimSpace(out)
	splitOutput := strings.Split(trimmedOutput,"\n")
	if len(splitOutput) < 1 {
		return nil, errors.New("no contexts found")
	}
	selected, err := fuzzyfinder.Find(splitOutput,
		func(i int) string {
			return splitOutput[i]
		})
	if err != nil {
		return nil, errors.New("no context selected, aborting")
	}
	return &splitOutput[selected], nil
}

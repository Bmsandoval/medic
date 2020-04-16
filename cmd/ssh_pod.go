/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bmsandoval/medic/bash"
	"github.com/bmsandoval/medic/services/kube_svc"
	"github.com/bmsandoval/medic/utils"
	"os/exec"
)

var kubeSshPodCmd = &cobra.Command{
	Use:   "ssh",
	Aliases: []string{"sh"},
	Short: "SSH into a pod",
	Long: `kubectl exec ${pod} -it /bin/sh`,
	Run: KubeSshPod,
}


func KubeSshPod(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("kubectl"); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Select a pod
	pod, err := kube_svc.SelectPod()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	shell, err := kube_svc.GetChoiceShell(*pod)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// SSH into pod
	cmd, args := bash.KubeSshPod(*pod, shell)
	if err := utils.ExecNotCapturingOutput(cmd, args); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(kubeSshPodCmd)
}

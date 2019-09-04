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
	"medic/bash"
	"medic/services/kube_svc"
	"medic/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var kubeSshPodCmd = &cobra.Command{
	Use:   "ssh",
	Aliases: []string{"sh"},
	Short: "SSH into a pod",
	Long: `kubectl exec ${pod} -it /bin/sh`,
	Run: KubeSshPod,
}


func KubeSshPod(_ *cobra.Command, _ []string) {
	// Select a pod
	pod, err := kube_svc.SelectPod()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// Delete a Pod
	sshCmd := bash.KubeSshPod(*pod)
	if err := utils.Exec(sshCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(kubeSshPodCmd)
}

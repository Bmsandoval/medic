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
	"log"
	"github.com/bmsandoval/medic/bash"
	"github.com/bmsandoval/medic/services/kube_svc"
	"github.com/bmsandoval/medic/utils"
	"os/exec"
)

var kubeDelPodCmd = &cobra.Command{
	Use:   "delete",
	Aliases: []string{"del"},
	Short: "Delete a pod",
	Long: `kubectl delete pod ${pod}`,
	Run: KubeDelPod,
}


func KubeDelPod(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("kubectl"); err != nil {
		log.Println(err.Error())
		return
	}

	// Select a pod
	pod, err := kube_svc.SelectPod()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Delete a Pod
	delCmd := bash.KubeDeletePod(*pod)


	if err := utils.Exec(delCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(kubeDelPodCmd)
}

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
	"medic/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var kubeStartCmd = &cobra.Command{
	Use:   "start",
	Aliases: []string{"stt"},
	Short: "Start minikube",
	Long: `. ~/projects/k8s-cluster-management/scripts/minikube.sh`,
	Run: KubeStart,
}


func KubeStart(_ *cobra.Command, _ []string) {
	// Delete a Pod
	startCmd := bash.KubeStart
	if err := utils.Exec(startCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(kubeStartCmd)
}

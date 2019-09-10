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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"medic/bash"
	"medic/config"
	"medic/utils"
)

var kubeStartCmd = &cobra.Command{
	Use:   "start",
	Aliases: []string{"stt"},
	Short: "Start minikube",
	Long: `. ~/projects/k8s-cluster-management/scripts/minikube.sh`,
	Run: KubeStart,
}


func KubeStart(_ *cobra.Command, _ []string) {
	configs := config.GetConfigFromViper()
	filePath := configs.MinikubeConfig.FilePath
	if filePath == "" {
		fmt.Println("Minikube start script's file path not found in config file\n aborting...")
		return
	}
	// Delete a Pod
	startCmd := bash.KubeStart(filePath)
	if err := utils.Exec(startCmd); err != nil {
		fmt.Println(err.Error())
	}
	success := color.New(color.FgGreen)
	if _, err := success.Println("SUCCESS"); err != nil {
		fmt.Println("FAIL")
	}
}

func init() {
	rootCmd.AddCommand(kubeStartCmd)
}

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
	"os/exec"
)

const (
	DefaultMapping = "8080:8080"
)

var forwardPortCmd = &cobra.Command{
	Use:   "port",
	Aliases: []string{"pt"},
	Short: "Forward local port to kubernetes pod",
	Long: `kubectl port-forward "${pod}" "${2}"`,
	Run: ForwardPort,
}


func ForwardPort(cmd *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("kubectl"); err != nil {
		log.Println(err.Error())
		return
	}

	// Get flag arguments
	ports, _:= cmd.Flags().GetString("port")
	if ports == "" {
		log.Println(fmt.Sprintf("No port provided, assuming %q", DefaultMapping))
		ports = "8080:8080"
	}
	// Select a pod
	pod, err := kube_svc.SelectPod()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// Forward the port
	forwardCmd := bash.KubePortForward(*pod, ports)
	if err := utils.Exec(forwardCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(forwardPortCmd)
	// add command line flags
	forwardPortCmd.Flags().StringP("port", "p", "", "local:remote")
}

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
	"github.com/bmsandoval/medic/bash"
	"github.com/bmsandoval/medic/services/kube_svc"
	"github.com/bmsandoval/medic/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var kubeSetContextCmd = &cobra.Command{
	Use:   "context",
	Aliases: []string{"ctx"},
	Short: "Set context",
	Long: `kubectl exec -it "${pod}" -- contextall main`,
	Run: KubeSetContext,
}

func KubeSetContext(_ *cobra.Command, _ []string) {
	// Check for necessary stuff
	if _, err := exec.LookPath("kubectl"); err != nil {
		log.Println(err.Error())
		return
	}

	// Select a context
	ctx, err := kube_svc.SelectContext()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// Kill a Pod
	contextCmd := bash.KubeSetContext(*ctx)
	if err := utils.Exec(contextCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(kubeSetContextCmd)
}

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
	"github.com/spf13/cobra"
	"log"
)

var helmPurgeCommand = &cobra.Command{
	Use:   "purge",
	Aliases: []string{"pg"},
	Short: "Purge a given release",
	Long: `helm delete --purge ${pod}`,
	Run: HelmPurge,
}

func HelmPurge(_ *cobra.Command, _ []string) {
	// Select a pod
	release, err := kube_svc.SelectRelease()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// Get command
	purgeCmd := bash.HelmDeletePurge(*release)
	// Run command
	err = utils.Exec(purgeCmd)
	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(helmPurgeCommand)
}

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
	"log"
	"strings"
)

var boatswainReleaseCmd = &cobra.Command{
	Use:   "release",
	Aliases: []string{"rls"},
	Short: "Release a given repo",
	Long: `boatswain release ${repo} ${cmdOptions} --assume-yes`,
	Run: BoatswainRelease,
}


func BoatswainRelease(cmd *cobra.Command, _ []string) {
	// Get flag arguments
	repo, _:= cmd.Flags().GetString("repo")
	if repo == "" {
		log.Println("No repo provided, please include a repo")
	}
	// Get Context
	err, ctx, errout := utils.ExecGetOutput(bash.KubeGetCurrentContext)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if errout != "" {
		log.Println("error getting context. maybe kubernetes hasn't started yet")
		return
	}
	// Set release options based on context
	var cmdOptions string
	if ctx == "minikube" {
		cmdOptions = ""
	} else if strings.Contains(ctx, "staging") {
		cmdOptions = "-e staging"
	}
	// Perform Release
	releaseCmd := bash.BoatswainRelease(repo, cmdOptions)
	if err := utils.Exec(releaseCmd); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(boatswainReleaseCmd)
	// add command line flags
	boatswainReleaseCmd.Flags().StringP("repo", "r", "", "odin")
}

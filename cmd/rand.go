/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// randCmd represents the rand command
var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "Generates a random time from future",
	Long: `Generates a random time from future that has range from now till 10 days.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		window := 10 * 24 * time.Hour
		s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
    	randomSeconds := r1.Intn(int(window.Seconds()))
    	randomTime := now.Add(time.Duration(randomSeconds) * time.Second)
    	zone, _ := now.Local().Zone()
		fmt.Println("UTC :", randomTime.UTC().Format(time.RFC3339))
		fmt.Println(zone, ":", randomTime.Format(time.RFC3339))
    	fmt.Println("Unix(seconds) :", randomTime.Unix())
	},
}

func init() {
	rootCmd.AddCommand(randCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

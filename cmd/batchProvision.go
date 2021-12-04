/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/fosshostorg/teardrop/provisioning"
	"github.com/spf13/cobra"
)

// batchProvisionCmd represents the batchProvision command
var batchProvisionCmd = &cobra.Command{
	Use:   "batch-provision",
	Short: "USE ONLY FOR TESTING - batch provision 3 master vms in dfw ",

	Run: func(cmd *cobra.Command, args []string) {
		println("Batch Provisioning VMs")
		client := provisioning.NewA64Client(os.Getenv("AARCH64_APIKEY"))
		client.BatchProvisionVMs("dfw", true, 3)
	},
}

func init() {
	rootCmd.AddCommand(batchProvisionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batchProvisionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batchProvisionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

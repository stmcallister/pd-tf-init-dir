/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/cobra"
	"github.com/zclconf/go-cty/cty"
)

// mkdirCmd represents the mkdir command
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Create directory with TF files inside",
	Long: `Create a directory with a main.tf and terraform.tf file inside. 
	The provider.tf file will be populated with PagerDuty as the required provider.`,
	Run: func(cmd *cobra.Command, args []string) {
		var dirName string

		if len(args) >= 1 && args[0] != "" {
			dirName = args[0]
		}

		if dirName == "" {
			fmt.Println("mkdir requires a directory name to be passed")
		} else {
			// create directory in the current directory
			if err := os.Mkdir(dirName, os.ModePerm); err != nil {
				fmt.Println(err.Error())
			}
			// create main.tf
			mainFile, err := os.Create(fmt.Sprintf("%s/main.tf", dirName))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer mainFile.Close()

			// create terraform.tf
			// create new empty hcl file object
			f := hclwrite.NewEmptyFile()

			// create new file on system
			tfFile, err := os.Create(fmt.Sprintf("%s/terraform.tf", dirName))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer tfFile.Close()

			// initialize the body of the new file object
			rootBody := f.Body()

			// initialize terraform object and set provider version
			tfBlock := rootBody.AppendNewBlock("terraform", nil)
			tfBlockBody := tfBlock.Body()
			reqProvsBlock := tfBlockBody.AppendNewBlock("required_providers",
				nil)
			reqProvsBlockBody := reqProvsBlock.Body()

			reqProvsBlockBody.SetAttributeValue("pagerduty", cty.ObjectVal(map[string]cty.Value{
				"source": cty.StringVal("PagerDuty/pagerduty"),
			}))
			rootBody.AppendNewline()
			tfFile.Write(f.Bytes())

			// change to new directory
			err = os.Chdir(dirName)
			if err != nil {
				fmt.Println(err.Error())
			}
			// open vs code
			cmd := exec.Command("code", ".")
			err = cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mkdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

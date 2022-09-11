package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/piavgh/alina/internal/signature"
)

var generateSignatureCmd = &cobra.Command{
	Use:   "generate-signature",
	Short: "Generate the bytecode of smart contract signature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Usage: alina generate-signature <signature>")

			return
		}

		fmt.Println("Input: ", args[0])

		fmt.Println("Output: ", signature.Generate(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(generateSignatureCmd)
}

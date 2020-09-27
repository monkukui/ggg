package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ggg",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("hello ggg")
		// フラグ名で値を取得する
		indexed, err := c.PersistentFlags().GetBool("indexed")
		if err != nil {
			fmt.Println(err)
			return
		}
		directed, err := c.PersistentFlags().GetBool("directed")
		if err != nil {
			fmt.Println(err)
			return
		}
		weighted, err := c.PersistentFlags().GetBool("weighted")
		if err != nil {
			fmt.Println(err)
			return
		}
		format, err := c.PersistentFlags().GetBool("format")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("indexed: ", indexed)
		fmt.Println("directed: ", directed)
		fmt.Println("weighted: ", weighted)
		fmt.Println("format: ", format)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	rootCmd.PersistentFlags().BoolP("indexed", "i", true, "graph is 1-indexed")
	rootCmd.PersistentFlags().BoolP("directed", "d", false, "graph is directed")
	rootCmd.PersistentFlags().BoolP("weighted", "w", false, "graph is weighted")
	rootCmd.PersistentFlags().BoolP("format", "f", true, "graph format is normal")
}

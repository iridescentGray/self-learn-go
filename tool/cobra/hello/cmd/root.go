package cmd

import (
	"fmt"
	"hello/cmd/sub1"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

func init() {
	// 全局参数 所有commond都能用
	// go run main.go verbose  -v true      go run main.go  print -v true
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose explain")
	rootCmd.AddCommand(tryCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(sub1.Sub1Cmd)

	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source explain")
}

func someFunc() error {
	fmt.Println("try error")
	return nil
}

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(Verbose)
		if err := someFunc(); err != nil {
			return err
		}
		return nil
	},
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Verbose)
	},
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd run")
		fmt.Println(Verbose)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

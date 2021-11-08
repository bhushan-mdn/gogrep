package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-grep PATTERN FILE",
	Short: "go-grep is an implementation of grep in go",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Not enough arguments")
			os.Exit(3)
		}
		pattern := args[0]
		file := args[1]
		grep(pattern, file)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func grep(pattern string, file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		if strings.Contains(text, pattern) {
			red := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(strings.ReplaceAll(text, pattern, red(pattern)))
		}
	}
}

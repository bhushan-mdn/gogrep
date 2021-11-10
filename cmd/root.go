package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-grep PATTERN FILE",
	Short: "go-grep is an implementation of grep in go",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("Missing required arguments.")
		os.Exit(1)
	}
	pattern := args[0]
	file := args[1]
	Search(pattern, file)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func Search(pattern string, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("File not found. Error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Invalid pattern expression. Error: ", err)
		os.Exit(1)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		if match := r.FindString(text); match != "" {
			red := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(r.ReplaceAllString(text, red(match)))
		}
	}
}

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
	Use:   "gogrep PATTERN [FILE]...",
	Short: "gogrep is an implementation of grep in go",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("Missing required arguments.")
		os.Exit(1)
	} else if len(args) < 2 {
		pattern := args[0]
		r := GetRegexp(pattern)
		GrepStdin(r)
	} else {
		pattern := args[0]
		files := args[1:]
		r := GetRegexp(pattern)
		GrepFile(r, files)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func GetRegexp(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Invalid pattern expression. Error: ", err)
		os.Exit(1)
	}
	return r
}

func GrepFile(r *regexp.Regexp, files []string) {
	if len(files) == 1 {
		filename := files[0]
		text := ReadFile(filename)
		PrintMatches(text, r)
	} else {
		for _, filename := range files {
			text := ReadFile(filename)
			PrintMatchesWithFilename(text, r, filename)
		}
	}
}

func ReadFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("File not found. Error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	text := Textify(s)
	return text
}

func GrepStdin(r *regexp.Regexp) {
	s := bufio.NewScanner(os.Stdin)
	text := Textify(s)
	PrintMatches(text, r)
}

func Textify(s *bufio.Scanner) []string {
	s.Split(bufio.ScanLines)
	var text []string

	for s.Scan() {
		text = append(text, s.Text())
	}
	return text
}

func PrintMatches(text []string, r *regexp.Regexp) {
	for _, line := range text {
		if match := r.FindString(line); match != "" {
			red := color.New(color.FgRed, color.Bold).SprintFunc()
			fmt.Println(r.ReplaceAllString(line, red(match)))
		}
	}
}

func PrintMatchesWithFilename(text []string, r *regexp.Regexp, filename string) {
	for _, line := range text {
		if match := r.FindString(line); match != "" {
			red := color.New(color.FgRed, color.Bold).SprintFunc()
			magenta := color.New(color.FgMagenta).SprintFunc()
			fmt.Printf("%s:%s\n", magenta(filename), r.ReplaceAllString(line, red(match)))
		}
	}
}

package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Alex-Waring/AoC/utils"
	"github.com/spf13/cobra"
)

var newInput = &cobra.Command{
	Use:     "input",
	Aliases: []string{"i"},
	Short:   "Download the input for a given year and day",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		GetInput(utils.IntegerOf(args[0]), utils.IntegerOf(args[1]))
	},
}

var baseFile = &cobra.Command{
	Use:     "file",
	Aliases: []string{"f"},
	Short:   "Prepare a base file",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		PrepareFile(utils.IntegerOf(args[0]), utils.IntegerOf(args[1]))
	},
}

func init() {
	rootCmd.AddCommand(newInput)
	rootCmd.AddCommand(baseFile)
}

// Requires the session cookie to be in the repo
func GetInput(year int, day int) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	utils.Check(err)

	cookie, err := os.ReadFile("session.cookie")
	utils.Check(err)

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: strings.TrimRight(string(cookie), "\n"),
	}
	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	utils.Check(err)

	body, err := io.ReadAll(res.Body)
	utils.Check(err)

	// specific error message from AOC site
	if strings.HasPrefix(string(body), "Please don't repeatedly") {
		log.Fatalf("Repeated request error")
	}

	// Write to file
	filename := filepath.Join(fmt.Sprintf("%d/Day%02d/input.txt", year, day))

	err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	utils.Check(err)
	err = os.WriteFile(filename, body, os.FileMode(0644))
	utils.Check(err)
}

// Prepares a basic golang template file
func PrepareFile(year int, day int) {
	filename := filepath.Join(fmt.Sprintf("%d/Day%02d/main.go", year, day))

	if _, err := os.Stat(filename); err == nil {
		return
	}

	base_file := `package main

import "github.com/Alex-Waring/AoC/utils"
	
func part1() {
	defer utils.Timer("part1")()
}
	
func part2() {
	defer utils.Timer("part2")()
}

func main() {
	input := utils.ReadInput("input.txt")
}`

	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	utils.Check(err)
	file, err := os.Create(filename)
	file.WriteString(base_file)
	utils.Check(err)
}

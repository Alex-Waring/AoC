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

var rustBaseFile = &cobra.Command{
	Use:     "rustfile",
	Aliases: []string{"rf"},
	Short:   "Prepare a rust base file",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		PrepareRustFile(utils.IntegerOf(args[0]), utils.IntegerOf(args[1]))
	},
}

func init() {
	rootCmd.AddCommand(newInput)
	rootCmd.AddCommand(baseFile)
	rootCmd.AddCommand(rustBaseFile)
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
	
func part1(input []string) {
	defer utils.Timer("part1")()
}
	
func part2(input []string) {
	defer utils.Timer("part2")()
}

func main() {
	input := utils.ReadInput("input.txt")
	part1(input)
	part2(input)
}`

	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	utils.Check(err)
	file, err := os.Create(filename)
	file.WriteString(base_file)
	utils.Check(err)

	_, err = os.Create(filepath.Join(fmt.Sprintf("%d/Day%02d/example.txt", year, day)))
	utils.Check(err)
}

func PrepareRustFile(year int, day int) {
	filename := filepath.Join(fmt.Sprintf("%d/day%02d.rs", year, day))

	if _, err := os.Stat(filename); err == nil {
		return
	}

	base_file := `use std::fs;

pub fn main() {
    println!("Part1: {}", part1());
    println!("Part2: {}", part2());
}

pub fn part1() -> String {
    let input = fs::read_to_string("./%d/day%02d/input.txt").expect("Error reading input.txt");

    return format!("");
}

pub fn part2() -> String {
    let input = fs::read_to_string("./%d/day%02d/input.txt").expect("Error reading input.txt");

    return format!("");
}
`

	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	utils.Check(err)
	file, err := os.Create(filename)
	file.WriteString(fmt.Sprintf(base_file, year, day, year, day))
	utils.Check(err)

	_, err = os.Create(filepath.Join(fmt.Sprintf("%d/Day%02d/example.txt", year, day)))
	utils.Check(err)
}

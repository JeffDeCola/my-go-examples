package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func makeTOC(heading string, headingNumber string, inputFilename string) {
	//fmt.Println("Working on heading", heading, line)

	// STEP 1 ***************************
	// FIX HEADING
	// Replace withspace with -
	headingLower := strings.Replace(heading, " ", "-", -1)
	// Remove all special characters except -
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		log.Fatal(err)

	}
	headingLower = reg.ReplaceAllString(headingLower, "")
	// Make all lowercase
	headingLower = strings.ToLower(headingLower)

	// STEP 2 *****************************
	// Get components to build link
	// This will be based on my directory structure
	githubURL := "https://github.com/JeffDeCola/"
	// What is repo name and path after that (if any)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Get everything after "jeff/cheatsheets/"
	if strings.Contains(dir, "jeff/cheatsheets/") {
		parts := strings.Split(dir, "jeff/cheatsheets/")
		dir = parts[1]
	}
	// Get everything after "jeff/other/"
	if strings.Contains(dir, "jeff/other/") {
		parts := strings.Split(dir, "jeff/other/")
		dir = parts[1]
	}
	// Get everything after "jeff/python/"
	if strings.Contains(dir, "jeff/python/") {
		parts := strings.Split(dir, "jeff/python/")
		dir = parts[1]
	}
	// Get everything after "github.com/JeffDeCola/"
	if strings.Contains(dir, "github.com/JeffDeCola/") {
		parts := strings.Split(dir, "github.com/JeffDeCola/")
		dir = parts[1]
	}

	// Extract repo name - get everything before /
	parts := strings.Split(dir, "/")
	repoName := parts[0]
	// Get everything after repo name
	parts = strings.Split(dir, repoName)
	dir = parts[1]

	// STEP 3 Build link
	// DO NOT add /tree/master if the dir string is empty
	link := ""
	if dir == "" {
		link = githubURL + repoName + inputFilename + "#" + headingLower
	} else {
		link = githubURL + repoName + "/tree/master" + dir + inputFilename + "#" + headingLower
	}

	// OUTPUT
	if headingNumber == "h2" {
		fmt.Print("* ")
	}
	if headingNumber == "h3" {
		fmt.Print("  * ")
	}
	fmt.Print("[", heading, "](", link, ")\n")
}

func main() {

	// Flags - Will default to README.md if no input giving
	inputFilenamePtr := flag.String("i", "README.md", "input file")
	heading3Ptr := flag.Bool("h3", false, "a bool for heading2")
	flag.Parse()

	// Do we put this in the link?
	inputFilename := *inputFilenamePtr
	if inputFilename == "README.md" {
		inputFilename = ""
	} else {
		inputFilename = "/" + inputFilename
	}

	heading2 := "## "
	heading3 := "### "

	// Open input file
	inputFile, err := os.Open(*inputFilenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fmt.Println("")

	// Start scanning the input file
	scanner := bufio.NewScanner(inputFile) // Increment the token
	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()
		// fmt.Println("Working on:", line)

		// Find heading 2.
		if strings.Contains(line, heading2) {

			// Is it ## with a space
			if string(line[0:3]) == heading2 {
				line = line[3:]
				makeTOC(line, "h2", inputFilename)
			}

			// Find heading 3
			if strings.Contains(line, heading3) && *heading3Ptr {

				// Is it ### with a space
				if string(line[0:4]) == heading3 {
					line = line[4:]
					makeTOC(line, "h3", inputFilename)
				}
			}

		}

	}

	fmt.Println("")

}

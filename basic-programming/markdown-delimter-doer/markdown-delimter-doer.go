package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type table struct {
	numColumns int
	colWidths  [20]string
	headings   [20]string
	rows       [10]string
	rownumber  int
}

func printLine(line string, outputFile *os.File) {
	fmt.Print(line)
	_, err := outputFile.WriteString(line)
	if err != nil {
		fmt.Println(err)
		outputFile.Close()
		return
	}
}

// Build the struct (One line at a time)
// If blank line it will be ignored.
func buildStruct(t *table, line string) {
	if strings.Contains(line, "tableNumCols") {
		line := strings.Replace(line, "tableNumCols: ", "", -1)
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.numColumns = i
	}
	if strings.Contains(line, "tableClWidth") {
		line := strings.Replace(line, "tableClWidth: ", "", -1)
		foo := strings.Split(line, ",")
		for i, v := range foo {
			t.colWidths[i] = v
		}
	}
	if strings.Contains(line, "tableHeading") {
		line := strings.Replace(line, "tableHeading: ", "", -1)
		foo := strings.Split(line, ",")
		for i, v := range foo {
			t.headings[i] = v
		}
	}
}

// create the html for table
func createTable(t *table, outputFile *os.File) {

	// <table>
	line := "<table style=\"font-size:.8em\">" + "\n"
	printLine(line, outputFile)

	// <col width> - Based on number col
	for i := 0; i < t.numColumns; i++ {
		line := "  <col width=\"" + t.colWidths[i] + "\">" + "\n"
		printLine(line, outputFile)
	}

	// HEADING
	line = "  <tr>" + "\n"
	printLine(line, outputFile)
	for i := 0; i < t.numColumns; i++ {
		line := "    <th>" + t.headings[i] + "</th>" + "\n"
		printLine(line, outputFile)
	}
	line = "  </tr>" + "\n"
	printLine(line, outputFile)

	// ROWS
	line = "  <tr>" + "\n"
	printLine(line, outputFile)

	// COLUMNS
	line = "    <td align=\"left\" valign=\"top\">" + "\n"
	printLine(line, outputFile)

	line = "    </td>" + "\n"
	printLine(line, outputFile)

	line = "  </tr>" + "\n"
	printLine(line, outputFile)

	// <\table>
	line = "</table>" + "\n"
	printLine(line, outputFile)

	/* line = "<table style="font-size:.8em">"
	  _, err := outputFile.WriteString(line)
	  if err != nil {
	      fmt.Println(err)
	      outputFile.Close()
	      return
	  }

	  // </table>
	  line = "</table>"
	  _, err := outputFile.WriteString(line)
	  if err != nil {
	      fmt.Println(err)
	      outputFile.Close()
	      return
	  }

	<tr>
	  <th></th>
	  <th></th>
	  <th>BOUGHT</th>
	  <th>REG</th>
	  <th>WAR</th>
	  <th>GOOD TIL</th>
	</tr>
	*/
}

func main() {

	t := table{}
	t.rownumber = 1

	// Open input file
	inputFile, err := os.Open("input.md")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create("output.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start scanning the file
	scanner := bufio.NewScanner(inputFile)

	// Increment the token
	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()
		// fmt.Println("Working on:", line)

		// If you find a delimiter, do something with it
		if line == "$$" {

			// Stay in here until you find another delimiter
			fmt.Println("START DELIMITER ************************************")
			for scanner.Scan() {

				// Read a line in file
				line := scanner.Text()

				// Exit and build table when you find another delimiter
				if line == "$$" {
					// BUILD TABLE
					createTable(&t, outputFile)
					fmt.Println("END DELIMITER************************************")
					break
				}

				// THE MAGIC - Build your struct
				buildStruct(&t, line)
			}

		} else {

			// Write line to output file
			line := line + "\n"
			fmt.Print(line)
			_, err := outputFile.WriteString(line)
			if err != nil {
				fmt.Println(err)
				outputFile.Close()
				return
			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

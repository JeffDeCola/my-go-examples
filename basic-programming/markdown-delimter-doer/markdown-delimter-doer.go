package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type table struct {
	columns       int
	columnW       [20]string
	columnA       [20]string
	headers       [20]string
	rowColumnLine [40][20][10]string
	rows          int
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

func buildTableStruct(t *table, stuff []string) {

	rowNumber := 0
	columnNumber := 0
	lineNumber := 0
	line := "stuff"

	// Loop over the array
	// Will get line by line
	for i := 0; i < len(stuff); i++ {

		line = stuff[i]

		if strings.Contains(line, "columns") {
			line = strings.Replace(line, "columns: ", "", -1)
			i, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			t.columns = i
		}
		if strings.Contains(line, "columnW") {
			line = strings.Replace(line, "columnW: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.columnW[i] = v
			}
		}
		if strings.Contains(line, "columnA") {
			line = strings.Replace(line, "columnA: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.columnA[i] = v
			}
		}
		if strings.Contains(line, "headers") {
			line = strings.Replace(line, "headers: ", "", -1)
			foo := strings.Split(line, ",")
			for i, v := range foo {
				t.headers[i] = v
			}
		}
		if strings.Contains(line, "rowcol") {
			// Get the column number
			tempcolumnNumberStr := string(line[6])
			tempcolumnNumber, _ := strconv.Atoi(tempcolumnNumberStr)
			if tempcolumnNumber < columnNumber {
				rowNumber++
				lineNumber = 0
			}
			if tempcolumnNumber == columnNumber {
				lineNumber++
			} else {
				lineNumber = 0
			}
			columnNumber = tempcolumnNumber
			replace := "rowcol" + tempcolumnNumberStr + ": "
			line := strings.Replace(line, replace, "", -1)
			fmt.Println(replace, "rowNumber:", rowNumber, "columnNumber:", columnNumber, "lineNumber:", lineNumber, "----", line)
			// Place in 3-D array
			t.rowColumnLine[rowNumber][columnNumber][lineNumber] = line
		}

	}
	t.rows = rowNumber

}

// Make an html table from the stuff
func makeHTMLTABLE(stuff []string, outputFile *os.File) {

	t := table{}

	buildTableStruct(&t, stuff)

	// <table>
	line := "<table style=\"font-size:.8em\">" + "\n"
	printLine(line, outputFile)

	// <col width> - Based on numbers of columns
	line = "  <!-- COLUMN WIDTHS -->" + "\n"
	printLine(line, outputFile)
	for i := 0; i < t.columns; i++ {
		line := "  <col width=\"" + t.columnW[i] + "\">" + "\n"
		printLine(line, outputFile)
	}

	// <tr> - Heading row - Based on numbers of columns
	line = "<!-- HEADING ROW -->" + "\n"
	printLine(line, outputFile)
	line = "  <tr>" + "\n"
	printLine(line, outputFile)
	for i := 0; i < t.columns; i++ {
		line := "    <th>" + t.headers[i] + "</th>" + "\n"
		printLine(line, outputFile)
	}
	line = "  </tr>" + "\n"
	printLine(line, outputFile)

	// LETS DO THE ROWS - YEAH
	// ROWS - ITERATE OVER
	for r := 0; r < t.rows+1; r++ {
		line = "  <!-- ROW -->" + "\n"
		printLine(line, outputFile)
		line = "  <tr>" + "\n"
		printLine(line, outputFile)
		// COLUMNS - ITERATE OVER
		for c := 1; c < t.columns+1; c++ {
			alignment := t.columnA[c-1]
			line = "    <td  align=\"" + alignment + "\" valign=\"top\">" + "\n"
			printLine(line, outputFile)
			// LINES - ITERATE OVER
			// Keep iterating until you get nothing
			for l := 0; l < 100; l++ {
				linebreak := ""
				// fmt.Println("r=", r, "c=", c, "l=", l)
				preline := t.rowColumnLine[r][c][l]
				// If blank do not print
				if preline != "" {
					if (l > 0) || (t.rowColumnLine[r][c][l+1] != "") {
						linebreak = "<br>"
					}
					line = "      " + preline + linebreak + "\n"
					printLine(line, outputFile)
				} else {
					line = "    </td>" + "\n"
					printLine(line, outputFile)
					break
				}
			}
		}
		line = "  </tr>" + "\n"
		printLine(line, outputFile)
	}

	// <\table>
	line = "</table>" + "\n"
	printLine(line, outputFile)
}

func main() {

	// Flags
	delimeterPtr := flag.String("delimeter", "DELIMETER", "what is the delimeter")
	inputFilenamePtr := flag.String("i", "INPUT", "input file")
	outputFilenamePtr := flag.String("o", "OUTPUT", "output file")
	htmlTableBoolPtr := flag.Bool("htmltable", false, "a bool")
	flag.Parse()

	// Temp storage for what you want to process between the delimeters
	var stuff []string

	// Open input file
	inputFile, err := os.Open(*inputFilenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create(*outputFilenamePtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start scanning the input file
	scanner := bufio.NewScanner(inputFile) // Increment the token
	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()
		// fmt.Println("Working on:", line)

		// If you find a delimiter, get all the lines in between and place in a table.
		if line == *delimeterPtr {

			// Stay in here until you find another delimiter
			for scanner.Scan() {

				// Read a line in file
				line := scanner.Text()

				// Exit and build table when you find another delimiter
				if line == *delimeterPtr {
					break
				}

				// Place the line the line array
				stuff = append(stuff, line)
			}

			// OK WE HAVE THE LINE ARRAY (Stuff between the delimeters)
			// htmltable switch
			if *htmlTableBoolPtr {

				fmt.Println("MAKE HTML TABLE")
				makeHTMLTABLE(stuff, outputFile)
				fmt.Println("END MAKE HTML TABLE")
			}

			// Did not find a delimiter on this line
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

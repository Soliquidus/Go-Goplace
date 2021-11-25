package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine replaces old occurrences in line by new ones
// If found is true and the pattern found, res with the resulting string and
// occ returns the number of occurrences.
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

// FindReplaceFile reads the source file and uses ProcessLine function in order to process text and
// returning new text
func FindReplaceFile(src string, old string, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	// Prevent text manipulation errors like a remplacement for Google
	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Println(res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	o := "Go"
	n := "Python"
	occ, lines, err := FindReplaceFile("wikigo.txt", o, n)
	if err != nil {
		fmt.Printf("Error while executing FindReplaceFile: %v\n", err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==")
	fmt.Printf("Number of occurences of %v: %v\n", o, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Print("Lines : [ ")
	leng := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < leng-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")
}

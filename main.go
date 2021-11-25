package main

import (
	"fmt"
	"strings"
)

// ProcessLine replace old occurrences in line by new ones
// If found is true and the pattern found, res with the resulting string and
// occ returns the number of occurrences.
func ProcessLine(line, old, new string) (found bool, res string, occ int){
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) ||strings.Contains(line, oldLower){
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func main() {
	found, res, occ := ProcessLine("Go was conceived in 2007 to improve programming productivity at Google, in an era of multicore processors, computer networks, and large codebases.[17] The designers wanted to resolve criticisms of other languages, while retaining many of their useful characteristics:[18]",
		"Go", "Python")
	fmt.Println(found, res, occ)
}

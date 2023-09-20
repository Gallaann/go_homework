package main

import (
	"flag"
	"fmt"
	"strings"
)

type flags struct {
	count      bool
	duplicates bool
	unique     bool
	ignoreCase bool
	fields     int
	chars      int
}

func parseFlags() flags {
	var flags flags
	flag.BoolVar(&flags.count, "c", false, "count repeated lines")
	flag.BoolVar(&flags.duplicates, "d", false, "output only duplicate lines")
	flag.BoolVar(&flags.unique, "u", false, "output only unique lines")
	flag.BoolVar(&flags.ignoreCase, "i", false, "ignore case when comparing lines")
	flag.IntVar(&flags.fields, "f", 0, "skip first num fields in each line")
	flag.IntVar(&flags.chars, "s", 0, "skip first num chars in each line")
	flag.Parse()
	return flags
}

func parseArguments() (string, string) {
	args := flag.Args()

	switch len(args) {
	case 0:
		return "", ""
	case 1:
		return args[0], ""
	default:
		return args[0], args[1]
	}
}

func checkFlags(flags flags) bool {
	sum := 0
	if flags.count {
		sum++
	}
	if flags.duplicates {
		sum++
	}
	if flags.unique {
		sum++
	}
	if sum != 1 {
		return false
	}

	return true
}

func skipFields(line string, numFields int) string {
	fields := strings.Fields(line)
	if len(fields) > numFields {
		return strings.Join(fields[numFields:], " ")
	}
	return ""
}

func skipChars(line string, numChars int) string {
	if len(line) > numChars {
		return line[numChars:]
	}
	return ""
}

func main() {
	flags := parseFlags()
	inputFile, outputFile := parseArguments()

	fmt.Println(flags, inputFile, outputFile)
}

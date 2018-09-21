package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"github.com/araddon/dateparse"
	"flag"
	"strings"
)

var (
	optInputFile            = ""
	optDeltaFormat          = "%9d"
	optDeltaThreshold int64 = 0
	optDateLength           = 23
)
func main() {
	flag.StringVar(&optInputFile, "input", "", "Input file path. Stdin will be used if its option is ommited")
	flag.StringVar(&optDeltaFormat, "delta-format", optDeltaFormat, "Timestamp delta output format in golang Sprinf() syntax")
	flag.Int64Var(&optDeltaThreshold, "threshold", optDeltaThreshold, "Filter lines with timestamp delta bigger than specified threshold")
	flag.IntVar(&optDateLength, "date-length", optDateLength, "Length of date string from line start")
	flag.Parse()

	var reader *bufio.Reader
	if optInputFile != "" {
		f, err := os.Open(optInputFile)
		if err != nil {
			panic(fmt.Sprintf("Error open file: %s", optInputFile))
		}
		defer f.Close()
		reader = bufio.NewReader(f)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}


	var prevTs int64 = -1
	spacer := fmt.Sprintf(strings.Replace(optDeltaFormat, "d", "s", -1), "")
	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %v", err))
			os.Exit(1)
		}

		var delta int64
		var deltaStr string
		if ts, ok := parseDate(text, optDateLength); ok {
			if prevTs == -1 {
				delta = 0
			} else {
				delta = ts - prevTs
			}
			prevTs = ts

			deltaStr = fmt.Sprintf(optDeltaFormat, delta)
		} else {
			deltaStr = spacer
		}

		if delta >= optDeltaThreshold {
			fmt.Print(deltaStr)
			fmt.Print(" ")
			fmt.Print(text)
		}
	}
}

func parseDate(logLine string, dateLength int) (int64, bool) {
	if len(logLine) < dateLength {
		return 0, false
	}

	if t, err := dateparse.ParseLocal(logLine[0 : dateLength]); err != nil {
		return 0, false
	} else {
		return t.UnixNano()/1E6, true // convert to milliseconds
	}
}
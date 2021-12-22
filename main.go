package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type DSMainState struct {
    records [][]string // could be string or int or double mostly
    headers []string
}

var VALID_HEADERS = map[string]int{"open": 1, "high": 1, "low": 1, "close": 1, "volume": 1}

func main() {
    var filepath *string
    filepath = flag.String("csv", "input.csv", "Specify path to the stock csv file")
    flag.Parse()
    f, err := os.Open(*filepath)
    if err != nil {
        fmt.Printf("Can not open file %s\n", *filepath)
        log.Fatal(err)
    }
    r := csv.NewReader(f)
    lines, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
}

// parse in lines
func (d *DSMainState) New(lines [][]string) error {
    if err := ValidateHeaders(lines[0]); err != nil {
        return err
    }
    d.headers = lines[0]
    d.records = lines[1:]
    return nil
}

// validate if the line is a valid header line
func ValidateHeaders(line []string) (error) {
    for _, h := range line {
        if lkup, ok := VALID_HEADERS[h]; lkup != 1 || !ok {
            return errors.New("Invalid headers")
        }
    }
    return nil
}

// Print out data
func (d DSMainState) Print() {
    io.WriteString(os.Stdout, "   " + strings.Join(d.headers, " | ") + "\n")
    for _, r := range d.records {
        io.WriteString(os.Stdout,  " | " + strings.Join(r, " | ") + "\n")
    }
}

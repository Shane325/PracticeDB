package main

import (
    "encoding/csv"
    "io"
    "log"
    "os"
)

type Scanner struct {
    tuples []Tuple
    idx int
}

func newScanner(filename string) *Scanner {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    csvReader := csv.NewReader(file)
    headers, err := csvReader.Read()
    if err != nil {
        log.Fatal(err)
    }

    var tuples []Tuple
    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        var tuple Tuple
        for idx, value := range record {
            tuple.values = append(tuple.values, Value{name: headers[idx], value: value})
        }
        tuples = append(tuples, tuple)
    }

    return &Scanner{tuples: tuples, idx: -1}
}

func (s *Scanner) next() bool {
    s.idx++
    return s.idx < len(s.tuples)
}

func (s *Scanner) execute() Tuple {
    return s.tuples[s.idx]
}

func (s *Scanner) close() {
    s.idx = -1
    s.tuples = nil
}

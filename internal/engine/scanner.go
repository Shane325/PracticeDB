package engine

import "github.com/shane325/PracticeDB/internal/plan"

import (
    "encoding/csv"
    "io"
    "log"
    "os"
)

type Scanner struct {
    tuples []plan.Tuple
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

    var tuples []plan.Tuple
    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        var tuple plan.Tuple
        for idx, Value := range record {
            tuple.Values = append(tuple.Values, plan.Value{Name: headers[idx], Value: Value})
        }
        tuples = append(tuples, tuple)
    }

    return &Scanner{tuples: tuples, idx: -1}
}

func (s *Scanner) next() bool {
    s.idx++
    return s.idx < len(s.tuples)
}

func (s *Scanner) execute() plan.Tuple {
    return s.tuples[s.idx]
}

func (s *Scanner) close() {
    s.idx = -1
    s.tuples = nil
}

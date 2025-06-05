package execution

import (
    "github.com/shane325/PracticeDB/internal/plan"
    "sort"
    "strings"
)

type Sort struct {
    source Iterator
    tuples []plan.Tuple
    idx int
}

func NewSort(field string, desc bool, source Iterator) *Sort {
    var tuples []plan.Tuple
    for source.Next() {
        tuple := source.Execute()
        tuples = append(tuples, tuple)
    }

    sortTuplesByField(tuples, field, desc)

    return &Sort{source: source, tuples: tuples, idx: -1}
}

func getValueByName(t plan.Tuple, name string) (string, bool) {
    for _, v := range t.Values {
        if v.Name == name {
            return v.Value, true
        }
    }
    return "", false
}

func sortTuplesByField(tuples []plan.Tuple, fieldName string, desc bool) {
    sort.Slice(tuples, func(i, j int) bool {
        vi, _ := getValueByName(tuples[i], fieldName)
        vj, _ := getValueByName(tuples[j], fieldName)
        // Just handling strings for now
        if (desc) {
            return strings.Compare(vi, vj) > 0
        } else {
            return strings.Compare(vi, vj) < 0
        }
    })
}

func (s *Sort) Next() bool {
    s.idx++
    return s.idx < len(s.tuples)
}

func (s *Sort) Execute() plan.Tuple {
    return s.tuples[s.idx]
}

func (s *Sort) Close() {
    s.source = nil
    s.tuples = nil
    s.idx = -1
}

package expressions

import "github.com/shane325/PracticeDB/internal/plan"

type Expression interface {
    Execute(plan.Tuple) bool
}

type Equals struct {
    field string
    value string
}

type NotEquals struct {
    field string
    value string
}

func NewEquals(field string, value string) *Equals {
    return &Equals{field: field, value: value}
}

func (e *Equals) Execute(tuple plan.Tuple) bool {
    for _, val := range tuple.Values {
        if (val.Name == e.field && val.Value == e.value) {
            return true
        }
    }
    return false
}

func NewNotEquals(field string, value string) *NotEquals {
    return &NotEquals{field: field, value: value}
}

func (ne *NotEquals) Execute(tuple plan.Tuple) bool {
    for _, val := range tuple.Values {
        if (val.Name == ne.field && val.Value != ne.value) {
            return true
        }
    }
    return false
}

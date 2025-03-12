package main

type Expression interface {
    execute(Tuple) bool
}

type Equals struct {
    field string
    value string
}

type NotEquals struct {
    field string
    value string
}

func newEquals(field string, value string) *Equals {
    return &Equals{field: field, value: value}
}

func (e *Equals) execute(tuple Tuple) bool {
    for _, val := range tuple.values {
        if (val.name == e.field && val.value == e.value) {
            return true
        }
    }
    return false
}

func newNotEquals(field string, value string) *NotEquals {
    return &NotEquals{field: field, value: value}
}

func (ne *NotEquals) execute(tuple Tuple) bool {
    for _, val := range tuple.values {
        if (val.name == ne.field && val.value != ne.value) {
            return true
        }
    }
    return false
}

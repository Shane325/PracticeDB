package main

type Iterator interface {
    next() bool
    execute() Tuple
    close()
}

type Tuple struct {
    values []Value
}

type Value struct {
    name string
    value string
}

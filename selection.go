package main

type Selection struct {
    expression Expression
    source Iterator
    current Tuple
}

func newSelection(expression Expression, source Iterator) *Selection {
    return &Selection{expression: expression, source: source, current: Tuple{}}
}

func (s *Selection) next() bool {
    for s.source.next() {
        tuple := s.source.execute()
        if (s.expression.execute(tuple)) {
            s.current = tuple
            return true
        }
    }
    return false
}

func (s *Selection) execute() Tuple {
    return s.current
}

func (s *Selection) close() {
    s.source = nil
    s.current = Tuple{}
}

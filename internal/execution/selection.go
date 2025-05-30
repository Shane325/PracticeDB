package execution

import (
    "github.com/shane325/PracticeDB/internal/expressions"
    "github.com/shane325/PracticeDB/internal/plan"
)

type Selection struct {
    expression expressions.Expression
    source Iterator
    current plan.Tuple
}

func newSelection(expression expressions.Expression, source Iterator) *Selection {
    return &Selection{expression: expression, source: source, current: plan.Tuple{}}
}

func (s *Selection) next() bool {
    for s.source.next() {
        tuple := s.source.execute()
        if (s.expression.Execute(tuple)) {
            s.current = tuple
            return true
        }
    }
    return false
}

func (s *Selection) execute() plan.Tuple {
    return s.current
}

func (s *Selection) close() {
    s.source = nil
    s.current = plan.Tuple{}
}

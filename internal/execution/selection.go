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

func NewSelection(expression expressions.Expression, source Iterator) *Selection {
    return &Selection{expression: expression, source: source, current: plan.Tuple{}}
}

func (s *Selection) Next() bool {
    for s.source.Next() {
        tuple := s.source.Execute()
        if (s.expression.Execute(tuple)) {
            s.current = tuple
            return true
        }
    }
    return false
}

func (s *Selection) Execute() plan.Tuple {
    return s.current
}

func (s *Selection) Close() {
    s.source = nil
    s.current = plan.Tuple{}
}

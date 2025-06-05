package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Limit struct {
    limit int
    count int
    source Iterator
    current plan.Tuple
}

func NewLimit(limit int, source Iterator) *Limit {
    return &Limit{limit: limit, count: 0, source: source, current: plan.Tuple{}}
}

func (l *Limit) Next() bool {
    if (l.source.Next() && l.count < l.limit) {
        l.count++
        tuple := l.source.Execute()
        l.current = tuple
        return true
    }
    return false
}

func (l *Limit) Execute() plan.Tuple {
    return l.current
}

func (l *Limit) Close() {
    l.limit = 0
    l.count = 0
    l.source = nil
    l.current = plan.Tuple{}
}

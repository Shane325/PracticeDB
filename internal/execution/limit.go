package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Limit struct {
    limit int
    count int
    source Iterator
    current plan.Tuple
}

func newLimit(limit int, source Iterator) *Limit {
    return &Limit{limit: limit, count: 0, source: source, current: plan.Tuple{}}
}

func (l *Limit) next() bool {
    if (l.source.next() && l.count < l.limit) {
        l.count++
        tuple := l.source.execute()
        l.current = tuple
        return true
    }
    return false
}

func (l *Limit) execute() plan.Tuple {
    return l.current
}

func (l *Limit) close() {
    l.limit = 0
    l.count = 0
    l.source = nil
    l.current = plan.Tuple{}
}

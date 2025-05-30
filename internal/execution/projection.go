package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Projection struct {
    field string
    source Iterator
    current plan.Tuple
}

func newProjection(field string, source Iterator) *Projection {
    return &Projection{field: field, source: source, current: plan.Tuple{}}
}

func (p *Projection) next() bool {
    if (p.source.next()) {
        tuple := p.source.execute()
        p.current = tuple
        return true
    }
    return false
}

func (p *Projection) execute() plan.Tuple {
    for _, val := range p.current.Values {
        if val.Name == p.field {
            newTuple := plan.Tuple{}
            newTuple.Values = append(newTuple.Values, plan.Value{Name: val.Name, Value: val.Value})
            return newTuple
        }
    }
    return plan.Tuple{}
}

func (p *Projection) close() {
    p.field = ""
    p.source = nil
    p.current = plan.Tuple{}
}

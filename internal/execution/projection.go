package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Projection struct {
    field string
    source Iterator
    current plan.Tuple
}

func NewProjection(field string, source Iterator) *Projection {
    return &Projection{field: field, source: source, current: plan.Tuple{}}
}

func (p *Projection) Next() bool {
    if (p.source.Next()) {
        tuple := p.source.Execute()
        p.current = tuple
        return true
    }
    return false
}

func (p *Projection) Execute() plan.Tuple {
    for _, val := range p.current.Values {
        if val.Name == p.field {
            newTuple := plan.Tuple{}
            newTuple.Values = append(newTuple.Values, plan.Value{Name: val.Name, Value: val.Value})
            return newTuple
        }
    }
    return plan.Tuple{}
}

func (p *Projection) Close() {
    p.field = ""
    p.source = nil
    p.current = plan.Tuple{}
}

package main

type Projection struct {
    field string
    source Iterator
    current Tuple
}

func newProjection(field string, source Iterator) *Projection {
    return &Projection{field: field, source: source, current: Tuple{}}
}

func (p *Projection) next() bool {
    if (p.source.next()) {
        tuple := p.source.execute()
        p.current = tuple
        return true
    }
    return false
}

func (p *Projection) execute() Tuple {
    for _, val := range p.current.values {
        if val.name == p.field {
            newTuple := Tuple{}
            newTuple.values = append(newTuple.values, Value{name: val.name, value: val.value})
            return newTuple
        }
    }
    return Tuple{}
}

func (p *Projection) close() {
    p.field = ""
    p.source = nil
    p.current = Tuple{}
}

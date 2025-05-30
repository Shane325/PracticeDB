package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Iterator interface {
    next() bool
    execute() plan.Tuple
    close()
}

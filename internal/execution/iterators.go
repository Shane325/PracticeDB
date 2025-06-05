package execution

import "github.com/shane325/PracticeDB/internal/plan"

type Iterator interface {
    Next() bool
    Execute() plan.Tuple
    Close()
}

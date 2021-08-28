package golden_gate

import (
	"fmt"

	"github.com/Open-Science-Global/poly"
	"github.com/Open-Science-Global/poly/clone"
)

func SimulateGoldenGate(sequences []poly.Sequence, enzymeStr string) []clone.Part {
	var fragments []clone.Part
	for _, sequence := range sequences {
		fragments = append(fragments, clone.Part{sequence.Sequence, sequence.Meta.Locus.Circular})
	}

	clonedFragments, err := clone.GoldenGate(fragments, enzymeStr)
	if err != nil {
		fmt.Println(err)
	}
	if len(clonedFragments) != 1 {
		fmt.Println("More than 1 cloned sequence simulated in basic genes")
	}
	return clonedFragments
}

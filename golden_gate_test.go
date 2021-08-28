package golden_gate

import (
	"fmt"
	"io/ioutil"

	"github.com/Open-Science-Global/poly"
	"github.com/Open-Science-Global/poly/io/genbank"
)

func ExampleSimulateGoldenGateLinearFragments() {

	dir, _ := ioutil.ReadDir("./data")

	var files []string
	for _, f := range dir {
		files = append(files, f.Name())
	}

	var sequences []poly.Sequence
	for _, file := range files {
		sequences = append(sequences, genbank.Read("./data/"+file))
	}
	finalPlasmid := SimulateGoldenGateLinearFragments(sequences, "BsaI")
	fmt.Println(len(finalPlasmid[0].Sequence), finalPlasmid[0].Circular)
	// Output: 10363 true
}

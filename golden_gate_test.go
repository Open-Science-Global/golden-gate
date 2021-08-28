package golden_gate

import (
	"fmt"
	"io/ioutil"

	"github.com/Open-Science-Global/poly"
	"github.com/Open-Science-Global/poly/io/genbank"
)

func ExampleSimulateGoldenGateLinearFragments() {

	root := "./data/linearSequences/"

	dir, _ := ioutil.ReadDir(root)

	var files []string
	for _, f := range dir {
		files = append(files, f.Name())
	}

	var sequences []poly.Sequence
	for _, file := range files {
		sequences = append(sequences, genbank.Read(root+file))
	}
	finalPlasmid := SimulateGoldenGate(sequences, "BsaI")
	fmt.Println(len(finalPlasmid[0].Sequence), finalPlasmid[0].Circular)
	// Output: 10363 true
}

func ExampleSimulateGoldenGateCircularFragments() {

	root := "./data/circularSequences/"

	dir, _ := ioutil.ReadDir(root)

	var files []string
	for _, f := range dir {
		files = append(files, f.Name())
	}

	var sequences []poly.Sequence
	for _, file := range files {
		sequences = append(sequences, genbank.Read(root+file))
	}
	finalPlasmid := SimulateGoldenGate(sequences, "BbsI")
	fmt.Println(finalPlasmid)
	// Output: More than 1 cloned sequence simulated in basic genes
	// []
}

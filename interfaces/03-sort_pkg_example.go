//Looking inside the golang standard packages you can found examples
//If we focus on sort package there is an interface that is used to implement sorting algorithms
/*type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

//So we need to implement these three methods in owr types in order to use the sort.Sort function
//and inside the sort package there is this gem example making different sorting criteria for different types
package sort_test

import (
	"fmt"
	"sort"
)

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

type Organ struct {
	Name   string
	Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

// ByWeight implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func Example_sortWrapper() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight:")
	printOrgans(s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)

	// Output:
	// Organs by weight:
	// prostate (62g)
	// pancreas (131g)
	// spleen   (162g)
	// heart    (290g)
	// brain    (1340g)
	// liver    (1494g)
	// Organs by name:
	// brain    (1340g)
	// heart    (290g)
	// liver    (1494g)
	// pancreas (131g)
	// prostate (62g)
	// spleen   (162g)
}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}

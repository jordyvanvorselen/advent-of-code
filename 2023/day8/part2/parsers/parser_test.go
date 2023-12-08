package parsers

import (
	"github.com/emirpasic/gods/trees/avltree"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedInstructions := "RLRLLLR"
	expectedTree := avltree.NewWithStringComparator()

	expectedTree.Put("AAA", []string{"BBB", "CCC"})
	expectedTree.Put("BBB", []string{"DDD", "EEE"})
	expectedTree.Put("CCC", []string{"ZZZ", "GGG"})
	expectedTree.Put("DDD", []string{"DDD", "DDD"})
	expectedTree.Put("EEE", []string{"EEE", "EEE"})
	expectedTree.Put("GGG", []string{"GGG", "GGG"})
	expectedTree.Put("ZZZ", []string{"ZZZ", "ZZZ"})

	input := []string{
		"RLRLLLR",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	resultInstructions, resultTree := Parse(input)

	if expectedInstructions != resultInstructions {
		t.Errorf("got %s but expected %s", resultInstructions, expectedInstructions)
	}

	if !reflect.DeepEqual(resultTree.Keys(), expectedTree.Keys()) || !reflect.DeepEqual(resultTree.Values(), expectedTree.Values()) {
		t.Errorf("got %+v but expected %+v", resultTree, expectedTree)
	}
}

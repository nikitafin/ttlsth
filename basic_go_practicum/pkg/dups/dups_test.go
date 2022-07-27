package dups

import (
	"basic_go_practicum/pkg"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	res := RemoveDuplicates(input)
	if !pkg.Equal(res, []string{"cat", "dog", "bird", "parrot"}) {
		t.Error("RemoveDup Invalid Output")
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(input)
	}
}

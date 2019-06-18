package main

import (
	"fmt"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

func Test(t *testing.T) {
	a := `
	hoge
	fuga
	piyo
	`

	b := `
	hoga
	fuga
	pipa
	piyo
	`

	diff := difflib.UnifiedDiff{
		A: difflib.SplitLines(a),
		B: difflib.SplitLines(b),
	}

	text, _ := difflib.GetUnifiedDiffString(diff)
	fmt.Printf(text)

}

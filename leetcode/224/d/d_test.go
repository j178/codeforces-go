// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`["####F","#C...","M...."]`, `1`, `2`,
			`true`,
		},
		{
			`["M.C...F"]`, `1`, `4`,
			`true`,
		},
		{
			`["M.C...F"]`, `1`, `3`,
			`false`,
		},
		{
			`["C...#","...#F","....#","M...."]`, `2`, `5`,
			`false`,
		},
		{
			`[".M...","..#..","#..#.","C#.#.","...#F"]`, `3`, `1`,
			`true`,
		},
		// TODO 测试入参最小的情况
		{
			`["CMF"]`, `1`, `1`,
			`true`,
		},
		{
			`["...CM.F"]`, `1`, `2`,
			`true`,
		},
		{
			`["..##F","...#.","MC..."]`, `1`, `3`,
			`true`,
		},
		{
			`["CM......","#######.","........",".#######","........","#######.","........","F#######"]`, `1`, `1`,
			`true`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, canMouseWin, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-224/problems/cat-and-mouse-ii/

// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`1`, `10`, 
			`2`,
		},
		{
			`5`, `15`, 
			`2`,
		},
		{
			`19`, `28`, 
			`2`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countBalls, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-226/problems/maximum-number-of-balls-in-a-box/
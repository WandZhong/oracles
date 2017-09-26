package dbu

import (
	//	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type ArraySuite struct{}

func checkNestedArray(src string, expected []string, c *C, comment interface{}) {
	var resb = SplitNestedArray([]byte(src))
	var res = make([]string, len(resb))
	for i := range resb {
		res[i] = string(resb[i])
	}
	c.Assert(res, DeepEquals, expected, comment)
}

func (suite *ArraySuite) TestSplitNestedArray(c *C) {
	checkNestedArray("{}", []string{}, c,
		Commentf("Empty array should work"))

	checkNestedArray("{1,2,3}", []string{"1", "2", "3"}, c,
		Commentf("1-dimension int array should work"))
	checkNestedArray("{1, 2,3}", []string{"1", " 2", "3"}, c,
		Commentf("1-dimension int array should work"))
	checkNestedArray("{10,12,   3  }", []string{"10", "12", "   3  "}, c,
		Commentf("1-dimension int array should work"))
	checkNestedArray("{ 0   ,  001200}", []string{" 0   ", "  001200"}, c,
		Commentf("1-dimension int array should work"))

	checkNestedArray("{ {0}   ,  { 001200,1}}", []string{" {0}   ", "  { 001200,1}"}, c,
		Commentf("2-dimension int array should work"))
	checkNestedArray("{ {0}   ,  { 001200,1} }", []string{" {0}   ", "  { 001200,1} "}, c,
		Commentf("2-dimension int array should work"))

	checkNestedArray("{ {{ 1 },{1 2, 3}}, {}, {{ 001200,1}} }", []string{" {{ 1 },{1 2, 3}}", " {}", " {{ 001200,1}} "}, c,
		Commentf("3-dimension int array should work"))
}

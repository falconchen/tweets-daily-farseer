package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	t.Log("Start regex testing")
	{
		testString := "aa99d3caa88aaaaab99dac99d"
		re, err := regexp.Compile(`[a-z]{1,4}(\d+)(d)`)
		if err != nil {
			t.Fatal(err)
		}
		all := re.FindAllStringSubmatch(testString, -1)
		fmt.Println(all)

		all = re.FindAllStringSubmatch(testString, 1)
		fmt.Println(all)
		//for _, b := range all {
		//	fmt.Println(string(b))
		//}

		//re, err = regexp.CompilePOSIX(`[a-z]{2}`)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//str = re.FindAllString(testString, -1)
		//fmt.Println(str)

	}
}
